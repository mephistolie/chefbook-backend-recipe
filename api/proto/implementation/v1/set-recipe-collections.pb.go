// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.27.0
// source: v1/set-recipe-collections.proto

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

type AddRecipeToCollectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecipeId     string `protobuf:"bytes,1,opt,name=recipeId,proto3" json:"recipeId,omitempty"`
	CollectionId string `protobuf:"bytes,2,opt,name=collectionId,proto3" json:"collectionId,omitempty"`
	UserId       string `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *AddRecipeToCollectionRequest) Reset() {
	*x = AddRecipeToCollectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_set_recipe_collections_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRecipeToCollectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRecipeToCollectionRequest) ProtoMessage() {}

func (x *AddRecipeToCollectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_set_recipe_collections_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRecipeToCollectionRequest.ProtoReflect.Descriptor instead.
func (*AddRecipeToCollectionRequest) Descriptor() ([]byte, []int) {
	return file_v1_set_recipe_collections_proto_rawDescGZIP(), []int{0}
}

func (x *AddRecipeToCollectionRequest) GetRecipeId() string {
	if x != nil {
		return x.RecipeId
	}
	return ""
}

func (x *AddRecipeToCollectionRequest) GetCollectionId() string {
	if x != nil {
		return x.CollectionId
	}
	return ""
}

func (x *AddRecipeToCollectionRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type AddRecipeToCollectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AddRecipeToCollectionResponse) Reset() {
	*x = AddRecipeToCollectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_set_recipe_collections_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRecipeToCollectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRecipeToCollectionResponse) ProtoMessage() {}

func (x *AddRecipeToCollectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_set_recipe_collections_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRecipeToCollectionResponse.ProtoReflect.Descriptor instead.
func (*AddRecipeToCollectionResponse) Descriptor() ([]byte, []int) {
	return file_v1_set_recipe_collections_proto_rawDescGZIP(), []int{1}
}

func (x *AddRecipeToCollectionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type RemoveRecipeFromCollectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecipeId     string `protobuf:"bytes,1,opt,name=recipeId,proto3" json:"recipeId,omitempty"`
	CollectionId string `protobuf:"bytes,2,opt,name=collectionId,proto3" json:"collectionId,omitempty"`
	UserId       string `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *RemoveRecipeFromCollectionRequest) Reset() {
	*x = RemoveRecipeFromCollectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_set_recipe_collections_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRecipeFromCollectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRecipeFromCollectionRequest) ProtoMessage() {}

func (x *RemoveRecipeFromCollectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_set_recipe_collections_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRecipeFromCollectionRequest.ProtoReflect.Descriptor instead.
func (*RemoveRecipeFromCollectionRequest) Descriptor() ([]byte, []int) {
	return file_v1_set_recipe_collections_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveRecipeFromCollectionRequest) GetRecipeId() string {
	if x != nil {
		return x.RecipeId
	}
	return ""
}

func (x *RemoveRecipeFromCollectionRequest) GetCollectionId() string {
	if x != nil {
		return x.CollectionId
	}
	return ""
}

func (x *RemoveRecipeFromCollectionRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type RemoveRecipeFromCollectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RemoveRecipeFromCollectionResponse) Reset() {
	*x = RemoveRecipeFromCollectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_set_recipe_collections_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRecipeFromCollectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRecipeFromCollectionResponse) ProtoMessage() {}

func (x *RemoveRecipeFromCollectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_set_recipe_collections_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRecipeFromCollectionResponse.ProtoReflect.Descriptor instead.
func (*RemoveRecipeFromCollectionResponse) Descriptor() ([]byte, []int) {
	return file_v1_set_recipe_collections_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveRecipeFromCollectionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SetRecipeCollectionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecipeId      string   `protobuf:"bytes,1,opt,name=recipeId,proto3" json:"recipeId,omitempty"`
	UserId        string   `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	CollectionIds []string `protobuf:"bytes,3,rep,name=collectionIds,proto3" json:"collectionIds,omitempty"`
}

func (x *SetRecipeCollectionsRequest) Reset() {
	*x = SetRecipeCollectionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_set_recipe_collections_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRecipeCollectionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRecipeCollectionsRequest) ProtoMessage() {}

func (x *SetRecipeCollectionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_set_recipe_collections_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRecipeCollectionsRequest.ProtoReflect.Descriptor instead.
func (*SetRecipeCollectionsRequest) Descriptor() ([]byte, []int) {
	return file_v1_set_recipe_collections_proto_rawDescGZIP(), []int{4}
}

func (x *SetRecipeCollectionsRequest) GetRecipeId() string {
	if x != nil {
		return x.RecipeId
	}
	return ""
}

func (x *SetRecipeCollectionsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SetRecipeCollectionsRequest) GetCollectionIds() []string {
	if x != nil {
		return x.CollectionIds
	}
	return nil
}

type SetRecipeCollectionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SetRecipeCollectionsResponse) Reset() {
	*x = SetRecipeCollectionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_set_recipe_collections_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRecipeCollectionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRecipeCollectionsResponse) ProtoMessage() {}

func (x *SetRecipeCollectionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_set_recipe_collections_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRecipeCollectionsResponse.ProtoReflect.Descriptor instead.
func (*SetRecipeCollectionsResponse) Descriptor() ([]byte, []int) {
	return file_v1_set_recipe_collections_proto_rawDescGZIP(), []int{5}
}

func (x *SetRecipeCollectionsResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_v1_set_recipe_collections_proto protoreflect.FileDescriptor

var file_v1_set_recipe_collections_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x74, 0x2d, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2d,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x76, 0x31, 0x22, 0x76, 0x0a, 0x1c, 0x41, 0x64, 0x64, 0x52, 0x65, 0x63, 0x69,
	0x70, 0x65, 0x54, 0x6f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x49,
	0x64, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x39, 0x0a,
	0x1d, 0x41, 0x64, 0x64, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x54, 0x6f, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x7b, 0x0a, 0x21, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x22, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x77, 0x0a, 0x1b, 0x53, 0x65, 0x74, 0x52, 0x65, 0x63, 0x69,
	0x70, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x22, 0x38,
	0x0a, 0x1c, 0x53, 0x65, 0x74, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x70, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x6c,
	0x69, 0x65, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x62, 0x6f, 0x6f, 0x6b, 0x2d, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2d, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_set_recipe_collections_proto_rawDescOnce sync.Once
	file_v1_set_recipe_collections_proto_rawDescData = file_v1_set_recipe_collections_proto_rawDesc
)

func file_v1_set_recipe_collections_proto_rawDescGZIP() []byte {
	file_v1_set_recipe_collections_proto_rawDescOnce.Do(func() {
		file_v1_set_recipe_collections_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_set_recipe_collections_proto_rawDescData)
	})
	return file_v1_set_recipe_collections_proto_rawDescData
}

var file_v1_set_recipe_collections_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v1_set_recipe_collections_proto_goTypes = []interface{}{
	(*AddRecipeToCollectionRequest)(nil),       // 0: v1.AddRecipeToCollectionRequest
	(*AddRecipeToCollectionResponse)(nil),      // 1: v1.AddRecipeToCollectionResponse
	(*RemoveRecipeFromCollectionRequest)(nil),  // 2: v1.RemoveRecipeFromCollectionRequest
	(*RemoveRecipeFromCollectionResponse)(nil), // 3: v1.RemoveRecipeFromCollectionResponse
	(*SetRecipeCollectionsRequest)(nil),        // 4: v1.SetRecipeCollectionsRequest
	(*SetRecipeCollectionsResponse)(nil),       // 5: v1.SetRecipeCollectionsResponse
}
var file_v1_set_recipe_collections_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_set_recipe_collections_proto_init() }
func file_v1_set_recipe_collections_proto_init() {
	if File_v1_set_recipe_collections_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_set_recipe_collections_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRecipeToCollectionRequest); i {
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
		file_v1_set_recipe_collections_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRecipeToCollectionResponse); i {
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
		file_v1_set_recipe_collections_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRecipeFromCollectionRequest); i {
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
		file_v1_set_recipe_collections_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRecipeFromCollectionResponse); i {
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
		file_v1_set_recipe_collections_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRecipeCollectionsRequest); i {
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
		file_v1_set_recipe_collections_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRecipeCollectionsResponse); i {
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
			RawDescriptor: file_v1_set_recipe_collections_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_set_recipe_collections_proto_goTypes,
		DependencyIndexes: file_v1_set_recipe_collections_proto_depIdxs,
		MessageInfos:      file_v1_set_recipe_collections_proto_msgTypes,
	}.Build()
	File_v1_set_recipe_collections_proto = out.File
	file_v1_set_recipe_collections_proto_rawDesc = nil
	file_v1_set_recipe_collections_proto_goTypes = nil
	file_v1_set_recipe_collections_proto_depIdxs = nil
}
