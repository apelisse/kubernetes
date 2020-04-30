/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fieldmanager_test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apiserver/pkg/endpoints/handlers/fieldmanager"
	"k8s.io/apiserver/pkg/endpoints/handlers/fieldmanager/internal"

	"k8s.io/kube-openapi/pkg/util/proto"
	prototesting "k8s.io/kube-openapi/pkg/util/proto/testing"
	"sigs.k8s.io/structured-merge-diff/v3/fieldpath"
	"sigs.k8s.io/structured-merge-diff/v3/merge"
	"sigs.k8s.io/structured-merge-diff/v3/typed"
	"sigs.k8s.io/yaml"
)

var fakeSchema = prototesting.Fake{
	Path: filepath.Join(
		strings.Repeat(".."+string(filepath.Separator), 8),
		"api", "openapi-spec", "swagger.json"),
}

type fakeObjectConvertor struct {
	converter  merge.Converter
	apiVersion fieldpath.APIVersion
}

func (c *fakeObjectConvertor) Convert(in, out, context interface{}) error {
	if typedValue, ok := in.(*typed.TypedValue); ok {
		var err error
		out, err = c.converter.Convert(typedValue, c.apiVersion)
		return err
	}
	out = in
	return nil
}

func (c *fakeObjectConvertor) ConvertToVersion(in runtime.Object, _ runtime.GroupVersioner) (runtime.Object, error) {
	return in, nil
}

func (c *fakeObjectConvertor) ConvertFieldLabel(_ schema.GroupVersionKind, _, _ string) (string, string, error) {
	return "", "", errors.New("not implemented")
}

type fakeObjectDefaulter struct{}

func (d *fakeObjectDefaulter) Default(in runtime.Object) {}

type TestFieldManager struct {
	fieldManager fieldmanager.Manager
	emptyObj     runtime.Object
	liveObj      runtime.Object
}

func NewTestFieldManager(gvk schema.GroupVersionKind) TestFieldManager {
	m := NewFakeOpenAPIModels()
	tc := NewFakeTypeConverter(m)

	converter := internal.NewVersionConverter(tc, &fakeObjectConvertor{}, gvk.GroupVersion())
	apiVersion := fieldpath.APIVersion(gvk.GroupVersion().String())
	f, err := fieldmanager.NewStructuredMergeManager(
		m,
		&fakeObjectConvertor{converter, apiVersion},
		&fakeObjectDefaulter{},
		gvk.GroupVersion(),
		gvk.GroupVersion(),
	)
	if err != nil {
		panic(err)
	}
	live := &unstructured.Unstructured{}
	live.SetKind(gvk.Kind)
	live.SetAPIVersion(gvk.GroupVersion().String())
	f = fieldmanager.NewStripMetaManager(f)
	f = fieldmanager.NewManagedFieldsUpdater(f)
	f = fieldmanager.NewBuildManagerInfoManager(f, gvk.GroupVersion())
	return TestFieldManager{
		fieldManager: f,
		emptyObj:     live,
		liveObj:      live.DeepCopyObject(),
	}
}

func NewFakeTypeConverter(m proto.Models) internal.TypeConverter {
	tc, err := internal.NewTypeConverter(m, false)
	if err != nil {
		panic(fmt.Sprintf("Failed to build TypeConverter: %v", err))
	}
	return tc
}

func NewFakeOpenAPIModels() proto.Models {
	d, err := fakeSchema.OpenAPISchema()
	if err != nil {
		panic(err)
	}
	m, err := proto.NewOpenAPIData(d)
	if err != nil {
		panic(err)
	}
	return m
}

func (f *TestFieldManager) Reset() {
	f.liveObj = f.emptyObj.DeepCopyObject()
}

func (f *TestFieldManager) Apply(obj runtime.Object, manager string, force bool) error {
	out, err := fieldmanager.NewFieldManager(f.fieldManager).Apply(f.liveObj, obj, manager, force)
	if err == nil {
		f.liveObj = out
	}
	return err
}

