// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: grpcreplay.proto

package grpcreplay

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Entry_Kind int32

const (
	Entry_TYPE_UNSPECIFIED Entry_Kind = 0
	// A unary request.
	// method: the full name of the method
	// message: the request proto
	// is_error: false
	// ref_index: 0
	Entry_REQUEST Entry_Kind = 1
	// A unary response.
	// method: the full name of the method
	// message:
	//   if is_error: a google.rpc.Status proto
	//   else:        the response proto
	// ref_index: index in the sequence of Entries of matching request (1-based)
	Entry_RESPONSE Entry_Kind = 2
	// A method that creates a stream.
	// method: the full name of the method
	// message:
	//   if is_error: a google.rpc.Status proto
	//   else:        nil
	// ref_index: 0
	Entry_CREATE_STREAM Entry_Kind = 3
	// A call to Send on the client returned by a stream-creating method.
	// method: unset
	// message: the proto being sent
	// is_error: false
	// ref_index: index of matching CREATE_STREAM entry (1-based)
	Entry_SEND Entry_Kind = 4
	// A call to Recv on the client returned by a stream-creating method.
	// method: unset
	// message:
	//   if is_error: a google.rpc.Status proto, or nil on EOF
	//   else:        the received message
	// ref_index: index of matching CREATE_STREAM entry
	Entry_RECV Entry_Kind = 5
)

var Entry_Kind_name = map[int32]string{
	0: "TYPE_UNSPECIFIED",
	1: "REQUEST",
	2: "RESPONSE",
	3: "CREATE_STREAM",
	4: "SEND",
	5: "RECV",
}

var Entry_Kind_value = map[string]int32{
	"TYPE_UNSPECIFIED": 0,
	"REQUEST":          1,
	"RESPONSE":         2,
	"CREATE_STREAM":    3,
	"SEND":             4,
	"RECV":             5,
}

func (x Entry_Kind) String() string {
	return proto.EnumName(Entry_Kind_name, int32(x))
}

func (Entry_Kind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a4aab7311a0847da, []int{0, 0}
}

// An Entry represents a single RPC activity, typically a request or response.
type Entry struct {
	Kind                 Entry_Kind `protobuf:"varint,1,opt,name=kind,proto3,enum=grpcreplay.Entry_Kind" json:"kind,omitempty"`
	Method               string     `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Message              *types.Any `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	IsError              bool       `protobuf:"varint,4,opt,name=is_error,json=isError,proto3" json:"is_error,omitempty"`
	RefIndex             int32      `protobuf:"varint,5,opt,name=ref_index,json=refIndex,proto3" json:"ref_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4aab7311a0847da, []int{0}
}
func (m *Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entry.Unmarshal(m, b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entry.Marshal(b, m, deterministic)
}
func (m *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(m, src)
}
func (m *Entry) XXX_Size() int {
	return xxx_messageInfo_Entry.Size(m)
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

func (m *Entry) GetKind() Entry_Kind {
	if m != nil {
		return m.Kind
	}
	return Entry_TYPE_UNSPECIFIED
}

func (m *Entry) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Entry) GetMessage() *types.Any {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Entry) GetIsError() bool {
	if m != nil {
		return m.IsError
	}
	return false
}

func (m *Entry) GetRefIndex() int32 {
	if m != nil {
		return m.RefIndex
	}
	return 0
}

func init() {
	proto.RegisterEnum("grpcreplay.Entry_Kind", Entry_Kind_name, Entry_Kind_value)
	proto.RegisterType((*Entry)(nil), "grpcreplay.Entry")
}

func init() { proto.RegisterFile("grpcreplay.proto", fileDescriptor_a4aab7311a0847da) }

var fileDescriptor_a4aab7311a0847da = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0x41, 0x4f, 0xf2, 0x40,
	0x10, 0x86, 0xbf, 0x85, 0x16, 0xca, 0xf0, 0x69, 0xd6, 0x0d, 0x21, 0x45, 0x2f, 0x0d, 0xa7, 0xc6,
	0xc3, 0x92, 0xe0, 0x2f, 0x20, 0x30, 0x26, 0xc4, 0x88, 0xb8, 0x5b, 0x4c, 0xbc, 0xd8, 0x80, 0x5d,
	0x6a, 0x23, 0xec, 0x92, 0x6d, 0x4d, 0xec, 0xef, 0xf0, 0x0f, 0x9b, 0x56, 0x88, 0xde, 0xe6, 0x99,
	0x79, 0x32, 0xef, 0x0b, 0x34, 0xb5, 0x87, 0x57, 0xab, 0x0e, 0xbb, 0x75, 0xc9, 0x0f, 0xd6, 0x14,
	0x86, 0xc1, 0xef, 0xe6, 0x72, 0x90, 0x1a, 0x93, 0xee, 0xd4, 0xa8, 0xbe, 0x6c, 0x3e, 0xb6, 0xa3,
	0xb5, 0x3e, 0x6a, 0xc3, 0xaf, 0x06, 0xb8, 0xa8, 0x0b, 0x5b, 0xb2, 0x6b, 0x70, 0xde, 0x33, 0x9d,
	0xf8, 0x24, 0x20, 0xe1, 0xf9, 0xb8, 0xcf, 0xff, 0x7c, 0xac, 0x05, 0x7e, 0x97, 0xe9, 0x44, 0xd4,
	0x0e, 0xeb, 0x43, 0x6b, 0xaf, 0x8a, 0x37, 0x93, 0xf8, 0x8d, 0x80, 0x84, 0x1d, 0x71, 0x24, 0xc6,
	0xa1, 0xbd, 0x57, 0x79, 0xbe, 0x4e, 0x95, 0xdf, 0x0c, 0x48, 0xd8, 0x1d, 0xf7, 0xf8, 0x4f, 0x34,
	0x3f, 0x45, 0xf3, 0x89, 0x2e, 0xc5, 0x49, 0x62, 0x03, 0xf0, 0xb2, 0x3c, 0x56, 0xd6, 0x1a, 0xeb,
	0x3b, 0x01, 0x09, 0x3d, 0xd1, 0xce, 0x72, 0xac, 0x90, 0x5d, 0x41, 0xc7, 0xaa, 0x6d, 0x9c, 0xe9,
	0x44, 0x7d, 0xfa, 0x6e, 0x40, 0x42, 0x57, 0x78, 0x56, 0x6d, 0xe7, 0x15, 0x0f, 0x5f, 0xc0, 0xa9,
	0xda, 0xb0, 0x1e, 0xd0, 0xe8, 0x79, 0x89, 0xf1, 0x6a, 0x21, 0x97, 0x38, 0x9d, 0xdf, 0xce, 0x71,
	0x46, 0xff, 0xb1, 0x2e, 0xb4, 0x05, 0x3e, 0xae, 0x50, 0x46, 0x94, 0xb0, 0xff, 0xe0, 0x09, 0x94,
	0xcb, 0x87, 0x85, 0x44, 0xda, 0x60, 0x17, 0x70, 0x36, 0x15, 0x38, 0x89, 0x30, 0x96, 0x91, 0xc0,
	0xc9, 0x3d, 0x6d, 0x32, 0x0f, 0x1c, 0x89, 0x8b, 0x19, 0x75, 0xaa, 0x49, 0xe0, 0xf4, 0x89, 0xba,
	0x9b, 0x56, 0x5d, 0xf7, 0xe6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x58, 0x8a, 0x3c, 0x2d, 0x57, 0x01,
	0x00, 0x00,
}
