// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.27.0
// source: v1/macronutrients.proto

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

type Macronutrients struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protein       *int32 `protobuf:"varint,1,opt,name=Protein,proto3,oneof" json:"Protein,omitempty"`
	Fats          *int32 `protobuf:"varint,2,opt,name=Fats,proto3,oneof" json:"Fats,omitempty"`
	Carbohydrates *int32 `protobuf:"varint,3,opt,name=Carbohydrates,proto3,oneof" json:"Carbohydrates,omitempty"`
}

func (x *Macronutrients) Reset() {
	*x = Macronutrients{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_macronutrients_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Macronutrients) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Macronutrients) ProtoMessage() {}

func (x *Macronutrients) ProtoReflect() protoreflect.Message {
	mi := &file_v1_macronutrients_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Macronutrients.ProtoReflect.Descriptor instead.
func (*Macronutrients) Descriptor() ([]byte, []int) {
	return file_v1_macronutrients_proto_rawDescGZIP(), []int{0}
}

func (x *Macronutrients) GetProtein() int32 {
	if x != nil && x.Protein != nil {
		return *x.Protein
	}
	return 0
}

func (x *Macronutrients) GetFats() int32 {
	if x != nil && x.Fats != nil {
		return *x.Fats
	}
	return 0
}

func (x *Macronutrients) GetCarbohydrates() int32 {
	if x != nil && x.Carbohydrates != nil {
		return *x.Carbohydrates
	}
	return 0
}

var File_v1_macronutrients_proto protoreflect.FileDescriptor

var file_v1_macronutrients_proto_rawDesc = []byte{
	0x0a, 0x17, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x63, 0x72, 0x6f, 0x6e, 0x75, 0x74, 0x72, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x22, 0x9a, 0x01,
	0x0a, 0x0e, 0x4d, 0x61, 0x63, 0x72, 0x6f, 0x6e, 0x75, 0x74, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x1d, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x00, 0x52, 0x07, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e, 0x88, 0x01, 0x01, 0x12,
	0x17, 0x0a, 0x04, 0x46, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52,
	0x04, 0x46, 0x61, 0x74, 0x73, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0d, 0x43, 0x61, 0x72, 0x62,
	0x6f, 0x68, 0x79, 0x64, 0x72, 0x61, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x02, 0x52, 0x0d, 0x43, 0x61, 0x72, 0x62, 0x6f, 0x68, 0x79, 0x64, 0x72, 0x61, 0x74, 0x65, 0x73,
	0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x46, 0x61, 0x74, 0x73, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x43, 0x61, 0x72,
	0x62, 0x6f, 0x68, 0x79, 0x64, 0x72, 0x61, 0x74, 0x65, 0x73, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x70, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x6c, 0x69, 0x65, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x62, 0x6f, 0x6f, 0x6b, 0x2d, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_v1_macronutrients_proto_rawDescOnce sync.Once
	file_v1_macronutrients_proto_rawDescData = file_v1_macronutrients_proto_rawDesc
)

func file_v1_macronutrients_proto_rawDescGZIP() []byte {
	file_v1_macronutrients_proto_rawDescOnce.Do(func() {
		file_v1_macronutrients_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_macronutrients_proto_rawDescData)
	})
	return file_v1_macronutrients_proto_rawDescData
}

var file_v1_macronutrients_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_v1_macronutrients_proto_goTypes = []interface{}{
	(*Macronutrients)(nil), // 0: v1.Macronutrients
}
var file_v1_macronutrients_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_macronutrients_proto_init() }
func file_v1_macronutrients_proto_init() {
	if File_v1_macronutrients_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_macronutrients_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Macronutrients); i {
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
	file_v1_macronutrients_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_macronutrients_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_macronutrients_proto_goTypes,
		DependencyIndexes: file_v1_macronutrients_proto_depIdxs,
		MessageInfos:      file_v1_macronutrients_proto_msgTypes,
	}.Build()
	File_v1_macronutrients_proto = out.File
	file_v1_macronutrients_proto_rawDesc = nil
	file_v1_macronutrients_proto_goTypes = nil
	file_v1_macronutrients_proto_depIdxs = nil
}
