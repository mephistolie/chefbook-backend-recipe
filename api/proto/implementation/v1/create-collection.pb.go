// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.27.0
// source: v1/create-collection.proto

package v1

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

type CreateCollectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string  `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	CollectionId *string `protobuf:"bytes,2,opt,name=collectionId,proto3,oneof" json:"collectionId,omitempty"`
	Name         string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Visibility   string  `protobuf:"bytes,4,opt,name=visibility,proto3" json:"visibility,omitempty"`
}

func (x *CreateCollectionRequest) Reset() {
	*x = CreateCollectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_create_collection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCollectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCollectionRequest) ProtoMessage() {}

func (x *CreateCollectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_create_collection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCollectionRequest.ProtoReflect.Descriptor instead.
func (*CreateCollectionRequest) Descriptor() ([]byte, []int) {
	return file_v1_create_collection_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCollectionRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateCollectionRequest) GetCollectionId() string {
	if x != nil && x.CollectionId != nil {
		return *x.CollectionId
	}
	return ""
}

func (x *CreateCollectionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCollectionRequest) GetVisibility() string {
	if x != nil {
		return x.Visibility
	}
	return ""
}

type CreateCollectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CollectionId string `protobuf:"bytes,1,opt,name=collectionId,proto3" json:"collectionId,omitempty"`
}

func (x *CreateCollectionResponse) Reset() {
	*x = CreateCollectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_create_collection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCollectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCollectionResponse) ProtoMessage() {}

func (x *CreateCollectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_create_collection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCollectionResponse.ProtoReflect.Descriptor instead.
func (*CreateCollectionResponse) Descriptor() ([]byte, []int) {
	return file_v1_create_collection_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCollectionResponse) GetCollectionId() string {
	if x != nil {
		return x.CollectionId
	}
	return ""
}

var File_v1_create_collection_proto protoreflect.FileDescriptor

var file_v1_create_collection_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2d, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31,
	0x22, 0x9f, 0x01, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x22, 0x3e, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x65, 0x70, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x6c, 0x69, 0x65, 0x2f, 0x63, 0x68, 0x65,
	0x66, 0x62, 0x6f, 0x6f, 0x6b, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_create_collection_proto_rawDescOnce sync.Once
	file_v1_create_collection_proto_rawDescData = file_v1_create_collection_proto_rawDesc
)

func file_v1_create_collection_proto_rawDescGZIP() []byte {
	file_v1_create_collection_proto_rawDescOnce.Do(func() {
		file_v1_create_collection_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_create_collection_proto_rawDescData)
	})
	return file_v1_create_collection_proto_rawDescData
}

var file_v1_create_collection_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_create_collection_proto_goTypes = []interface{}{
	(*CreateCollectionRequest)(nil),  // 0: v1.CreateCollectionRequest
	(*CreateCollectionResponse)(nil), // 1: v1.CreateCollectionResponse
}
var file_v1_create_collection_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_create_collection_proto_init() }
func file_v1_create_collection_proto_init() {
	if File_v1_create_collection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_create_collection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCollectionRequest); i {
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
		file_v1_create_collection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCollectionResponse); i {
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
	file_v1_create_collection_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_create_collection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_create_collection_proto_goTypes,
		DependencyIndexes: file_v1_create_collection_proto_depIdxs,
		MessageInfos:      file_v1_create_collection_proto_msgTypes,
	}.Build()
	File_v1_create_collection_proto = out.File
	file_v1_create_collection_proto_rawDesc = nil
	file_v1_create_collection_proto_goTypes = nil
	file_v1_create_collection_proto_depIdxs = nil
}
