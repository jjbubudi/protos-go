// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tides/events.proto

package tides

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TideRecorded struct {
	Time                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Meters               float64              `protobuf:"fixed64,2,opt,name=meters,proto3" json:"meters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TideRecorded) Reset()         { *m = TideRecorded{} }
func (m *TideRecorded) String() string { return proto.CompactTextString(m) }
func (*TideRecorded) ProtoMessage()    {}
func (*TideRecorded) Descriptor() ([]byte, []int) {
	return fileDescriptor_6122246a85abeb25, []int{0}
}

func (m *TideRecorded) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TideRecorded.Unmarshal(m, b)
}
func (m *TideRecorded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TideRecorded.Marshal(b, m, deterministic)
}
func (m *TideRecorded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TideRecorded.Merge(m, src)
}
func (m *TideRecorded) XXX_Size() int {
	return xxx_messageInfo_TideRecorded.Size(m)
}
func (m *TideRecorded) XXX_DiscardUnknown() {
	xxx_messageInfo_TideRecorded.DiscardUnknown(m)
}

var xxx_messageInfo_TideRecorded proto.InternalMessageInfo

func (m *TideRecorded) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *TideRecorded) GetMeters() float64 {
	if m != nil {
		return m.Meters
	}
	return 0
}

type TidePredicted struct {
	Time                 *timestamp.Timestamp        `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Predictions          []*TidePredicted_Prediction `protobuf:"bytes,2,rep,name=predictions,proto3" json:"predictions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *TidePredicted) Reset()         { *m = TidePredicted{} }
func (m *TidePredicted) String() string { return proto.CompactTextString(m) }
func (*TidePredicted) ProtoMessage()    {}
func (*TidePredicted) Descriptor() ([]byte, []int) {
	return fileDescriptor_6122246a85abeb25, []int{1}
}

func (m *TidePredicted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TidePredicted.Unmarshal(m, b)
}
func (m *TidePredicted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TidePredicted.Marshal(b, m, deterministic)
}
func (m *TidePredicted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TidePredicted.Merge(m, src)
}
func (m *TidePredicted) XXX_Size() int {
	return xxx_messageInfo_TidePredicted.Size(m)
}
func (m *TidePredicted) XXX_DiscardUnknown() {
	xxx_messageInfo_TidePredicted.DiscardUnknown(m)
}

var xxx_messageInfo_TidePredicted proto.InternalMessageInfo

func (m *TidePredicted) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *TidePredicted) GetPredictions() []*TidePredicted_Prediction {
	if m != nil {
		return m.Predictions
	}
	return nil
}

type TidePredicted_Prediction struct {
	Time                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Meters               float64              `protobuf:"fixed64,2,opt,name=meters,proto3" json:"meters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TidePredicted_Prediction) Reset()         { *m = TidePredicted_Prediction{} }
func (m *TidePredicted_Prediction) String() string { return proto.CompactTextString(m) }
func (*TidePredicted_Prediction) ProtoMessage()    {}
func (*TidePredicted_Prediction) Descriptor() ([]byte, []int) {
	return fileDescriptor_6122246a85abeb25, []int{1, 0}
}

func (m *TidePredicted_Prediction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TidePredicted_Prediction.Unmarshal(m, b)
}
func (m *TidePredicted_Prediction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TidePredicted_Prediction.Marshal(b, m, deterministic)
}
func (m *TidePredicted_Prediction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TidePredicted_Prediction.Merge(m, src)
}
func (m *TidePredicted_Prediction) XXX_Size() int {
	return xxx_messageInfo_TidePredicted_Prediction.Size(m)
}
func (m *TidePredicted_Prediction) XXX_DiscardUnknown() {
	xxx_messageInfo_TidePredicted_Prediction.DiscardUnknown(m)
}

var xxx_messageInfo_TidePredicted_Prediction proto.InternalMessageInfo

func (m *TidePredicted_Prediction) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *TidePredicted_Prediction) GetMeters() float64 {
	if m != nil {
		return m.Meters
	}
	return 0
}

