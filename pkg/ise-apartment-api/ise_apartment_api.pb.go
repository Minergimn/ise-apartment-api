// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: ozonmp/ise_apartment_api/v1/ise_apartment_api.proto

package ise_apartment_api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Apartment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Object string `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
	Owner  string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *Apartment) Reset() {
	*x = Apartment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Apartment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Apartment) ProtoMessage() {}

func (x *Apartment) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Apartment.ProtoReflect.Descriptor instead.
func (*Apartment) Descriptor() ([]byte, []int) {
	return file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_rawDescGZIP(), []int{0}
}

func (x *Apartment) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Apartment) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *Apartment) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

type DescribeApartmentV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApartmentId uint64 `protobuf:"varint,1,opt,name=apartment_id,json=apartmentId,proto3" json:"apartment_id,omitempty"`
}

func (x *DescribeApartmentV1Request) Reset() {
	*x = DescribeApartmentV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeApartmentV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeApartmentV1Request) ProtoMessage() {}

func (x *DescribeApartmentV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeApartmentV1Request.ProtoReflect.Descriptor instead.
func (*DescribeApartmentV1Request) Descriptor() ([]byte, []int) {
	return file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeApartmentV1Request) GetApartmentId() uint64 {
	if x != nil {
		return x.ApartmentId
	}
	return 0
}

type DescribeApartmentV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *Apartment `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *DescribeApartmentV1Response) Reset() {
	*x = DescribeApartmentV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeApartmentV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeApartmentV1Response) ProtoMessage() {}

func (x *DescribeApartmentV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeApartmentV1Response.ProtoReflect.Descriptor instead.
func (*DescribeApartmentV1Response) Descriptor() ([]byte, []int) {
	return file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeApartmentV1Response) GetValue() *Apartment {
	if x != nil {
		return x.Value
	}
	return nil
}

type CreateApartmentV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *Apartment `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *CreateApartmentV1Request) Reset() {
	*x = CreateApartmentV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateApartmentV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateApartmentV1Request) ProtoMessage() {}

func (x *CreateApartmentV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateApartmentV1Request.ProtoReflect.Descriptor instead.
func (*CreateApartmentV1Request) Descriptor() ([]byte, []int) {
	return file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_rawDescGZIP(), []int{3}
}

func (x *CreateApartmentV1Request) GetValue() *Apartment {
	if x != nil {
		return x.Value
	}
	return nil
}

type CreateApartmentV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApartmentId uint64 `protobuf:"varint,1,opt,name=apartment_id,json=apartmentId,proto3" json:"apartment_id,omitempty"`
}

func (x *CreateApartmentV1Response) Reset() {
	*x = CreateApartmentV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateApartmentV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateApartmentV1Response) ProtoMessage() {}

func (x *CreateApartmentV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateApartmentV1Response.ProtoReflect.Descriptor instead.
func (*CreateApartmentV1Response) Descriptor() ([]byte, []int) {
	return file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_rawDescGZIP(), []int{4}
}

func (x *CreateApartmentV1Response) GetApartmentId() uint64 {
	if x != nil {
		return x.ApartmentId
	}
	return 0
}

type ListApartmentsV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params *ListApartmentsV1Request_ListApartmentsParams `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *ListApartmentsV1Request) Reset() {
	*x = ListApartmentsV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListApartmentsV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApartmentsV1Request) ProtoMessage() {}

func (x *ListApartmentsV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApartmentsV1Request.ProtoReflect.Descriptor instead.
func (*ListApartmentsV1Request) Descriptor() ([]byte, []int) {
	return file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_rawDescGZIP(), []int{5}
}

func (x *ListApartmentsV1Request) GetParams() *ListApartmentsV1Request_ListApartmentsParams {
	if x != nil {
		return x.Params
	}
	return nil
}

type ListApartmentsV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Apartment `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListApartmentsV1Response) Reset() {
	*x = ListApartmentsV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListApartmentsV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApartmentsV1Response) ProtoMessage() {}

func (x *ListApartmentsV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApartmentsV1Response.ProtoReflect.Descriptor instead.
func (*ListApartmentsV1Response) Descriptor() ([]byte, []int) {
	return file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_rawDescGZIP(), []int{6}
}

func (x *ListApartmentsV1Response) GetItems() []*Apartment {
	if x != nil {
		return x.Items
	}
	return nil
}

type RemoveApartmentV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApartmentId uint64 `protobuf:"varint,1,opt,name=apartment_id,json=apartmentId,proto3" json:"apartment_id,omitempty"`
}

func (x *RemoveApartmentV1Request) Reset() {
	*x = RemoveApartmentV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveApartmentV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveApartmentV1Request) ProtoMessage() {}

func (x *RemoveApartmentV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveApartmentV1Request.ProtoReflect.Descriptor instead.
func (*RemoveApartmentV1Request) Descriptor() ([]byte, []int) {
	return file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveApartmentV1Request) GetApartmentId() uint64 {
	if x != nil {
		return x.ApartmentId
	}
	return 0
}

type RemoveApartmentV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Found bool `protobuf:"varint,1,opt,name=found,proto3" json:"found,omitempty"`
}

func (x *RemoveApartmentV1Response) Reset() {
	*x = RemoveApartmentV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveApartmentV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveApartmentV1Response) ProtoMessage() {}

func (x *RemoveApartmentV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveApartmentV1Response.ProtoReflect.Descriptor instead.
func (*RemoveApartmentV1Response) Descriptor() ([]byte, []int) {
	return file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveApartmentV1Response) GetFound() bool {
	if x != nil {
		return x.Found
	}
	return false
}

type ListApartmentsV1Request_ListApartmentsParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids    []uint64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Offset uint64   `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  uint64   `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Object string   `protobuf:"bytes,4,opt,name=object,proto3" json:"object,omitempty"`
	Owner  string   `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *ListApartmentsV1Request_ListApartmentsParams) Reset() {
	*x = ListApartmentsV1Request_ListApartmentsParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListApartmentsV1Request_ListApartmentsParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApartmentsV1Request_ListApartmentsParams) ProtoMessage() {}

