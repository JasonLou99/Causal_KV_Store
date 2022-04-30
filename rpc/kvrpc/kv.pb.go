// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        v3.20.0
// source: kv.proto

package kvproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PutAppendArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	Op    string `protobuf:"bytes,3,opt,name=Op,proto3" json:"Op,omitempty"`
	Id    int64  `protobuf:"varint,4,opt,name=Id,proto3" json:"Id,omitempty"`
	Seq   int64  `protobuf:"varint,5,opt,name=Seq,proto3" json:"Seq,omitempty"`
}

func (x *PutAppendArgs) Reset() {
	*x = PutAppendArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutAppendArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutAppendArgs) ProtoMessage() {}

func (x *PutAppendArgs) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutAppendArgs.ProtoReflect.Descriptor instead.
func (*PutAppendArgs) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{0}
}

func (x *PutAppendArgs) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PutAppendArgs) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *PutAppendArgs) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *PutAppendArgs) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PutAppendArgs) GetSeq() int64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

type PutAppendReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsLeader bool `protobuf:"varint,1,opt,name=IsLeader,proto3" json:"IsLeader,omitempty"`
	Success  bool `protobuf:"varint,2,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *PutAppendReply) Reset() {
	*x = PutAppendReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutAppendReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutAppendReply) ProtoMessage() {}

func (x *PutAppendReply) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutAppendReply.ProtoReflect.Descriptor instead.
func (*PutAppendReply) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{1}
}

func (x *PutAppendReply) GetIsLeader() bool {
	if x != nil {
		return x.IsLeader
	}
	return false
}

func (x *PutAppendReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type PutAppendInCausalArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key         string           `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Value       string           `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	Op          string           `protobuf:"bytes,3,opt,name=Op,proto3" json:"Op,omitempty"`
	Id          int64            `protobuf:"varint,4,opt,name=Id,proto3" json:"Id,omitempty"`
	Seq         int64            `protobuf:"varint,5,opt,name=Seq,proto3" json:"Seq,omitempty"`
	VectorClock map[string]int32 `protobuf:"bytes,6,rep,name=VectorClock,proto3" json:"VectorClock,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *PutAppendInCausalArgs) Reset() {
	*x = PutAppendInCausalArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutAppendInCausalArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutAppendInCausalArgs) ProtoMessage() {}

func (x *PutAppendInCausalArgs) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutAppendInCausalArgs.ProtoReflect.Descriptor instead.
func (*PutAppendInCausalArgs) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{2}
}

func (x *PutAppendInCausalArgs) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PutAppendInCausalArgs) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *PutAppendInCausalArgs) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *PutAppendInCausalArgs) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PutAppendInCausalArgs) GetSeq() int64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *PutAppendInCausalArgs) GetVectorClock() map[string]int32 {
	if x != nil {
		return x.VectorClock
	}
	return nil
}

type PutAppendInCausalReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success     bool             `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	VectorClock map[string]int32 `protobuf:"bytes,2,rep,name=VectorClock,proto3" json:"VectorClock,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *PutAppendInCausalReply) Reset() {
	*x = PutAppendInCausalReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutAppendInCausalReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutAppendInCausalReply) ProtoMessage() {}

func (x *PutAppendInCausalReply) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutAppendInCausalReply.ProtoReflect.Descriptor instead.
func (*PutAppendInCausalReply) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{3}
}

func (x *PutAppendInCausalReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *PutAppendInCausalReply) GetVectorClock() map[string]int32 {
	if x != nil {
		return x.VectorClock
	}
	return nil
}

type GetArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
}

func (x *GetArgs) Reset() {
	*x = GetArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArgs) ProtoMessage() {}

func (x *GetArgs) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArgs.ProtoReflect.Descriptor instead.
func (*GetArgs) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{4}
}

