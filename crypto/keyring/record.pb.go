// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/crypto/keyring/v1/record.proto

package keyring

import (
	fmt "fmt"
	types "github.com/opzlabs/cosmos-sdk/codec/types"
	hd "github.com/opzlabs/cosmos-sdk/crypto/hd"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// Record is used for representing a key in the keyring.
type Record struct {
	// name represents a name of Record
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// pub_key represents a public key in any format
	PubKey *types.Any `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	// Record contains one of the following items
	//
	// Types that are valid to be assigned to Item:
	//	*Record_Local_
	//	*Record_Ledger_
	//	*Record_Multi_
	//	*Record_Offline_
	Item isRecord_Item `protobuf_oneof:"item"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_36d640103edea005, []int{0}
}
func (m *Record) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Record.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return m.Size()
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

type isRecord_Item interface {
	isRecord_Item()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Record_Local_ struct {
	Local *Record_Local `protobuf:"bytes,3,opt,name=local,proto3,oneof" json:"local,omitempty"`
}
type Record_Ledger_ struct {
	Ledger *Record_Ledger `protobuf:"bytes,4,opt,name=ledger,proto3,oneof" json:"ledger,omitempty"`
}
type Record_Multi_ struct {
	Multi *Record_Multi `protobuf:"bytes,5,opt,name=multi,proto3,oneof" json:"multi,omitempty"`
}
type Record_Offline_ struct {
	Offline *Record_Offline `protobuf:"bytes,6,opt,name=offline,proto3,oneof" json:"offline,omitempty"`
}

func (*Record_Local_) isRecord_Item()   {}
func (*Record_Ledger_) isRecord_Item()  {}
func (*Record_Multi_) isRecord_Item()   {}
func (*Record_Offline_) isRecord_Item() {}

func (m *Record) GetItem() isRecord_Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *Record) GetLocal() *Record_Local {
	if x, ok := m.GetItem().(*Record_Local_); ok {
		return x.Local
	}
	return nil
}

func (m *Record) GetLedger() *Record_Ledger {
	if x, ok := m.GetItem().(*Record_Ledger_); ok {
		return x.Ledger
	}
	return nil
}

func (m *Record) GetMulti() *Record_Multi {
	if x, ok := m.GetItem().(*Record_Multi_); ok {
		return x.Multi
	}
	return nil
}

func (m *Record) GetOffline() *Record_Offline {
	if x, ok := m.GetItem().(*Record_Offline_); ok {
		return x.Offline
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Record) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Record_Local_)(nil),
		(*Record_Ledger_)(nil),
		(*Record_Multi_)(nil),
		(*Record_Offline_)(nil),
	}
}

// Item is a keyring item stored in a keyring backend.
// Local item
type Record_Local struct {
	PrivKey *types.Any `protobuf:"bytes,1,opt,name=priv_key,json=privKey,proto3" json:"priv_key,omitempty"`
}

func (m *Record_Local) Reset()         { *m = Record_Local{} }
func (m *Record_Local) String() string { return proto.CompactTextString(m) }
func (*Record_Local) ProtoMessage()    {}
func (*Record_Local) Descriptor() ([]byte, []int) {
	return fileDescriptor_36d640103edea005, []int{0, 0}
}
func (m *Record_Local) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Record_Local) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Record_Local.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Record_Local) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record_Local.Merge(m, src)
}
func (m *Record_Local) XXX_Size() int {
	return m.Size()
}
func (m *Record_Local) XXX_DiscardUnknown() {
	xxx_messageInfo_Record_Local.DiscardUnknown(m)
}

var xxx_messageInfo_Record_Local proto.InternalMessageInfo

