// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.3
// source: dockersinventory.proto

package pb

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

type DockerImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Cmd   string `protobuf:"bytes,2,opt,name=cmd,proto3" json:"cmd,omitempty"`
}

func (x *DockerImage) Reset() {
	*x = DockerImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dockersinventory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DockerImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DockerImage) ProtoMessage() {}

func (x *DockerImage) ProtoReflect() protoreflect.Message {
	mi := &file_dockersinventory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DockerImage.ProtoReflect.Descriptor instead.
func (*DockerImage) Descriptor() ([]byte, []int) {
	return file_dockersinventory_proto_rawDescGZIP(), []int{0}
}

func (x *DockerImage) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *DockerImage) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

type GetDockerImagesListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Cmd   string `protobuf:"bytes,2,opt,name=cmd,proto3" json:"cmd,omitempty"`
}

func (x *GetDockerImagesListRequest) Reset() {
	*x = GetDockerImagesListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dockersinventory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDockerImagesListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDockerImagesListRequest) ProtoMessage() {}

func (x *GetDockerImagesListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dockersinventory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDockerImagesListRequest.ProtoReflect.Descriptor instead.
func (*GetDockerImagesListRequest) Descriptor() ([]byte, []int) {
	return file_dockersinventory_proto_rawDescGZIP(), []int{1}
}

func (x *GetDockerImagesListRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *GetDockerImagesListRequest) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

type GetDockerImagesListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DockerImages []*DockerImage `protobuf:"bytes,1,rep,name=dockerImages,proto3" json:"dockerImages,omitempty"`
}

func (x *GetDockerImagesListResponse) Reset() {
	*x = GetDockerImagesListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dockersinventory_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDockerImagesListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDockerImagesListResponse) ProtoMessage() {}

func (x *GetDockerImagesListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dockersinventory_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDockerImagesListResponse.ProtoReflect.Descriptor instead.
func (*GetDockerImagesListResponse) Descriptor() ([]byte, []int) {
	return file_dockersinventory_proto_rawDescGZIP(), []int{2}
}

func (x *GetDockerImagesListResponse) GetDockerImages() []*DockerImage {
	if x != nil {
		return x.DockerImages
	}
	return nil
}

var File_dockersinventory_proto protoreflect.FileDescriptor

var file_dockersinventory_proto_rawDesc = []byte{
	0x0a, 0x16, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x73, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x0b, 0x44, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x63, 0x6d, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x22,
	0x44, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x63, 0x6d, 0x64, 0x22, 0x4f, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0c, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x44, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x0c, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x32, 0x5f, 0x0a, 0x09, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x12, 0x52, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x47, 0x65, 0x74,
	0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13, 0x64, 0x6f, 0x63, 0x6b, 0x65,
	0x72, 0x73, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dockersinventory_proto_rawDescOnce sync.Once
	file_dockersinventory_proto_rawDescData = file_dockersinventory_proto_rawDesc
)

func file_dockersinventory_proto_rawDescGZIP() []byte {
	file_dockersinventory_proto_rawDescOnce.Do(func() {
		file_dockersinventory_proto_rawDescData = protoimpl.X.CompressGZIP(file_dockersinventory_proto_rawDescData)
	})
	return file_dockersinventory_proto_rawDescData
}

var file_dockersinventory_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_dockersinventory_proto_goTypes = []interface{}{
	(*DockerImage)(nil),                 // 0: DockerImage
	(*GetDockerImagesListRequest)(nil),  // 1: GetDockerImagesListRequest
	(*GetDockerImagesListResponse)(nil), // 2: GetDockerImagesListResponse
}
var file_dockersinventory_proto_depIdxs = []int32{
	0, // 0: GetDockerImagesListResponse.dockerImages:type_name -> DockerImage
	1, // 1: Inventory.GetDockerImagesList:input_type -> GetDockerImagesListRequest
	2, // 2: Inventory.GetDockerImagesList:output_type -> GetDockerImagesListResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_dockersinventory_proto_init() }
func file_dockersinventory_proto_init() {
	if File_dockersinventory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dockersinventory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DockerImage); i {
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
		file_dockersinventory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDockerImagesListRequest); i {
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
		file_dockersinventory_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDockerImagesListResponse); i {
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
			RawDescriptor: file_dockersinventory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dockersinventory_proto_goTypes,
		DependencyIndexes: file_dockersinventory_proto_depIdxs,
		MessageInfos:      file_dockersinventory_proto_msgTypes,
	}.Build()
	File_dockersinventory_proto = out.File
	file_dockersinventory_proto_rawDesc = nil
	file_dockersinventory_proto_goTypes = nil
	file_dockersinventory_proto_depIdxs = nil
}