func init() {
	proto.RegisterType((*TideRecorded)(nil), "jimmyau.tides.TideRecorded")
	proto.RegisterType((*TidePredicted)(nil), "jimmyau.tides.TidePredicted")
	proto.RegisterType((*TidePredicted_Prediction)(nil), "jimmyau.tides.TidePredicted.Prediction")
}

func init() { proto.RegisterFile("tides/events.proto", fileDescriptor_6122246a85abeb25) }

var fileDescriptor_6122246a85abeb25 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x89, 0xca, 0x1e, 0xa6, 0xee, 0xc1, 0x1e, 0x64, 0xe9, 0xc5, 0x65, 0x3d, 0xd8, 0x8b,
	0x13, 0x58, 0xdf, 0x60, 0x41, 0xc4, 0x5b, 0x29, 0xc5, 0x83, 0xb7, 0x4d, 0x33, 0xd6, 0x14, 0xb3,
	0x53, 0x9a, 0x44, 0xf0, 0x6d, 0x7d, 0x14, 0x69, 0x6a, 0xd5, 0x5e, 0x65, 0x8f, 0xc9, 0xff, 0x7f,
	0x5f, 0x32, 0x0c, 0xa4, 0xde, 0x68, 0x72, 0x92, 0xde, 0xe9, 0xe0, 0x1d, 0x76, 0x3d, 0x7b, 0x4e,
	0x97, 0xad, 0xb1, 0xf6, 0x63, 0x1f, 0x30, 0x66, 0xd9, 0x55, 0xc3, 0xdc, 0xbc, 0x91, 0x8c, 0xa1,
	0x0a, 0x2f, 0xd2, 0x1b, 0x4b, 0xce, 0xef, 0x6d, 0x37, 0xf6, 0x37, 0x4f, 0x70, 0x5e, 0x19, 0x4d,
	0x25, 0xd5, 0xdc, 0x6b, 0xd2, 0x29, 0xc2, 0xd9, 0x50, 0x59, 0x89, 0xb5, 0xc8, 0x93, 0x6d, 0x86,
	0x23, 0x8f, 0x13, 0x8f, 0xd5, 0xc4, 0x97, 0xb1, 0x97, 0x5e, 0xc2, 0xc2, 0x92, 0xa7, 0xde, 0xad,
	0x4e, 0xd6, 0x22, 0x17, 0xe5, 0xf7, 0x69, 0xf3, 0x29, 0x60, 0x39, 0x88, 0x8b, 0x9e, 0xb4, 0xa9,
	0xfd, 0x3f, 0xcc, 0x8f, 0x90, 0x74, 0x23, 0x6c, 0xf8, 0x30, 0xe8, 0x4f, 0xf3, 0x64, 0x7b, 0x83,
	0xb3, 0xf9, 0x70, 0xf6, 0x04, 0x16, 0x3f, 0xfd, 0xf2, 0x2f, 0x9b, 0x55, 0x00, 0xbf, 0xd1, 0xb1,
	0x46, 0xdc, 0x3d, 0xc0, 0x45, 0xcd, 0x76, 0xfe, 0xa1, 0x5d, 0x72, 0x1f, 0xb7, 0x51, 0x0c, 0xb2,
	0x42, 0x3c, 0x5f, 0x37, 0xc6, 0xbf, 0x06, 0x85, 0x35, 0x5b, 0xd9, 0xb6, 0x2a, 0xa8, 0xa0, 0xcd,
	0xb8, 0x0c, 0x77, 0xdb, 0xb0, 0x8c, 0x8c, 0x5a, 0xc4, 0x8b, 0xbb, 0xaf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xd2, 0xe8, 0x2a, 0xba, 0xd0, 0x01, 0x00, 0x00,
}
