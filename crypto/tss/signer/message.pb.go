// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/getamis/alice/crypto/tss/signer/message.proto

package signer

import (
	fmt "fmt"
	commitment "github.com/getamis/alice/crypto/commitment"
	zkproof "github.com/getamis/alice/crypto/zkproof"
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

type Type int32

const (
	Type_Pubkey       Type = 0
	Type_EncK         Type = 1
	Type_Mta          Type = 2
	Type_Delta        Type = 3
	Type_ProofAi      Type = 4
	Type_CommitViAi   Type = 5
	Type_DecommitViAi Type = 6
	Type_CommitUiTi   Type = 7
	Type_DecommitUiTi Type = 8
	Type_Si           Type = 9
)

var Type_name = map[int32]string{
	0: "Pubkey",
	1: "EncK",
	2: "Mta",
	3: "Delta",
	4: "ProofAi",
	5: "CommitViAi",
	6: "DecommitViAi",
	7: "CommitUiTi",
	8: "DecommitUiTi",
	9: "Si",
}

var Type_value = map[string]int32{
	"Pubkey":       0,
	"EncK":         1,
	"Mta":          2,
	"Delta":        3,
	"ProofAi":      4,
	"CommitViAi":   5,
	"DecommitViAi": 6,
	"CommitUiTi":   7,
	"DecommitUiTi": 8,
	"Si":           9,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ad801314df39a0f8, []int{0}
}

type Message struct {
	Type Type   `protobuf:"varint,1,opt,name=type,proto3,enum=signer.Type" json:"type,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Body:
	//	*Message_Pubkey
	//	*Message_EncK
	//	*Message_Mta
	//	*Message_Delta
	//	*Message_ProofAi
	//	*Message_CommitViAi
	//	*Message_DecommitViAi
	//	*Message_CommitUiTi
	//	*Message_DecommitUiTi
	//	*Message_Si
	Body                 isMessage_Body `protobuf_oneof:"body"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad801314df39a0f8, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_Pubkey
}

func (m *Message) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type isMessage_Body interface {
	isMessage_Body()
}

type Message_Pubkey struct {
	Pubkey *BodyPublicKey `protobuf:"bytes,3,opt,name=pubkey,proto3,oneof"`
}

type Message_EncK struct {
	EncK *BodyEncK `protobuf:"bytes,4,opt,name=encK,proto3,oneof"`
}

type Message_Mta struct {
	Mta *BodyMta `protobuf:"bytes,5,opt,name=mta,proto3,oneof"`
}

type Message_Delta struct {
	Delta *BodyDelta `protobuf:"bytes,6,opt,name=delta,proto3,oneof"`
}

type Message_ProofAi struct {
	ProofAi *BodyProofAi `protobuf:"bytes,7,opt,name=proofAi,proto3,oneof"`
}

type Message_CommitViAi struct {
	CommitViAi *BodyCommitViAi `protobuf:"bytes,8,opt,name=commitViAi,proto3,oneof"`
}

type Message_DecommitViAi struct {
	DecommitViAi *BodyDecommitViAi `protobuf:"bytes,9,opt,name=decommitViAi,proto3,oneof"`
}

type Message_CommitUiTi struct {
	CommitUiTi *BodyCommitUiTi `protobuf:"bytes,10,opt,name=commitUiTi,proto3,oneof"`
}

type Message_DecommitUiTi struct {
	DecommitUiTi *BodyDecommitUiTi `protobuf:"bytes,11,opt,name=decommitUiTi,proto3,oneof"`
}

type Message_Si struct {
	Si *BodySi `protobuf:"bytes,12,opt,name=si,proto3,oneof"`
}

func (*Message_Pubkey) isMessage_Body() {}

func (*Message_EncK) isMessage_Body() {}

func (*Message_Mta) isMessage_Body() {}

func (*Message_Delta) isMessage_Body() {}

func (*Message_ProofAi) isMessage_Body() {}

func (*Message_CommitViAi) isMessage_Body() {}

func (*Message_DecommitViAi) isMessage_Body() {}

func (*Message_CommitUiTi) isMessage_Body() {}

func (*Message_DecommitUiTi) isMessage_Body() {}

func (*Message_Si) isMessage_Body() {}

func (m *Message) GetBody() isMessage_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Message) GetPubkey() *BodyPublicKey {
	if x, ok := m.GetBody().(*Message_Pubkey); ok {
		return x.Pubkey
	}
	return nil
}

