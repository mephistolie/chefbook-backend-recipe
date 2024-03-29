// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.2
// source: v1/get-recipe-book.proto

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

type RecipeState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecipeId     string   `protobuf:"bytes,1,opt,name=recipeId,proto3" json:"recipeId,omitempty"`
	Version      int32    `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	OwnerName    *string  `protobuf:"bytes,3,opt,name=ownerName,proto3,oneof" json:"ownerName,omitempty"`
	OwnerAvatar  *string  `protobuf:"bytes,4,opt,name=ownerAvatar,proto3,oneof" json:"ownerAvatar,omitempty"`
	Translations []string `protobuf:"bytes,5,rep,name=translations,proto3" json:"translations,omitempty"`
	Rating       float32  `protobuf:"fixed32,6,opt,name=rating,proto3" json:"rating,omitempty"`
	Score        *int32   `protobuf:"varint,7,opt,name=score,proto3,oneof" json:"score,omitempty"`
	Votes        int32    `protobuf:"varint,8,opt,name=votes,proto3" json:"votes,omitempty"`
	Categories   []string `protobuf:"bytes,9,rep,name=categories,proto3" json:"categories,omitempty"`
	IsFavourite  bool     `protobuf:"varint,10,opt,name=isFavourite,proto3" json:"isFavourite,omitempty"`
	Tags         []string `protobuf:"bytes,11,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *RecipeState) Reset() {
	*x = RecipeState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_get_recipe_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecipeState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecipeState) ProtoMessage() {}

func (x *RecipeState) ProtoReflect() protoreflect.Message {
	mi := &file_v1_get_recipe_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecipeState.ProtoReflect.Descriptor instead.
func (*RecipeState) Descriptor() ([]byte, []int) {
	return file_v1_get_recipe_book_proto_rawDescGZIP(), []int{0}
}

func (x *RecipeState) GetRecipeId() string {
	if x != nil {
		return x.RecipeId
	}
	return ""
}

func (x *RecipeState) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *RecipeState) GetOwnerName() string {
	if x != nil && x.OwnerName != nil {
		return *x.OwnerName
	}
	return ""
}

func (x *RecipeState) GetOwnerAvatar() string {
	if x != nil && x.OwnerAvatar != nil {
		return *x.OwnerAvatar
	}
	return ""
}

func (x *RecipeState) GetTranslations() []string {
	if x != nil {
		return x.Translations
	}
	return nil
}

func (x *RecipeState) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *RecipeState) GetScore() int32 {
	if x != nil && x.Score != nil {
		return *x.Score
	}
	return 0
}

func (x *RecipeState) GetVotes() int32 {
	if x != nil {
		return x.Votes
	}
	return 0
}

func (x *RecipeState) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *RecipeState) GetIsFavourite() bool {
	if x != nil {
		return x.IsFavourite
	}
	return false
}

