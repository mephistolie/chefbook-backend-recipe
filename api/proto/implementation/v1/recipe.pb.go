// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.2
// source: v1/recipe.proto

package v1

import (
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

type Recipe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecipeId          string                         `protobuf:"bytes,1,opt,name=recipeId,proto3" json:"recipeId,omitempty"`
	Name              string                         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	OwnerId           string                         `protobuf:"bytes,3,opt,name=ownerId,proto3" json:"ownerId,omitempty"`
	OwnerName         *string                        `protobuf:"bytes,4,opt,name=ownerName,proto3,oneof" json:"ownerName,omitempty"`
	OwnerAvatar       *string                        `protobuf:"bytes,5,opt,name=ownerAvatar,proto3,oneof" json:"ownerAvatar,omitempty"`
	IsOwned           bool                           `protobuf:"varint,6,opt,name=isOwned,proto3" json:"isOwned,omitempty"`
	IsSaved           bool                           `protobuf:"varint,7,opt,name=isSaved,proto3" json:"isSaved,omitempty"`
	Visibility        string                         `protobuf:"bytes,8,opt,name=visibility,proto3" json:"visibility,omitempty"`
	IsEncrypted       bool                           `protobuf:"varint,9,opt,name=isEncrypted,proto3" json:"isEncrypted,omitempty"`
	Language          string                         `protobuf:"bytes,10,opt,name=language,proto3" json:"language,omitempty"`
	Translation       *string                        `protobuf:"bytes,11,opt,name=translation,proto3,oneof" json:"translation,omitempty"`
	Translations      map[string]*RecipeTranslations `protobuf:"bytes,12,rep,name=translations,proto3" json:"translations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Description       *string                        `protobuf:"bytes,13,opt,name=description,proto3,oneof" json:"description,omitempty"`
	CreationTimestamp *timestamppb.Timestamp         `protobuf:"bytes,14,opt,name=creationTimestamp,proto3" json:"creationTimestamp,omitempty"`
	UpdateTimestamp   *timestamppb.Timestamp         `protobuf:"bytes,15,opt,name=updateTimestamp,proto3" json:"updateTimestamp,omitempty"`
	Version           int32                          `protobuf:"varint,16,opt,name=version,proto3" json:"version,omitempty"`
	Rating            float32                        `protobuf:"fixed32,17,opt,name=rating,proto3" json:"rating,omitempty"`
	Score             *int32                         `protobuf:"varint,18,opt,name=score,proto3,oneof" json:"score,omitempty"`
	Votes             int32                          `protobuf:"varint,19,opt,name=votes,proto3" json:"votes,omitempty"`
	Tags              []string                       `protobuf:"bytes,20,rep,name=tags,proto3" json:"tags,omitempty"`
	Categories        []string                       `protobuf:"bytes,21,rep,name=categories,proto3" json:"categories,omitempty"`
	IsFavourite       bool                           `protobuf:"varint,22,opt,name=isFavourite,proto3" json:"isFavourite,omitempty"`
	Servings          *int32                         `protobuf:"varint,23,opt,name=servings,proto3,oneof" json:"servings,omitempty"`
	Time              *int32                         `protobuf:"varint,24,opt,name=time,proto3,oneof" json:"time,omitempty"`
	Calories          *int32                         `protobuf:"varint,25,opt,name=calories,proto3,oneof" json:"calories,omitempty"`
	Macronutrients    *Macronutrients                `protobuf:"bytes,26,opt,name=macronutrients,proto3,oneof" json:"macronutrients,omitempty"`
	Ingredients       []*IngredientItem              `protobuf:"bytes,27,rep,name=ingredients,proto3" json:"ingredients,omitempty"`
	Cooking           []*CookingItem                 `protobuf:"bytes,28,rep,name=cooking,proto3" json:"cooking,omitempty"`
	Pictures          *RecipePictures                `protobuf:"bytes,29,opt,name=pictures,proto3" json:"pictures,omitempty"`
}

func (x *Recipe) Reset() {
	*x = Recipe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_recipe_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Recipe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recipe) ProtoMessage() {}

func (x *Recipe) ProtoReflect() protoreflect.Message {
	mi := &file_v1_recipe_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recipe.ProtoReflect.Descriptor instead.
func (*Recipe) Descriptor() ([]byte, []int) {
	return file_v1_recipe_proto_rawDescGZIP(), []int{0}
}

func (x *Recipe) GetRecipeId() string {
	if x != nil {
		return x.RecipeId
	}
	return ""
}

func (x *Recipe) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Recipe) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *Recipe) GetOwnerName() string {
	if x != nil && x.OwnerName != nil {
		return *x.OwnerName
	}
	return ""
}

func (x *Recipe) GetOwnerAvatar() string {
	if x != nil && x.OwnerAvatar != nil {
		return *x.OwnerAvatar
	}
	return ""
}

func (x *Recipe) GetIsOwned() bool {
	if x != nil {
		return x.IsOwned
	}
	return false
}

func (x *Recipe) GetIsSaved() bool {
	if x != nil {
		return x.IsSaved
	}
	return false
}

func (x *Recipe) GetVisibility() string {
	if x != nil {
		return x.Visibility
	}
	return ""
}

func (x *Recipe) GetIsEncrypted() bool {
	if x != nil {
		return x.IsEncrypted
	}
	return false
}

func (x *Recipe) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *Recipe) GetTranslation() string {
	if x != nil && x.Translation != nil {
		return *x.Translation
	}
	return ""
}

func (x *Recipe) GetTranslations() map[string]*RecipeTranslations {
	if x != nil {
		return x.Translations
	}
	return nil
}

func (x *Recipe) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *Recipe) GetCreationTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.CreationTimestamp
	}
	return nil
}

func (x *Recipe) GetUpdateTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTimestamp
	}
	return nil
}

func (x *Recipe) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Recipe) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Recipe) GetScore() int32 {
	if x != nil && x.Score != nil {
		return *x.Score
	}
	return 0
}

func (x *Recipe) GetVotes() int32 {
	if x != nil {
		return x.Votes
	}
	return 0
}

func (x *Recipe) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Recipe) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *Recipe) GetIsFavourite() bool {
	if x != nil {
		return x.IsFavourite
	}
	return false
}

func (x *Recipe) GetServings() int32 {
	if x != nil && x.Servings != nil {
		return *x.Servings
	}
	return 0
}

func (x *Recipe) GetTime() int32 {
	if x != nil && x.Time != nil {
		return *x.Time
	}
	return 0
}

func (x *Recipe) GetCalories() int32 {
	if x != nil && x.Calories != nil {
		return *x.Calories
	}
	return 0
}

func (x *Recipe) GetMacronutrients() *Macronutrients {
	if x != nil {
		return x.Macronutrients
	}
	return nil
}

func (x *Recipe) GetIngredients() []*IngredientItem {
	if x != nil {
		return x.Ingredients
	}
	return nil
}

func (x *Recipe) GetCooking() []*CookingItem {
	if x != nil {
		return x.Cooking
	}
	return nil
}

func (x *Recipe) GetPictures() *RecipePictures {
	if x != nil {
		return x.Pictures
	}
	return nil
}

type RecipeTranslations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Translations []*RecipeTranslationInfo `protobuf:"bytes,1,rep,name=translations,proto3" json:"translations,omitempty"`
}

func (x *RecipeTranslations) Reset() {
	*x = RecipeTranslations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_recipe_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecipeTranslations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecipeTranslations) ProtoMessage() {}

func (x *RecipeTranslations) ProtoReflect() protoreflect.Message {
	mi := &file_v1_recipe_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecipeTranslations.ProtoReflect.Descriptor instead.
func (*RecipeTranslations) Descriptor() ([]byte, []int) {
	return file_v1_recipe_proto_rawDescGZIP(), []int{1}
}

func (x *RecipeTranslations) GetTranslations() []*RecipeTranslationInfo {
	if x != nil {
		return x.Translations
	}
	return nil
}

type RecipeTranslationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorId     string  `protobuf:"bytes,1,opt,name=authorId,proto3" json:"authorId,omitempty"`
	AuthorName   *string `protobuf:"bytes,2,opt,name=authorName,proto3,oneof" json:"authorName,omitempty"`
	AuthorAvatar *string `protobuf:"bytes,3,opt,name=authorAvatar,proto3,oneof" json:"authorAvatar,omitempty"`
}

func (x *RecipeTranslationInfo) Reset() {
	*x = RecipeTranslationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_recipe_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecipeTranslationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecipeTranslationInfo) ProtoMessage() {}

func (x *RecipeTranslationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_v1_recipe_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecipeTranslationInfo.ProtoReflect.Descriptor instead.
func (*RecipeTranslationInfo) Descriptor() ([]byte, []int) {
	return file_v1_recipe_proto_rawDescGZIP(), []int{2}
}

func (x *RecipeTranslationInfo) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *RecipeTranslationInfo) GetAuthorName() string {
	if x != nil && x.AuthorName != nil {
		return *x.AuthorName
	}
	return ""
}

func (x *RecipeTranslationInfo) GetAuthorAvatar() string {
	if x != nil && x.AuthorAvatar != nil {
		return *x.AuthorAvatar
	}
	return ""
}

var File_v1_recipe_proto protoreflect.FileDescriptor

var file_v1_recipe_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x63, 0x72, 0x6f,
	0x6e, 0x75, 0x74, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x13, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x0a, 0x0a, 0x06, 0x52, 0x65,
	0x63, 0x69, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x25, 0x0a, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x4f, 0x77,
	0x6e, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x4f, 0x77, 0x6e,
	0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x53, 0x61, 0x76, 0x65, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x53, 0x61, 0x76, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x20, 0x0a, 0x0b,
	0x69, 0x73, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x69, 0x73, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x02, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01,
	0x01, 0x12, 0x40, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x48, 0x0a, 0x11, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x44, 0x0a, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x19, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05, 0x48, 0x04, 0x52, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x88, 0x01, 0x01, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x73,
	0x18, 0x13, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x15, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x69, 0x74, 0x65,
	0x18, 0x16, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x75, 0x72,
	0x69, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x17, 0x20, 0x01, 0x28, 0x05, 0x48, 0x05, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67,
	0x73, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x18, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x06, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a,
	0x08, 0x63, 0x61, 0x6c, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x19, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x07, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x3f,
	0x0a, 0x0e, 0x6d, 0x61, 0x63, 0x72, 0x6f, 0x6e, 0x75, 0x74, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x1a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x63, 0x72,
	0x6f, 0x6e, 0x75, 0x74, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x48, 0x08, 0x52, 0x0e, 0x6d, 0x61,
	0x63, 0x72, 0x6f, 0x6e, 0x75, 0x74, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x88, 0x01, 0x01, 0x12,
	0x34, 0x0a, 0x0b, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x1b,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0b, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x29, 0x0a, 0x07, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x18, 0x1c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x07, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x12, 0x2e, 0x0a, 0x08, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x1d, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x50, 0x69,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52, 0x08, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x1a, 0x57, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x69,
	0x70, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x63, 0x61, 0x6c, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x6d, 0x61, 0x63, 0x72, 0x6f, 0x6e, 0x75,
	0x74, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x53, 0x0a, 0x12, 0x52, 0x65, 0x63, 0x69, 0x70,
	0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3d, 0x0a,
	0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xa1, 0x01, 0x0a,
	0x15, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x49, 0x64, 0x12, 0x23, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x0c, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x88, 0x01, 0x01,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x42,
	0x0f, 0x0a, 0x0d, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x65, 0x70, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x6c, 0x69, 0x65, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x62,
	0x6f, 0x6f, 0x6b, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x72, 0x65, 0x63, 0x69,
	0x70, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_recipe_proto_rawDescOnce sync.Once
	file_v1_recipe_proto_rawDescData = file_v1_recipe_proto_rawDesc
)

func file_v1_recipe_proto_rawDescGZIP() []byte {
	file_v1_recipe_proto_rawDescOnce.Do(func() {
		file_v1_recipe_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_recipe_proto_rawDescData)
	})
	return file_v1_recipe_proto_rawDescData
}

var file_v1_recipe_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v1_recipe_proto_goTypes = []interface{}{
	(*Recipe)(nil),                // 0: v1.Recipe
	(*RecipeTranslations)(nil),    // 1: v1.RecipeTranslations
	(*RecipeTranslationInfo)(nil), // 2: v1.RecipeTranslationInfo
	nil,                           // 3: v1.Recipe.TranslationsEntry
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*Macronutrients)(nil),        // 5: v1.Macronutrients
	(*IngredientItem)(nil),        // 6: v1.IngredientItem
	(*CookingItem)(nil),           // 7: v1.CookingItem
	(*RecipePictures)(nil),        // 8: v1.RecipePictures
}
var file_v1_recipe_proto_depIdxs = []int32{
	3, // 0: v1.Recipe.translations:type_name -> v1.Recipe.TranslationsEntry
	4, // 1: v1.Recipe.creationTimestamp:type_name -> google.protobuf.Timestamp
	4, // 2: v1.Recipe.updateTimestamp:type_name -> google.protobuf.Timestamp
	5, // 3: v1.Recipe.macronutrients:type_name -> v1.Macronutrients
	6, // 4: v1.Recipe.ingredients:type_name -> v1.IngredientItem
	7, // 5: v1.Recipe.cooking:type_name -> v1.CookingItem
	8, // 6: v1.Recipe.pictures:type_name -> v1.RecipePictures
	2, // 7: v1.RecipeTranslations.translations:type_name -> v1.RecipeTranslationInfo
	1, // 8: v1.Recipe.TranslationsEntry.value:type_name -> v1.RecipeTranslations
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_v1_recipe_proto_init() }
func file_v1_recipe_proto_init() {
	if File_v1_recipe_proto != nil {
		return
	}
	file_v1_macronutrients_proto_init()
	file_v1_ingredient_proto_init()
	file_v1_cooking_proto_init()
	file_v1_pictures_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_recipe_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Recipe); i {
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
		file_v1_recipe_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecipeTranslations); i {
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
		file_v1_recipe_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecipeTranslationInfo); i {
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
	file_v1_recipe_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_v1_recipe_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_recipe_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_recipe_proto_goTypes,
		DependencyIndexes: file_v1_recipe_proto_depIdxs,
		MessageInfos:      file_v1_recipe_proto_msgTypes,
	}.Build()
	File_v1_recipe_proto = out.File
	file_v1_recipe_proto_rawDesc = nil
	file_v1_recipe_proto_goTypes = nil
	file_v1_recipe_proto_depIdxs = nil
}