func (m *Message) GetEncK() *BodyEncK {
	if x, ok := m.GetBody().(*Message_EncK); ok {
		return x.EncK
	}
	return nil
}

func (m *Message) GetMta() *BodyMta {
	if x, ok := m.GetBody().(*Message_Mta); ok {
		return x.Mta
	}
	return nil
}

func (m *Message) GetDelta() *BodyDelta {
	if x, ok := m.GetBody().(*Message_Delta); ok {
		return x.Delta
	}
	return nil
}

func (m *Message) GetProofAi() *BodyProofAi {
	if x, ok := m.GetBody().(*Message_ProofAi); ok {
		return x.ProofAi
	}
	return nil
}

func (m *Message) GetCommitViAi() *BodyCommitViAi {
	if x, ok := m.GetBody().(*Message_CommitViAi); ok {
		return x.CommitViAi
	}
	return nil
}

func (m *Message) GetDecommitViAi() *BodyDecommitViAi {
	if x, ok := m.GetBody().(*Message_DecommitViAi); ok {
		return x.DecommitViAi
	}
	return nil
}

func (m *Message) GetCommitUiTi() *BodyCommitUiTi {
	if x, ok := m.GetBody().(*Message_CommitUiTi); ok {
		return x.CommitUiTi
	}
	return nil
}

func (m *Message) GetDecommitUiTi() *BodyDecommitUiTi {
	if x, ok := m.GetBody().(*Message_DecommitUiTi); ok {
		return x.DecommitUiTi
	}
	return nil
}

func (m *Message) GetSi() *BodySi {
	if x, ok := m.GetBody().(*Message_Si); ok {
		return x.Si
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Message) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Message_Pubkey)(nil),
		(*Message_EncK)(nil),
		(*Message_Mta)(nil),
		(*Message_Delta)(nil),
		(*Message_ProofAi)(nil),
		(*Message_CommitViAi)(nil),
		(*Message_DecommitViAi)(nil),
		(*Message_CommitUiTi)(nil),
		(*Message_DecommitUiTi)(nil),
		(*Message_Si)(nil),
	}
}

