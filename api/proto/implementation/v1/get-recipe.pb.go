// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.2
// source: v1/get-recipe.proto

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

type GetRecipeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecipeId         string  `protobuf:"bytes,1,opt,name=recipeId,proto3" json:"recipeId,omitempty"`
	UserId           string  `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Language         *string `protobuf:"bytes,3,opt,name=language,proto3,oneof" json:"language,omitempty"`
	TranslatorId     *string `protobuf:"bytes,4,opt,name=translatorId,proto3,oneof" json:"translatorId,omitempty"`
	Translate        bool    `protobuf:"varint,5,opt,name=translate,proto3" json:"translate,omitempty"`
	SubscriptionPlan string  `protobuf:"bytes,6,opt,name=subscriptionPlan,proto3" json:"subscriptionPlan,omitempty"`
}

func (x *GetRecipeRequest) Reset() {
	*x = GetRecipeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_get_recipe_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecipeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecipeRequest) ProtoMessage() {}

func (x *GetRecipeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_get_recipe_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecipeRequest.ProtoReflect.Descriptor instead.
func (*GetRecipeRequest) Descriptor() ([]byte, []int) {
	return file_v1_get_recipe_proto_rawDescGZIP(), []int{0}
}

func (x *GetRecipeRequest) GetRecipeId() string {
	if x != nil {
		return x.RecipeId
	}
	return ""
}

func (x *GetRecipeRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetRecipeRequest) GetLanguage() string {
	if x != nil && x.Language != nil {
		return *x.Language
	}
	return ""
}

func (x *GetRecipeRequest) GetTranslatorId() string {
	if x != nil && x.TranslatorId != nil {
		return *x.TranslatorId
	}
	return ""
}

func (x *GetRecipeRequest) GetTranslate() bool {
	if x != nil {
		return x.Translate
	}
	return false
}

func (x *GetRecipeRequest) GetSubscriptionPlan() string {
	if x != nil {
		return x.SubscriptionPlan
	}
	return ""
}

type GetRecipeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recipe     *Recipe                        `protobuf:"bytes,1,opt,name=recipe,proto3" json:"recipe,omitempty"`
	Categories map[string]*RecipeCategoryInfo `protobuf:"bytes,2,rep,name=categories,proto3" json:"categories,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Tags       map[string]*RecipeTag          `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	TagGroups  map[string]string              `protobuf:"bytes,4,rep,name=tagGroups,proto3" json:"tagGroups,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetRecipeResponse) Reset() {
	*x = GetRecipeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_get_recipe_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecipeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecipeResponse) ProtoMessage() {}

func (x *GetRecipeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_get_recipe_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecipeResponse.ProtoReflect.Descriptor instead.
func (*GetRecipeResponse) Descriptor() ([]byte, []int) {
	return file_v1_get_recipe_proto_rawDescGZIP(), []int{1}
}

func (x *GetRecipeResponse) GetRecipe() *Recipe {
	if x != nil {
		return x.Recipe
	}
	return nil
}

func (x *GetRecipeResponse) GetCategories() map[string]*RecipeCategoryInfo {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *GetRecipeResponse) GetTags() map[string]*RecipeTag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *GetRecipeResponse) GetTagGroups() map[string]string {
	if x != nil {
		return x.TagGroups
	}
	return nil
}

var File_v1_get_recipe_proto protoreflect.FileDescriptor

var file_v1_get_recipe_proto_rawDesc = []byte{
	0x0a, 0x13, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x0f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x65, 0x2d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2d,
	0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x01, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x6f,
	0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0c, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x49, 0x64, 0x22, 0xd4, 0x03, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x69,
	0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x06, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x12, 0x45,
	0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x69, 0x70,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x69,
	0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x42, 0x0a, 0x09, 0x74, 0x61,
	0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x61, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x09, 0x74, 0x61, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x1a, 0x55,
	0x0a, 0x0f, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x46, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x54,
	0x61, 0x67, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3c, 0x0a,
	0x0e, 0x54, 0x61, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x3d, 0x5a, 0x3b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x70, 0x68, 0x69, 0x73,
	0x74, 0x6f, 0x6c, 0x69, 0x65, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x62, 0x6f, 0x6f, 0x6b, 0x2d, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_v1_get_recipe_proto_rawDescOnce sync.Once
	file_v1_get_recipe_proto_rawDescData = file_v1_get_recipe_proto_rawDesc
)

func file_v1_get_recipe_proto_rawDescGZIP() []byte {
	file_v1_get_recipe_proto_rawDescOnce.Do(func() {
		file_v1_get_recipe_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_get_recipe_proto_rawDescData)
	})
	return file_v1_get_recipe_proto_rawDescData
}

var file_v1_get_recipe_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_v1_get_recipe_proto_goTypes = []interface{}{
	(*GetRecipeRequest)(nil),   // 0: v1.GetRecipeRequest
	(*GetRecipeResponse)(nil),  // 1: v1.GetRecipeResponse
	nil,                        // 2: v1.GetRecipeResponse.CategoriesEntry
	nil,                        // 3: v1.GetRecipeResponse.TagsEntry
	nil,                        // 4: v1.GetRecipeResponse.TagGroupsEntry
	(*Recipe)(nil),             // 5: v1.Recipe
	(*RecipeCategoryInfo)(nil), // 6: v1.RecipeCategoryInfo
	(*RecipeTag)(nil),          // 7: v1.RecipeTag
}
var file_v1_get_recipe_proto_depIdxs = []int32{
	5, // 0: v1.GetRecipeResponse.recipe:type_name -> v1.Recipe
	2, // 1: v1.GetRecipeResponse.categories:type_name -> v1.GetRecipeResponse.CategoriesEntry
	3, // 2: v1.GetRecipeResponse.tags:type_name -> v1.GetRecipeResponse.TagsEntry
	4, // 3: v1.GetRecipeResponse.tagGroups:type_name -> v1.GetRecipeResponse.TagGroupsEntry
	6, // 4: v1.GetRecipeResponse.CategoriesEntry.value:type_name -> v1.RecipeCategoryInfo
	7, // 5: v1.GetRecipeResponse.TagsEntry.value:type_name -> v1.RecipeTag
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_v1_get_recipe_proto_init() }
func file_v1_get_recipe_proto_init() {
	if File_v1_get_recipe_proto != nil {
		return
	}
	file_v1_recipe_proto_init()
	file_v1_recipe_category_proto_init()
	file_v1_recipe_tag_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_get_recipe_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecipeRequest); i {
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
		file_v1_get_recipe_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecipeResponse); i {
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
	file_v1_get_recipe_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_get_recipe_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_get_recipe_proto_goTypes,
		DependencyIndexes: file_v1_get_recipe_proto_depIdxs,
		MessageInfos:      file_v1_get_recipe_proto_msgTypes,
	}.Build()
	File_v1_get_recipe_proto = out.File
	file_v1_get_recipe_proto_rawDesc = nil
	file_v1_get_recipe_proto_goTypes = nil
	file_v1_get_recipe_proto_depIdxs = nil
}
