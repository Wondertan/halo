// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: network.proto

package consensus

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReceiveResponse_Step int32

const (
	ReceiveResponse_UNKNOWN   ReceiveResponse_Step = 0
	ReceiveResponse_PROPOSE   ReceiveResponse_Step = 1
	ReceiveResponse_PREVOTE   ReceiveResponse_Step = 2
	ReceiveResponse_PRECOMMIT ReceiveResponse_Step = 3
	ReceiveResponse_COMMIT    ReceiveResponse_Step = 4
)

// Enum value maps for ReceiveResponse_Step.
var (
	ReceiveResponse_Step_name = map[int32]string{
		0: "UNKNOWN",
		1: "PROPOSE",
		2: "PREVOTE",
		3: "PRECOMMIT",
		4: "COMMIT",
	}
	ReceiveResponse_Step_value = map[string]int32{
		"UNKNOWN":   0,
		"PROPOSE":   1,
		"PREVOTE":   2,
		"PRECOMMIT": 3,
		"COMMIT":    4,
	}
)

func (x ReceiveResponse_Step) Enum() *ReceiveResponse_Step {
	p := new(ReceiveResponse_Step)
	*p = x
	return p
}

func (x ReceiveResponse_Step) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReceiveResponse_Step) Descriptor() protoreflect.EnumDescriptor {
	return file_network_proto_enumTypes[0].Descriptor()
}

func (ReceiveResponse_Step) Type() protoreflect.EnumType {
	return &file_network_proto_enumTypes[0]
}

func (x ReceiveResponse_Step) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReceiveResponse_Step.Descriptor instead.
func (ReceiveResponse_Step) EnumDescriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{1, 0}
}

type Vote_Type int32

const (
	Vote_UNKNOWN   Vote_Type = 0
	Vote_PREVOTE   Vote_Type = 1
	Vote_PRECOMMIT Vote_Type = 2
)

// Enum value maps for Vote_Type.
var (
	Vote_Type_name = map[int32]string{
		0: "UNKNOWN",
		1: "PREVOTE",
		2: "PRECOMMIT",
	}
	Vote_Type_value = map[string]int32{
		"UNKNOWN":   0,
		"PREVOTE":   1,
		"PRECOMMIT": 2,
	}
)

func (x Vote_Type) Enum() *Vote_Type {
	p := new(Vote_Type)
	*p = x
	return p
}

func (x Vote_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Vote_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_network_proto_enumTypes[1].Descriptor()
}

func (Vote_Type) Type() protoreflect.EnumType {
	return &file_network_proto_enumTypes[1]
}

func (x Vote_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Vote_Type.Descriptor instead.
func (Vote_Type) EnumDescriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{6, 0}
}

type ReceiveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg *Msg `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ReceiveRequest) Reset() {
	*x = ReceiveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveRequest) ProtoMessage() {}

func (x *ReceiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveRequest.ProtoReflect.Descriptor instead.
func (*ReceiveRequest) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{0}
}

func (x *ReceiveRequest) GetMsg() *Msg {
	if x != nil {
		return x.Msg
	}
	return nil
}

type ReceiveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height uint64               `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Round  uint32               `protobuf:"varint,2,opt,name=round,proto3" json:"round,omitempty"`
	Step   ReceiveResponse_Step `protobuf:"varint,3,opt,name=step,proto3,enum=consensus.ReceiveResponse_Step" json:"step,omitempty"`
}

func (x *ReceiveResponse) Reset() {
	*x = ReceiveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveResponse) ProtoMessage() {}

func (x *ReceiveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveResponse.ProtoReflect.Descriptor instead.
func (*ReceiveResponse) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{1}
}

func (x *ReceiveResponse) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *ReceiveResponse) GetRound() uint32 {
	if x != nil {
		return x.Round
	}
	return 0
}

func (x *ReceiveResponse) GetStep() ReceiveResponse_Step {
	if x != nil {
		return x.Step
	}
	return ReceiveResponse_UNKNOWN
}

type BroadcastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg *Msg `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *BroadcastRequest) Reset() {
	*x = BroadcastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastRequest) ProtoMessage() {}

func (x *BroadcastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastRequest.ProtoReflect.Descriptor instead.
func (*BroadcastRequest) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{2}
}

func (x *BroadcastRequest) GetMsg() *Msg {
	if x != nil {
		return x.Msg
	}
	return nil
}

type BroadcastResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BroadcastResponse) Reset() {
	*x = BroadcastResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastResponse) ProtoMessage() {}

func (x *BroadcastResponse) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastResponse.ProtoReflect.Descriptor instead.
func (*BroadcastResponse) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{3}
}

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Sum:
	//
	//	*Msg_Payload
	//	*Msg_Vote
	Sum isMsg_Sum `protobuf_oneof:"sum"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{4}
}

func (m *Msg) GetSum() isMsg_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (x *Msg) GetPayload() *Payload {
	if x, ok := x.GetSum().(*Msg_Payload); ok {
		return x.Payload
	}
	return nil
}

func (x *Msg) GetVote() *Vote {
	if x, ok := x.GetSum().(*Msg_Vote); ok {
		return x.Vote
	}
	return nil
}

type isMsg_Sum interface {
	isMsg_Sum()
}

type Msg_Payload struct {
	Payload *Payload `protobuf:"bytes,1,opt,name=payload,proto3,oneof"`
}

type Msg_Vote struct {
	Vote *Vote `protobuf:"bytes,2,opt,name=vote,proto3,oneof"`
}

func (*Msg_Payload) isMsg_Sum() {}

func (*Msg_Vote) isMsg_Sum() {}