type BodyPublicKey struct {
	Pubkey               []byte                            `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	AgCommitment         *commitment.HashCommitmentMessage `protobuf:"bytes,3,opt,name=agCommitment,proto3" json:"agCommitment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *BodyPublicKey) Reset()         { *m = BodyPublicKey{} }
func (m *BodyPublicKey) String() string { return proto.CompactTextString(m) }
func (*BodyPublicKey) ProtoMessage()    {}
func (*BodyPublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad801314df39a0f8, []int{1}
}

func (m *BodyPublicKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BodyPublicKey.Unmarshal(m, b)
}
func (m *BodyPublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BodyPublicKey.Marshal(b, m, deterministic)
}
func (m *BodyPublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BodyPublicKey.Merge(m, src)
}
func (m *BodyPublicKey) XXX_Size() int {
	return xxx_messageInfo_BodyPublicKey.Size(m)
}
func (m *BodyPublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_BodyPublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_BodyPublicKey proto.InternalMessageInfo

func (m *BodyPublicKey) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *BodyPublicKey) GetAgCommitment() *commitment.HashCommitmentMessage {
	if m != nil {
		return m.AgCommitment
	}
	return nil
}

type BodyEncK struct {
	Enck                 []byte   `protobuf:"bytes,2,opt,name=enck,proto3" json:"enck,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BodyEncK) Reset()         { *m = BodyEncK{} }
func (m *BodyEncK) String() string { return proto.CompactTextString(m) }
func (*BodyEncK) ProtoMessage()    {}
func (*BodyEncK) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad801314df39a0f8, []int{2}
}

func (m *BodyEncK) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BodyEncK.Unmarshal(m, b)
}
func (m *BodyEncK) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BodyEncK.Marshal(b, m, deterministic)
}
func (m *BodyEncK) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BodyEncK.Merge(m, src)
}
func (m *BodyEncK) XXX_Size() int {
	return xxx_messageInfo_BodyEncK.Size(m)
}
func (m *BodyEncK) XXX_DiscardUnknown() {
	xxx_messageInfo_BodyEncK.DiscardUnknown(m)
}

var xxx_messageInfo_BodyEncK proto.InternalMessageInfo

func (m *BodyEncK) GetEnck() []byte {
	if m != nil {
		return m.Enck
	}
	return nil
}

type BodyMta struct {
	EncAiAlpha           []byte   `protobuf:"bytes,1,opt,name=encAiAlpha,proto3" json:"encAiAlpha,omitempty"`
	EncWiAlpha           []byte   `protobuf:"bytes,2,opt,name=encWiAlpha,proto3" json:"encWiAlpha,omitempty"`
	WiProof              []byte   `protobuf:"bytes,3,opt,name=wiProof,proto3" json:"wiProof,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BodyMta) Reset()         { *m = BodyMta{} }
func (m *BodyMta) String() string { return proto.CompactTextString(m) }
func (*BodyMta) ProtoMessage()    {}
func (*BodyMta) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad801314df39a0f8, []int{3}
}

func (m *BodyMta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BodyMta.Unmarshal(m, b)
}
func (m *BodyMta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BodyMta.Marshal(b, m, deterministic)
}
func (m *BodyMta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BodyMta.Merge(m, src)
}
func (m *BodyMta) XXX_Size() int {
	return xxx_messageInfo_BodyMta.Size(m)
}
func (m *BodyMta) XXX_DiscardUnknown() {
	xxx_messageInfo_BodyMta.DiscardUnknown(m)
}

var xxx_messageInfo_BodyMta proto.InternalMessageInfo

func (m *BodyMta) GetEncAiAlpha() []byte {
	if m != nil {
		return m.EncAiAlpha
	}
	return nil
}

func (m *BodyMta) GetEncWiAlpha() []byte {
	if m != nil {
		return m.EncWiAlpha
	}
	return nil
}

func (m *BodyMta) GetWiProof() []byte {
	if m != nil {
		return m.WiProof
	}
	return nil
}

type BodyDelta struct {
	Delta                []byte   `protobuf:"bytes,1,opt,name=delta,proto3" json:"delta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BodyDelta) Reset()         { *m = BodyDelta{} }
func (m *BodyDelta) String() string { return proto.CompactTextString(m) }
func (*BodyDelta) ProtoMessage()    {}
func (*BodyDelta) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad801314df39a0f8, []int{4}
}

func (m *BodyDelta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BodyDelta.Unmarshal(m, b)
}
func (m *BodyDelta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BodyDelta.Marshal(b, m, deterministic)
}
func (m *BodyDelta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BodyDelta.Merge(m, src)
}
func (m *BodyDelta) XXX_Size() int {
	return xxx_messageInfo_BodyDelta.Size(m)
}
func (m *BodyDelta) XXX_DiscardUnknown() {
	xxx_messageInfo_BodyDelta.DiscardUnknown(m)
}

var xxx_messageInfo_BodyDelta proto.InternalMessageInfo

func (m *BodyDelta) GetDelta() []byte {
	if m != nil {
		return m.Delta
	}
	return nil
}

type BodyProofAi struct {
	AgDecommitment       *commitment.HashDecommitmentMessage `protobuf:"bytes,1,opt,name=agDecommitment,proto3" json:"agDecommitment,omitempty"`
	AiProof              *zkproof.SchnorrProofMessage        `protobuf:"bytes,2,opt,name=aiProof,proto3" json:"aiProof,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *BodyProofAi) Reset()         { *m = BodyProofAi{} }
func (m *BodyProofAi) String() string { return proto.CompactTextString(m) }
func (*BodyProofAi) ProtoMessage()    {}
func (*BodyProofAi) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad801314df39a0f8, []int{5}
}

func (m *BodyProofAi) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BodyProofAi.Unmarshal(m, b)
}
func (m *BodyProofAi) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BodyProofAi.Marshal(b, m, deterministic)
}
func (m *BodyProofAi) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BodyProofAi.Merge(m, src)
}
func (m *BodyProofAi) XXX_Size() int {
	return xxx_messageInfo_BodyProofAi.Size(m)
}
func (m *BodyProofAi) XXX_DiscardUnknown() {
	xxx_messageInfo_BodyProofAi.DiscardUnknown(m)
}

