// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.27.0
// source: v1/get-collection.proto

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

type GetCollectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	CollectionId string `protobuf:"bytes,2,opt,name=collectionId,proto3" json:"collectionId,omitempty"`
}

func (x *GetCollectionRequest) Reset() {
	*x = GetCollectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_get_collection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionRequest) ProtoMessage() {}

func (x *GetCollectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_get_collection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCollectionRequest.ProtoReflect.Descriptor instead.
func (*GetCollectionRequest) Descriptor() ([]byte, []int) {
	return file_v1_get_collection_proto_rawDescGZIP(), []int{0}
}

func (x *GetCollectionRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetCollectionRequest) GetCollectionId() string {
	if x != nil {
		return x.CollectionId
	}
	return ""
}

type GetCollectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Collection   *Collection                   `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
	ProfilesInfo map[string]*RecipeProfileInfo `protobuf:"bytes,2,rep,name=profilesInfo,proto3" json:"profilesInfo,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetCollectionResponse) Reset() {
	*x = GetCollectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_get_collection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionResponse) ProtoMessage() {}

func (x *GetCollectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_get_collection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCollectionResponse.ProtoReflect.Descriptor instead.
func (*GetCollectionResponse) Descriptor() ([]byte, []int) {
	return file_v1_get_collection_proto_rawDescGZIP(), []int{1}
}

func (x *GetCollectionResponse) GetCollection() *Collection {
	if x != nil {
		return x.Collection
	}
	return nil
}

func (x *GetCollectionResponse) GetProfilesInfo() map[string]*RecipeProfileInfo {
	if x != nil {
		return x.ProfilesInfo
	}
	return nil
}

var File_v1_get_collection_proto protoreflect.FileDescriptor

var file_v1_get_collection_proto_rawDesc = []byte{
	0x0a, 0x17, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x13, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x15, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2d, 0x69,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xf0, 0x01,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4f, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x56, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x65, 0x70, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x6c, 0x69, 0x65, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x62,
	0x6f, 0x6f, 0x6b, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x72, 0x65, 0x63, 0x69,
	0x70, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_get_collection_proto_rawDescOnce sync.Once
	file_v1_get_collection_proto_rawDescData = file_v1_get_collection_proto_rawDesc
)

func file_v1_get_collection_proto_rawDescGZIP() []byte {
	file_v1_get_collection_proto_rawDescOnce.Do(func() {
		file_v1_get_collection_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_get_collection_proto_rawDescData)
	})
	return file_v1_get_collection_proto_rawDescData
}

var file_v1_get_collection_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1_get_collection_proto_goTypes = []interface{}{
	(*GetCollectionRequest)(nil),  // 0: v1.GetCollectionRequest
	(*GetCollectionResponse)(nil), // 1: v1.GetCollectionResponse
	nil,                           // 2: v1.GetCollectionResponse.ProfilesInfoEntry
	(*Collection)(nil),            // 3: v1.Collection
	(*RecipeProfileInfo)(nil),     // 4: v1.RecipeProfileInfo
}
var file_v1_get_collection_proto_depIdxs = []int32{
	3, // 0: v1.GetCollectionResponse.collection:type_name -> v1.Collection
	2, // 1: v1.GetCollectionResponse.profilesInfo:type_name -> v1.GetCollectionResponse.ProfilesInfoEntry
	4, // 2: v1.GetCollectionResponse.ProfilesInfoEntry.value:type_name -> v1.RecipeProfileInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v1_get_collection_proto_init() }
func file_v1_get_collection_proto_init() {
	if File_v1_get_collection_proto != nil {
		return
	}
	file_v1_collection_proto_init()
	file_v1_profile_info_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_get_collection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCollectionRequest); i {
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
		file_v1_get_collection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCollectionResponse); i {
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
			RawDescriptor: file_v1_get_collection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_get_collection_proto_goTypes,
		DependencyIndexes: file_v1_get_collection_proto_depIdxs,
		MessageInfos:      file_v1_get_collection_proto_msgTypes,
	}.Build()
	File_v1_get_collection_proto = out.File
	file_v1_get_collection_proto_rawDesc = nil
	file_v1_get_collection_proto_goTypes = nil
	file_v1_get_collection_proto_depIdxs = nil
}