func (f *TestFieldManager) Update(obj runtime.Object, manager string) error {
	out, err := fieldmanager.NewFieldManager(f.fieldManager).Update(f.liveObj, obj, manager)
	if err == nil {
		f.liveObj = out
	}
	return err
}

func (f *TestFieldManager) ManagedFields() []metav1.ManagedFieldsEntry {
	accessor, err := meta.Accessor(f.liveObj)
	if err != nil {
		panic(fmt.Errorf("couldn't get accessor: %v", err))
	}

	return accessor.GetManagedFields()
}

// TestUpdateApplyConflict tests that applying to an object, which
// wasn't created by apply, will give conflicts
func TestUpdateApplyConflict(t *testing.T) {
	f := NewTestFieldManager(schema.FromAPIVersionAndKind("apps/v1", "Deployment"))

	patch := []byte(`{
		"apiVersion": "apps/v1",
		"kind": "Deployment",
		"metadata": {
			"name": "deployment",
			"labels": {"app": "nginx"}
		},
		"spec": {
                        "replicas": 3,
                        "selector": {
                                "matchLabels": {
                                         "app": "nginx"
                                }
                        },
                        "template": {
                                "metadata": {
                                        "labels": {
                                                "app": "nginx"
                                        }
                                },
                                "spec": {
				        "containers": [{
					        "name":  "nginx",
					        "image": "nginx:latest"
				        }]
                                }
                        }
		}
	}`)
	newObj := &unstructured.Unstructured{Object: map[string]interface{}{}}
	if err := yaml.Unmarshal(patch, &newObj.Object); err != nil {
		t.Fatalf("error decoding YAML: %v", err)
	}

	if err := f.Update(newObj, "fieldmanager_test"); err != nil {
		t.Fatalf("failed to apply object: %v", err)
	}

	appliedObj := &unstructured.Unstructured{Object: map[string]interface{}{}}
	if err := yaml.Unmarshal([]byte(`{
		"apiVersion": "apps/v1",
		"kind": "Deployment",
		"metadata": {
			"name": "deployment",
		},
		"spec": {
			"replicas": 101,
		}
	}`), &appliedObj.Object); err != nil {
		t.Fatalf("error decoding YAML: %v", err)
	}

	err := f.Apply(appliedObj, "fieldmanager_conflict", false)
	if err == nil || !apierrors.IsConflict(err) {
		t.Fatalf("Expecting to get conflicts but got %v", err)
	}
}

func TestInvalidManagedFieldsIsDropped(t *testing.T) {
	f := NewTestFieldManager(schema.FromAPIVersionAndKind("apps/v1", "Deployment"))

	patch := []byte(`{
		"apiVersion": "apps/v1",
		"kind": "Deployment",
		"metadata": {
			"name": "deployment",
			"labels": {"app": "nginx"}
		},
		"spec": {
                        "replicas": 3,
                        "selector": {
                                "matchLabels": {
                                         "app": "nginx"
                                }
                        },
                        "template": {
                                "metadata": {
                                        "labels": {
                                                "app": "nginx"
                                        }
                                },
                                "spec": {
				        "containers": [{
					        "name":  "nginx",
					        "image": "nginx:latest"
				        }]
                                }
                        }
		}
	}`)
	newObj := &unstructured.Unstructured{Object: map[string]interface{}{}}
	if err := yaml.Unmarshal(patch, &newObj.Object); err != nil {
		t.Fatalf("error decoding YAML: %v", err)
	}

	if err := f.Update(newObj, "fieldmanager_test"); err != nil {
		t.Fatalf("failed to update object: %v", err)
	}

	liveObj := f.liveObj.DeepCopyObject()
	accessor, err := meta.Accessor(liveObj)
	if err != nil {
		panic(fmt.Errorf("couldn't get accessor: %v", err))
	}
	// Make managed fields invalid (missing FieldsType)
	managedFields := accessor.GetManagedFields()
	managedFields[0].FieldsType = ""
	managedFields[0].Manager = "replacement"
	accessor.SetManagedFields(managedFields)
	// Make a change to the label.
	labels := accessor.GetLabels()
	labels["app"] = "my-nginx"
	accessor.SetLabels(labels)

	if err := f.Update(liveObj, "fieldmanager_test_2"); err != nil {
		t.Fatalf("failed to update object: %v", err)
	}

	managedFields = f.ManagedFields()
	if len(managedFields) != 2 {
		t.Fatalf("Expected 2 managedfields got: %v", len(managedFields))
	}
	if managedFields[0].Manager != "fieldmanager_test" {
		t.Fatalf("Expected first manager to be 'fieldmanager_test', got: %v", managedFields[0].Manager)
	}
	if managedFields[1].Manager != "fieldmanager_test_2" {
		t.Fatalf("Expected second manager to be 'fieldmanager_test_2', got: %v", managedFields[0].Manager)
	}
}