// Ledger item
type Record_Ledger struct {
	Path *hd.BIP44Params `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (m *Record_Ledger) Reset()         { *m = Record_Ledger{} }
func (m *Record_Ledger) String() string { return proto.CompactTextString(m) }
func (*Record_Ledger) ProtoMessage()    {}
func (*Record_Ledger) Descriptor() ([]byte, []int) {
	return fileDescriptor_36d640103edea005, []int{0, 1}
}
func (m *Record_Ledger) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Record_Ledger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Record_Ledger.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Record_Ledger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record_Ledger.Merge(m, src)
}
func (m *Record_Ledger) XXX_Size() int {
	return m.Size()
}
func (m *Record_Ledger) XXX_DiscardUnknown() {
	xxx_messageInfo_Record_Ledger.DiscardUnknown(m)
}

var xxx_messageInfo_Record_Ledger proto.InternalMessageInfo

// Multi item
type Record_Multi struct {
}

func (m *Record_Multi) Reset()         { *m = Record_Multi{} }
func (m *Record_Multi) String() string { return proto.CompactTextString(m) }
func (*Record_Multi) ProtoMessage()    {}
func (*Record_Multi) Descriptor() ([]byte, []int) {
	return fileDescriptor_36d640103edea005, []int{0, 2}
}
func (m *Record_Multi) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Record_Multi) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Record_Multi.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Record_Multi) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record_Multi.Merge(m, src)
}
func (m *Record_Multi) XXX_Size() int {
	return m.Size()
}
func (m *Record_Multi) XXX_DiscardUnknown() {
	xxx_messageInfo_Record_Multi.DiscardUnknown(m)
}

var xxx_messageInfo_Record_Multi proto.InternalMessageInfo

// Offline item
type Record_Offline struct {
}

func (m *Record_Offline) Reset()         { *m = Record_Offline{} }
func (m *Record_Offline) String() string { return proto.CompactTextString(m) }
func (*Record_Offline) ProtoMessage()    {}
func (*Record_Offline) Descriptor() ([]byte, []int) {
	return fileDescriptor_36d640103edea005, []int{0, 3}
}
func (m *Record_Offline) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Record_Offline) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Record_Offline.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Record_Offline) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record_Offline.Merge(m, src)
}
func (m *Record_Offline) XXX_Size() int {
	return m.Size()
}
func (m *Record_Offline) XXX_DiscardUnknown() {
	xxx_messageInfo_Record_Offline.DiscardUnknown(m)
}

var xxx_messageInfo_Record_Offline proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Record)(nil), "cosmos.crypto.keyring.v1.Record")
	proto.RegisterType((*Record_Local)(nil), "cosmos.crypto.keyring.v1.Record.Local")
	proto.RegisterType((*Record_Ledger)(nil), "cosmos.crypto.keyring.v1.Record.Ledger")
	proto.RegisterType((*Record_Multi)(nil), "cosmos.crypto.keyring.v1.Record.Multi")
	proto.RegisterType((*Record_Offline)(nil), "cosmos.crypto.keyring.v1.Record.Offline")
}

func init() {
	proto.RegisterFile("cosmos/crypto/keyring/v1/record.proto", fileDescriptor_36d640103edea005)
}

var fileDescriptor_36d640103edea005 = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x3d, 0x6b, 0xdb, 0x40,
	0x1c, 0xc6, 0xa5, 0x5a, 0x2f, 0xf5, 0x75, 0x3b, 0x3c, 0xa8, 0xa2, 0x08, 0x53, 0x68, 0x6b, 0x28,
	0xbe, 0xc3, 0xad, 0x87, 0x4e, 0x06, 0x9b, 0x0e, 0x36, 0x4e, 0x88, 0xd1, 0x98, 0x25, 0xe8, 0xe5,
	0x2c, 0x09, 0x4b, 0x3a, 0x71, 0x92, 0x0c, 0xfa, 0x16, 0xf9, 0x58, 0x1e, 0x3d, 0x66, 0x4c, 0xec,
	0x2d, 0x9f, 0x22, 0xdc, 0x9d, 0x3c, 0xc4, 0x90, 0x38, 0x93, 0x4e, 0xdc, 0xef, 0xf9, 0x3f, 0xcf,
	0x73, 0xfc, 0xc1, 0x8f, 0x80, 0x96, 0x19, 0x2d, 0x71, 0xc0, 0x9a, 0xa2, 0xa2, 0x78, 0x43, 0x1a,
	0x96, 0xe4, 0x11, 0xde, 0x8e, 0x30, 0x23, 0x01, 0x65, 0x21, 0x2a, 0x18, 0xad, 0x28, 0xb4, 0x24,
	0x86, 0x24, 0x86, 0x5a, 0x0c, 0x6d, 0x47, 0x76, 0x2f, 0xa2, 0x11, 0x15, 0x10, 0xe6, 0x27, 0xc9,
	0xdb, 0x5f, 0x23, 0x4a, 0xa3, 0x94, 0x60, 0xf1, 0xe7, 0xd7, 0x6b, 0xec, 0xe5, 0x4d, 0x7b, 0xf5,
	0xed, 0xb5, 0x63, 0x1c, 0x72, 0xb3, 0xb8, 0x35, 0xfa, 0xfe, 0xdc, 0x01, 0x86, 0x2b, 0x9c, 0x21,
	0x04, 0x5a, 0xee, 0x65, 0xc4, 0x52, 0xfb, 0xea, 0xa0, 0xeb, 0x8a, 0x33, 0x1c, 0x02, 0xb3, 0xa8,
	0xfd, 0xbb, 0x0d, 0x69, 0xac, 0x4f, 0x7d, 0x75, 0xf0, 0xe5, 0x4f, 0x0f, 0x49, 0x27, 0x74, 0x72,
	0x42, 0xd3, 0xbc, 0x71, 0x8d, 0xa2, 0xf6, 0x97, 0xa4, 0x81, 0x13, 0xa0, 0xa7, 0x34, 0xf0, 0x52,
	0xab, 0x23, 0xe0, 0x9f, 0xe8, 0xad, 0x1a, 0x48, 0x7a, 0xa2, 0x2b, 0x4e, 0xcf, 0x15, 0x57, 0xca,
	0xe0, 0x14, 0x18, 0x29, 0x09, 0x23, 0xc2, 0x2c, 0x4d, 0x0c, 0xf8, 0x75, 0x79, 0x80, 0xc0, 0xe7,
	0x8a, 0xdb, 0x0a, 0x79, 0x84, 0xac, 0x4e, 0xab, 0xc4, 0xd2, 0x3f, 0x18, 0xe1, 0x9a, 0xd3, 0x3c,
	0x82, 0x90, 0xc1, 0xff, 0xc0, 0xa4, 0xeb, 0x75, 0x9a, 0xe4, 0xc4, 0x32, 0xc4, 0x84, 0xc1, 0xc5,
	0x09, 0x37, 0x92, 0x9f, 0x2b, 0xee, 0x49, 0x6a, 0xff, 0x03, 0xba, 0xa8, 0x06, 0x31, 0xf8, 0x5c,
	0xb0, 0x64, 0x2b, 0x5e, 0x50, 0x7d, 0xe7, 0x05, 0x4d, 0x4e, 0x2d, 0x49, 0x63, 0x4f, 0x80, 0x21,
	0x3b, 0xc1, 0x31, 0xd0, 0x0a, 0xaf, 0x8a, 0x5b, 0x59, 0xff, 0x2c, 0x46, 0x1c, 0xf2, 0x04, 0xb3,
	0xc5, 0x6a, 0x3c, 0x5e, 0x79, 0xcc, 0xcb, 0x4a, 0x57, 0xd0, 0xb6, 0x09, 0x74, 0xd1, 0xc8, 0xee,
	0x02, 0xb3, 0x0d, 0x36, 0x33, 0x80, 0x96, 0x54, 0x24, 0x9b, 0x2d, 0x76, 0x4f, 0x8e, 0xb2, 0x3b,
	0x38, 0xea, 0xfe, 0xe0, 0xa8, 0x8f, 0x07, 0x47, 0xbd, 0x3f, 0x3a, 0xca, 0xfe, 0xe8, 0x28, 0x0f,
	0x47, 0x47, 0xb9, 0xfd, 0x1d, 0x25, 0x55, 0x5c, 0xfb, 0x28, 0xa0, 0x19, 0x3e, 0xed, 0x8c, 0xf8,
	0x0c, 0xcb, 0x70, 0x73, 0xb6, 0xb0, 0xbe, 0x21, 0xd2, 0xff, 0x7d, 0x09, 0x00, 0x00, 0xff, 0xff,
	0x64, 0x83, 0x0c, 0x89, 0xd0, 0x02, 0x00, 0x00,
}

func (m *Record) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Record) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Record) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Item != nil {
		{
			size := m.Item.Size()
			i -= size
			if _, err := m.Item.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.PubKey != nil {
		{
			size, err := m.PubKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRecord(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintRecord(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Record_Local_) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Record_Local_) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Local != nil {
		{
			size, err := m.Local.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRecord(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *Record_Ledger_) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Record_Ledger_) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Ledger != nil {
		{
			size, err := m.Ledger.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRecord(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *Record_Multi_) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Record_Multi_) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Multi != nil {
		{
			size, err := m.Multi.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRecord(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *Record_Offline_) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Record_Offline_) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Offline != nil {
		{
			size, err := m.Offline.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRecord(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	return len(dAtA) - i, nil
}
func (m *Record_Local) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Record_Local) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Record_Local) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PrivKey != nil {
		{
			size, err := m.PrivKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRecord(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Record_Ledger) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Record_Ledger) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Record_Ledger) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Path != nil {
		{
			size, err := m.Path.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRecord(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Record_Multi) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Record_Multi) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Record_Multi) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *Record_Offline) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Record_Offline) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Record_Offline) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintRecord(dAtA []byte, offset int, v uint64) int {
	offset -= sovRecord(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Record) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRecord(uint64(l))
	}
	if m.PubKey != nil {
		l = m.PubKey.Size()
		n += 1 + l + sovRecord(uint64(l))
	}
	if m.Item != nil {
		n += m.Item.Size()
	}
	return n
}

func (m *Record_Local_) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Local != nil {
		l = m.Local.Size()
		n += 1 + l + sovRecord(uint64(l))
	}
	return n
}
func (m *Record_Ledger_) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ledger != nil {
		l = m.Ledger.Size()
		n += 1 + l + sovRecord(uint64(l))
	}
	return n
}
func (m *Record_Multi_) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Multi != nil {
		l = m.Multi.Size()
		n += 1 + l + sovRecord(uint64(l))
	}
	return n
}
func (m *Record_Offline_) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Offline != nil {
		l = m.Offline.Size()
		n += 1 + l + sovRecord(uint64(l))
	}
	return n
}
func (m *Record_Local) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PrivKey != nil {
		l = m.PrivKey.Size()
		n += 1 + l + sovRecord(uint64(l))
	}
	return n
}

func (m *Record_Ledger) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Path != nil {
		l = m.Path.Size()
		n += 1 + l + sovRecord(uint64(l))
	}
	return n
}

func (m *Record_Multi) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *Record_Offline) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovRecord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRecord(x uint64) (n int) {
	return sovRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Record) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecord
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
			return fmt.Errorf("proto: Record: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Record: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
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
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
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
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PubKey == nil {
				m.PubKey = &types.Any{}
			}
			if err := m.PubKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Local", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
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
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Record_Local{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Item = &Record_Local_{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ledger", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
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
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Record_Ledger{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Item = &Record_Ledger_{v}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Multi", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
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
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Record_Multi{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Item = &Record_Multi_{v}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offline", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
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
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Record_Offline{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Item = &Record_Offline_{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRecord
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
func (m *Record_Local) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecord
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
			return fmt.Errorf("proto: Local: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Local: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrivKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
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
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PrivKey == nil {
				m.PrivKey = &types.Any{}
			}
			if err := m.PrivKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRecord
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
func (m *Record_Ledger) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecord
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
			return fmt.Errorf("proto: Ledger: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ledger: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
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
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Path == nil {
				m.Path = &hd.BIP44Params{}
			}
			if err := m.Path.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRecord
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
func (m *Record_Multi) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecord
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
			return fmt.Errorf("proto: Multi: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Multi: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRecord
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
func (m *Record_Offline) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecord
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
			return fmt.Errorf("proto: Offline: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Offline: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRecord
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
func skipRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRecord
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
					return 0, ErrIntOverflowRecord
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
					return 0, ErrIntOverflowRecord
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
				return 0, ErrInvalidLengthRecord
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRecord
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRecord
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRecord        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRecord          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRecord = fmt.Errorf("proto: unexpected end of group")
)