func (x *ListApartmentsV1Request_ListApartmentsParams) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApartmentsV1Request_ListApartmentsParams.ProtoReflect.Descriptor instead.
func (*ListApartmentsV1Request_ListApartmentsParams) Descriptor() ([]byte, []int) {
	return file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_rawDescGZIP(), []int{5, 0}
}

func (x *ListApartmentsV1Request_ListApartmentsParams) GetIds() []uint64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *ListApartmentsV1Request_ListApartmentsParams) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListApartmentsV1Request_ListApartmentsParams) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListApartmentsV1Request_ListApartmentsParams) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *ListApartmentsV1Request_ListApartmentsParams) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

var File_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto protoreflect.FileDescriptor

var file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_rawDesc = []byte{
	0x0a, 0x33, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2f, 0x69, 0x73, 0x65, 0x5f, 0x61, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x73,
	0x65, 0x5f, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x69, 0x73,
	0x65, 0x5f, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x09, 0x41, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18,
	0xe8, 0x07, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x20, 0x0a, 0x05, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05,
	0x10, 0x01, 0x18, 0xe8, 0x07, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x48, 0x0a, 0x1a,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x0c, 0x61, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x0b, 0x61, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x5b, 0x0a, 0x1b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x41, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x69, 0x73,
	0x65, 0x5f, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x62, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x46, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x69, 0x73, 0x65, 0x5f, 0x61, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3e, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x61, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x83, 0x02, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x61, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x49, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x69, 0x73, 0x65,
	0x5f, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x84, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x03, 0x69, 0x64,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x58, 0x0a,
	0x18, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x56,
	0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d,
	0x70, 0x2e, 0x69, 0x73, 0x65, 0x5f, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x46, 0x0a, 0x18, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x41, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x0c, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02,
	0x20, 0x00, 0x52, 0x0b, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22,
	0x31, 0x0a, 0x19, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x75,
	0x6e, 0x64, 0x32, 0xd8, 0x05, 0x0a, 0x16, 0x49, 0x73, 0x65, 0x41, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xb3, 0x01,
	0x0a, 0x13, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x56, 0x31, 0x12, 0x37, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x69,
	0x73, 0x65, 0x5f, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38,
	0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x69, 0x73, 0x65, 0x5f, 0x61, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x31,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23,
	0x12, 0x21, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x7d, 0x12, 0xac, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x31, 0x12, 0x35, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e,
	0x6d, 0x70, 0x2e, 0x69, 0x73, 0x65, 0x5f, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x36, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x69, 0x73, 0x65, 0x5f, 0x61, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x31,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22,
	0x22, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0xa8, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x56, 0x31, 0x12, 0x34, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70,
	0x2e, 0x69, 0x73, 0x65, 0x5f, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e,
	0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x69, 0x73, 0x65, 0x5f, 0x61, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x17, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x3a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0xad, 0x01,
	0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x56, 0x31, 0x12, 0x35, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x69, 0x73, 0x65,
	0x5f, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x6f, 0x7a, 0x6f,
	0x6e, 0x6d, 0x70, 0x2e, 0x69, 0x73, 0x65, 0x5f, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x2a, 0x21, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b,
	0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x4d, 0x5a,
	0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x7a, 0x6f, 0x6e,
	0x6d, 0x70, 0x2f, 0x69, 0x73, 0x65, 0x2d, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x73, 0x65, 0x2d, 0x61, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x3b, 0x69, 0x73, 0x65, 0x5f, 0x61,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_rawDescOnce sync.Once
	file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_rawDescData = file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_rawDesc
)

