// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: meme.proto

package memeService

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

type TagMatchType int32

const (
	TagMatchType_ANY TagMatchType = 0 // Match memes that have any of the specified tags
	TagMatchType_ALL TagMatchType = 1 // Match memes that have all of the specified tags
)

// Enum value maps for TagMatchType.
var (
	TagMatchType_name = map[int32]string{
		0: "ANY",
		1: "ALL",
	}
	TagMatchType_value = map[string]int32{
		"ANY": 0,
		"ALL": 1,
	}
)

func (x TagMatchType) Enum() *TagMatchType {
	p := new(TagMatchType)
	*p = x
	return p
}

func (x TagMatchType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TagMatchType) Descriptor() protoreflect.EnumDescriptor {
	return file_meme_proto_enumTypes[0].Descriptor()
}

func (TagMatchType) Type() protoreflect.EnumType {
	return &file_meme_proto_enumTypes[0]
}

func (x TagMatchType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TagMatchType.Descriptor instead.
func (TagMatchType) EnumDescriptor() ([]byte, []int) {
	return file_meme_proto_rawDescGZIP(), []int{0}
}

type SortOrder int32

const (
	SortOrder_NEWEST      SortOrder = 0
	SortOrder_OLDEST      SortOrder = 1
	SortOrder_MOST_TAGGED SortOrder = 2
)

// Enum value maps for SortOrder.
var (
	SortOrder_name = map[int32]string{
		0: "NEWEST",
		1: "OLDEST",
		2: "MOST_TAGGED",
	}
	SortOrder_value = map[string]int32{
		"NEWEST":      0,
		"OLDEST":      1,
		"MOST_TAGGED": 2,
	}
)

func (x SortOrder) Enum() *SortOrder {
	p := new(SortOrder)
	*p = x
	return p
}

func (x SortOrder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortOrder) Descriptor() protoreflect.EnumDescriptor {
	return file_meme_proto_enumTypes[1].Descriptor()
}

func (SortOrder) Type() protoreflect.EnumType {
	return &file_meme_proto_enumTypes[1]
}

func (x SortOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortOrder.Descriptor instead.
func (SortOrder) EnumDescriptor() ([]byte, []int) {
	return file_meme_proto_rawDescGZIP(), []int{1}
}

type UploadMemeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image      []byte   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	MediaType  string   `protobuf:"bytes,2,opt,name=media_type,json=mediaType,proto3" json:"media_type,omitempty"`
	Tags       []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	Name       string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Dimensions []int32  `protobuf:"varint,5,rep,packed,name=dimensions,proto3" json:"dimensions,omitempty"`
}

func (x *UploadMemeRequest) Reset() {
	*x = UploadMemeRequest{}
	mi := &file_meme_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadMemeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadMemeRequest) ProtoMessage() {}

func (x *UploadMemeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_meme_proto_msgTypes[0]
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
	return file_meme_proto_rawDescGZIP(), []int{0}
}

func (x *UploadMemeRequest) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *UploadMemeRequest) GetMediaType() string {
	if x != nil {
		return x.MediaType
	}
	return ""
}

func (x *UploadMemeRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *UploadMemeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UploadMemeRequest) GetDimensions() []int32 {
	if x != nil {
		return x.Dimensions
	}
	return nil
}

type GetMemeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMemeRequest) Reset() {
	*x = GetMemeRequest{}
	mi := &file_meme_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMemeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMemeRequest) ProtoMessage() {}

