/*
Copyright 2018 The Kubernetes Authors.

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

package apply

import (
	"reflect"
	"strings"
	"testing"

	"sigs.k8s.io/structured-merge-diff/fieldpath"
	"sigs.k8s.io/structured-merge-diff/value"
)

// TestRoundTripManagedFields will roundtrip ManagedFields from the format used by
// https://github.com/kubernetes-sigs/structured-merge-diff to the wire format (api format) and back
func TestRoundTripManagedFields(t *testing.T) {
	tests := []struct {
		yaml      string
		errString string
	}{
		{
			yaml: `spec:
  float: 3.1415`,
		},
		{
			yaml: `spec:
  containers:
  - name: c
    image: i`,
		},
		{
			yaml: `apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.15.4
        ports:
        - containerPort: 80`,
		},
		{
			yaml: `kind: StorageClass
apiVersion: storage.k8s.io/v1
metadata:
  name: gluster-vol-default
provisioner: kubernetes.io/glusterfs
parameters:
  resturl: "http://192.168.10.100:8080"
  restuser: ""
  secretNamespace: ""
  secretName: ""
allowVolumeExpansion: true`,
		},
		{
			yaml: `apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  # name must match the spec fields below, and be in the form: <plural>.<group>
  name: crontabs.stable.example.com
spec:
  # group name to use for REST API: /apis/<group>/<version>
  group: stable.example.com
  # list of versions supported by this CustomResourceDefinition
  versions:
    - name: v1
      # Each version can be enabled/disabled by Served flag.
      served: true
      # One and only one version must be marked as the storage version.
      storage: true
  # either Namespaced or Cluster
  scope: Namespaced
  names:
    # plural name to be used in the URL: /apis/<group>/<version>/<plural>
    plural: crontabs
    # singular name to be used as an alias on the CLI and for display
    singular: crontab
    # kind is normally the CamelCased singular type. Your resource manifests use this.
    kind: CronTab
    # shortNames allow shorter string to match your resource on the CLI
    shortNames:
    - ct`,
		},
	}

	for i, tc := range tests {
		v, _ := value.FromYAML([]byte(tc.yaml))
		original := fieldpath.ManagedFields(map[string]*fieldpath.VersionedSet{
			"foo": {
				APIVersion: fieldpath.APIVersion("v1"),
				Set:        fieldpath.SetFromValue(v),
			},
		})
		encoded, err := EncodeManagedFields(original)
		if err == nil && len(tc.errString) > 0 {
			t.Errorf("[%v]expected error but got none.", i)
			continue
		}
		if err != nil && len(tc.errString) == 0 {
			t.Errorf("[%v]did not expect error but got: %v", i, err)
			continue
		}
		if err != nil && len(tc.errString) > 0 && !strings.Contains(err.Error(), tc.errString) {
			t.Errorf("[%v]expected error with %q but got: %v", i, tc.errString, err)
			continue
		}
		decoded, err := DecodeManagedFields(encoded)
		if err != nil {
			t.Errorf("[%v]did not expect round trip error but got: %v", i, err)
			continue
		}
		if !reflect.DeepEqual(decoded, original) {
			t.Errorf("[%v]expected:\n\t%+v\nbut got:\n\t%+v", i, original, decoded)
		}
	}
}