// Payload is the value, v, proposed by the proposer in any given round
type Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Round  uint32 `protobuf:"varint,2,opt,name=round,proto3" json:"round,omitempty"`
	// this is the signature from the proposer of this payload
	Signature []byte   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	Data      [][]byte `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *Payload) Reset() {
	*x = Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payload) ProtoMessage() {}

func (x *Payload) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payload.ProtoReflect.Descriptor instead.
func (*Payload) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{5}
}

func (x *Payload) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Payload) GetRound() uint32 {
	if x != nil {
		return x.Round
	}
	return 0
}

func (x *Payload) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *Payload) GetData() [][]byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Vote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   Vote_Type `protobuf:"varint,1,opt,name=type,proto3,enum=consensus.Vote_Type" json:"type,omitempty"`
	Height uint64    `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Round  uint32    `protobuf:"varint,3,opt,name=round,proto3" json:"round,omitempty"`
	// a voter may refer to a payload from a previous round that the voter is
	// locked on. This is used to find out the payload_id that the signature
	// refers to.
	ReferenceRound uint32 `protobuf:"varint,4,opt,name=reference_round,json=referenceRound,proto3" json:"reference_round,omitempty"`
	// validators are ordered in a set by voting power. Only the index is necessary
	// to retrieve the validators public key
	ValidatorIndex uint32 `protobuf:"varint,5,opt,name=validator_index,json=validatorIndex,proto3" json:"validator_index,omitempty"`
	Signature      []byte `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Vote) Reset() {
	*x = Vote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vote) ProtoMessage() {}

func (x *Vote) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vote.ProtoReflect.Descriptor instead.
func (*Vote) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{6}
}

func (x *Vote) GetType() Vote_Type {
	if x != nil {
		return x.Type
	}
	return Vote_UNKNOWN
}

func (x *Vote) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Vote) GetRound() uint32 {
	if x != nil {
		return x.Round
	}
	return 0
}

func (x *Vote) GetReferenceRound() uint32 {
	if x != nil {
		return x.ReferenceRound
	}
	return 0
}

func (x *Vote) GetValidatorIndex() uint32 {
	if x != nil {
		return x.ValidatorIndex
	}
	return 0
}

func (x *Vote) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_network_proto protoreflect.FileDescriptor

var file_network_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x22, 0x32, 0x0a, 0x0e, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0xbe,
	0x01, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x12, 0x33, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x52,
	0x04, 0x73, 0x74, 0x65, 0x70, 0x22, 0x48, 0x0a, 0x04, 0x53, 0x74, 0x65, 0x70, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52,
	0x4f, 0x50, 0x4f, 0x53, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x45, 0x56, 0x4f,
	0x54, 0x45, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x52, 0x45, 0x43, 0x4f, 0x4d, 0x4d, 0x49,
	0x54, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x10, 0x04, 0x22,
	0x34, 0x0a, 0x10, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x4d, 0x73, 0x67,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x13, 0x0a, 0x11, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x63, 0x0a, 0x03, 0x4d, 0x73,
	0x67, 0x12, 0x2e, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x00, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x25, 0x0a, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x56, 0x6f, 0x74, 0x65,
	0x48, 0x00, 0x52, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x42, 0x05, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x22,
	0x69, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xff, 0x01, 0x0a, 0x04, 0x56,
	0x6f, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x56, 0x6f,
	0x74, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x6f, 0x75, 0x6e, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x2f, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x45, 0x56, 0x4f, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0d, 0x0a,
	0x09, 0x50, 0x52, 0x45, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x10, 0x02, 0x32, 0x4c, 0x0a, 0x08,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x50, 0x0a, 0x06, 0x53, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73,
	0x74, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x42, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1f, 0x5a, 0x1d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6d, 0x77, 0x61, 0x74,
	0x65, 0x72, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_network_proto_rawDescOnce sync.Once
	file_network_proto_rawDescData = file_network_proto_rawDesc
)

func file_network_proto_rawDescGZIP() []byte {
	file_network_proto_rawDescOnce.Do(func() {
		file_network_proto_rawDescData = protoimpl.X.CompressGZIP(file_network_proto_rawDescData)
	})
	return file_network_proto_rawDescData
}

var file_network_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_network_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_network_proto_goTypes = []interface{}{
	(ReceiveResponse_Step)(0), // 0: consensus.ReceiveResponse.Step
	(Vote_Type)(0),            // 1: consensus.Vote.Type
	(*ReceiveRequest)(nil),    // 2: consensus.ReceiveRequest
	(*ReceiveResponse)(nil),   // 3: consensus.ReceiveResponse
	(*BroadcastRequest)(nil),  // 4: consensus.BroadcastRequest
	(*BroadcastResponse)(nil), // 5: consensus.BroadcastResponse
	(*Msg)(nil),               // 6: consensus.Msg
	(*Payload)(nil),           // 7: consensus.Payload
	(*Vote)(nil),              // 8: consensus.Vote
}
var file_network_proto_depIdxs = []int32{
	6, // 0: consensus.ReceiveRequest.msg:type_name -> consensus.Msg
	0, // 1: consensus.ReceiveResponse.step:type_name -> consensus.ReceiveResponse.Step
	6, // 2: consensus.BroadcastRequest.msg:type_name -> consensus.Msg
	7, // 3: consensus.Msg.payload:type_name -> consensus.Payload
	8, // 4: consensus.Msg.vote:type_name -> consensus.Vote
	1, // 5: consensus.Vote.type:type_name -> consensus.Vote.Type
	2, // 6: consensus.Receiver.Receive:input_type -> consensus.ReceiveRequest
	4, // 7: consensus.Sender.Broadcast:input_type -> consensus.BroadcastRequest
	3, // 8: consensus.Receiver.Receive:output_type -> consensus.ReceiveResponse
	5, // 9: consensus.Sender.Broadcast:output_type -> consensus.BroadcastResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_network_proto_init() }
func file_network_proto_init() {
	if File_network_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_network_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_network_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_network_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_network_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_network_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_network_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_network_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vote); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_network_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Msg_Payload)(nil),
		(*Msg_Vote)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_network_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_network_proto_goTypes,
		DependencyIndexes: file_network_proto_depIdxs,
		EnumInfos:         file_network_proto_enumTypes,
		MessageInfos:      file_network_proto_msgTypes,
	}.Build()
	File_network_proto = out.File
	file_network_proto_rawDesc = nil
	file_network_proto_goTypes = nil
	file_network_proto_depIdxs = nil
}