func (x *GetMemeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_meme_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMemeRequest.ProtoReflect.Descriptor instead.
func (*GetMemeRequest) Descriptor() ([]byte, []int) {
	return file_meme_proto_rawDescGZIP(), []int{1}
}

func (x *GetMemeRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteMemeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteMemeRequest) Reset() {
	*x = DeleteMemeRequest{}
	mi := &file_meme_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteMemeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMemeRequest) ProtoMessage() {}

func (x *DeleteMemeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_meme_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMemeRequest.ProtoReflect.Descriptor instead.
func (*DeleteMemeRequest) Descriptor() ([]byte, []int) {
	return file_meme_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteMemeRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetTimelineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page      int32     `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int32     `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	SortOrder SortOrder `protobuf:"varint,3,opt,name=sort_order,json=sortOrder,proto3,enum=meme.SortOrder" json:"sort_order,omitempty"`
}

func (x *GetTimelineRequest) Reset() {
	*x = GetTimelineRequest{}
	mi := &file_meme_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTimelineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTimelineRequest) ProtoMessage() {}

func (x *GetTimelineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_meme_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTimelineRequest.ProtoReflect.Descriptor instead.
func (*GetTimelineRequest) Descriptor() ([]byte, []int) {
	return file_meme_proto_rawDescGZIP(), []int{3}
}

func (x *GetTimelineRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetTimelineRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetTimelineRequest) GetSortOrder() SortOrder {
	if x != nil {
		return x.SortOrder
	}
	return SortOrder_NEWEST
}

type FilterMemesByTagsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags      []string     `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	MatchType TagMatchType `protobuf:"varint,2,opt,name=match_type,json=matchType,proto3,enum=meme.TagMatchType" json:"match_type,omitempty"`
	Page      int32        `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int32        `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	SortOrder SortOrder    `protobuf:"varint,5,opt,name=sort_order,json=sortOrder,proto3,enum=meme.SortOrder" json:"sort_order,omitempty"`
}

func (x *FilterMemesByTagsRequest) Reset() {
	*x = FilterMemesByTagsRequest{}
	mi := &file_meme_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FilterMemesByTagsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterMemesByTagsRequest) ProtoMessage() {}

func (x *FilterMemesByTagsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_meme_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterMemesByTagsRequest.ProtoReflect.Descriptor instead.
func (*FilterMemesByTagsRequest) Descriptor() ([]byte, []int) {
	return file_meme_proto_rawDescGZIP(), []int{4}
}

func (x *FilterMemesByTagsRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *FilterMemesByTagsRequest) GetMatchType() TagMatchType {
	if x != nil {
		return x.MatchType
	}
	return TagMatchType_ANY
}

func (x *FilterMemesByTagsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *FilterMemesByTagsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *FilterMemesByTagsRequest) GetSortOrder() SortOrder {
	if x != nil {
		return x.SortOrder
	}
	return SortOrder_NEWEST
}

type SearchMemesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query    string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Page     int32  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32  `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *SearchMemesRequest) Reset() {
	*x = SearchMemesRequest{}
	mi := &file_meme_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchMemesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchMemesRequest) ProtoMessage() {}

func (x *SearchMemesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_meme_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchMemesRequest.ProtoReflect.Descriptor instead.
func (*SearchMemesRequest) Descriptor() ([]byte, []int) {
	return file_meme_proto_rawDescGZIP(), []int{5}
}

func (x *SearchMemesRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SearchMemesRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchMemesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// Response messages
type MemeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MediaUrl   string   `protobuf:"bytes,2,opt,name=media_url,json=mediaUrl,proto3" json:"media_url,omitempty"`
	MediaType  string   `protobuf:"bytes,3,opt,name=media_type,json=mediaType,proto3" json:"media_type,omitempty"`
	Name       string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Tags       []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	Dimensions []int32  `protobuf:"varint,6,rep,packed,name=dimensions,proto3" json:"dimensions,omitempty"`
}

func (x *MemeResponse) Reset() {
	*x = MemeResponse{}
	mi := &file_meme_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MemeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemeResponse) ProtoMessage() {}

func (x *MemeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_meme_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemeResponse.ProtoReflect.Descriptor instead.
func (*MemeResponse) Descriptor() ([]byte, []int) {
	return file_meme_proto_rawDescGZIP(), []int{6}
}

func (x *MemeResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MemeResponse) GetMediaUrl() string {
	if x != nil {
		return x.MediaUrl
	}
	return ""
}

func (x *MemeResponse) GetMediaType() string {
	if x != nil {
		return x.MediaType
	}
	return ""
}

func (x *MemeResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MemeResponse) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *MemeResponse) GetDimensions() []int32 {
	if x != nil {
		return x.Dimensions
	}
	return nil
}

type DeleteMemeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteMemeResponse) Reset() {
	*x = DeleteMemeResponse{}
	mi := &file_meme_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteMemeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMemeResponse) ProtoMessage() {}

func (x *DeleteMemeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_meme_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMemeResponse.ProtoReflect.Descriptor instead.
func (*DeleteMemeResponse) Descriptor() ([]byte, []int) {
	return file_meme_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteMemeResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type MemesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Memes      []*MemeResponse `protobuf:"bytes,1,rep,name=memes,proto3" json:"memes,omitempty"`
	TotalCount int32           `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	Page       int32           `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	TotalPages int32           `protobuf:"varint,4,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
}

func (x *MemesResponse) Reset() {
	*x = MemesResponse{}
	mi := &file_meme_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MemesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemesResponse) ProtoMessage() {}

func (x *MemesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_meme_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemesResponse.ProtoReflect.Descriptor instead.
func (*MemesResponse) Descriptor() ([]byte, []int) {
	return file_meme_proto_rawDescGZIP(), []int{8}
}

func (x *MemesResponse) GetMemes() []*MemeResponse {
	if x != nil {
		return x.Memes
	}
	return nil
}

func (x *MemesResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *MemesResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *MemesResponse) GetTotalPages() int32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

var File_meme_proto protoreflect.FileDescriptor

var file_meme_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x65, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x65,
	0x6d, 0x65, 0x22, 0x90, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x65, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a, 0x64, 0x69, 0x6d, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4d, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x75, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x6d, 0x65, 0x6d, 0x65, 0x2e, 0x53,
	0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x09, 0x73, 0x6f, 0x72, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x22, 0xc2, 0x01, 0x0a, 0x18, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4d, 0x65,
	0x6d, 0x65, 0x73, 0x42, 0x79, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x12, 0x31, 0x0a, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x6d, 0x65, 0x6d, 0x65, 0x2e,
	0x54, 0x61, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74,
	0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x6d,
	0x65, 0x6d, 0x65, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x09, 0x73,
	0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x5b, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x4d, 0x65, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xa2, 0x01, 0x0a, 0x0c, 0x4d, 0x65, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69,
	0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a,
	0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x2e, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x8f, 0x01, 0x0a, 0x0d, 0x4d,
	0x65, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05,
	0x6d, 0x65, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x65,
	0x6d, 0x65, 0x2e, 0x4d, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x05, 0x6d, 0x65, 0x6d, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x2a, 0x20, 0x0a, 0x0c,
	0x54, 0x61, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03,
	0x41, 0x4e, 0x59, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4c, 0x4c, 0x10, 0x01, 0x2a, 0x34,
	0x0a, 0x09, 0x53, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0a, 0x0a, 0x06, 0x4e,
	0x45, 0x57, 0x45, 0x53, 0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x4c, 0x44, 0x45, 0x53,
	0x54, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x4f, 0x53, 0x54, 0x5f, 0x54, 0x41, 0x47, 0x47,
	0x45, 0x44, 0x10, 0x02, 0x32, 0x89, 0x03, 0x0a, 0x0b, 0x4d, 0x65, 0x6d, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x65,
	0x6d, 0x65, 0x12, 0x17, 0x2e, 0x6d, 0x65, 0x6d, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x4d, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x65,
	0x6d, 0x65, 0x2e, 0x4d, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x33, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x65, 0x12, 0x14, 0x2e, 0x6d, 0x65, 0x6d,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x6d, 0x65, 0x6d, 0x65, 0x2e, 0x4d, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65,
	0x6d, 0x65, 0x12, 0x17, 0x2e, 0x6d, 0x65, 0x6d, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4d, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x65,
	0x6d, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x4d, 0x65, 0x6d, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x6d, 0x65, 0x6d, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6d, 0x65, 0x6d, 0x65, 0x2e, 0x4d, 0x65, 0x6d, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x11, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x4d, 0x65, 0x6d, 0x65, 0x73, 0x42, 0x79, 0x54, 0x61, 0x67, 0x73, 0x12, 0x1e, 0x2e,
	0x6d, 0x65, 0x6d, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x6d, 0x65, 0x73,
	0x42, 0x79, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x6d, 0x65, 0x6d, 0x65, 0x2e, 0x4d, 0x65, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x65, 0x6d, 0x65,
	0x73, 0x12, 0x18, 0x2e, 0x6d, 0x65, 0x6d, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d,
	0x65, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6d, 0x65,
	0x6d, 0x65, 0x2e, 0x4d, 0x65, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42,
	0x61, 0x73, 0x73, 0x65, 0x6d, 0x48, 0x61, 0x6c, 0x69, 0x6d, 0x2f, 0x6d, 0x65, 0x6d, 0x65, 0x44,
	0x42, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meme_proto_rawDescOnce sync.Once
	file_meme_proto_rawDescData = file_meme_proto_rawDesc
)

func file_meme_proto_rawDescGZIP() []byte {
	file_meme_proto_rawDescOnce.Do(func() {
		file_meme_proto_rawDescData = protoimpl.X.CompressGZIP(file_meme_proto_rawDescData)
	})
	return file_meme_proto_rawDescData
}

var file_meme_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_meme_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_meme_proto_goTypes = []any{
	(TagMatchType)(0),                // 0: meme.TagMatchType
	(SortOrder)(0),                   // 1: meme.SortOrder
	(*UploadMemeRequest)(nil),        // 2: meme.UploadMemeRequest
	(*GetMemeRequest)(nil),           // 3: meme.GetMemeRequest
	(*DeleteMemeRequest)(nil),        // 4: meme.DeleteMemeRequest
	(*GetTimelineRequest)(nil),       // 5: meme.GetTimelineRequest
	(*FilterMemesByTagsRequest)(nil), // 6: meme.FilterMemesByTagsRequest
	(*SearchMemesRequest)(nil),       // 7: meme.SearchMemesRequest
	(*MemeResponse)(nil),             // 8: meme.MemeResponse
	(*DeleteMemeResponse)(nil),       // 9: meme.DeleteMemeResponse
	(*MemesResponse)(nil),            // 10: meme.MemesResponse
}
var file_meme_proto_depIdxs = []int32{
	1,  // 0: meme.GetTimelineRequest.sort_order:type_name -> meme.SortOrder
	0,  // 1: meme.FilterMemesByTagsRequest.match_type:type_name -> meme.TagMatchType
	1,  // 2: meme.FilterMemesByTagsRequest.sort_order:type_name -> meme.SortOrder
	8,  // 3: meme.MemesResponse.memes:type_name -> meme.MemeResponse
	2,  // 4: meme.MemeService.UploadMeme:input_type -> meme.UploadMemeRequest
	3,  // 5: meme.MemeService.GetMeme:input_type -> meme.GetMemeRequest
	4,  // 6: meme.MemeService.DeleteMeme:input_type -> meme.DeleteMemeRequest
	5,  // 7: meme.MemeService.GetTimelineMemes:input_type -> meme.GetTimelineRequest
	6,  // 8: meme.MemeService.FilterMemesByTags:input_type -> meme.FilterMemesByTagsRequest
	7,  // 9: meme.MemeService.SearchMemes:input_type -> meme.SearchMemesRequest
	8,  // 10: meme.MemeService.UploadMeme:output_type -> meme.MemeResponse
	8,  // 11: meme.MemeService.GetMeme:output_type -> meme.MemeResponse
	9,  // 12: meme.MemeService.DeleteMeme:output_type -> meme.DeleteMemeResponse
	10, // 13: meme.MemeService.GetTimelineMemes:output_type -> meme.MemesResponse
	10, // 14: meme.MemeService.FilterMemesByTags:output_type -> meme.MemesResponse
	10, // 15: meme.MemeService.SearchMemes:output_type -> meme.MemesResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_meme_proto_init() }
func file_meme_proto_init() {
	if File_meme_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_meme_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_meme_proto_goTypes,
		DependencyIndexes: file_meme_proto_depIdxs,
		EnumInfos:         file_meme_proto_enumTypes,
		MessageInfos:      file_meme_proto_msgTypes,
	}.Build()
	File_meme_proto = out.File
	file_meme_proto_rawDesc = nil
	file_meme_proto_goTypes = nil
	file_meme_proto_depIdxs = nil
}
