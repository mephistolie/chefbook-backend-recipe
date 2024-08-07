// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.27.0
// source: v1/generate-recipe-pictures-upload-links.proto

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

type RecipePictureUploadLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PictureLink string            `protobuf:"bytes,1,opt,name=pictureLink,proto3" json:"pictureLink,omitempty"`
	UploadLink  string            `protobuf:"bytes,2,opt,name=uploadLink,proto3" json:"uploadLink,omitempty"`
	FormData    map[string]string `protobuf:"bytes,3,rep,name=formData,proto3" json:"formData,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MaxSize     int64             `protobuf:"varint,4,opt,name=maxSize,proto3" json:"maxSize,omitempty"`
}

func (x *RecipePictureUploadLink) Reset() {
	*x = RecipePictureUploadLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_generate_recipe_pictures_upload_links_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecipePictureUploadLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecipePictureUploadLink) ProtoMessage() {}

func (x *RecipePictureUploadLink) ProtoReflect() protoreflect.Message {
	mi := &file_v1_generate_recipe_pictures_upload_links_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecipePictureUploadLink.ProtoReflect.Descriptor instead.
func (*RecipePictureUploadLink) Descriptor() ([]byte, []int) {
	return file_v1_generate_recipe_pictures_upload_links_proto_rawDescGZIP(), []int{0}
}

func (x *RecipePictureUploadLink) GetPictureLink() string {
	if x != nil {
		return x.PictureLink
	}
	return ""
}

func (x *RecipePictureUploadLink) GetUploadLink() string {
	if x != nil {
		return x.UploadLink
	}
	return ""
}

func (x *RecipePictureUploadLink) GetFormData() map[string]string {
	if x != nil {
		return x.FormData
	}
	return nil
}

func (x *RecipePictureUploadLink) GetMaxSize() int64 {
	if x != nil {
		return x.MaxSize
	}
	return 0
}

type GenerateRecipePicturesUploadLinksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecipeId      string `protobuf:"bytes,1,opt,name=recipeId,proto3" json:"recipeId,omitempty"`
	UserId        string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	PicturesCount int32  `protobuf:"varint,3,opt,name=picturesCount,proto3" json:"picturesCount,omitempty"`
	Subscription  string `protobuf:"bytes,4,opt,name=subscription,proto3" json:"subscription,omitempty"`
}

func (x *GenerateRecipePicturesUploadLinksRequest) Reset() {
	*x = GenerateRecipePicturesUploadLinksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_generate_recipe_pictures_upload_links_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateRecipePicturesUploadLinksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateRecipePicturesUploadLinksRequest) ProtoMessage() {}

func (x *GenerateRecipePicturesUploadLinksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_generate_recipe_pictures_upload_links_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateRecipePicturesUploadLinksRequest.ProtoReflect.Descriptor instead.
func (*GenerateRecipePicturesUploadLinksRequest) Descriptor() ([]byte, []int) {
	return file_v1_generate_recipe_pictures_upload_links_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateRecipePicturesUploadLinksRequest) GetRecipeId() string {
	if x != nil {
		return x.RecipeId
	}
	return ""
}

func (x *GenerateRecipePicturesUploadLinksRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GenerateRecipePicturesUploadLinksRequest) GetPicturesCount() int32 {
	if x != nil {
		return x.PicturesCount
	}
	return 0
}

func (x *GenerateRecipePicturesUploadLinksRequest) GetSubscription() string {
	if x != nil {
		return x.Subscription
	}
	return ""
}

type GenerateRecipePicturesUploadLinksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Links []*RecipePictureUploadLink `protobuf:"bytes,1,rep,name=links,proto3" json:"links,omitempty"`
}

func (x *GenerateRecipePicturesUploadLinksResponse) Reset() {
	*x = GenerateRecipePicturesUploadLinksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_generate_recipe_pictures_upload_links_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateRecipePicturesUploadLinksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateRecipePicturesUploadLinksResponse) ProtoMessage() {}

func (x *GenerateRecipePicturesUploadLinksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_generate_recipe_pictures_upload_links_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateRecipePicturesUploadLinksResponse.ProtoReflect.Descriptor instead.
func (*GenerateRecipePicturesUploadLinksResponse) Descriptor() ([]byte, []int) {
	return file_v1_generate_recipe_pictures_upload_links_proto_rawDescGZIP(), []int{2}
}

func (x *GenerateRecipePicturesUploadLinksResponse) GetLinks() []*RecipePictureUploadLink {
	if x != nil {
		return x.Links
	}
	return nil
}

var File_v1_generate_recipe_pictures_upload_links_proto protoreflect.FileDescriptor

var file_v1_generate_recipe_pictures_upload_links_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2d, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x65, 0x2d, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2d, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x2d, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x76, 0x31, 0x22, 0xf9, 0x01, 0x0a, 0x17, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x50,
	0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b,
	0x12, 0x20, 0x0a, 0x0b, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x4c, 0x69,
	0x6e, 0x6b, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69,
	0x6e, 0x6b, 0x12, 0x45, 0x0a, 0x08, 0x66, 0x6f, 0x72, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65,
	0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e,
	0x6b, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x08, 0x66, 0x6f, 0x72, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x78,
	0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x53,
	0x69, 0x7a, 0x65, 0x1a, 0x3b, 0x0a, 0x0d, 0x46, 0x6f, 0x72, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xa8, 0x01, 0x0a, 0x28, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5e, 0x0a, 0x29, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x50, 0x69, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x42, 0x3d, 0x5a, 0x3b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x70, 0x68, 0x69, 0x73,
	0x74, 0x6f, 0x6c, 0x69, 0x65, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x62, 0x6f, 0x6f, 0x6b, 0x2d, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_v1_generate_recipe_pictures_upload_links_proto_rawDescOnce sync.Once
	file_v1_generate_recipe_pictures_upload_links_proto_rawDescData = file_v1_generate_recipe_pictures_upload_links_proto_rawDesc
)

func file_v1_generate_recipe_pictures_upload_links_proto_rawDescGZIP() []byte {
	file_v1_generate_recipe_pictures_upload_links_proto_rawDescOnce.Do(func() {
		file_v1_generate_recipe_pictures_upload_links_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_generate_recipe_pictures_upload_links_proto_rawDescData)
	})
	return file_v1_generate_recipe_pictures_upload_links_proto_rawDescData
}

var file_v1_generate_recipe_pictures_upload_links_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v1_generate_recipe_pictures_upload_links_proto_goTypes = []interface{}{
	(*RecipePictureUploadLink)(nil),                   // 0: v1.RecipePictureUploadLink
	(*GenerateRecipePicturesUploadLinksRequest)(nil),  // 1: v1.GenerateRecipePicturesUploadLinksRequest
	(*GenerateRecipePicturesUploadLinksResponse)(nil), // 2: v1.GenerateRecipePicturesUploadLinksResponse
	nil, // 3: v1.RecipePictureUploadLink.FormDataEntry
}
var file_v1_generate_recipe_pictures_upload_links_proto_depIdxs = []int32{
	3, // 0: v1.RecipePictureUploadLink.formData:type_name -> v1.RecipePictureUploadLink.FormDataEntry
	0, // 1: v1.GenerateRecipePicturesUploadLinksResponse.links:type_name -> v1.RecipePictureUploadLink
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_generate_recipe_pictures_upload_links_proto_init() }
func file_v1_generate_recipe_pictures_upload_links_proto_init() {
	if File_v1_generate_recipe_pictures_upload_links_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_generate_recipe_pictures_upload_links_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecipePictureUploadLink); i {
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
		file_v1_generate_recipe_pictures_upload_links_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateRecipePicturesUploadLinksRequest); i {
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
		file_v1_generate_recipe_pictures_upload_links_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateRecipePicturesUploadLinksResponse); i {
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
			RawDescriptor: file_v1_generate_recipe_pictures_upload_links_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_generate_recipe_pictures_upload_links_proto_goTypes,
		DependencyIndexes: file_v1_generate_recipe_pictures_upload_links_proto_depIdxs,
		MessageInfos:      file_v1_generate_recipe_pictures_upload_links_proto_msgTypes,
	}.Build()
	File_v1_generate_recipe_pictures_upload_links_proto = out.File
	file_v1_generate_recipe_pictures_upload_links_proto_rawDesc = nil
	file_v1_generate_recipe_pictures_upload_links_proto_goTypes = nil
	file_v1_generate_recipe_pictures_upload_links_proto_depIdxs = nil
}