func (x *GetArgs) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value    string `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
	IsLeader bool   `protobuf:"varint,2,opt,name=IsLeader,proto3" json:"IsLeader,omitempty"`
}

func (x *GetReply) Reset() {
	*x = GetReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReply) ProtoMessage() {}

func (x *GetReply) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReply.ProtoReflect.Descriptor instead.
func (*GetReply) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{5}
}

func (x *GetReply) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *GetReply) GetIsLeader() bool {
	if x != nil {
		return x.IsLeader
	}
	return false
}

type GetInCausalArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key         string           `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	VectorClock map[string]int32 `protobuf:"bytes,2,rep,name=VectorClock,proto3" json:"VectorClock,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *GetInCausalArgs) Reset() {
	*x = GetInCausalArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInCausalArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInCausalArgs) ProtoMessage() {}

func (x *GetInCausalArgs) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInCausalArgs.ProtoReflect.Descriptor instead.
func (*GetInCausalArgs) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{6}
}

func (x *GetInCausalArgs) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *GetInCausalArgs) GetVectorClock() map[string]int32 {
	if x != nil {
		return x.VectorClock
	}
	return nil
}

type GetInCausalReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	Value   string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *GetInCausalReply) Reset() {
	*x = GetInCausalReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInCausalReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInCausalReply) ProtoMessage() {}

func (x *GetInCausalReply) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInCausalReply.ProtoReflect.Descriptor instead.
func (*GetInCausalReply) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{7}
}

func (x *GetInCausalReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetInCausalReply) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_kv_proto protoreflect.FileDescriptor

var file_kv_proto_rawDesc = []byte{
	0x0a, 0x08, 0x6b, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x0d, 0x50, 0x75,
	0x74, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x41, 0x72, 0x67, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x4b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x4f, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x65, 0x71, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x53, 0x65, 0x71, 0x22, 0x46, 0x0a, 0x0e, 0x50, 0x75, 0x74, 0x41, 0x70, 0x70, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x73, 0x4c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x49, 0x73, 0x4c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0xfc, 0x01,
	0x0a, 0x15, 0x50, 0x75, 0x74, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x43, 0x61, 0x75,
	0x73, 0x61, 0x6c, 0x41, 0x72, 0x67, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x4f, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x4f, 0x70, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x53, 0x65, 0x71, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x53, 0x65,
	0x71, 0x12, 0x49, 0x0a, 0x0b, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x50, 0x75, 0x74, 0x41, 0x70, 0x70, 0x65,
	0x6e, 0x64, 0x49, 0x6e, 0x43, 0x61, 0x75, 0x73, 0x61, 0x6c, 0x41, 0x72, 0x67, 0x73, 0x2e, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0b, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x1a, 0x3e, 0x0a, 0x10,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xbe, 0x01, 0x0a,
	0x16, 0x50, 0x75, 0x74, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x43, 0x61, 0x75, 0x73,
	0x61, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x4a, 0x0a, 0x0b, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x50, 0x75, 0x74, 0x41, 0x70, 0x70, 0x65,
	0x6e, 0x64, 0x49, 0x6e, 0x43, 0x61, 0x75, 0x73, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0b, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x1a, 0x3e, 0x0a,
	0x10, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x1b, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x41, 0x72, 0x67, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x22, 0x3c, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x49, 0x73, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x49, 0x73, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0xa8, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x43, 0x61, 0x75, 0x73, 0x61, 0x6c, 0x41, 0x72, 0x67, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x43,
	0x0a, 0x0b, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x43, 0x61, 0x75, 0x73, 0x61,
	0x6c, 0x41, 0x72, 0x67, 0x73, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63,
	0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c,
	0x6f, 0x63, 0x6b, 0x1a, 0x3e, 0x0a, 0x10, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f,
	0x63, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x42, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x43, 0x61, 0x75, 0x73,
	0x61, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xd0, 0x01, 0x0a, 0x02, 0x4b, 0x56, 0x12, 0x2e,
	0x0a, 0x09, 0x50, 0x75, 0x74, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x12, 0x0e, 0x2e, 0x50, 0x75,
	0x74, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x0f, 0x2e, 0x50, 0x75,
	0x74, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x1c,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x08, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x67, 0x73, 0x1a,
	0x09, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x11,
	0x50, 0x75, 0x74, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x43, 0x61, 0x75, 0x73, 0x61,
	0x6c, 0x12, 0x16, 0x2e, 0x50, 0x75, 0x74, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x43,
	0x61, 0x75, 0x73, 0x61, 0x6c, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x17, 0x2e, 0x50, 0x75, 0x74, 0x41,
	0x70, 0x70, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x43, 0x61, 0x75, 0x73, 0x61, 0x6c, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x43, 0x61, 0x75,
	0x73, 0x61, 0x6c, 0x12, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x43, 0x61, 0x75, 0x73, 0x61,
	0x6c, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x43, 0x61, 0x75,
	0x73, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f,
	0x3b, 0x6b, 0x76, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kv_proto_rawDescOnce sync.Once
	file_kv_proto_rawDescData = file_kv_proto_rawDesc
)

func file_kv_proto_rawDescGZIP() []byte {
	file_kv_proto_rawDescOnce.Do(func() {
		file_kv_proto_rawDescData = protoimpl.X.CompressGZIP(file_kv_proto_rawDescData)
	})
	return file_kv_proto_rawDescData
}

var file_kv_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_kv_proto_goTypes = []interface{}{
	(*PutAppendArgs)(nil),          // 0: PutAppendArgs
	(*PutAppendReply)(nil),         // 1: PutAppendReply
	(*PutAppendInCausalArgs)(nil),  // 2: PutAppendInCausalArgs
	(*PutAppendInCausalReply)(nil), // 3: PutAppendInCausalReply
	(*GetArgs)(nil),                // 4: GetArgs
	(*GetReply)(nil),               // 5: GetReply
	(*GetInCausalArgs)(nil),        // 6: GetInCausalArgs
	(*GetInCausalReply)(nil),       // 7: GetInCausalReply
	nil,                            // 8: PutAppendInCausalArgs.VectorClockEntry
	nil,                            // 9: PutAppendInCausalReply.VectorClockEntry
	nil,                            // 10: GetInCausalArgs.VectorClockEntry
}
var file_kv_proto_depIdxs = []int32{
	8,  // 0: PutAppendInCausalArgs.VectorClock:type_name -> PutAppendInCausalArgs.VectorClockEntry
	9,  // 1: PutAppendInCausalReply.VectorClock:type_name -> PutAppendInCausalReply.VectorClockEntry
	10, // 2: GetInCausalArgs.VectorClock:type_name -> GetInCausalArgs.VectorClockEntry
	0,  // 3: KV.PutAppend:input_type -> PutAppendArgs
	4,  // 4: KV.Get:input_type -> GetArgs
	2,  // 5: KV.PutAppendInCausal:input_type -> PutAppendInCausalArgs
	6,  // 6: KV.GetInCausal:input_type -> GetInCausalArgs
	1,  // 7: KV.PutAppend:output_type -> PutAppendReply
	5,  // 8: KV.Get:output_type -> GetReply
	3,  // 9: KV.PutAppendInCausal:output_type -> PutAppendInCausalReply
	7,  // 10: KV.GetInCausal:output_type -> GetInCausalReply
	7,  // [7:11] is the sub-list for method output_type
	3,  // [3:7] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_kv_proto_init() }
func file_kv_proto_init() {
	if File_kv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutAppendArgs); i {
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
		file_kv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutAppendReply); i {
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
		file_kv_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutAppendInCausalArgs); i {
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
		file_kv_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutAppendInCausalReply); i {
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
		file_kv_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArgs); i {
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
		file_kv_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReply); i {
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
		file_kv_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInCausalArgs); i {
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
		file_kv_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInCausalReply); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kv_proto_goTypes,
		DependencyIndexes: file_kv_proto_depIdxs,
		MessageInfos:      file_kv_proto_msgTypes,
	}.Build()
	File_kv_proto = out.File
	file_kv_proto_rawDesc = nil
	file_kv_proto_goTypes = nil
	file_kv_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KVClient is the client API for KV service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KVClient interface {
	PutAppend(ctx context.Context, in *PutAppendArgs, opts ...grpc.CallOption) (*PutAppendReply, error)
	Get(ctx context.Context, in *GetArgs, opts ...grpc.CallOption) (*GetReply, error)
	PutAppendInCausal(ctx context.Context, in *PutAppendInCausalArgs, opts ...grpc.CallOption) (*PutAppendInCausalReply, error)
	GetInCausal(ctx context.Context, in *GetInCausalArgs, opts ...grpc.CallOption) (*GetInCausalReply, error)
}

type kVClient struct {
	cc grpc.ClientConnInterface
}

func NewKVClient(cc grpc.ClientConnInterface) KVClient {
	return &kVClient{cc}
}

func (c *kVClient) PutAppend(ctx context.Context, in *PutAppendArgs, opts ...grpc.CallOption) (*PutAppendReply, error) {
	out := new(PutAppendReply)
	err := c.cc.Invoke(ctx, "/KV/PutAppend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) Get(ctx context.Context, in *GetArgs, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := c.cc.Invoke(ctx, "/KV/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) PutAppendInCausal(ctx context.Context, in *PutAppendInCausalArgs, opts ...grpc.CallOption) (*PutAppendInCausalReply, error) {
	out := new(PutAppendInCausalReply)
	err := c.cc.Invoke(ctx, "/KV/PutAppendInCausal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) GetInCausal(ctx context.Context, in *GetInCausalArgs, opts ...grpc.CallOption) (*GetInCausalReply, error) {
	out := new(GetInCausalReply)
	err := c.cc.Invoke(ctx, "/KV/GetInCausal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KVServer is the server API for KV service.
type KVServer interface {
	PutAppend(context.Context, *PutAppendArgs) (*PutAppendReply, error)
	Get(context.Context, *GetArgs) (*GetReply, error)
	PutAppendInCausal(context.Context, *PutAppendInCausalArgs) (*PutAppendInCausalReply, error)
	GetInCausal(context.Context, *GetInCausalArgs) (*GetInCausalReply, error)
}

// UnimplementedKVServer can be embedded to have forward compatible implementations.
type UnimplementedKVServer struct {
}

func (*UnimplementedKVServer) PutAppend(context.Context, *PutAppendArgs) (*PutAppendReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutAppend not implemented")
}
func (*UnimplementedKVServer) Get(context.Context, *GetArgs) (*GetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedKVServer) PutAppendInCausal(context.Context, *PutAppendInCausalArgs) (*PutAppendInCausalReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutAppendInCausal not implemented")
}
func (*UnimplementedKVServer) GetInCausal(context.Context, *GetInCausalArgs) (*GetInCausalReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInCausal not implemented")
}

func RegisterKVServer(s *grpc.Server, srv KVServer) {
	s.RegisterService(&_KV_serviceDesc, srv)
}

func _KV_PutAppend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutAppendArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).PutAppend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KV/PutAppend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).PutAppend(ctx, req.(*PutAppendArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KV/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Get(ctx, req.(*GetArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_PutAppendInCausal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutAppendInCausalArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).PutAppendInCausal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KV/PutAppendInCausal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).PutAppendInCausal(ctx, req.(*PutAppendInCausalArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_GetInCausal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInCausalArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).GetInCausal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KV/GetInCausal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).GetInCausal(ctx, req.(*GetInCausalArgs))
	}
	return interceptor(ctx, in, info, handler)
}

var _KV_serviceDesc = grpc.ServiceDesc{
	ServiceName: "KV",
	HandlerType: (*KVServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutAppend",
			Handler:    _KV_PutAppend_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _KV_Get_Handler,
		},
		{
			MethodName: "PutAppendInCausal",
			Handler:    _KV_PutAppendInCausal_Handler,
		},
		{
			MethodName: "GetInCausal",
			Handler:    _KV_GetInCausal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kv.proto",
}
