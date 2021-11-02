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
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

	Id      uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Object  string                 `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
	Owner   string                 `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Created *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created,proto3" json:"created,omitempty"`
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

func (x *Apartment) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
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
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x09, 0x41, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0x48, 0x0a, 0x1a, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x0c, 0x61, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x0b, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x5b, 0x0a, 0x1b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x41, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x69, 0x73, 0x65,
	0x5f, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x32, 0xca, 0x01, 0x0a, 0x16, 0x49, 0x73, 0x65, 0x41, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xaf, 0x01,
	0x0a, 0x13, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x56, 0x31, 0x12, 0x37, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x69,
	0x73, 0x65, 0x5f, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38,
	0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x69, 0x73, 0x65, 0x5f, 0x61, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x31,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f,
	0x12, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x2f, 0x7b, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x42,
	0x4d, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x7a,
	0x6f, 0x6e, 0x6d, 0x70, 0x2f, 0x69, 0x73, 0x65, 0x2d, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x73, 0x65, 0x2d, 0x61,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x3b, 0x69, 0x73, 0x65,
	0x5f, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_goTypes = []interface{}{
	(*Apartment)(nil),                   // 0: ozonmp.ise_apartment_api.v1.Apartment
	(*DescribeApartmentV1Request)(nil),  // 1: ozonmp.ise_apartment_api.v1.DescribeApartmentV1Request
	(*DescribeApartmentV1Response)(nil), // 2: ozonmp.ise_apartment_api.v1.DescribeApartmentV1Response
	(*timestamppb.Timestamp)(nil),       // 3: google.protobuf.Timestamp
}
var file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_depIdxs = []int32{
	3, // 0: ozonmp.ise_apartment_api.v1.Apartment.created:type_name -> google.protobuf.Timestamp
	0, // 1: ozonmp.ise_apartment_api.v1.DescribeApartmentV1Response.value:type_name -> ozonmp.ise_apartment_api.v1.Apartment
	1, // 2: ozonmp.ise_apartment_api.v1.IseApartmentApiService.DescribeApartmentV1:input_type -> ozonmp.ise_apartment_api.v1.DescribeApartmentV1Request
	2, // 3: ozonmp.ise_apartment_api.v1.IseApartmentApiService.DescribeApartmentV1:output_type -> ozonmp.ise_apartment_api.v1.DescribeApartmentV1Response
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ozonmp_ise_apartment_api_v1_ise_apartment_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
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