func TestApplyStripsFields(t *testing.T) {
	f := NewTestFieldManager(schema.FromAPIVersionAndKind("apps/v1", "Deployment"))

	newObj := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
		},
	}

	newObj.SetName("b")
	newObj.SetNamespace("b")
	newObj.SetUID("b")
	newObj.SetClusterName("b")
	newObj.SetGeneration(0)
	newObj.SetResourceVersion("b")
	newObj.SetCreationTimestamp(metav1.NewTime(time.Now()))
	newObj.SetManagedFields([]metav1.ManagedFieldsEntry{
		{
			Manager:    "update",
			Operation:  metav1.ManagedFieldsOperationApply,
			APIVersion: "apps/v1",
		},
	})
	if err := f.Update(newObj, "fieldmanager_test"); err != nil {
		t.Fatalf("failed to apply object: %v", err)
	}

	if m := f.ManagedFields(); len(m) != 0 {
		t.Fatalf("fields did not get stripped: %v", m)
	}
}

func TestVersionCheck(t *testing.T) {
	f := NewTestFieldManager(schema.FromAPIVersionAndKind("apps/v1", "Deployment"))

	appliedObj := &unstructured.Unstructured{Object: map[string]interface{}{}}
	if err := yaml.Unmarshal([]byte(`{
		"apiVersion": "apps/v1",
		"kind": "Deployment",
	}`), &appliedObj.Object); err != nil {
		t.Fatalf("error decoding YAML: %v", err)
	}

	// patch has 'apiVersion: apps/v1' and live version is apps/v1 -> no errors
	err := f.Apply(appliedObj, "fieldmanager_test", false)
	if err != nil {
		t.Fatalf("failed to apply object: %v", err)
	}

	appliedObj = &unstructured.Unstructured{Object: map[string]interface{}{}}
	if err := yaml.Unmarshal([]byte(`{
		"apiVersion": "apps/v1beta1",
		"kind": "Deployment",
	}`), &appliedObj.Object); err != nil {
		t.Fatalf("error decoding YAML: %v", err)
	}

	// patch has 'apiVersion: apps/v2' but live version is apps/v1 -> error
	err = f.Apply(appliedObj, "fieldmanager_test", false)
	if err == nil {
		t.Fatalf("expected an error from mismatched patch and live versions")
	}
	switch typ := err.(type) {
	default:
		t.Fatalf("expected error to be of type %T was %T", apierrors.StatusError{}, typ)
	case apierrors.APIStatus:
		if typ.Status().Code != http.StatusBadRequest {
			t.Fatalf("expected status code to be %d but was %d",
				http.StatusBadRequest, typ.Status().Code)
		}
	}
}
func TestVersionCheckDoesNotPanic(t *testing.T) {
	f := NewTestFieldManager(schema.FromAPIVersionAndKind("apps/v1", "Deployment"))

	appliedObj := &unstructured.Unstructured{Object: map[string]interface{}{}}
	if err := yaml.Unmarshal([]byte(`{
		"apiVersion": "apps/v1",
		"kind": "Deployment",
	}`), &appliedObj.Object); err != nil {
		t.Fatalf("error decoding YAML: %v", err)
	}

	// patch has 'apiVersion: apps/v1' and live version is apps/v1 -> no errors
	err := f.Apply(appliedObj, "fieldmanager_test", false)
	if err != nil {
		t.Fatalf("failed to apply object: %v", err)
	}

	appliedObj = &unstructured.Unstructured{Object: map[string]interface{}{}}
	if err := yaml.Unmarshal([]byte(`{
		}`), &appliedObj.Object); err != nil {
		t.Fatalf("error decoding YAML: %v", err)
	}

	// patch has 'apiVersion: apps/v2' but live version is apps/v1 -> error
	err = f.Apply(appliedObj, "fieldmanager_test", false)
	if err == nil {
		t.Fatalf("expected an error from mismatched patch and live versions")
	}
	switch typ := err.(type) {
	default:
		t.Fatalf("expected error to be of type %T was %T", apierrors.StatusError{}, typ)
	case apierrors.APIStatus:
		if typ.Status().Code != http.StatusBadRequest {
			t.Fatalf("expected status code to be %d but was %d",
				http.StatusBadRequest, typ.Status().Code)
		}
	}
}

