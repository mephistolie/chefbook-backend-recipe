// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: v1/recipe-service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RecipeService_GetRecipes_FullMethodName                        = "/v1.RecipeService/GetRecipes"
	RecipeService_GetRandomRecipe_FullMethodName                   = "/v1.RecipeService/GetRandomRecipe"
	RecipeService_GetRecipeBook_FullMethodName                     = "/v1.RecipeService/GetRecipeBook"
	RecipeService_CreateRecipe_FullMethodName                      = "/v1.RecipeService/CreateRecipe"
	RecipeService_GetRecipe_FullMethodName                         = "/v1.RecipeService/GetRecipe"
	RecipeService_UpdateRecipe_FullMethodName                      = "/v1.RecipeService/UpdateRecipe"
	RecipeService_DeleteRecipe_FullMethodName                      = "/v1.RecipeService/DeleteRecipe"
	RecipeService_GenerateRecipePicturesUploadLinks_FullMethodName = "/v1.RecipeService/GenerateRecipePicturesUploadLinks"
	RecipeService_SetRecipePictures_FullMethodName                 = "/v1.RecipeService/SetRecipePictures"
	RecipeService_RateRecipe_FullMethodName                        = "/v1.RecipeService/RateRecipe"
	RecipeService_SaveToRecipeBook_FullMethodName                  = "/v1.RecipeService/SaveToRecipeBook"
	RecipeService_RemoveFromRecipeBook_FullMethodName              = "/v1.RecipeService/RemoveFromRecipeBook"
	RecipeService_SetRecipeFavouriteStatus_FullMethodName          = "/v1.RecipeService/SetRecipeFavouriteStatus"
	RecipeService_SetRecipeCategories_FullMethodName               = "/v1.RecipeService/SetRecipeCategories"
	RecipeService_TranslateRecipe_FullMethodName                   = "/v1.RecipeService/TranslateRecipe"
	RecipeService_DeleteRecipeTranslation_FullMethodName           = "/v1.RecipeService/DeleteRecipeTranslation"
	RecipeService_GetRecipePolicy_FullMethodName                   = "/v1.RecipeService/GetRecipePolicy"
	RecipeService_GetRecipeNames_FullMethodName                    = "/v1.RecipeService/GetRecipeNames"
)

// RecipeServiceClient is the client API for RecipeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecipeServiceClient interface {
	GetRecipes(ctx context.Context, in *GetRecipesRequest, opts ...grpc.CallOption) (*GetRecipesResponse, error)
	GetRandomRecipe(ctx context.Context, in *GetRandomRecipeRequest, opts ...grpc.CallOption) (*GetRecipeResponse, error)
	GetRecipeBook(ctx context.Context, in *GetRecipeBookRequest, opts ...grpc.CallOption) (*GetRecipeBookResponse, error)
	CreateRecipe(ctx context.Context, in *RecipeInput, opts ...grpc.CallOption) (*CreateRecipeResponse, error)
	GetRecipe(ctx context.Context, in *GetRecipeRequest, opts ...grpc.CallOption) (*GetRecipeResponse, error)
	UpdateRecipe(ctx context.Context, in *RecipeInput, opts ...grpc.CallOption) (*UpdateRecipeResponse, error)
	DeleteRecipe(ctx context.Context, in *DeleteRecipeRequest, opts ...grpc.CallOption) (*DeleteRecipeResponse, error)
	GenerateRecipePicturesUploadLinks(ctx context.Context, in *GenerateRecipePicturesUploadLinksRequest, opts ...grpc.CallOption) (*GenerateRecipePicturesUploadLinksResponse, error)
	SetRecipePictures(ctx context.Context, in *SetRecipePicturesRequest, opts ...grpc.CallOption) (*SetRecipePicturesResponse, error)
	RateRecipe(ctx context.Context, in *RateRecipeRequest, opts ...grpc.CallOption) (*RateRecipeResponse, error)
	SaveToRecipeBook(ctx context.Context, in *SaveToRecipeBookRequest, opts ...grpc.CallOption) (*SaveToRecipeBookResponse, error)
	RemoveFromRecipeBook(ctx context.Context, in *RemoveFromRecipeBookRequest, opts ...grpc.CallOption) (*RemoveFromRecipeBookResponse, error)
	SetRecipeFavouriteStatus(ctx context.Context, in *SetRecipeFavouriteStatusRequest, opts ...grpc.CallOption) (*SetRecipeFavouriteStatusResponse, error)
	SetRecipeCategories(ctx context.Context, in *SetRecipeCategoriesRequest, opts ...grpc.CallOption) (*SetRecipeCategoriesResponse, error)
	TranslateRecipe(ctx context.Context, in *TranslateRecipeRequest, opts ...grpc.CallOption) (*TranslateRecipeResponse, error)
	DeleteRecipeTranslation(ctx context.Context, in *DeleteRecipeTranslationRequest, opts ...grpc.CallOption) (*DeleteRecipeTranslationResponse, error)
	GetRecipePolicy(ctx context.Context, in *GetRecipePolicyRequest, opts ...grpc.CallOption) (*GetRecipePolicyResponse, error)
	GetRecipeNames(ctx context.Context, in *GetRecipeNamesRequest, opts ...grpc.CallOption) (*GetRecipeNamesResponse, error)
}

type recipeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRecipeServiceClient(cc grpc.ClientConnInterface) RecipeServiceClient {
	return &recipeServiceClient{cc}
}

func (c *recipeServiceClient) GetRecipes(ctx context.Context, in *GetRecipesRequest, opts ...grpc.CallOption) (*GetRecipesResponse, error) {
	out := new(GetRecipesResponse)
	err := c.cc.Invoke(ctx, RecipeService_GetRecipes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeServiceClient) GetRandomRecipe(ctx context.Context, in *GetRandomRecipeRequest, opts ...grpc.CallOption) (*GetRecipeResponse, error) {
	out := new(GetRecipeResponse)
	err := c.cc.Invoke(ctx, RecipeService_GetRandomRecipe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeServiceClient) GetRecipeBook(ctx context.Context, in *GetRecipeBookRequest, opts ...grpc.CallOption) (*GetRecipeBookResponse, error) {
	out := new(GetRecipeBookResponse)
	err := c.cc.Invoke(ctx, RecipeService_GetRecipeBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeServiceClient) CreateRecipe(ctx context.Context, in *RecipeInput, opts ...grpc.CallOption) (*CreateRecipeResponse, error) {
	out := new(CreateRecipeResponse)
	err := c.cc.Invoke(ctx, RecipeService_CreateRecipe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeServiceClient) GetRecipe(ctx context.Context, in *GetRecipeRequest, opts ...grpc.CallOption) (*GetRecipeResponse, error) {
	out := new(GetRecipeResponse)
	err := c.cc.Invoke(ctx, RecipeService_GetRecipe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeServiceClient) UpdateRecipe(ctx context.Context, in *RecipeInput, opts ...grpc.CallOption) (*UpdateRecipeResponse, error) {
	out := new(UpdateRecipeResponse)
	err := c.cc.Invoke(ctx, RecipeService_UpdateRecipe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeServiceClient) DeleteRecipe(ctx context.Context, in *DeleteRecipeRequest, opts ...grpc.CallOption) (*DeleteRecipeResponse, error) {
	out := new(DeleteRecipeResponse)
	err := c.cc.Invoke(ctx, RecipeService_DeleteRecipe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeServiceClient) GenerateRecipePicturesUploadLinks(ctx context.Context, in *GenerateRecipePicturesUploadLinksRequest, opts ...grpc.CallOption) (*GenerateRecipePicturesUploadLinksResponse, error) {
	out := new(GenerateRecipePicturesUploadLinksResponse)
	err := c.cc.Invoke(ctx, RecipeService_GenerateRecipePicturesUploadLinks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeServiceClient) SetRecipePictures(ctx context.Context, in *SetRecipePicturesRequest, opts ...grpc.CallOption) (*SetRecipePicturesResponse, error) {
	out := new(SetRecipePicturesResponse)
	err := c.cc.Invoke(ctx, RecipeService_SetRecipePictures_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeServiceClient) RateRecipe(ctx context.Context, in *RateRecipeRequest, opts ...grpc.CallOption) (*RateRecipeResponse, error) {
	out := new(RateRecipeResponse)
	err := c.cc.Invoke(ctx, RecipeService_RateRecipe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeServiceClient) SaveToRecipeBook(ctx context.Context, in *SaveToRecipeBookRequest, opts ...grpc.CallOption) (*SaveToRecipeBookResponse, error) {
	out := new(SaveToRecipeBookResponse)
	err := c.cc.Invoke(ctx, RecipeService_SaveToRecipeBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeServiceClient) RemoveFromRecipeBook(ctx context.Context, in *RemoveFromRecipeBookRequest, opts ...grpc.CallOption) (*RemoveFromRecipeBookResponse, error) {
	out := new(RemoveFromRecipeBookResponse)
	err := c.cc.Invoke(ctx, RecipeService_RemoveFromRecipeBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeServiceClient) SetRecipeFavouriteStatus(ctx context.Context, in *SetRecipeFavouriteStatusRequest, opts ...grpc.CallOption) (*SetRecipeFavouriteStatusResponse, error) {
	out := new(SetRecipeFavouriteStatusResponse)
	err := c.cc.Invoke(ctx, RecipeService_SetRecipeFavouriteStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeServiceClient) SetRecipeCategories(ctx context.Context, in *SetRecipeCategoriesRequest, opts ...grpc.CallOption) (*SetRecipeCategoriesResponse, error) {
	out := new(SetRecipeCategoriesResponse)
	err := c.cc.Invoke(ctx, RecipeService_SetRecipeCategories_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeServiceClient) TranslateRecipe(ctx context.Context, in *TranslateRecipeRequest, opts ...grpc.CallOption) (*TranslateRecipeResponse, error) {
	out := new(TranslateRecipeResponse)
	err := c.cc.Invoke(ctx, RecipeService_TranslateRecipe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeServiceClient) DeleteRecipeTranslation(ctx context.Context, in *DeleteRecipeTranslationRequest, opts ...grpc.CallOption) (*DeleteRecipeTranslationResponse, error) {
	out := new(DeleteRecipeTranslationResponse)
	err := c.cc.Invoke(ctx, RecipeService_DeleteRecipeTranslation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeServiceClient) GetRecipePolicy(ctx context.Context, in *GetRecipePolicyRequest, opts ...grpc.CallOption) (*GetRecipePolicyResponse, error) {
	out := new(GetRecipePolicyResponse)
	err := c.cc.Invoke(ctx, RecipeService_GetRecipePolicy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeServiceClient) GetRecipeNames(ctx context.Context, in *GetRecipeNamesRequest, opts ...grpc.CallOption) (*GetRecipeNamesResponse, error) {
	out := new(GetRecipeNamesResponse)
	err := c.cc.Invoke(ctx, RecipeService_GetRecipeNames_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecipeServiceServer is the server API for RecipeService service.
// All implementations must embed UnimplementedRecipeServiceServer
// for forward compatibility
type RecipeServiceServer interface {
	GetRecipes(context.Context, *GetRecipesRequest) (*GetRecipesResponse, error)
	GetRandomRecipe(context.Context, *GetRandomRecipeRequest) (*GetRecipeResponse, error)
	GetRecipeBook(context.Context, *GetRecipeBookRequest) (*GetRecipeBookResponse, error)
	CreateRecipe(context.Context, *RecipeInput) (*CreateRecipeResponse, error)
	GetRecipe(context.Context, *GetRecipeRequest) (*GetRecipeResponse, error)
	UpdateRecipe(context.Context, *RecipeInput) (*UpdateRecipeResponse, error)
	DeleteRecipe(context.Context, *DeleteRecipeRequest) (*DeleteRecipeResponse, error)
	GenerateRecipePicturesUploadLinks(context.Context, *GenerateRecipePicturesUploadLinksRequest) (*GenerateRecipePicturesUploadLinksResponse, error)
	SetRecipePictures(context.Context, *SetRecipePicturesRequest) (*SetRecipePicturesResponse, error)
	RateRecipe(context.Context, *RateRecipeRequest) (*RateRecipeResponse, error)
	SaveToRecipeBook(context.Context, *SaveToRecipeBookRequest) (*SaveToRecipeBookResponse, error)
	RemoveFromRecipeBook(context.Context, *RemoveFromRecipeBookRequest) (*RemoveFromRecipeBookResponse, error)
	SetRecipeFavouriteStatus(context.Context, *SetRecipeFavouriteStatusRequest) (*SetRecipeFavouriteStatusResponse, error)
	SetRecipeCategories(context.Context, *SetRecipeCategoriesRequest) (*SetRecipeCategoriesResponse, error)
	TranslateRecipe(context.Context, *TranslateRecipeRequest) (*TranslateRecipeResponse, error)
	DeleteRecipeTranslation(context.Context, *DeleteRecipeTranslationRequest) (*DeleteRecipeTranslationResponse, error)
	GetRecipePolicy(context.Context, *GetRecipePolicyRequest) (*GetRecipePolicyResponse, error)
	GetRecipeNames(context.Context, *GetRecipeNamesRequest) (*GetRecipeNamesResponse, error)
	mustEmbedUnimplementedRecipeServiceServer()
}

// UnimplementedRecipeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRecipeServiceServer struct {
}

func (UnimplementedRecipeServiceServer) GetRecipes(context.Context, *GetRecipesRequest) (*GetRecipesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecipes not implemented")
}
func (UnimplementedRecipeServiceServer) GetRandomRecipe(context.Context, *GetRandomRecipeRequest) (*GetRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRandomRecipe not implemented")
}
func (UnimplementedRecipeServiceServer) GetRecipeBook(context.Context, *GetRecipeBookRequest) (*GetRecipeBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecipeBook not implemented")
}
func (UnimplementedRecipeServiceServer) CreateRecipe(context.Context, *RecipeInput) (*CreateRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRecipe not implemented")
}
func (UnimplementedRecipeServiceServer) GetRecipe(context.Context, *GetRecipeRequest) (*GetRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecipe not implemented")
}
func (UnimplementedRecipeServiceServer) UpdateRecipe(context.Context, *RecipeInput) (*UpdateRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRecipe not implemented")
}
func (UnimplementedRecipeServiceServer) DeleteRecipe(context.Context, *DeleteRecipeRequest) (*DeleteRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRecipe not implemented")
}
func (UnimplementedRecipeServiceServer) GenerateRecipePicturesUploadLinks(context.Context, *GenerateRecipePicturesUploadLinksRequest) (*GenerateRecipePicturesUploadLinksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateRecipePicturesUploadLinks not implemented")
}
func (UnimplementedRecipeServiceServer) SetRecipePictures(context.Context, *SetRecipePicturesRequest) (*SetRecipePicturesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRecipePictures not implemented")
}
func (UnimplementedRecipeServiceServer) RateRecipe(context.Context, *RateRecipeRequest) (*RateRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RateRecipe not implemented")
}
func (UnimplementedRecipeServiceServer) SaveToRecipeBook(context.Context, *SaveToRecipeBookRequest) (*SaveToRecipeBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveToRecipeBook not implemented")
}
func (UnimplementedRecipeServiceServer) RemoveFromRecipeBook(context.Context, *RemoveFromRecipeBookRequest) (*RemoveFromRecipeBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromRecipeBook not implemented")
}
func (UnimplementedRecipeServiceServer) SetRecipeFavouriteStatus(context.Context, *SetRecipeFavouriteStatusRequest) (*SetRecipeFavouriteStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRecipeFavouriteStatus not implemented")
}
func (UnimplementedRecipeServiceServer) SetRecipeCategories(context.Context, *SetRecipeCategoriesRequest) (*SetRecipeCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRecipeCategories not implemented")
}
func (UnimplementedRecipeServiceServer) TranslateRecipe(context.Context, *TranslateRecipeRequest) (*TranslateRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TranslateRecipe not implemented")
}
func (UnimplementedRecipeServiceServer) DeleteRecipeTranslation(context.Context, *DeleteRecipeTranslationRequest) (*DeleteRecipeTranslationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRecipeTranslation not implemented")
}
func (UnimplementedRecipeServiceServer) GetRecipePolicy(context.Context, *GetRecipePolicyRequest) (*GetRecipePolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecipePolicy not implemented")
}
func (UnimplementedRecipeServiceServer) GetRecipeNames(context.Context, *GetRecipeNamesRequest) (*GetRecipeNamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecipeNames not implemented")
}
func (UnimplementedRecipeServiceServer) mustEmbedUnimplementedRecipeServiceServer() {}

// UnsafeRecipeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecipeServiceServer will
// result in compilation errors.
type UnsafeRecipeServiceServer interface {
	mustEmbedUnimplementedRecipeServiceServer()
}

func RegisterRecipeServiceServer(s grpc.ServiceRegistrar, srv RecipeServiceServer) {
	s.RegisterService(&RecipeService_ServiceDesc, srv)
}

func _RecipeService_GetRecipes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecipesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeServiceServer).GetRecipes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecipeService_GetRecipes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeServiceServer).GetRecipes(ctx, req.(*GetRecipesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeService_GetRandomRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRandomRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeServiceServer).GetRandomRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecipeService_GetRandomRecipe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeServiceServer).GetRandomRecipe(ctx, req.(*GetRandomRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeService_GetRecipeBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecipeBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeServiceServer).GetRecipeBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecipeService_GetRecipeBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeServiceServer).GetRecipeBook(ctx, req.(*GetRecipeBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeService_CreateRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecipeInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeServiceServer).CreateRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecipeService_CreateRecipe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeServiceServer).CreateRecipe(ctx, req.(*RecipeInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeService_GetRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeServiceServer).GetRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecipeService_GetRecipe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeServiceServer).GetRecipe(ctx, req.(*GetRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeService_UpdateRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecipeInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeServiceServer).UpdateRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecipeService_UpdateRecipe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeServiceServer).UpdateRecipe(ctx, req.(*RecipeInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeService_DeleteRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeServiceServer).DeleteRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecipeService_DeleteRecipe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeServiceServer).DeleteRecipe(ctx, req.(*DeleteRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeService_GenerateRecipePicturesUploadLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateRecipePicturesUploadLinksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeServiceServer).GenerateRecipePicturesUploadLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecipeService_GenerateRecipePicturesUploadLinks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeServiceServer).GenerateRecipePicturesUploadLinks(ctx, req.(*GenerateRecipePicturesUploadLinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeService_SetRecipePictures_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRecipePicturesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeServiceServer).SetRecipePictures(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecipeService_SetRecipePictures_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeServiceServer).SetRecipePictures(ctx, req.(*SetRecipePicturesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeService_RateRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeServiceServer).RateRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecipeService_RateRecipe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeServiceServer).RateRecipe(ctx, req.(*RateRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeService_SaveToRecipeBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveToRecipeBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeServiceServer).SaveToRecipeBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecipeService_SaveToRecipeBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeServiceServer).SaveToRecipeBook(ctx, req.(*SaveToRecipeBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeService_RemoveFromRecipeBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFromRecipeBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeServiceServer).RemoveFromRecipeBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecipeService_RemoveFromRecipeBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeServiceServer).RemoveFromRecipeBook(ctx, req.(*RemoveFromRecipeBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeService_SetRecipeFavouriteStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRecipeFavouriteStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeServiceServer).SetRecipeFavouriteStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecipeService_SetRecipeFavouriteStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeServiceServer).SetRecipeFavouriteStatus(ctx, req.(*SetRecipeFavouriteStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeService_SetRecipeCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRecipeCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeServiceServer).SetRecipeCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecipeService_SetRecipeCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeServiceServer).SetRecipeCategories(ctx, req.(*SetRecipeCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeService_TranslateRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranslateRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeServiceServer).TranslateRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecipeService_TranslateRecipe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeServiceServer).TranslateRecipe(ctx, req.(*TranslateRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeService_DeleteRecipeTranslation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRecipeTranslationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeServiceServer).DeleteRecipeTranslation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecipeService_DeleteRecipeTranslation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeServiceServer).DeleteRecipeTranslation(ctx, req.(*DeleteRecipeTranslationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeService_GetRecipePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecipePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeServiceServer).GetRecipePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecipeService_GetRecipePolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeServiceServer).GetRecipePolicy(ctx, req.(*GetRecipePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeService_GetRecipeNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecipeNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeServiceServer).GetRecipeNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecipeService_GetRecipeNames_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeServiceServer).GetRecipeNames(ctx, req.(*GetRecipeNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RecipeService_ServiceDesc is the grpc.ServiceDesc for RecipeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecipeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.RecipeService",
	HandlerType: (*RecipeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRecipes",
			Handler:    _RecipeService_GetRecipes_Handler,
		},
		{
			MethodName: "GetRandomRecipe",
			Handler:    _RecipeService_GetRandomRecipe_Handler,
		},
		{
			MethodName: "GetRecipeBook",
			Handler:    _RecipeService_GetRecipeBook_Handler,
		},
		{
			MethodName: "CreateRecipe",
			Handler:    _RecipeService_CreateRecipe_Handler,
		},
		{
			MethodName: "GetRecipe",
			Handler:    _RecipeService_GetRecipe_Handler,
		},
		{
			MethodName: "UpdateRecipe",
			Handler:    _RecipeService_UpdateRecipe_Handler,
		},
		{
			MethodName: "DeleteRecipe",
			Handler:    _RecipeService_DeleteRecipe_Handler,
		},
		{
			MethodName: "GenerateRecipePicturesUploadLinks",
			Handler:    _RecipeService_GenerateRecipePicturesUploadLinks_Handler,
		},
		{
			MethodName: "SetRecipePictures",
			Handler:    _RecipeService_SetRecipePictures_Handler,
		},
		{
			MethodName: "RateRecipe",
			Handler:    _RecipeService_RateRecipe_Handler,
		},
		{
			MethodName: "SaveToRecipeBook",
			Handler:    _RecipeService_SaveToRecipeBook_Handler,
		},
		{
			MethodName: "RemoveFromRecipeBook",
			Handler:    _RecipeService_RemoveFromRecipeBook_Handler,
		},
		{
			MethodName: "SetRecipeFavouriteStatus",
			Handler:    _RecipeService_SetRecipeFavouriteStatus_Handler,
		},
		{
			MethodName: "SetRecipeCategories",
			Handler:    _RecipeService_SetRecipeCategories_Handler,
		},
		{
			MethodName: "TranslateRecipe",
			Handler:    _RecipeService_TranslateRecipe_Handler,
		},
		{
			MethodName: "DeleteRecipeTranslation",
			Handler:    _RecipeService_DeleteRecipeTranslation_Handler,
		},
		{
			MethodName: "GetRecipePolicy",
			Handler:    _RecipeService_GetRecipePolicy_Handler,
		},
		{
			MethodName: "GetRecipeNames",
			Handler:    _RecipeService_GetRecipeNames_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/recipe-service.proto",
}