func (x *RecipeState) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type GetRecipeBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string  `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	UserLanguage *string `protobuf:"bytes,2,opt,name=userLanguage,proto3,oneof" json:"userLanguage,omitempty"`
}

func (x *GetRecipeBookRequest) Reset() {
	*x = GetRecipeBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_get_recipe_book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecipeBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecipeBookRequest) ProtoMessage() {}

func (x *GetRecipeBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_get_recipe_book_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecipeBookRequest.ProtoReflect.Descriptor instead.
func (*GetRecipeBookRequest) Descriptor() ([]byte, []int) {
	return file_v1_get_recipe_book_proto_rawDescGZIP(), []int{1}
}

func (x *GetRecipeBookRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetRecipeBookRequest) GetUserLanguage() string {
	if x != nil && x.UserLanguage != nil {
		return *x.UserLanguage
	}
	return ""
}

type GetRecipeBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recipes           []*RecipeState        `protobuf:"bytes,1,rep,name=recipes,proto3" json:"recipes,omitempty"`
	Categories        []*RecipeCategory     `protobuf:"bytes,2,rep,name=categories,proto3" json:"categories,omitempty"`
	Tags              map[string]*RecipeTag `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	TagGroups         map[string]string     `protobuf:"bytes,4,rep,name=tagGroups,proto3" json:"tagGroups,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	HasEncryptedVault bool                  `protobuf:"varint,5,opt,name=hasEncryptedVault,proto3" json:"hasEncryptedVault,omitempty"`
}

func (x *GetRecipeBookResponse) Reset() {
	*x = GetRecipeBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_get_recipe_book_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecipeBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecipeBookResponse) ProtoMessage() {}

func (x *GetRecipeBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_get_recipe_book_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecipeBookResponse.ProtoReflect.Descriptor instead.
func (*GetRecipeBookResponse) Descriptor() ([]byte, []int) {
	return file_v1_get_recipe_book_proto_rawDescGZIP(), []int{2}
}

func (x *GetRecipeBookResponse) GetRecipes() []*RecipeState {
	if x != nil {
		return x.Recipes
	}
	return nil
}

func (x *GetRecipeBookResponse) GetCategories() []*RecipeCategory {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *GetRecipeBookResponse) GetTags() map[string]*RecipeTag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *GetRecipeBookResponse) GetTagGroups() map[string]string {
	if x != nil {
		return x.TagGroups
	}
	return nil
}

func (x *GetRecipeBookResponse) GetHasEncryptedVault() bool {
	if x != nil {
		return x.HasEncryptedVault
	}
	return false
}

var File_v1_get_recipe_book_proto protoreflect.FileDescriptor

var file_v1_get_recipe_book_proto_rawDesc = []byte{
	0x0a, 0x18, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2d,
	0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x18,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x2d, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x02,
	0x0a, 0x0b, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a,
	0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x19, 0x0a, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73,
	0x46, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x69, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x69, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x69, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x0e,
	0x0a, 0x0c, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x68, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72,
	0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x88, 0x01,
	0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x22, 0xab, 0x03, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65,
	0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x07,
	0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x07,
	0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x12, 0x46, 0x0a, 0x09, 0x74, 0x61, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x54, 0x61, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x09, 0x74, 0x61, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x2c, 0x0a, 0x11,
	0x68, 0x61, 0x73, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x56, 0x61, 0x75, 0x6c,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x68, 0x61, 0x73, 0x45, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x65, 0x64, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x1a, 0x46, 0x0a, 0x09, 0x54, 0x61,
	0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x63, 0x69, 0x70, 0x65, 0x54, 0x61, 0x67, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x3c, 0x0a, 0x0e, 0x54, 0x61, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x65, 0x70, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x6c, 0x69, 0x65, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x62,
	0x6f, 0x6f, 0x6b, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x72, 0x65, 0x63, 0x69,
	0x70, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_get_recipe_book_proto_rawDescOnce sync.Once
	file_v1_get_recipe_book_proto_rawDescData = file_v1_get_recipe_book_proto_rawDesc
)

func file_v1_get_recipe_book_proto_rawDescGZIP() []byte {
	file_v1_get_recipe_book_proto_rawDescOnce.Do(func() {
		file_v1_get_recipe_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_get_recipe_book_proto_rawDescData)
	})
	return file_v1_get_recipe_book_proto_rawDescData
}

var file_v1_get_recipe_book_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_v1_get_recipe_book_proto_goTypes = []interface{}{
	(*RecipeState)(nil),           // 0: v1.RecipeState
	(*GetRecipeBookRequest)(nil),  // 1: v1.GetRecipeBookRequest
	(*GetRecipeBookResponse)(nil), // 2: v1.GetRecipeBookResponse
	nil,                           // 3: v1.GetRecipeBookResponse.TagsEntry
	nil,                           // 4: v1.GetRecipeBookResponse.TagGroupsEntry
	(*RecipeCategory)(nil),        // 5: v1.RecipeCategory
	(*RecipeTag)(nil),             // 6: v1.RecipeTag
}
var file_v1_get_recipe_book_proto_depIdxs = []int32{
	0, // 0: v1.GetRecipeBookResponse.recipes:type_name -> v1.RecipeState
	5, // 1: v1.GetRecipeBookResponse.categories:type_name -> v1.RecipeCategory
	3, // 2: v1.GetRecipeBookResponse.tags:type_name -> v1.GetRecipeBookResponse.TagsEntry
	4, // 3: v1.GetRecipeBookResponse.tagGroups:type_name -> v1.GetRecipeBookResponse.TagGroupsEntry
	6, // 4: v1.GetRecipeBookResponse.TagsEntry.value:type_name -> v1.RecipeTag
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_v1_get_recipe_book_proto_init() }
func file_v1_get_recipe_book_proto_init() {
	if File_v1_get_recipe_book_proto != nil {
		return
	}
	file_v1_recipe_category_proto_init()
	file_v1_recipe_tag_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_get_recipe_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecipeState); i {
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
		file_v1_get_recipe_book_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecipeBookRequest); i {
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
		file_v1_get_recipe_book_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecipeBookResponse); i {
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
	file_v1_get_recipe_book_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_v1_get_recipe_book_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_get_recipe_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_get_recipe_book_proto_goTypes,
		DependencyIndexes: file_v1_get_recipe_book_proto_depIdxs,
		MessageInfos:      file_v1_get_recipe_book_proto_msgTypes,
	}.Build()
	File_v1_get_recipe_book_proto = out.File
	file_v1_get_recipe_book_proto_rawDesc = nil
	file_v1_get_recipe_book_proto_goTypes = nil
	file_v1_get_recipe_book_proto_depIdxs = nil
}