func TestApplyDoesNotStripLabels(t *testing.T) {
	f := NewTestFieldManager(schema.FromAPIVersionAndKind("v1", "Pod"))

	appliedObj := &unstructured.Unstructured{Object: map[string]interface{}{}}
	if err := yaml.Unmarshal([]byte(`{
		"apiVersion": "v1",
		"kind": "Pod",
		"metadata": {
			"labels": {
				"a": "b"
			},
		}
	}`), &appliedObj.Object); err != nil {
		t.Fatalf("error decoding YAML: %v", err)
	}

	err := f.Apply(appliedObj, "fieldmanager_test", false)
	if err != nil {
		t.Fatalf("failed to apply object: %v", err)
	}

	if m := f.ManagedFields(); len(m) != 1 {
		t.Fatalf("labels shouldn't get stripped on apply: %v", m)
	}
}

func getObjectBytes(file string) []byte {
	s, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return s
}

func TestApplyNewObject(t *testing.T) {
	tests := []struct {
		gvk schema.GroupVersionKind
		obj []byte
	}{
		{
			gvk: schema.FromAPIVersionAndKind("v1", "Pod"),
			obj: getObjectBytes("pod.yaml"),
		},
		{
			gvk: schema.FromAPIVersionAndKind("v1", "Node"),
			obj: getObjectBytes("node.yaml"),
		},
		{
			gvk: schema.FromAPIVersionAndKind("v1", "Endpoints"),
			obj: getObjectBytes("endpoints.yaml"),
		},
	}

	for _, test := range tests {
		t.Run(test.gvk.String(), func(t *testing.T) {
			f := NewTestFieldManager(test.gvk)

			appliedObj := &unstructured.Unstructured{Object: map[string]interface{}{}}
			if err := yaml.Unmarshal(test.obj, &appliedObj.Object); err != nil {
				t.Fatalf("error decoding YAML: %v", err)
			}

			if err := f.Apply(appliedObj, "fieldmanager_test", false); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func BenchmarkNewObject(b *testing.B) {
	tests := []struct {
		gvk schema.GroupVersionKind
		obj []byte
	}{
		{
			gvk: schema.FromAPIVersionAndKind("v1", "Pod"),
			obj: getObjectBytes("pod.yaml"),
		},
		{
			gvk: schema.FromAPIVersionAndKind("v1", "Node"),
			obj: getObjectBytes("node.yaml"),
		},
		{
			gvk: schema.FromAPIVersionAndKind("v1", "Endpoints"),
			obj: getObjectBytes("endpoints.yaml"),
		},
	}
	scheme := runtime.NewScheme()
	if err := corev1.AddToScheme(scheme); err != nil {
		b.Fatalf("Failed to add to scheme: %v", err)
	}
	for _, test := range tests {
		b.Run(test.gvk.Kind, func(b *testing.B) {
			f := NewTestFieldManager(test.gvk)

			decoder := serializer.NewCodecFactory(scheme).UniversalDecoder(test.gvk.GroupVersion())
			newObj, err := runtime.Decode(decoder, test.obj)
			if err != nil {
				b.Fatalf("Failed to parse yaml object: %v", err)
			}
			objMeta, err := meta.Accessor(newObj)
			if err != nil {
				b.Fatalf("Failed to get object meta: %v", err)
			}
			objMeta.SetManagedFields([]metav1.ManagedFieldsEntry{
				{
					Manager:    "default",
					Operation:  "Update",
					APIVersion: "v1",
				},
			})
			appliedObj := &unstructured.Unstructured{Object: map[string]interface{}{}}
			if err := yaml.Unmarshal(test.obj, &appliedObj.Object); err != nil {
				b.Fatalf("Failed to parse yaml object: %v", err)
			}
			b.Run("Update", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for n := 0; n < b.N; n++ {
					if err := f.Update(newObj, "fieldmanager_test"); err != nil {
						b.Fatal(err)
					}
					f.Reset()
				}
			})
			b.Run("UpdateTwice", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for n := 0; n < b.N; n++ {
					if err := f.Update(newObj, "fieldmanager_test"); err != nil {
						b.Fatal(err)
					}
					if err := f.Update(newObj, "fieldmanager_test_2"); err != nil {
						b.Fatal(err)
					}
					f.Reset()
				}
			})
			b.Run("Apply", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for n := 0; n < b.N; n++ {
					if err := f.Apply(appliedObj, "fieldmanager_test", false); err != nil {
						b.Fatal(err)
					}
					f.Reset()
				}
			})
			b.Run("UpdateApply", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for n := 0; n < b.N; n++ {
					if err := f.Update(newObj, "fieldmanager_test"); err != nil {
						b.Fatal(err)
					}
					if err := f.Apply(appliedObj, "fieldmanager_test", false); err != nil {
						b.Fatal(err)
					}
					f.Reset()
				}
			})
		})
	}
}

