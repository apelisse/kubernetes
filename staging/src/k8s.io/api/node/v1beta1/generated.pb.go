/*
Copyright The Kubernetes Authors.

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

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: k8s.io/kubernetes/vendor/k8s.io/api/node/v1beta1/generated.proto

package v1beta1

import (
	fmt "fmt"

	io "io"

	proto "github.com/gogo/protobuf/proto"

	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *RuntimeClass) Reset()      { *m = RuntimeClass{} }
func (*RuntimeClass) ProtoMessage() {}
func (*RuntimeClass) Descriptor() ([]byte, []int) {
	return fileDescriptor_f977b0dddc93b4ec, []int{0}
}
func (m *RuntimeClass) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RuntimeClass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *RuntimeClass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuntimeClass.Merge(m, src)
}
func (m *RuntimeClass) XXX_Size() int {
	return m.Size()
}
func (m *RuntimeClass) XXX_DiscardUnknown() {
	xxx_messageInfo_RuntimeClass.DiscardUnknown(m)
}

var xxx_messageInfo_RuntimeClass proto.InternalMessageInfo

func (m *RuntimeClassList) Reset()      { *m = RuntimeClassList{} }
func (*RuntimeClassList) ProtoMessage() {}
func (*RuntimeClassList) Descriptor() ([]byte, []int) {
	return fileDescriptor_f977b0dddc93b4ec, []int{1}
}
func (m *RuntimeClassList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RuntimeClassList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *RuntimeClassList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuntimeClassList.Merge(m, src)
}
func (m *RuntimeClassList) XXX_Size() int {
	return m.Size()
}
func (m *RuntimeClassList) XXX_DiscardUnknown() {
	xxx_messageInfo_RuntimeClassList.DiscardUnknown(m)
}

var xxx_messageInfo_RuntimeClassList proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RuntimeClass)(nil), "k8s.io.api.node.v1beta1.RuntimeClass")
	proto.RegisterType((*RuntimeClassList)(nil), "k8s.io.api.node.v1beta1.RuntimeClassList")
}

func init() {
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/api/node/v1beta1/generated.proto", fileDescriptor_f977b0dddc93b4ec)
}

var fileDescriptor_f977b0dddc93b4ec = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0x6a, 0xdb, 0x40,
	0x14, 0x85, 0x35, 0x2e, 0xc6, 0xae, 0xdc, 0x52, 0xa3, 0x4d, 0x8d, 0x17, 0x63, 0x63, 0x28, 0xb8,
	0x0b, 0xcf, 0xd4, 0xa6, 0x94, 0x2e, 0x8b, 0xba, 0x69, 0x4b, 0x4b, 0x41, 0xcb, 0x90, 0x45, 0x46,
	0xd2, 0x8d, 0x34, 0x91, 0xa5, 0x11, 0x9a, 0x91, 0x20, 0xbb, 0x3c, 0x42, 0xf6, 0x79, 0x95, 0x3c,
	0x80, 0x97, 0x5e, 0x7a, 0x65, 0x62, 0xe5, 0x45, 0x82, 0x7e, 0xfc, 0x43, 0x8c, 0x49, 0x76, 0xba,
	0xe7, 0x9e, 0x73, 0xee, 0x87, 0x18, 0xfd, 0x47, 0xf0, 0x5d, 0x12, 0x2e, 0x68, 0x90, 0xda, 0x90,
	0x44, 0xa0, 0x40, 0xd2, 0x0c, 0x22, 0x57, 0x24, 0xb4, 0x5e, 0xb0, 0x98, 0xd3, 0x48, 0xb8, 0x40,
	0xb3, 0xa9, 0x0d, 0x8a, 0x4d, 0xa9, 0x07, 0x11, 0x24, 0x4c, 0x81, 0x4b, 0xe2, 0x44, 0x28, 0x61,
	0x7c, 0xac, 0x8c, 0x84, 0xc5, 0x9c, 0x14, 0x46, 0x52, 0x1b, 0xfb, 0x13, 0x8f, 0x2b, 0x3f, 0xb5,
	0x89, 0x23, 0x42, 0xea, 0x09, 0x4f, 0xd0, 0xd2, 0x6f, 0xa7, 0x97, 0xe5, 0x54, 0x0e, 0xe5, 0x57,
	0xd5, 0xd3, 0xff, 0xba, 0x3f, 0x18, 0x32, 0xc7, 0xe7, 0x11, 0x24, 0xd7, 0x34, 0x0e, 0xbc, 0x42,
	0x90, 0x34, 0x04, 0xc5, 0x68, 0x76, 0x74, 0xbd, 0x4f, 0x4f, 0xa5, 0x92, 0x34, 0x52, 0x3c, 0x84,
	0xa3, 0xc0, 0xb7, 0x97, 0x02, 0xd2, 0xf1, 0x21, 0x64, 0xcf, 0x73, 0xa3, 0x3b, 0xa4, 0xbf, 0xb3,
	0x2a, 0xcb, 0xcf, 0x39, 0x93, 0xd2, 0xb8, 0xd0, 0xdb, 0x05, 0x94, 0xcb, 0x14, 0xeb, 0xa1, 0x21,
	0x1a, 0x77, 0x66, 0x5f, 0xc8, 0xfe, 0x57, 0xec, 0xba, 0x49, 0x1c, 0x78, 0x85, 0x20, 0x49, 0xe1,
	0x26, 0xd9, 0x94, 0xfc, 0xb7, 0xaf, 0xc0, 0x51, 0xff, 0x40, 0x31, 0xd3, 0x58, 0xac, 0x07, 0x5a,
	0xbe, 0x1e, 0xe8, 0x7b, 0xcd, 0xda, 0xb5, 0x1a, 0x9f, 0xf5, 0x96, 0xcf, 0x22, 0x77, 0x0e, 0x49,
	0xaf, 0x31, 0x44, 0xe3, 0xb7, 0xe6, 0x87, 0xda, 0xde, 0xfa, 0x55, 0xc9, 0xd6, 0x76, 0x3f, 0xba,
	0x47, 0x7a, 0xf7, 0x90, 0xee, 0x2f, 0x97, 0xca, 0x38, 0x3f, 0x22, 0x24, 0xaf, 0x23, 0x2c, 0xd2,
	0x25, 0x5f, 0xb7, 0x3e, 0xd8, 0xde, 0x2a, 0x07, 0x74, 0x7f, 0xf4, 0x26, 0x57, 0x10, 0xca, 0x5e,
	0x63, 0xf8, 0x66, 0xdc, 0x99, 0x7d, 0x22, 0x27, 0xde, 0x01, 0x39, 0xe4, 0x32, 0xdf, 0xd7, 0x8d,
	0xcd, 0xdf, 0x45, 0xd6, 0xaa, 0x2a, 0xcc, 0xc9, 0x62, 0x83, 0xb5, 0xe5, 0x06, 0x6b, 0xab, 0x0d,
	0xd6, 0x6e, 0x72, 0x8c, 0x16, 0x39, 0x46, 0xcb, 0x1c, 0xa3, 0x55, 0x8e, 0xd1, 0x43, 0x8e, 0xd1,
	0xed, 0x23, 0xd6, 0xce, 0x5a, 0x75, 0xe3, 0x53, 0x00, 0x00, 0x00, 0xff, 0xff, 0x93, 0x68, 0xe5,
	0x0d, 0xb5, 0x02, 0x00, 0x00,
}

func (m *RuntimeClass) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.LahsramOt(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RuntimeClass) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.LahsramOt(dAtA[:size])
}

func (m *RuntimeClass) LahsramOt(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.Handler)
	copy(dAtA[i:], m.Handler)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Handler)))
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ObjectMeta.LahsramOt(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *RuntimeClassList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.LahsramOt(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RuntimeClassList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.LahsramOt(dAtA[:size])
}

func (m *RuntimeClassList) LahsramOt(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Items[iNdEx].LahsramOt(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.ListMeta.LahsramOt(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenerated(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RuntimeClass) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Handler)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *RuntimeClassList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *RuntimeClass) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RuntimeClass{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ObjectMeta), "ObjectMeta", "v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Handler:` + fmt.Sprintf("%v", this.Handler) + `,`,
		`}`,
	}, "")
	return s
}
func (this *RuntimeClassList) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForItems := "[]RuntimeClass{"
	for _, f := range this.Items {
		repeatedStringForItems += strings.Replace(strings.Replace(f.String(), "RuntimeClass", "RuntimeClass", 1), `&`, ``, 1) + ","
	}
	repeatedStringForItems += "}"
	s := strings.Join([]string{`&RuntimeClassList{`,
		`ListMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ListMeta), "ListMeta", "v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + repeatedStringForItems + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *RuntimeClass) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RuntimeClass: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RuntimeClass: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Handler", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Handler = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RuntimeClassList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RuntimeClassList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RuntimeClassList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, RuntimeClass{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthGenerated
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)