var xxx_messageInfo_BodyProofAi proto.InternalMessageInfo

func (m *BodyProofAi) GetAgDecommitment() *commitment.HashDecommitmentMessage {
	if m != nil {
		return m.AgDecommitment
	}
	return nil
}

func (m *BodyProofAi) GetAiProof() *zkproof.SchnorrProofMessage {
	if m != nil {
		return m.AiProof
	}
	return nil
}

type BodyCommitViAi struct {
	ViCommitment         *commitment.HashCommitmentMessage `protobuf:"bytes,1,opt,name=viCommitment,proto3" json:"viCommitment,omitempty"`
	AiCommitment         *commitment.HashCommitmentMessage `protobuf:"bytes,2,opt,name=aiCommitment,proto3" json:"aiCommitment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *BodyCommitViAi) Reset()         { *m = BodyCommitViAi{} }
func (m *BodyCommitViAi) String() string { return proto.CompactTextString(m) }
func (*BodyCommitViAi) ProtoMessage()    {}
func (*BodyCommitViAi) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad801314df39a0f8, []int{6}
}

func (m *BodyCommitViAi) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BodyCommitViAi.Unmarshal(m, b)
}
func (m *BodyCommitViAi) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BodyCommitViAi.Marshal(b, m, deterministic)
}
func (m *BodyCommitViAi) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BodyCommitViAi.Merge(m, src)
}
func (m *BodyCommitViAi) XXX_Size() int {
	return xxx_messageInfo_BodyCommitViAi.Size(m)
}
func (m *BodyCommitViAi) XXX_DiscardUnknown() {
	xxx_messageInfo_BodyCommitViAi.DiscardUnknown(m)
}

var xxx_messageInfo_BodyCommitViAi proto.InternalMessageInfo

func (m *BodyCommitViAi) GetViCommitment() *commitment.HashCommitmentMessage {
	if m != nil {
		return m.ViCommitment
	}
	return nil
}

func (m *BodyCommitViAi) GetAiCommitment() *commitment.HashCommitmentMessage {
	if m != nil {
		return m.AiCommitment
	}
	return nil
}

type BodyDecommitViAi struct {
	ViDecommitment       *commitment.HashDecommitmentMessage `protobuf:"bytes,1,opt,name=viDecommitment,proto3" json:"viDecommitment,omitempty"`
	AiDecommitment       *commitment.HashDecommitmentMessage `protobuf:"bytes,2,opt,name=aiDecommitment,proto3" json:"aiDecommitment,omitempty"`
	RhoIProof            *zkproof.SchnorrProofMessage        `protobuf:"bytes,3,opt,name=rhoIProof,proto3" json:"rhoIProof,omitempty"`
	LiProof              *zkproof.SchnorrProofMessage        `protobuf:"bytes,4,opt,name=liProof,proto3" json:"liProof,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *BodyDecommitViAi) Reset()         { *m = BodyDecommitViAi{} }
func (m *BodyDecommitViAi) String() string { return proto.CompactTextString(m) }
func (*BodyDecommitViAi) ProtoMessage()    {}
func (*BodyDecommitViAi) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad801314df39a0f8, []int{7}
}

func (m *BodyDecommitViAi) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BodyDecommitViAi.Unmarshal(m, b)
}
func (m *BodyDecommitViAi) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BodyDecommitViAi.Marshal(b, m, deterministic)
}
func (m *BodyDecommitViAi) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BodyDecommitViAi.Merge(m, src)
}
func (m *BodyDecommitViAi) XXX_Size() int {
	return xxx_messageInfo_BodyDecommitViAi.Size(m)
}
func (m *BodyDecommitViAi) XXX_DiscardUnknown() {
	xxx_messageInfo_BodyDecommitViAi.DiscardUnknown(m)
}

var xxx_messageInfo_BodyDecommitViAi proto.InternalMessageInfo

func (m *BodyDecommitViAi) GetViDecommitment() *commitment.HashDecommitmentMessage {
	if m != nil {
		return m.ViDecommitment
	}
	return nil
}