func toUnstructured(b *testing.B, o runtime.Object) *unstructured.Unstructured {
	u, err := runtime.DefaultUnstructuredConverter.ToUnstructured(o)
	if err != nil {
		b.Fatalf("Failed to unmarshal to json: %v", err)
	}
	return &unstructured.Unstructured{Object: u}
}

func BenchmarkConvertObjectToTyped(b *testing.B) {
	tests := []struct {
		gvk schema.GroupVersionKind
		obj []byte
	}{
		{
			gvk: schema.FromAPIVersionAndKind("v1", "Pod"),
			obj: getObjectBytes("pod.yaml"),
		},
		{
			gvk: schema.FromAPIVersionAndKind("v1", "Node"),
			obj: getObjectBytes("node.yaml"),
		},
		{
			gvk: schema.FromAPIVersionAndKind("v1", "Endpoints"),
			obj: getObjectBytes("endpoints.yaml"),
		},
	}
	scheme := runtime.NewScheme()
	if err := corev1.AddToScheme(scheme); err != nil {
		b.Fatalf("Failed to add to scheme: %v", err)
	}

	for _, test := range tests {
		b.Run(test.gvk.Kind, func(b *testing.B) {
			decoder := serializer.NewCodecFactory(scheme).UniversalDecoder(test.gvk.GroupVersion())
			m := NewFakeOpenAPIModels()
			typeConverter := NewFakeTypeConverter(m)

			structured, err := runtime.Decode(decoder, test.obj)
			if err != nil {
				b.Fatalf("Failed to parse yaml object: %v", err)
			}
			b.Run("structured", func(b *testing.B) {
				b.ReportAllocs()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						_, err := typeConverter.ObjectToTyped(structured)
						if err != nil {
							b.Errorf("Error in ObjectToTyped: %v", err)
						}
					}
				})
			})

			unstructured := toUnstructured(b, structured)
			b.Run("unstructured", func(b *testing.B) {
				b.ReportAllocs()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						_, err := typeConverter.ObjectToTyped(unstructured)
						if err != nil {
							b.Errorf("Error in ObjectToTyped: %v", err)
						}
					}
				})
			})
		})
	}
}

