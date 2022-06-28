// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: conflict/conflict_vote.proto

package types

import (
	fmt "fmt"
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

type Provider struct {
	Account  string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Response []byte `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (m *Provider) Reset()         { *m = Provider{} }
func (m *Provider) String() string { return proto.CompactTextString(m) }
func (*Provider) ProtoMessage()    {}
func (*Provider) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5ff6a0d8edaa7f1, []int{0}
}
func (m *Provider) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Provider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Provider.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Provider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Provider.Merge(m, src)
}
func (m *Provider) XXX_Size() int {
	return m.Size()
}
func (m *Provider) XXX_DiscardUnknown() {
	xxx_messageInfo_Provider.DiscardUnknown(m)
}

var xxx_messageInfo_Provider proto.InternalMessageInfo

func (m *Provider) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *Provider) GetResponse() []byte {
	if m != nil {
		return m.Response
	}
	return nil
}

type Vote struct {
	Hash   []byte `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
	Result int64  `protobuf:"varint,2,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (m *Vote) Reset()         { *m = Vote{} }
func (m *Vote) String() string { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()    {}
func (*Vote) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5ff6a0d8edaa7f1, []int{1}
}
func (m *Vote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Vote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Vote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vote.Merge(m, src)
}
func (m *Vote) XXX_Size() int {
	return m.Size()
}
func (m *Vote) XXX_DiscardUnknown() {
	xxx_messageInfo_Vote.DiscardUnknown(m)
}

var xxx_messageInfo_Vote proto.InternalMessageInfo

func (m *Vote) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Vote) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

type ConflictVote struct {
	Index          string          `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	VoteDeadline   int64           `protobuf:"varint,3,opt,name=voteDeadline,proto3" json:"voteDeadline,omitempty"`
	VoteStartBlock int64           `protobuf:"varint,4,opt,name=voteStartBlock,proto3" json:"voteStartBlock,omitempty"`
	VoteState      int64           `protobuf:"varint,5,opt,name=voteState,proto3" json:"voteState,omitempty"`
	ChainID        string          `protobuf:"bytes,6,opt,name=chainID,proto3" json:"chainID,omitempty"`
	ApiUrl         string          `protobuf:"bytes,7,opt,name=apiUrl,proto3" json:"apiUrl,omitempty"`
	RequestData    []byte          `protobuf:"bytes,8,opt,name=requestData,proto3" json:"requestData,omitempty"`
	RequestBlock   int64           `protobuf:"varint,9,opt,name=requestBlock,proto3" json:"requestBlock,omitempty"`
	FirstProvider  Provider        `protobuf:"bytes,10,opt,name=firstProvider,proto3" json:"firstProvider"`
	SecondProvider Provider        `protobuf:"bytes,11,opt,name=secondProvider,proto3" json:"secondProvider"`
	VotersHash     map[string]Vote `protobuf:"bytes,13,rep,name=votersHash,proto3" json:"votersHash" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *ConflictVote) Reset()         { *m = ConflictVote{} }
func (m *ConflictVote) String() string { return proto.CompactTextString(m) }
func (*ConflictVote) ProtoMessage()    {}
func (*ConflictVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5ff6a0d8edaa7f1, []int{2}
}
func (m *ConflictVote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConflictVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConflictVote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConflictVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConflictVote.Merge(m, src)
}
func (m *ConflictVote) XXX_Size() int {
	return m.Size()
}
func (m *ConflictVote) XXX_DiscardUnknown() {
	xxx_messageInfo_ConflictVote.DiscardUnknown(m)
}

var xxx_messageInfo_ConflictVote proto.InternalMessageInfo

func (m *ConflictVote) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *ConflictVote) GetVoteDeadline() int64 {
	if m != nil {
		return m.VoteDeadline
	}
	return 0
}

func (m *ConflictVote) GetVoteStartBlock() int64 {
	if m != nil {
		return m.VoteStartBlock
	}
	return 0
}

func (m *ConflictVote) GetVoteState() int64 {
	if m != nil {
		return m.VoteState
	}
	return 0
}

func (m *ConflictVote) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *ConflictVote) GetApiUrl() string {
	if m != nil {
		return m.ApiUrl
	}
	return ""
}

func (m *ConflictVote) GetRequestData() []byte {
	if m != nil {
		return m.RequestData
	}
	return nil
}

func (m *ConflictVote) GetRequestBlock() int64 {
	if m != nil {
		return m.RequestBlock
	}
	return 0
}

func (m *ConflictVote) GetFirstProvider() Provider {
	if m != nil {
		return m.FirstProvider
	}
	return Provider{}
}

func (m *ConflictVote) GetSecondProvider() Provider {
	if m != nil {
		return m.SecondProvider
	}
	return Provider{}
}

func (m *ConflictVote) GetVotersHash() map[string]Vote {
	if m != nil {
		return m.VotersHash
	}
	return nil
}

func init() {
	proto.RegisterType((*Provider)(nil), "lavanet.lava.conflict.Provider")
	proto.RegisterType((*Vote)(nil), "lavanet.lava.conflict.Vote")
	proto.RegisterType((*ConflictVote)(nil), "lavanet.lava.conflict.ConflictVote")
	proto.RegisterMapType((map[string]Vote)(nil), "lavanet.lava.conflict.ConflictVote.VotersHashEntry")
}

func init() { proto.RegisterFile("conflict/conflict_vote.proto", fileDescriptor_c5ff6a0d8edaa7f1) }

var fileDescriptor_c5ff6a0d8edaa7f1 = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x6e, 0xd6, 0xb4, 0x6b, 0x5f, 0xba, 0x81, 0xac, 0x81, 0xac, 0x32, 0x65, 0x51, 0x0f, 0x28,
	0xa7, 0x44, 0x74, 0x17, 0xc4, 0x09, 0x95, 0x22, 0x81, 0x10, 0x12, 0x0a, 0x62, 0x12, 0xbb, 0x20,
	0x2f, 0xf5, 0x5a, 0x6b, 0xc1, 0x2e, 0xb6, 0x53, 0xad, 0xff, 0x82, 0x9f, 0xb5, 0xe3, 0x8e, 0x70,
	0x41, 0xa8, 0xfd, 0x23, 0xc8, 0x8e, 0x53, 0xba, 0x8a, 0x1d, 0x38, 0xf9, 0x7d, 0x5f, 0xbf, 0xef,
	0xf3, 0x7b, 0x7d, 0x31, 0x1c, 0xe7, 0x82, 0x5f, 0x16, 0x2c, 0xd7, 0x69, 0x5d, 0x7c, 0x59, 0x08,
	0x4d, 0x93, 0xb9, 0x14, 0x5a, 0xa0, 0x47, 0x05, 0x59, 0x10, 0x4e, 0x75, 0x62, 0xce, 0xa4, 0x56,
	0xf4, 0x8f, 0xa6, 0x62, 0x2a, 0xac, 0x22, 0x35, 0x55, 0x25, 0x1e, 0xbc, 0x84, 0xce, 0x07, 0x29,
	0x16, 0x6c, 0x42, 0x25, 0xc2, 0xb0, 0x4f, 0xf2, 0x5c, 0x94, 0x5c, 0x63, 0x2f, 0xf2, 0xe2, 0x6e,
	0x56, 0x43, 0xd4, 0x87, 0x8e, 0xa4, 0x6a, 0x2e, 0xb8, 0xa2, 0x78, 0x2f, 0xf2, 0xe2, 0x5e, 0xb6,
	0xc1, 0x83, 0x21, 0xf8, 0x67, 0x42, 0x53, 0x84, 0xc0, 0x7f, 0x43, 0xd4, 0xcc, 0x5a, 0x7b, 0x99,
	0xad, 0xd1, 0x63, 0x68, 0x67, 0x54, 0x95, 0x85, 0xb6, 0xae, 0x66, 0xe6, 0xd0, 0xe0, 0xa7, 0x0f,
	0xbd, 0x57, 0xae, 0x31, 0x6b, 0x3e, 0x82, 0x16, 0xe3, 0x13, 0x7a, 0xed, 0x2e, 0xae, 0x00, 0x1a,
	0x40, 0xcf, 0xcc, 0x35, 0xa6, 0x64, 0x52, 0x30, 0x4e, 0x71, 0xd3, 0x86, 0xdc, 0xe1, 0xd0, 0x53,
	0x38, 0x34, 0xf8, 0xa3, 0x26, 0x52, 0x8f, 0x0a, 0x91, 0x5f, 0x61, 0xdf, 0xaa, 0x76, 0x58, 0x74,
	0x0c, 0x5d, 0xc7, 0x68, 0x8a, 0x5b, 0x56, 0xf2, 0x97, 0x30, 0xa3, 0xe7, 0x33, 0xc2, 0xf8, 0xdb,
	0x31, 0x6e, 0x57, 0xa3, 0x3b, 0x68, 0x46, 0x20, 0x73, 0xf6, 0x49, 0x16, 0x78, 0xdf, 0xfe, 0xe0,
	0x10, 0x8a, 0x20, 0x90, 0xf4, 0x5b, 0x49, 0x95, 0x1e, 0x13, 0x4d, 0x70, 0xc7, 0x4e, 0xbd, 0x4d,
	0x99, 0xee, 0x1d, 0xac, 0xfa, 0xea, 0x56, 0xdd, 0x6f, 0x73, 0xe8, 0x1d, 0x1c, 0x5c, 0x32, 0xa9,
	0x74, 0xbd, 0x03, 0x0c, 0x91, 0x17, 0x07, 0xc3, 0x93, 0xe4, 0x9f, 0x3b, 0x4c, 0x6a, 0xd9, 0xc8,
	0xbf, 0xf9, 0x75, 0xd2, 0xc8, 0xee, 0x7a, 0xd1, 0x7b, 0x38, 0x54, 0x34, 0x17, 0x7c, 0xb2, 0x49,
	0x0b, 0xfe, 0x27, 0x6d, 0xc7, 0x8c, 0x3e, 0x03, 0x98, 0x3f, 0x48, 0x2a, 0xbb, 0xd6, 0x83, 0xa8,
	0x19, 0x07, 0xc3, 0xd3, 0x7b, 0xa2, 0xb6, 0x97, 0x99, 0x9c, 0x6d, 0x5c, 0xaf, 0xb9, 0x96, 0x4b,
	0x17, 0xbf, 0x15, 0xd6, 0x3f, 0x87, 0x07, 0x3b, 0x22, 0xf4, 0x10, 0x9a, 0x57, 0x74, 0xe9, 0xf6,
	0x6f, 0x4a, 0xf4, 0x0c, 0x5a, 0x0b, 0x52, 0x94, 0xd5, 0x17, 0x17, 0x0c, 0x9f, 0xdc, 0x73, 0xb5,
	0x09, 0xca, 0x2a, 0xe5, 0x8b, 0xbd, 0xe7, 0xde, 0x68, 0x74, 0xb3, 0x0a, 0xbd, 0xdb, 0x55, 0xe8,
	0xfd, 0x5e, 0x85, 0xde, 0xf7, 0x75, 0xd8, 0xb8, 0x5d, 0x87, 0x8d, 0x1f, 0xeb, 0xb0, 0x71, 0x1e,
	0x4f, 0x99, 0x9e, 0x95, 0x17, 0x49, 0x2e, 0xbe, 0xa6, 0x2e, 0xcb, 0x9e, 0xe9, 0xf5, 0xe6, 0x1d,
	0xa5, 0x7a, 0x39, 0xa7, 0xea, 0xa2, 0x6d, 0x1f, 0xc7, 0xe9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x5c, 0x24, 0x5d, 0xcf, 0x69, 0x03, 0x00, 0x00,
}

func (m *Provider) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Provider) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Provider) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Response) > 0 {
		i -= len(m.Response)
		copy(dAtA[i:], m.Response)
		i = encodeVarintConflictVote(dAtA, i, uint64(len(m.Response)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintConflictVote(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Vote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Vote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Vote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Result != 0 {
		i = encodeVarintConflictVote(dAtA, i, uint64(m.Result))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintConflictVote(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ConflictVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConflictVote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConflictVote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VotersHash) > 0 {
		for k := range m.VotersHash {
			v := m.VotersHash[k]
			baseI := i
			{
				size, err := (&v).MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintConflictVote(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintConflictVote(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintConflictVote(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x6a
		}
	}
	{
		size, err := m.SecondProvider.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintConflictVote(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	{
		size, err := m.FirstProvider.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintConflictVote(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	if m.RequestBlock != 0 {
		i = encodeVarintConflictVote(dAtA, i, uint64(m.RequestBlock))
		i--
		dAtA[i] = 0x48
	}
	if len(m.RequestData) > 0 {
		i -= len(m.RequestData)
		copy(dAtA[i:], m.RequestData)
		i = encodeVarintConflictVote(dAtA, i, uint64(len(m.RequestData)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.ApiUrl) > 0 {
		i -= len(m.ApiUrl)
		copy(dAtA[i:], m.ApiUrl)
		i = encodeVarintConflictVote(dAtA, i, uint64(len(m.ApiUrl)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ChainID) > 0 {
		i -= len(m.ChainID)
		copy(dAtA[i:], m.ChainID)
		i = encodeVarintConflictVote(dAtA, i, uint64(len(m.ChainID)))
		i--
		dAtA[i] = 0x32
	}
	if m.VoteState != 0 {
		i = encodeVarintConflictVote(dAtA, i, uint64(m.VoteState))
		i--
		dAtA[i] = 0x28
	}
	if m.VoteStartBlock != 0 {
		i = encodeVarintConflictVote(dAtA, i, uint64(m.VoteStartBlock))
		i--
		dAtA[i] = 0x20
	}
	if m.VoteDeadline != 0 {
		i = encodeVarintConflictVote(dAtA, i, uint64(m.VoteDeadline))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintConflictVote(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintConflictVote(dAtA []byte, offset int, v uint64) int {
	offset -= sovConflictVote(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Provider) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovConflictVote(uint64(l))
	}
	l = len(m.Response)
	if l > 0 {
		n += 1 + l + sovConflictVote(uint64(l))
	}
	return n
}

func (m *Vote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovConflictVote(uint64(l))
	}
	if m.Result != 0 {
		n += 1 + sovConflictVote(uint64(m.Result))
	}
	return n
}

func (m *ConflictVote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovConflictVote(uint64(l))
	}
	if m.VoteDeadline != 0 {
		n += 1 + sovConflictVote(uint64(m.VoteDeadline))
	}
	if m.VoteStartBlock != 0 {
		n += 1 + sovConflictVote(uint64(m.VoteStartBlock))
	}
	if m.VoteState != 0 {
		n += 1 + sovConflictVote(uint64(m.VoteState))
	}
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovConflictVote(uint64(l))
	}
	l = len(m.ApiUrl)
	if l > 0 {
		n += 1 + l + sovConflictVote(uint64(l))
	}
	l = len(m.RequestData)
	if l > 0 {
		n += 1 + l + sovConflictVote(uint64(l))
	}
	if m.RequestBlock != 0 {
		n += 1 + sovConflictVote(uint64(m.RequestBlock))
	}
	l = m.FirstProvider.Size()
	n += 1 + l + sovConflictVote(uint64(l))
	l = m.SecondProvider.Size()
	n += 1 + l + sovConflictVote(uint64(l))
	if len(m.VotersHash) > 0 {
		for k, v := range m.VotersHash {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + len(k) + sovConflictVote(uint64(len(k))) + 1 + l + sovConflictVote(uint64(l))
			n += mapEntrySize + 1 + sovConflictVote(uint64(mapEntrySize))
		}
	}
	return n
}

func sovConflictVote(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConflictVote(x uint64) (n int) {
	return sovConflictVote(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Provider) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConflictVote
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
			return fmt.Errorf("proto: Provider: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Provider: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictVote
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
				return ErrInvalidLengthConflictVote
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConflictVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Response", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthConflictVote
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthConflictVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Response = append(m.Response[:0], dAtA[iNdEx:postIndex]...)
			if m.Response == nil {
				m.Response = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConflictVote(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConflictVote
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
func (m *Vote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConflictVote
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
			return fmt.Errorf("proto: Vote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Vote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthConflictVote
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthConflictVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			m.Result = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Result |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConflictVote(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConflictVote
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
func (m *ConflictVote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConflictVote
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
			return fmt.Errorf("proto: ConflictVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConflictVote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictVote
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
				return ErrInvalidLengthConflictVote
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConflictVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteDeadline", wireType)
			}
			m.VoteDeadline = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VoteDeadline |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteStartBlock", wireType)
			}
			m.VoteStartBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VoteStartBlock |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteState", wireType)
			}
			m.VoteState = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VoteState |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictVote
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
				return ErrInvalidLengthConflictVote
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConflictVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictVote
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
				return ErrInvalidLengthConflictVote
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConflictVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApiUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthConflictVote
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthConflictVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestData = append(m.RequestData[:0], dAtA[iNdEx:postIndex]...)
			if m.RequestData == nil {
				m.RequestData = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestBlock", wireType)
			}
			m.RequestBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestBlock |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstProvider", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictVote
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
				return ErrInvalidLengthConflictVote
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConflictVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FirstProvider.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecondProvider", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictVote
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
				return ErrInvalidLengthConflictVote
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConflictVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SecondProvider.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotersHash", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictVote
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
				return ErrInvalidLengthConflictVote
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConflictVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.VotersHash == nil {
				m.VotersHash = make(map[string]Vote)
			}
			var mapkey string
			mapvalue := &Vote{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowConflictVote
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
							return ErrIntOverflowConflictVote
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
						return ErrInvalidLengthConflictVote
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthConflictVote
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowConflictVote
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthConflictVote
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthConflictVote
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Vote{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipConflictVote(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthConflictVote
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.VotersHash[mapkey] = *mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConflictVote(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConflictVote
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
func skipConflictVote(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConflictVote
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
					return 0, ErrIntOverflowConflictVote
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
					return 0, ErrIntOverflowConflictVote
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
				return 0, ErrInvalidLengthConflictVote
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConflictVote
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConflictVote
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConflictVote        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConflictVote          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConflictVote = fmt.Errorf("proto: unexpected end of group")
)
