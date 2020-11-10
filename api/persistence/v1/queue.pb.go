// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/server/api/persistence/v1/queue.proto

package persistence

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"

	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
	_ "github.com/gogo/protobuf/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// message_payload column
type QueueMessage struct {
}

func (m *QueueMessage) Reset()      { *m = QueueMessage{} }
func (*QueueMessage) ProtoMessage() {}
func (*QueueMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_3582553309991064, []int{0}
}
func (m *QueueMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueueMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueueMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueueMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueMessage.Merge(m, src)
}
func (m *QueueMessage) XXX_Size() int {
	return m.Size()
}
func (m *QueueMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueMessage.DiscardUnknown(m)
}

var xxx_messageInfo_QueueMessage proto.InternalMessageInfo

// data column
type QueueMetadata struct {
	ClusterAckLevels map[string]int64 `protobuf:"bytes,1,rep,name=cluster_ack_levels,json=clusterAckLevels,proto3" json:"cluster_ack_levels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (m *QueueMetadata) Reset()      { *m = QueueMetadata{} }
func (*QueueMetadata) ProtoMessage() {}
func (*QueueMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_3582553309991064, []int{1}
}
func (m *QueueMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueueMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueueMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueueMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueMetadata.Merge(m, src)
}
func (m *QueueMetadata) XXX_Size() int {
	return m.Size()
}
func (m *QueueMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_QueueMetadata proto.InternalMessageInfo

func (m *QueueMetadata) GetClusterAckLevels() map[string]int64 {
	if m != nil {
		return m.ClusterAckLevels
	}
	return nil
}

func init() {
	proto.RegisterType((*QueueMessage)(nil), "temporal.server.api.persistence.v1.QueueMessage")
	proto.RegisterType((*QueueMetadata)(nil), "temporal.server.api.persistence.v1.QueueMetadata")
	proto.RegisterMapType((map[string]int64)(nil), "temporal.server.api.persistence.v1.QueueMetadata.ClusterAckLevelsEntry")
}

func init() {
	proto.RegisterFile("temporal/server/api/persistence/v1/queue.proto", fileDescriptor_3582553309991064)
}

var fileDescriptor_3582553309991064 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xbb, 0x4e, 0x02, 0x41,
	0x14, 0x86, 0xf7, 0x40, 0x34, 0x71, 0xbd, 0x84, 0x6c, 0x34, 0x21, 0x14, 0x47, 0x42, 0x45, 0x35,
	0x13, 0xd4, 0xc2, 0x68, 0xa5, 0xc4, 0xd8, 0x68, 0x21, 0xa5, 0x0d, 0x19, 0x96, 0xe3, 0x66, 0x65,
	0x61, 0xd6, 0xb9, 0x6c, 0x42, 0xe7, 0x23, 0xf8, 0x18, 0x3e, 0x8a, 0x89, 0x0d, 0x25, 0xa5, 0x0c,
	0x8d, 0x25, 0x8f, 0x60, 0x60, 0x35, 0x22, 0xd1, 0xd8, 0xcd, 0xff, 0xcf, 0x37, 0xdf, 0xdc, 0x7c,
	0x66, 0xa8, 0x9f, 0x4a, 0x25, 0x12, 0xae, 0x49, 0x65, 0xa4, 0xb8, 0x48, 0x63, 0x9e, 0x92, 0xd2,
	0xb1, 0x36, 0x34, 0x08, 0x89, 0x67, 0x0d, 0xfe, 0x60, 0xc9, 0x12, 0x4b, 0x95, 0x34, 0x32, 0xa8,
	0x7d, 0xf1, 0x2c, 0xe7, 0x99, 0x48, 0x63, 0xb6, 0xc4, 0xb3, 0xac, 0x51, 0xc1, 0x48, 0xca, 0x28,
	0x21, 0xbe, 0x58, 0xd1, 0xb1, 0x77, 0xbc, 0x6b, 0x95, 0x30, 0xb1, 0x1c, 0xe4, 0x8e, 0xca, 0xfe,
	0xea, 0xbc, 0x89, 0xfb, 0xa4, 0x8d, 0xe8, 0xa7, 0x39, 0x50, 0xdb, 0xf1, 0xb7, 0x6e, 0xe6, 0x7b,
	0x5e, 0x93, 0xd6, 0x22, 0xa2, 0xda, 0x2b, 0xf8, 0xdb, 0x9f, 0x85, 0x11, 0x5d, 0x61, 0x44, 0x60,
	0xfd, 0x20, 0x4c, 0xac, 0x36, 0xa4, 0xda, 0x22, 0xec, 0xb5, 0x13, 0xca, 0x28, 0xd1, 0x65, 0xa8,
	0x16, 0xeb, 0x9b, 0x07, 0x97, 0xec, 0xff, 0x33, 0xb2, 0x1f, 0x3a, 0xd6, 0xcc, 0x5d, 0x67, 0x61,
	0xef, 0x6a, 0x61, 0xba, 0x18, 0x18, 0x35, 0x6c, 0x95, 0xc2, 0x95, 0xba, 0xd2, 0xf4, 0xf7, 0x7e,
	0x45, 0x83, 0x92, 0x5f, 0xec, 0xd1, 0xb0, 0x0c, 0x55, 0xa8, 0x6f, 0xb4, 0xe6, 0xc3, 0x60, 0xd7,
	0x5f, 0xcb, 0x44, 0x62, 0xa9, 0x5c, 0xa8, 0x42, 0xbd, 0xd8, 0xca, 0xc3, 0x49, 0xe1, 0x18, 0xce,
	0xef, 0x47, 0x13, 0xf4, 0xc6, 0x13, 0xf4, 0x66, 0x13, 0x84, 0x47, 0x87, 0xf0, 0xec, 0x10, 0x5e,
	0x1c, 0xc2, 0xc8, 0x21, 0xbc, 0x39, 0x84, 0x77, 0x87, 0xde, 0xcc, 0x21, 0x3c, 0x4d, 0xd1, 0x1b,
	0x4d, 0xd1, 0x1b, 0x4f, 0xd1, 0xbb, 0x3d, 0x8a, 0xe4, 0xf7, 0xbd, 0x62, 0xf9, 0xf7, 0x77, 0x9d,
	0x2e, 0xc5, 0xce, 0xfa, 0xe2, 0x41, 0x0f, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x79, 0x96,
	0x18, 0xe7, 0x01, 0x00, 0x00,
}

func (this *QueueMessage) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueueMessage)
	if !ok {
		that2, ok := that.(QueueMessage)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *QueueMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueueMetadata)
	if !ok {
		that2, ok := that.(QueueMetadata)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.ClusterAckLevels) != len(that1.ClusterAckLevels) {
		return false
	}
	for i := range this.ClusterAckLevels {
		if this.ClusterAckLevels[i] != that1.ClusterAckLevels[i] {
			return false
		}
	}
	return true
}
func (this *QueueMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&persistence.QueueMessage{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *QueueMetadata) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&persistence.QueueMetadata{")
	keysForClusterAckLevels := make([]string, 0, len(this.ClusterAckLevels))
	for k, _ := range this.ClusterAckLevels {
		keysForClusterAckLevels = append(keysForClusterAckLevels, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForClusterAckLevels)
	mapStringForClusterAckLevels := "map[string]int64{"
	for _, k := range keysForClusterAckLevels {
		mapStringForClusterAckLevels += fmt.Sprintf("%#v: %#v,", k, this.ClusterAckLevels[k])
	}
	mapStringForClusterAckLevels += "}"
	if this.ClusterAckLevels != nil {
		s = append(s, "ClusterAckLevels: "+mapStringForClusterAckLevels+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringQueue(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *QueueMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueueMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueueMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueueMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueueMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueueMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClusterAckLevels) > 0 {
		for k := range m.ClusterAckLevels {
			v := m.ClusterAckLevels[k]
			baseI := i
			i = encodeVarintQueue(dAtA, i, uint64(v))
			i--
			dAtA[i] = 0x10
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintQueue(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintQueue(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQueue(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueue(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueueMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueueMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ClusterAckLevels) > 0 {
		for k, v := range m.ClusterAckLevels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovQueue(uint64(len(k))) + 1 + sovQueue(uint64(v))
			n += mapEntrySize + 1 + sovQueue(uint64(mapEntrySize))
		}
	}
	return n
}

func sovQueue(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueue(x uint64) (n int) {
	return sovQueue(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *QueueMessage) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&QueueMessage{`,
		`}`,
	}, "")
	return s
}
func (this *QueueMetadata) String() string {
	if this == nil {
		return "nil"
	}
	keysForClusterAckLevels := make([]string, 0, len(this.ClusterAckLevels))
	for k, _ := range this.ClusterAckLevels {
		keysForClusterAckLevels = append(keysForClusterAckLevels, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForClusterAckLevels)
	mapStringForClusterAckLevels := "map[string]int64{"
	for _, k := range keysForClusterAckLevels {
		mapStringForClusterAckLevels += fmt.Sprintf("%v: %v,", k, this.ClusterAckLevels[k])
	}
	mapStringForClusterAckLevels += "}"
	s := strings.Join([]string{`&QueueMetadata{`,
		`ClusterAckLevels:` + mapStringForClusterAckLevels + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringQueue(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *QueueMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueue
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
			return fmt.Errorf("proto: QueueMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueueMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQueue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQueue
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQueue
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
func (m *QueueMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueue
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
			return fmt.Errorf("proto: QueueMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueueMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterAckLevels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueue
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
				return ErrInvalidLengthQueue
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClusterAckLevels == nil {
				m.ClusterAckLevels = make(map[string]int64)
			}
			var mapkey string
			var mapvalue int64
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQueue
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQueue
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthQueue
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthQueue
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQueue
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipQueue(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthQueue
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.ClusterAckLevels[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQueue
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQueue
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
func skipQueue(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueue
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
					return 0, ErrIntOverflowQueue
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQueue
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
				return 0, ErrInvalidLengthQueue
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueue
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueue
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueue        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueue          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueue = fmt.Errorf("proto: unexpected end of group")
)