func (m *BodyDecommitViAi) GetAiDecommitment() *commitment.HashDecommitmentMessage {
	if m != nil {
		return m.AiDecommitment
	}
	return nil
}

func (m *BodyDecommitViAi) GetRhoIProof() *zkproof.SchnorrProofMessage {
	if m != nil {
		return m.RhoIProof
	}
	return nil
}

func (m *BodyDecommitViAi) GetLiProof() *zkproof.SchnorrProofMessage {
	if m != nil {
		return m.LiProof
	}
	return nil
}

type BodyCommitUiTi struct {
	UiCommitment         *commitment.HashCommitmentMessage `protobuf:"bytes,1,opt,name=uiCommitment,proto3" json:"uiCommitment,omitempty"`
	TiCommitment         *commitment.HashCommitmentMessage `protobuf:"bytes,2,opt,name=tiCommitment,proto3" json:"tiCommitment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *BodyCommitUiTi) Reset()         { *m = BodyCommitUiTi{} }
func (m *BodyCommitUiTi) String() string { return proto.CompactTextString(m) }
func (*BodyCommitUiTi) ProtoMessage()    {}
func (*BodyCommitUiTi) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad801314df39a0f8, []int{8}
}

func (m *BodyCommitUiTi) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BodyCommitUiTi.Unmarshal(m, b)
}
func (m *BodyCommitUiTi) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BodyCommitUiTi.Marshal(b, m, deterministic)
}
func (m *BodyCommitUiTi) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BodyCommitUiTi.Merge(m, src)
}
func (m *BodyCommitUiTi) XXX_Size() int {
	return xxx_messageInfo_BodyCommitUiTi.Size(m)
}
func (m *BodyCommitUiTi) XXX_DiscardUnknown() {
	xxx_messageInfo_BodyCommitUiTi.DiscardUnknown(m)
}

var xxx_messageInfo_BodyCommitUiTi proto.InternalMessageInfo

func (m *BodyCommitUiTi) GetUiCommitment() *commitment.HashCommitmentMessage {
	if m != nil {
		return m.UiCommitment
	}
	return nil
}

func (m *BodyCommitUiTi) GetTiCommitment() *commitment.HashCommitmentMessage {
	if m != nil {
		return m.TiCommitment
	}
	return nil
}

type BodyDecommitUiTi struct {
	UiDecommitment       *commitment.HashDecommitmentMessage `protobuf:"bytes,1,opt,name=uiDecommitment,proto3" json:"uiDecommitment,omitempty"`
	TiDecommitment       *commitment.HashDecommitmentMessage `protobuf:"bytes,2,opt,name=tiDecommitment,proto3" json:"tiDecommitment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *BodyDecommitUiTi) Reset()         { *m = BodyDecommitUiTi{} }
func (m *BodyDecommitUiTi) String() string { return proto.CompactTextString(m) }
func (*BodyDecommitUiTi) ProtoMessage()    {}
func (*BodyDecommitUiTi) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad801314df39a0f8, []int{9}
}

func (m *BodyDecommitUiTi) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BodyDecommitUiTi.Unmarshal(m, b)
}
func (m *BodyDecommitUiTi) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BodyDecommitUiTi.Marshal(b, m, deterministic)
}
func (m *BodyDecommitUiTi) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BodyDecommitUiTi.Merge(m, src)
}
func (m *BodyDecommitUiTi) XXX_Size() int {
	return xxx_messageInfo_BodyDecommitUiTi.Size(m)
}
func (m *BodyDecommitUiTi) XXX_DiscardUnknown() {
	xxx_messageInfo_BodyDecommitUiTi.DiscardUnknown(m)
}

var xxx_messageInfo_BodyDecommitUiTi proto.InternalMessageInfo

func (m *BodyDecommitUiTi) GetUiDecommitment() *commitment.HashDecommitmentMessage {
	if m != nil {
		return m.UiDecommitment
	}
	return nil
}

func (m *BodyDecommitUiTi) GetTiDecommitment() *commitment.HashDecommitmentMessage {
	if m != nil {
		return m.TiDecommitment
	}
	return nil
}