func file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_rawDescGZIP() []byte {
	file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_rawDescOnce.Do(func() {
		file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_rawDescData)
	})
	return file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_rawDescData
}

var file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_goTypes = []interface{}{
	(*Apartment)(nil),                                    // 0: ozonmp.ise_apartment_api.v1.Apartment
	(*DescribeApartmentV1Request)(nil),                   // 1: ozonmp.ise_apartment_api.v1.DescribeApartmentV1Request
	(*DescribeApartmentV1Response)(nil),                  // 2: ozonmp.ise_apartment_api.v1.DescribeApartmentV1Response
	(*CreateApartmentV1Request)(nil),                     // 3: ozonmp.ise_apartment_api.v1.CreateApartmentV1Request
	(*CreateApartmentV1Response)(nil),                    // 4: ozonmp.ise_apartment_api.v1.CreateApartmentV1Response
	(*ListApartmentsV1Request)(nil),                      // 5: ozonmp.ise_apartment_api.v1.ListApartmentsV1Request
	(*ListApartmentsV1Response)(nil),                     // 6: ozonmp.ise_apartment_api.v1.ListApartmentsV1Response
	(*RemoveApartmentV1Request)(nil),                     // 7: ozonmp.ise_apartment_api.v1.RemoveApartmentV1Request
	(*RemoveApartmentV1Response)(nil),                    // 8: ozonmp.ise_apartment_api.v1.RemoveApartmentV1Response
	(*ListApartmentsV1Request_ListApartmentsParams)(nil), // 9: ozonmp.ise_apartment_api.v1.ListApartmentsV1Request.ListApartmentsParams
}
var file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_depIdxs = []int32{
	0, // 0: ozonmp.ise_apartment_api.v1.DescribeApartmentV1Response.value:type_name -> ozonmp.ise_apartment_api.v1.Apartment
	0, // 1: ozonmp.ise_apartment_api.v1.CreateApartmentV1Request.value:type_name -> ozonmp.ise_apartment_api.v1.Apartment
	9, // 2: ozonmp.ise_apartment_api.v1.ListApartmentsV1Request.params:type_name -> ozonmp.ise_apartment_api.v1.ListApartmentsV1Request.ListApartmentsParams
	0, // 3: ozonmp.ise_apartment_api.v1.ListApartmentsV1Response.items:type_name -> ozonmp.ise_apartment_api.v1.Apartment
	1, // 4: ozonmp.ise_apartment_api.v1.IseApartmentApiService.DescribeApartmentV1:input_type -> ozonmp.ise_apartment_api.v1.DescribeApartmentV1Request
	3, // 5: ozonmp.ise_apartment_api.v1.IseApartmentApiService.CreateApartmentV1:input_type -> ozonmp.ise_apartment_api.v1.CreateApartmentV1Request
	5, // 6: ozonmp.ise_apartment_api.v1.IseApartmentApiService.ListApartmentsV1:input_type -> ozonmp.ise_apartment_api.v1.ListApartmentsV1Request
	7, // 7: ozonmp.ise_apartment_api.v1.IseApartmentApiService.RemoveApartmentV1:input_type -> ozonmp.ise_apartment_api.v1.RemoveApartmentV1Request
	2, // 8: ozonmp.ise_apartment_api.v1.IseApartmentApiService.DescribeApartmentV1:output_type -> ozonmp.ise_apartment_api.v1.DescribeApartmentV1Response
	4, // 9: ozonmp.ise_apartment_api.v1.IseApartmentApiService.CreateApartmentV1:output_type -> ozonmp.ise_apartment_api.v1.CreateApartmentV1Response
	6, // 10: ozonmp.ise_apartment_api.v1.IseApartmentApiService.ListApartmentsV1:output_type -> ozonmp.ise_apartment_api.v1.ListApartmentsV1Response
	8, // 11: ozonmp.ise_apartment_api.v1.IseApartmentApiService.RemoveApartmentV1:output_type -> ozonmp.ise_apartment_api.v1.RemoveApartmentV1Response
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_init() }
func file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_init() {
	if File_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Apartment); i {
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
		file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeApartmentV1Request); i {
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
		file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeApartmentV1Response); i {
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
		file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateApartmentV1Request); i {
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
		file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateApartmentV1Response); i {
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
		file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListApartmentsV1Request); i {
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
		file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListApartmentsV1Response); i {
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
		file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveApartmentV1Request); i {
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
		file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveApartmentV1Response); i {
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
		file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListApartmentsV1Request_ListApartmentsParams); i {
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
			RawDescriptor: file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_goTypes,
		DependencyIndexes: file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_depIdxs,
		MessageInfos:      file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes,
	}.Build()
	File_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto = out.File
	file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_rawDesc = nil
	file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_goTypes = nil
	file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_depIdxs = nil
}
