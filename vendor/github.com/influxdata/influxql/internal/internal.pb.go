// Code generated by protoc-gen-gogo.
// source: internal/internal.proto
// DO NOT EDIT!

/*
Package influxql is a generated protocol buffer package.

It is generated from these files:
	internal/internal.proto

It has these top-level messages:
	Measurements
	Measurement
*/
package influxql

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Measurements struct {
	Items            []*Measurement `protobuf:"bytes,1,rep,name=Items" json:"Items,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Measurements) Reset()                    { *m = Measurements{} }
func (m *Measurements) String() string            { return proto.CompactTextString(m) }
func (*Measurements) ProtoMessage()               {}
func (*Measurements) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{0} }

func (m *Measurements) GetItems() []*Measurement {
	if m != nil {
		return m.Items
	}
	return nil
}

type Measurement struct {
	Database         *string `protobuf:"bytes,1,opt,name=Database" json:"Database,omitempty"`
	RetentionPolicy  *string `protobuf:"bytes,2,opt,name=RetentionPolicy" json:"RetentionPolicy,omitempty"`
	Name             *string `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	Regex            *string `protobuf:"bytes,4,opt,name=Regex" json:"Regex,omitempty"`
	IsTarget         *bool   `protobuf:"varint,5,opt,name=IsTarget" json:"IsTarget,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Measurement) Reset()                    { *m = Measurement{} }
func (m *Measurement) String() string            { return proto.CompactTextString(m) }
func (*Measurement) ProtoMessage()               {}
func (*Measurement) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{1} }

func (m *Measurement) GetDatabase() string {
	if m != nil && m.Database != nil {
		return *m.Database
	}
	return ""
}

func (m *Measurement) GetRetentionPolicy() string {
	if m != nil && m.RetentionPolicy != nil {
		return *m.RetentionPolicy
	}
	return ""
}

func (m *Measurement) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Measurement) GetRegex() string {
	if m != nil && m.Regex != nil {
		return *m.Regex
	}
	return ""
}

func (m *Measurement) GetIsTarget() bool {
	if m != nil && m.IsTarget != nil {
		return *m.IsTarget
	}
	return false
}

func init() {
	proto.RegisterType((*Measurements)(nil), "influxql.Measurements")
	proto.RegisterType((*Measurement)(nil), "influxql.Measurement")
}

func init() { proto.RegisterFile("internal/internal.proto", fileDescriptorInternal) }

var fileDescriptorInternal = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcf, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x87, 0x31, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38, 0x32,
	0xf3, 0xd2, 0x72, 0x4a, 0x2b, 0x0a, 0x73, 0x94, 0xac, 0xb9, 0x78, 0x7c, 0x53, 0x13, 0x8b, 0x4b,
	0x8b, 0x52, 0x73, 0x53, 0xf3, 0x4a, 0x8a, 0x85, 0xb4, 0xb9, 0x58, 0x3d, 0x4b, 0x52, 0x73, 0x8b,
	0x25, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x44, 0xf5, 0x60, 0x2a, 0xf5, 0x90, 0x94, 0x05, 0x41,
	0xd4, 0x28, 0xcd, 0x64, 0xe4, 0xe2, 0x46, 0x12, 0x16, 0x92, 0xe2, 0xe2, 0x70, 0x49, 0x2c, 0x49,
	0x4c, 0x4a, 0x2c, 0x4e, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x85, 0x34, 0xb8,
	0xf8, 0x83, 0x52, 0x4b, 0x52, 0xf3, 0x4a, 0x32, 0xf3, 0xf3, 0x02, 0xf2, 0x73, 0x32, 0x93, 0x2b,
	0x25, 0x98, 0xc0, 0x4a, 0xd0, 0x85, 0x85, 0x84, 0xb8, 0x58, 0xfc, 0x12, 0x73, 0x53, 0x25, 0x98,
	0xc1, 0xd2, 0x60, 0xb6, 0x90, 0x08, 0x17, 0x6b, 0x50, 0x6a, 0x7a, 0x6a, 0x85, 0x04, 0x0b, 0x58,
	0x10, 0xc2, 0x01, 0xd9, 0xe7, 0x59, 0x1c, 0x92, 0x58, 0x94, 0x9e, 0x5a, 0x22, 0xc1, 0xaa, 0xc0,
	0xa8, 0xc1, 0x11, 0x04, 0xe7, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x16, 0x06, 0x23, 0xfc,
	0x00, 0x00, 0x00,
}
