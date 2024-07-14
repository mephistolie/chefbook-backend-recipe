// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.27.0
// source: v1/translate-recipe.proto

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

type IngredientTranslation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string  `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Unit *string `protobuf:"bytes,2,opt,name=unit,proto3,oneof" json:"unit,omitempty"`
}

func (x *IngredientTranslation) Reset() {
	*x = IngredientTranslation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_translate_recipe_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngredientTranslation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngredientTranslation) ProtoMessage() {}

func (x *IngredientTranslation) ProtoReflect() protoreflect.Message {
	mi := &file_v1_translate_recipe_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngredientTranslation.ProtoReflect.Descriptor instead.
func (*IngredientTranslation) Descriptor() ([]byte, []int) {
	return file_v1_translate_recipe_proto_rawDescGZIP(), []int{0}
}

func (x *IngredientTranslation) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *IngredientTranslation) GetUnit() string {
	if x != nil && x.Unit != nil {
		return *x.Unit
	}
	return ""
}

type TranslateRecipeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecipeId     string                            `protobuf:"bytes,1,opt,name=recipeId,proto3" json:"recipeId,omitempty"`
	TranslatorId string                            `protobuf:"bytes,2,opt,name=translatorId,proto3" json:"translatorId,omitempty"`
	Language     string                            `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"`
	Name         string                            `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description  *string                           `protobuf:"bytes,5,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Ingredients  map[string]*IngredientTranslation `protobuf:"bytes,6,rep,name=ingredients,proto3" json:"ingredients,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Cooking      map[string]string                 `protobuf:"bytes,7,rep,name=cooking,proto3" json:"cooking,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TranslateRecipeRequest) Reset() {
	*x = TranslateRecipeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_translate_recipe_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranslateRecipeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslateRecipeRequest) ProtoMessage() {}

func (x *TranslateRecipeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_translate_recipe_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslateRecipeRequest.ProtoReflect.Descriptor instead.
func (*TranslateRecipeRequest) Descriptor() ([]byte, []int) {
	return file_v1_translate_recipe_proto_rawDescGZIP(), []int{1}
}

func (x *TranslateRecipeRequest) GetRecipeId() string {
	if x != nil {
		return x.RecipeId
	}
	return ""
}

func (x *TranslateRecipeRequest) GetTranslatorId() string {
	if x != nil {
		return x.TranslatorId
	}
	return ""
}

func (x *TranslateRecipeRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *TranslateRecipeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TranslateRecipeRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *TranslateRecipeRequest) GetIngredients() map[string]*IngredientTranslation {
	if x != nil {
		return x.Ingredients
	}
	return nil
}

func (x *TranslateRecipeRequest) GetCooking() map[string]string {
	if x != nil {
		return x.Cooking
	}
	return nil
}

type TranslateRecipeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TranslateRecipeResponse) Reset() {
	*x = TranslateRecipeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_translate_recipe_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranslateRecipeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslateRecipeResponse) ProtoMessage() {}

func (x *TranslateRecipeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_translate_recipe_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslateRecipeResponse.ProtoReflect.Descriptor instead.
func (*TranslateRecipeResponse) Descriptor() ([]byte, []int) {
	return file_v1_translate_recipe_proto_rawDescGZIP(), []int{2}
}

func (x *TranslateRecipeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_v1_translate_recipe_proto protoreflect.FileDescriptor

var file_v1_translate_recipe_proto_rawDesc = []byte{
	0x0a, 0x19, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x2d, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x22,
	0x4d, 0x0a, 0x15, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x17, 0x0a, 0x04,
	0x75, 0x6e, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x75, 0x6e,
	0x69, 0x74, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x22, 0xe8,
	0x03, 0x0a, 0x16, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x69,
	0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61,
	0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01,
	0x12, 0x4d, 0x0a, 0x0b, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0b, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x41, 0x0a, 0x07, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x1a, 0x59, 0x0a, 0x10, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x67,
	0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a,
	0x0c, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x33, 0x0a, 0x17, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x3d,
	0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x70,
	0x68, 0x69, 0x73, 0x74, 0x6f, 0x6c, 0x69, 0x65, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x62, 0x6f, 0x6f,
	0x6b, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_translate_recipe_proto_rawDescOnce sync.Once
	file_v1_translate_recipe_proto_rawDescData = file_v1_translate_recipe_proto_rawDesc
)

func file_v1_translate_recipe_proto_rawDescGZIP() []byte {
	file_v1_translate_recipe_proto_rawDescOnce.Do(func() {
		file_v1_translate_recipe_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_translate_recipe_proto_rawDescData)
	})
	return file_v1_translate_recipe_proto_rawDescData
}

var file_v1_translate_recipe_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_v1_translate_recipe_proto_goTypes = []interface{}{
	(*IngredientTranslation)(nil),   // 0: v1.IngredientTranslation
	(*TranslateRecipeRequest)(nil),  // 1: v1.TranslateRecipeRequest
	(*TranslateRecipeResponse)(nil), // 2: v1.TranslateRecipeResponse
	nil,                             // 3: v1.TranslateRecipeRequest.IngredientsEntry
	nil,                             // 4: v1.TranslateRecipeRequest.CookingEntry
}
var file_v1_translate_recipe_proto_depIdxs = []int32{
	3, // 0: v1.TranslateRecipeRequest.ingredients:type_name -> v1.TranslateRecipeRequest.IngredientsEntry
	4, // 1: v1.TranslateRecipeRequest.cooking:type_name -> v1.TranslateRecipeRequest.CookingEntry
	0, // 2: v1.TranslateRecipeRequest.IngredientsEntry.value:type_name -> v1.IngredientTranslation
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v1_translate_recipe_proto_init() }
func file_v1_translate_recipe_proto_init() {
	if File_v1_translate_recipe_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_translate_recipe_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngredientTranslation); i {
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
		file_v1_translate_recipe_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranslateRecipeRequest); i {
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
		file_v1_translate_recipe_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranslateRecipeResponse); i {
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
	file_v1_translate_recipe_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_v1_translate_recipe_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_translate_recipe_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_translate_recipe_proto_goTypes,
		DependencyIndexes: file_v1_translate_recipe_proto_depIdxs,
		MessageInfos:      file_v1_translate_recipe_proto_msgTypes,
	}.Build()
	File_v1_translate_recipe_proto = out.File
	file_v1_translate_recipe_proto_rawDesc = nil
	file_v1_translate_recipe_proto_goTypes = nil
	file_v1_translate_recipe_proto_depIdxs = nil
}
