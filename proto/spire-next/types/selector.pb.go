// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spire-next/types/selector.proto

package types

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type SelectorMatch_MatchBehavior int32

const (
	// Indicates that the selectors in this match are equal to the
	// candidate selectors, independent of ordering.
	SelectorMatch_MATCH_EXACT SelectorMatch_MatchBehavior = 0
	// Indicates that the selectors in this match are a subset of the
	// candidate selectors.
	SelectorMatch_MATCH_SUBSET SelectorMatch_MatchBehavior = 1
)

var SelectorMatch_MatchBehavior_name = map[int32]string{
	0: "MATCH_EXACT",
	1: "MATCH_SUBSET",
}

var SelectorMatch_MatchBehavior_value = map[string]int32{
	"MATCH_EXACT":  0,
	"MATCH_SUBSET": 1,
}

func (x SelectorMatch_MatchBehavior) String() string {
	return proto.EnumName(SelectorMatch_MatchBehavior_name, int32(x))
}

func (SelectorMatch_MatchBehavior) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_db3686cdbd811d0c, []int{1, 0}
}

type Selector struct {
	// The type of the selector. This is typically the name of the plugin that
	// produces the selector.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// The value of the selector.
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Selector) Reset()         { *m = Selector{} }
func (m *Selector) String() string { return proto.CompactTextString(m) }
func (*Selector) ProtoMessage()    {}
func (*Selector) Descriptor() ([]byte, []int) {
	return fileDescriptor_db3686cdbd811d0c, []int{0}
}

func (m *Selector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Selector.Unmarshal(m, b)
}
func (m *Selector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Selector.Marshal(b, m, deterministic)
}
func (m *Selector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Selector.Merge(m, src)
}
func (m *Selector) XXX_Size() int {
	return xxx_messageInfo_Selector.Size(m)
}
func (m *Selector) XXX_DiscardUnknown() {
	xxx_messageInfo_Selector.DiscardUnknown(m)
}

var xxx_messageInfo_Selector proto.InternalMessageInfo

func (m *Selector) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Selector) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SelectorMatch struct {
	// The set of selectors to match on.
	Selectors []*Selector `protobuf:"bytes,1,rep,name=selectors,proto3" json:"selectors,omitempty"`
	// How to match the selectors.
	Match                SelectorMatch_MatchBehavior `protobuf:"varint,2,opt,name=match,proto3,enum=spire.types.SelectorMatch_MatchBehavior" json:"match,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *SelectorMatch) Reset()         { *m = SelectorMatch{} }
func (m *SelectorMatch) String() string { return proto.CompactTextString(m) }
func (*SelectorMatch) ProtoMessage()    {}
func (*SelectorMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_db3686cdbd811d0c, []int{1}
}

func (m *SelectorMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SelectorMatch.Unmarshal(m, b)
}
func (m *SelectorMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SelectorMatch.Marshal(b, m, deterministic)
}
func (m *SelectorMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SelectorMatch.Merge(m, src)
}
func (m *SelectorMatch) XXX_Size() int {
	return xxx_messageInfo_SelectorMatch.Size(m)
}
func (m *SelectorMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_SelectorMatch.DiscardUnknown(m)
}

var xxx_messageInfo_SelectorMatch proto.InternalMessageInfo

func (m *SelectorMatch) GetSelectors() []*Selector {
	if m != nil {
		return m.Selectors
	}
	return nil
}

func (m *SelectorMatch) GetMatch() SelectorMatch_MatchBehavior {
	if m != nil {
		return m.Match
	}
	return SelectorMatch_MATCH_EXACT
}

func init() {
	proto.RegisterEnum("spire.types.SelectorMatch_MatchBehavior", SelectorMatch_MatchBehavior_name, SelectorMatch_MatchBehavior_value)
	proto.RegisterType((*Selector)(nil), "spire.types.Selector")
	proto.RegisterType((*SelectorMatch)(nil), "spire.types.SelectorMatch")
}

func init() { proto.RegisterFile("spire-next/types/selector.proto", fileDescriptor_db3686cdbd811d0c) }

var fileDescriptor_db3686cdbd811d0c = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0x2e, 0xc8, 0x2c,
	0x4a, 0xd5, 0xcd, 0x4b, 0xad, 0x28, 0xd1, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2f, 0x4e, 0xcd,
	0x49, 0x4d, 0x2e, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x06, 0x2b, 0xd0,
	0x03, 0xcb, 0x29, 0x99, 0x70, 0x71, 0x04, 0x43, 0xa5, 0x85, 0x84, 0xb8, 0x58, 0x40, 0x82, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x08, 0x17, 0x6b, 0x59, 0x62, 0x4e, 0x69,
	0xaa, 0x04, 0x13, 0x58, 0x10, 0xc2, 0x51, 0xda, 0xc1, 0xc8, 0xc5, 0x0b, 0xd3, 0xe6, 0x9b, 0x58,
	0x92, 0x9c, 0x21, 0x64, 0xcc, 0xc5, 0x09, 0xb3, 0xa6, 0x58, 0x82, 0x51, 0x81, 0x59, 0x83, 0xdb,
	0x48, 0x54, 0x0f, 0xc9, 0x22, 0x3d, 0x98, 0xf2, 0x20, 0x84, 0x3a, 0x21, 0x3b, 0x2e, 0xd6, 0x5c,
	0x90, 0x6e, 0xb0, 0xe1, 0x7c, 0x46, 0x1a, 0x58, 0x35, 0x80, 0xcd, 0xd7, 0x03, 0x93, 0x4e, 0xa9,
	0x19, 0x89, 0x65, 0x99, 0xf9, 0x45, 0x41, 0x10, 0x6d, 0x4a, 0x46, 0x5c, 0xbc, 0x28, 0xe2, 0x42,
	0xfc, 0x5c, 0xdc, 0xbe, 0x8e, 0x21, 0xce, 0x1e, 0xf1, 0xae, 0x11, 0x8e, 0xce, 0x21, 0x02, 0x0c,
	0x42, 0x02, 0x5c, 0x3c, 0x10, 0x81, 0xe0, 0x50, 0xa7, 0x60, 0xd7, 0x10, 0x01, 0x46, 0x27, 0x83,
	0x28, 0xbd, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xe2, 0x82, 0xcc,
	0xb4, 0xb4, 0x54, 0x7d, 0xb0, 0xbd, 0xfa, 0xe0, 0xe0, 0xd1, 0x47, 0x0f, 0xbe, 0x24, 0x36, 0xb0,
	0xb8, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x19, 0xe0, 0x7b, 0x59, 0x01, 0x00, 0x00,
}