func BenchmarkCompare(b *testing.B) {
	tests := []struct {
		gvk schema.GroupVersionKind
		obj []byte
	}{
		{
			gvk: schema.FromAPIVersionAndKind("v1", "Pod"),
			obj: getObjectBytes("pod.yaml"),
		},
		{
			gvk: schema.FromAPIVersionAndKind("v1", "Node"),
			obj: getObjectBytes("node.yaml"),
		},
		{
			gvk: schema.FromAPIVersionAndKind("v1", "Endpoints"),
			obj: getObjectBytes("endpoints.yaml"),
		},
	}

	scheme := runtime.NewScheme()
	if err := corev1.AddToScheme(scheme); err != nil {
		b.Fatalf("Failed to add to scheme: %v", err)
	}

	for _, test := range tests {
		b.Run(test.gvk.Kind, func(b *testing.B) {
			decoder := serializer.NewCodecFactory(scheme).UniversalDecoder(test.gvk.GroupVersion())
			m := NewFakeOpenAPIModels()
			typeConverter := NewFakeTypeConverter(m)

			structured, err := runtime.Decode(decoder, test.obj)
			if err != nil {
				b.Fatal(err)
			}
			tv1, err := typeConverter.ObjectToTyped(structured)
			if err != nil {
				b.Errorf("Error in ObjectToTyped: %v", err)
			}
			tv2, err := typeConverter.ObjectToTyped(structured)
			if err != nil {
				b.Errorf("Error in ObjectToTyped: %v", err)
			}

			b.Run("structured", func(b *testing.B) {
				b.ReportAllocs()
				for n := 0; n < b.N; n++ {
					_, err = tv1.Compare(tv2)
					if err != nil {
						b.Errorf("Error in ObjectToTyped: %v", err)
					}
				}
			})

			unstructured := toUnstructured(b, structured)
			utv1, err := typeConverter.ObjectToTyped(unstructured)
			if err != nil {
				b.Errorf("Error in ObjectToTyped: %v", err)
			}
			utv2, err := typeConverter.ObjectToTyped(unstructured)
			if err != nil {
				b.Errorf("Error in ObjectToTyped: %v", err)
			}
			b.Run("unstructured", func(b *testing.B) {
				b.ReportAllocs()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						_, err = utv1.Compare(utv2)
						if err != nil {
							b.Errorf("Error in ObjectToTyped: %v", err)
						}
					}
				})
			})
		})
	}
}

func BenchmarkRepeatedUpdate(b *testing.B) {
	f := NewTestFieldManager(schema.FromAPIVersionAndKind("v1", "Pod"))
	podBytes := getObjectBytes("pod.yaml")

	var obj *corev1.Pod
	if err := yaml.Unmarshal(podBytes, &obj); err != nil {
		b.Fatalf("Failed to parse yaml object: %v", err)
	}
	obj.Spec.Containers[0].Image = "nginx:latest"
	objs := []*corev1.Pod{obj}
	obj = obj.DeepCopy()
	obj.Spec.Containers[0].Image = "nginx:4.3"
	objs = append(objs, obj)

	appliedObj := &unstructured.Unstructured{Object: map[string]interface{}{}}
	if err := yaml.Unmarshal(podBytes, &appliedObj.Object); err != nil {
		b.Fatalf("error decoding YAML: %v", err)
	}

	err := f.Apply(appliedObj, "fieldmanager_apply", false)
	if err != nil {
		b.Fatal(err)
	}

	if err := f.Update(objs[1], "fieldmanager_1"); err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err := f.Update(objs[n%len(objs)], fmt.Sprintf("fieldmanager_%d", n%len(objs)))
		if err != nil {
			b.Fatal(err)
		}
		f.Reset()
	}
}

