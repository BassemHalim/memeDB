// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: meme_service.proto

package protobuff

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

// Error handling
type ErrorCode int32

const (
	ErrorCode_ERROR_UNSPECIFIED     ErrorCode = 0
	ErrorCode_ERROR_INVALID_IMAGE   ErrorCode = 1
	ErrorCode_ERROR_IMAGE_TOO_LARGE ErrorCode = 2
	ErrorCode_ERROR_INVALID_TAGS    ErrorCode = 3
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0: "ERROR_UNSPECIFIED",
		1: "ERROR_INVALID_IMAGE",
		2: "ERROR_IMAGE_TOO_LARGE",
		3: "ERROR_INVALID_TAGS",
	}
	ErrorCode_value = map[string]int32{
		"ERROR_UNSPECIFIED":     0,
		"ERROR_INVALID_IMAGE":   1,
		"ERROR_IMAGE_TOO_LARGE": 2,
		"ERROR_INVALID_TAGS":    3,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_meme_service_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_meme_service_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_meme_service_proto_rawDescGZIP(), []int{0}
}

// Request message for uploading a meme
type UploadMemeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Binary content of the image
	ImageData []byte `protobuf:"bytes,1,opt,name=image_data,json=imageData,proto3" json:"image_data,omitempty"`
	// Image metadata
	ImageName string `protobuf:"bytes,2,opt,name=image_name,json=imageName,proto3" json:"image_name,omitempty"`
	MimeType  string `protobuf:"bytes,3,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	// List of tags associated with the meme
	Tags []string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *UploadMemeRequest) Reset() {
	*x = UploadMemeRequest{}
	mi := &file_meme_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadMemeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadMemeRequest) ProtoMessage() {}

func (x *UploadMemeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_meme_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadMemeRequest.ProtoReflect.Descriptor instead.
func (*UploadMemeRequest) Descriptor() ([]byte, []int) {
	return file_meme_service_proto_rawDescGZIP(), []int{0}
}

func (x *UploadMemeRequest) GetImageData() []byte {
	if x != nil {
		return x.ImageData
	}
	return nil
}

func (x *UploadMemeRequest) GetImageName() string {
	if x != nil {
		return x.ImageName
	}
	return ""
}

func (x *UploadMemeRequest) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *UploadMemeRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

// Response message after uploading a meme
type UploadMemeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier for the uploaded meme
	MemeId string `protobuf:"bytes,1,opt,name=meme_id,json=memeId,proto3" json:"meme_id,omitempty"`
	// URL where the meme can be accessed
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// Timestamp of when the meme was uploaded
	UploadTimestamp int64 `protobuf:"varint,3,opt,name=upload_timestamp,json=uploadTimestamp,proto3" json:"upload_timestamp,omitempty"`
}

func (x *UploadMemeResponse) Reset() {
	*x = UploadMemeResponse{}
	mi := &file_meme_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadMemeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadMemeResponse) ProtoMessage() {}

func (x *UploadMemeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_meme_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadMemeResponse.ProtoReflect.Descriptor instead.
func (*UploadMemeResponse) Descriptor() ([]byte, []int) {
	return file_meme_service_proto_rawDescGZIP(), []int{1}
}

func (x *UploadMemeResponse) GetMemeId() string {
	if x != nil {
		return x.MemeId
	}
	return ""
}

func (x *UploadMemeResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UploadMemeResponse) GetUploadTimestamp() int64 {
	if x != nil {
		return x.UploadTimestamp
	}
	return 0
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    ErrorCode `protobuf:"varint,1,opt,name=code,proto3,enum=meme_service.ErrorCode" json:"code,omitempty"`
	Message string    `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	mi := &file_meme_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_meme_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_meme_service_proto_rawDescGZIP(), []int{2}
}

func (x *Error) GetCode() ErrorCode {
	if x != nil {
		return x.Code
	}
	return ErrorCode_ERROR_UNSPECIFIED
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_meme_service_proto protoreflect.FileDescriptor

var file_meme_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x65, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d, 0x65, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x22, 0x82, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x65, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x6a, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x4d, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x6d, 0x65, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6d, 0x65, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x29, 0x0a, 0x10, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x22, 0x4e, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2b, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x6d,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2a, 0x6e, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x15, 0x0a, 0x11, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x01,
	0x12, 0x19, 0x0a, 0x15, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f,
	0x54, 0x4f, 0x4f, 0x5f, 0x4c, 0x41, 0x52, 0x47, 0x45, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x54, 0x41, 0x47,
	0x53, 0x10, 0x03, 0x32, 0x5e, 0x0a, 0x0b, 0x4d, 0x65, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x65, 0x6d, 0x65,
	0x12, 0x1f, 0x2e, 0x6d, 0x65, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x6d, 0x65, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x42, 0x61, 0x73, 0x73, 0x65, 0x6d, 0x48, 0x61, 0x6c, 0x69, 0x6d, 0x2f, 0x6d, 0x65,
	0x6d, 0x65, 0x44, 0x42, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x66, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meme_service_proto_rawDescOnce sync.Once
	file_meme_service_proto_rawDescData = file_meme_service_proto_rawDesc
)

func file_meme_service_proto_rawDescGZIP() []byte {
	file_meme_service_proto_rawDescOnce.Do(func() {
		file_meme_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_meme_service_proto_rawDescData)
	})
	return file_meme_service_proto_rawDescData
}

var file_meme_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_meme_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_meme_service_proto_goTypes = []any{
	(ErrorCode)(0),             // 0: meme_service.ErrorCode
	(*UploadMemeRequest)(nil),  // 1: meme_service.UploadMemeRequest
	(*UploadMemeResponse)(nil), // 2: meme_service.UploadMemeResponse
	(*Error)(nil),              // 3: meme_service.Error
}
var file_meme_service_proto_depIdxs = []int32{
	0, // 0: meme_service.Error.code:type_name -> meme_service.ErrorCode
	1, // 1: meme_service.MemeService.UploadMeme:input_type -> meme_service.UploadMemeRequest
	2, // 2: meme_service.MemeService.UploadMeme:output_type -> meme_service.UploadMemeResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_meme_service_proto_init() }
func file_meme_service_proto_init() {
	if File_meme_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_meme_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_meme_service_proto_goTypes,
		DependencyIndexes: file_meme_service_proto_depIdxs,
		EnumInfos:         file_meme_service_proto_enumTypes,
		MessageInfos:      file_meme_service_proto_msgTypes,
	}.Build()
	File_meme_service_proto = out.File
	file_meme_service_proto_rawDesc = nil
	file_meme_service_proto_goTypes = nil
	file_meme_service_proto_depIdxs = nil
}