type BodySi struct {
	Si                   []byte   `protobuf:"bytes,1,opt,name=si,proto3" json:"si,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BodySi) Reset()         { *m = BodySi{} }
func (m *BodySi) String() string { return proto.CompactTextString(m) }
func (*BodySi) ProtoMessage()    {}
func (*BodySi) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad801314df39a0f8, []int{10}
}

func (m *BodySi) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BodySi.Unmarshal(m, b)
}
func (m *BodySi) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BodySi.Marshal(b, m, deterministic)
}
func (m *BodySi) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BodySi.Merge(m, src)
}
func (m *BodySi) XXX_Size() int {
	return xxx_messageInfo_BodySi.Size(m)
}
func (m *BodySi) XXX_DiscardUnknown() {
	xxx_messageInfo_BodySi.DiscardUnknown(m)
}

var xxx_messageInfo_BodySi proto.InternalMessageInfo

func (m *BodySi) GetSi() []byte {
	if m != nil {
		return m.Si
	}
	return nil
}

func init() {
	proto.RegisterEnum("signer.Type", Type_name, Type_value)
	proto.RegisterType((*Message)(nil), "signer.Message")
	proto.RegisterType((*BodyPublicKey)(nil), "signer.BodyPublicKey")
	proto.RegisterType((*BodyEncK)(nil), "signer.BodyEncK")
	proto.RegisterType((*BodyMta)(nil), "signer.BodyMta")
	proto.RegisterType((*BodyDelta)(nil), "signer.BodyDelta")
	proto.RegisterType((*BodyProofAi)(nil), "signer.BodyProofAi")
	proto.RegisterType((*BodyCommitViAi)(nil), "signer.BodyCommitViAi")
	proto.RegisterType((*BodyDecommitViAi)(nil), "signer.BodyDecommitViAi")
	proto.RegisterType((*BodyCommitUiTi)(nil), "signer.BodyCommitUiTi")
	proto.RegisterType((*BodyDecommitUiTi)(nil), "signer.BodyDecommitUiTi")
	proto.RegisterType((*BodySi)(nil), "signer.BodySi")
}

func init() {
	proto.RegisterFile("github.com/getamis/alice/crypto/tss/signer/message.proto", fileDescriptor_ad801314df39a0f8)
}

var fileDescriptor_ad801314df39a0f8 = []byte{
	// 747 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x6a, 0xdb, 0x4a,
	0x10, 0xb6, 0x64, 0x59, 0xb2, 0xc7, 0x3a, 0x3e, 0x3a, 0x7b, 0xce, 0x09, 0x22, 0x94, 0xe0, 0x28,
	0x50, 0xd2, 0x5e, 0x48, 0x90, 0xd2, 0x12, 0x7a, 0x51, 0x70, 0xd2, 0x80, 0x8b, 0x09, 0x04, 0x25,
	0x6d, 0xaf, 0x65, 0x79, 0x6b, 0x2f, 0xfe, 0x91, 0xb1, 0xd6, 0x29, 0xea, 0x1b, 0xf4, 0xb6, 0x0f,
	0x50, 0xfa, 0x00, 0xbd, 0xeb, 0x0b, 0x16, 0x8d, 0x56, 0xd6, 0xae, 0xdb, 0x90, 0xd4, 0x97, 0xda,
	0xf9, 0xbe, 0x99, 0xd1, 0x7c, 0xdf, 0xec, 0xc2, 0xe9, 0x98, 0xf1, 0xc9, 0x7a, 0xe8, 0xc7, 0xc9,
	0x3c, 0x18, 0x53, 0x1e, 0xcd, 0x59, 0x1a, 0x44, 0x33, 0x16, 0xd3, 0x20, 0x5e, 0x65, 0x4b, 0x9e,
	0x04, 0x3c, 0x4d, 0x83, 0x94, 0x8d, 0x17, 0x74, 0x15, 0xcc, 0x69, 0x9a, 0x46, 0x63, 0xea, 0x2f,
	0x57, 0x09, 0x4f, 0x88, 0x59, 0x9c, 0xee, 0xdf, 0x9b, 0x21, 0x4e, 0xe6, 0x73, 0xc6, 0xe7, 0x74,
	0xc1, 0xd5, 0x0c, 0xfb, 0xcf, 0xef, 0x63, 0x7e, 0x9a, 0x2e, 0x57, 0x49, 0xf2, 0x41, 0xa5, 0x79,
	0x3f, 0x0c, 0xb0, 0x2e, 0x8b, 0x13, 0xd2, 0x05, 0x83, 0x67, 0x4b, 0xea, 0x6a, 0x5d, 0xed, 0xb8,
	0x73, 0x62, 0xfb, 0x45, 0x4f, 0xfe, 0x4d, 0xb6, 0xa4, 0x21, 0x46, 0x48, 0x07, 0x74, 0x36, 0x72,
	0xf5, 0xae, 0x76, 0xdc, 0x0a, 0x75, 0x36, 0x22, 0x01, 0x98, 0xcb, 0xf5, 0x70, 0x4a, 0x33, 0xb7,
	0xde, 0xd5, 0x8e, 0xdb, 0x27, 0xff, 0x97, 0x9c, 0xb3, 0x64, 0x94, 0x5d, 0xad, 0x87, 0x33, 0x16,
	0x0f, 0x68, 0xd6, 0xaf, 0x85, 0x02, 0x46, 0x1e, 0x83, 0x41, 0x17, 0xf1, 0xc0, 0x35, 0x10, 0xee,
	0xc8, 0xf0, 0x8b, 0x45, 0x3c, 0xe8, 0xd7, 0x42, 0x8c, 0x93, 0x23, 0xa8, 0xcf, 0x79, 0xe4, 0x36,
	0x10, 0xf6, 0xb7, 0x0c, 0xbb, 0xe4, 0x51, 0xbf, 0x16, 0xe6, 0x51, 0xf2, 0x04, 0x1a, 0x23, 0x3a,
	0xe3, 0x91, 0x6b, 0x22, 0xec, 0x1f, 0x19, 0xf6, 0x3a, 0x0f, 0xf4, 0x6b, 0x61, 0x81, 0x20, 0x01,
	0x58, 0xf8, 0xf7, 0x3d, 0xe6, 0x5a, 0x08, 0xfe, 0x57, 0xe9, 0xb4, 0x08, 0xf5, 0x6b, 0x61, 0x89,
	0x22, 0xa7, 0x00, 0xc5, 0xa8, 0xdf, 0xb1, 0x1e, 0x73, 0x9b, 0xc8, 0xd9, 0x93, 0x39, 0xe7, 0x9b,
	0x68, 0xbf, 0x16, 0x4a, 0x58, 0xf2, 0x0a, 0xec, 0x11, 0x95, 0xb8, 0x2d, 0xe4, 0xba, 0x6a, 0x73,
	0xb1, 0xcc, 0x56, 0xf0, 0x55, 0xe5, 0xb7, 0xec, 0x86, 0xb9, 0x70, 0x57, 0xe5, 0x3c, 0x5a, 0x55,
	0xce, 0xbf, 0xe4, 0xca, 0xc8, 0x6d, 0xdf, 0x5d, 0x59, 0xb0, 0x15, 0x3c, 0xe9, 0x82, 0x9e, 0x32,
	0xd7, 0x46, 0x56, 0x47, 0x66, 0x5d, 0xe7, 0x58, 0x3d, 0x65, 0x67, 0x26, 0x18, 0xc3, 0x64, 0x94,
	0x79, 0x0b, 0xf8, 0x4b, 0x51, 0x98, 0xec, 0x6d, 0x8c, 0x90, 0x9b, 0xc7, 0xde, 0xe8, 0x7d, 0x01,
	0x76, 0x34, 0x3e, 0xdf, 0x78, 0x56, 0xd8, 0xe4, 0xd0, 0xaf, 0x6c, 0xec, 0xf7, 0xa3, 0x74, 0x52,
	0x21, 0x84, 0x17, 0x43, 0x85, 0xe6, 0x1d, 0x40, 0xb3, 0xb4, 0x08, 0x21, 0x68, 0xa1, 0x29, 0xba,
	0xd0, 0x46, 0xbb, 0x4c, 0xbd, 0x18, 0x2c, 0xe1, 0x0d, 0x72, 0x00, 0x40, 0x17, 0x71, 0x8f, 0xf5,
	0x66, 0xcb, 0x49, 0x24, 0xba, 0x91, 0x4e, 0x44, 0xfc, 0xbd, 0x88, 0xeb, 0x9b, 0xb8, 0x38, 0x21,
	0x2e, 0x58, 0x1f, 0x19, 0x1a, 0x02, 0x9b, 0xb5, 0xc3, 0xf2, 0xd3, 0x3b, 0x84, 0xd6, 0xc6, 0x59,
	0xe4, 0xbf, 0xd2, 0x7b, 0x45, 0x85, 0xe2, 0xc3, 0xfb, 0xa2, 0x41, 0x5b, 0x32, 0x14, 0x19, 0x40,
	0x27, 0x1a, 0x97, 0x33, 0xc7, 0x01, 0x68, 0x38, 0x80, 0xa3, 0xed, 0x01, 0xc8, 0x98, 0x72, 0x04,
	0x5b, 0x54, 0xf2, 0x02, 0xac, 0x48, 0x74, 0xa6, 0x63, 0x96, 0x47, 0xbe, 0xd8, 0x69, 0xff, 0x3a,
	0x9e, 0x2c, 0x92, 0xd5, 0x0a, 0x83, 0x25, 0xbd, 0x04, 0x7b, 0x5f, 0x35, 0xe8, 0xa8, 0x8e, 0xcd,
	0x65, 0xb9, 0x65, 0xe7, 0xdb, 0x5d, 0x3d, 0x44, 0x16, 0x99, 0x86, 0xea, 0xca, 0x69, 0xf4, 0x87,
	0xab, 0x2b, 0xd1, 0xbc, 0x6f, 0x3a, 0x38, 0xdb, 0x6b, 0x91, 0x8f, 0xee, 0x96, 0xed, 0x3c, 0x3a,
	0x95, 0x8a, 0x3a, 0xa8, 0xc9, 0xf4, 0x3f, 0xd1, 0x41, 0x4d, 0xf6, 0x12, 0x5a, 0xab, 0x49, 0xf2,
	0xa6, 0xf2, 0xc8, 0x7d, 0x4a, 0x54, 0xf0, 0x5c, 0xc3, 0x99, 0xd0, 0xd0, 0x78, 0x88, 0x86, 0xb3,
	0xdf, 0x6a, 0x88, 0xdb, 0x7a, 0x01, 0xf6, 0x7a, 0x37, 0x0d, 0xd7, 0x5b, 0x1a, 0xf2, 0xdd, 0x34,
	0x94, 0x69, 0xde, 0x77, 0x4d, 0xd5, 0x10, 0x5b, 0x1c, 0x40, 0x67, 0xbd, 0xbb, 0x86, 0xeb, 0x5f,
	0x34, 0xe4, 0xbb, 0x6b, 0xa8, 0x52, 0x3d, 0x17, 0xcc, 0xe2, 0x62, 0xcb, 0x9f, 0xb4, 0x94, 0x89,
	0x2d, 0xd6, 0x53, 0xf6, 0xf4, 0xb3, 0x06, 0x46, 0xfe, 0xe2, 0x11, 0x00, 0xf3, 0x0a, 0x2f, 0x31,
	0xa7, 0x46, 0x9a, 0x60, 0xe4, 0x77, 0x8f, 0xa3, 0x11, 0x0b, 0xea, 0x97, 0x3c, 0x72, 0x74, 0xd2,
	0x82, 0x06, 0xde, 0x04, 0x4e, 0x9d, 0xb4, 0xc1, 0x12, 0x0b, 0xef, 0x18, 0xa4, 0x03, 0x50, 0x2d,
	0x9a, 0xd3, 0x20, 0x0e, 0xd8, 0xb2, 0xaf, 0x1d, 0xb3, 0x42, 0xe4, 0x33, 0x72, 0x2c, 0x19, 0x81,
	0x27, 0x4d, 0x62, 0x82, 0x7e, 0xcd, 0x9c, 0xd6, 0xd0, 0xc4, 0x37, 0xfa, 0xd9, 0xcf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x65, 0x39, 0x06, 0x58, 0x58, 0x08, 0x00, 0x00,
}