func TestApplyFailsWithManagedFields(t *testing.T) {
	f := NewTestFieldManager(schema.FromAPIVersionAndKind("v1", "Pod"))

	appliedObj := &unstructured.Unstructured{Object: map[string]interface{}{}}
	if err := yaml.Unmarshal([]byte(`{
		"apiVersion": "v1",
		"kind": "Pod",
		"metadata": {
			"managedFields": [
				{
				  "manager": "test",
				}
			]
		}
	}`), &appliedObj.Object); err != nil {
		t.Fatalf("error decoding YAML: %v", err)
	}

	err := f.Apply(appliedObj, "fieldmanager_test", false)

	if err == nil {
		t.Fatalf("successfully applied with set managed fields")
	}
}

func TestApplySuccessWithNoManagedFields(t *testing.T) {
	f := NewTestFieldManager(schema.FromAPIVersionAndKind("v1", "Pod"))

	appliedObj := &unstructured.Unstructured{Object: map[string]interface{}{}}
	if err := yaml.Unmarshal([]byte(`{
		"apiVersion": "v1",
		"kind": "Pod",
		"metadata": {
			"labels": {
				"a": "b"
			},
		}
	}`), &appliedObj.Object); err != nil {
		t.Fatalf("error decoding YAML: %v", err)
	}
	err := f.Apply(appliedObj, "fieldmanager_test", false)

	if err != nil {
		t.Fatalf("failed to apply object: %v", err)
	}
}

// Run an update and apply, and make sure that nothing has changed.
func TestNoOpChanges(t *testing.T) {
	f := NewTestFieldManager(schema.FromAPIVersionAndKind("v1", "Pod"))

	obj := &unstructured.Unstructured{Object: map[string]interface{}{}}
	if err := yaml.Unmarshal([]byte(`{
		"apiVersion": "v1",
		"kind": "Pod",
		"metadata": {
			"labels": {
				"a": "b"
			},
		}
	}`), &obj.Object); err != nil {
		t.Fatalf("error decoding YAML: %v", err)
	}
	if err := f.Apply(obj, "fieldmanager_test_apply", false); err != nil {
		t.Fatalf("failed to apply object: %v", err)
	}
	before := f.liveObj.DeepCopyObject()
	// Wait to make sure the timestamp is different
	time.Sleep(time.Second)
	// Applying with a different fieldmanager will create an entry..
	if err := f.Apply(obj, "fieldmanager_test_apply_other", false); err != nil {
		t.Fatalf("failed to update object: %v", err)
	}
	if reflect.DeepEqual(before, f.liveObj) {
		t.Fatalf("Applying no-op apply with new manager didn't change object: \n%v", f.liveObj)
	}
	before = f.liveObj.DeepCopyObject()
	// Wait to make sure the timestamp is different
	time.Sleep(time.Second)
	if err := f.Update(obj, "fieldmanager_test_update"); err != nil {
		t.Fatalf("failed to update object: %v", err)
	}
	if !reflect.DeepEqual(before, f.liveObj) {
		t.Fatalf("No-op update has changed the object:\n%v\n---\n%v", before, f.liveObj)
	}
	before = f.liveObj.DeepCopyObject()
	// Wait to make sure the timestamp is different
	time.Sleep(time.Second)
	if err := f.Apply(obj, "fieldmanager_test_apply", true); err != nil {
		t.Fatalf("failed to re-apply object: %v", err)
	}
	if !reflect.DeepEqual(before, f.liveObj) {
		t.Fatalf("No-op apply has changed the object:\n%v\n---\n%v", before, f.liveObj)
	}
}
