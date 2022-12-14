// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.5
// source: items.proto

package items

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

type FetchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url       string                 `protobuf:"bytes,1,opt,name=Url,proto3" json:"Url,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *FetchRequest) Reset() {
	*x = FetchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_items_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRequest) ProtoMessage() {}

func (x *FetchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_items_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRequest.ProtoReflect.Descriptor instead.
func (*FetchRequest) Descriptor() ([]byte, []int) {
	return file_items_proto_rawDescGZIP(), []int{0}
}

func (x *FetchRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *FetchRequest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type SortField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string                 `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Price     float32                `protobuf:"fixed32,2,opt,name=Price,proto3" json:"Price,omitempty"`
	Changes   int32                  `protobuf:"varint,3,opt,name=Changes,proto3" json:"Changes,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *SortField) Reset() {
	*x = SortField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_items_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SortField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SortField) ProtoMessage() {}

func (x *SortField) ProtoReflect() protoreflect.Message {
	mi := &file_items_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SortField.ProtoReflect.Descriptor instead.
func (*SortField) Descriptor() ([]byte, []int) {
	return file_items_proto_rawDescGZIP(), []int{1}
}

func (x *SortField) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SortField) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *SortField) GetChanges() int32 {
	if x != nil {
		return x.Changes
	}
	return 0
}

func (x *SortField) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit        int64      `protobuf:"varint,1,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset       int64      `protobuf:"varint,2,opt,name=Offset,proto3" json:"Offset,omitempty"`
	SortingAsc   bool       `protobuf:"varint,3,opt,name=SortingAsc,proto3" json:"SortingAsc,omitempty"`
	SortingName  string     `protobuf:"bytes,4,opt,name=SortingName,proto3" json:"SortingName,omitempty"`
	SortingField *SortField `protobuf:"bytes,5,opt,name=SortingField,proto3" json:"SortingField,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_items_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_items_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_items_proto_rawDescGZIP(), []int{2}
}

func (x *ListRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListRequest) GetSortingAsc() bool {
	if x != nil {
		return x.SortingAsc
	}
	return false
}

func (x *ListRequest) GetSortingName() string {
	if x != nil {
		return x.SortingName
	}
	return ""
}

func (x *ListRequest) GetSortingField() *SortField {
	if x != nil {
		return x.SortingField
	}
	return nil
}

type FetchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Text string `protobuf:"bytes,2,opt,name=Text,proto3" json:"Text,omitempty"`
}

func (x *FetchResponse) Reset() {
	*x = FetchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_items_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchResponse) ProtoMessage() {}

func (x *FetchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_items_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchResponse.ProtoReflect.Descriptor instead.
func (*FetchResponse) Descriptor() ([]byte, []int) {
	return file_items_proto_rawDescGZIP(), []int{3}
}

func (x *FetchResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *FetchResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*ListResponseItems `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_items_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_items_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_items_proto_rawDescGZIP(), []int{4}
}

func (x *ListResponse) GetItems() []*ListResponseItems {
	if x != nil {
		return x.Items
	}
	return nil
}

type ListResponseItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string                 `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Price     int64                  `protobuf:"varint,2,opt,name=Price,proto3" json:"Price,omitempty"`
	Changes   int64                  `protobuf:"varint,3,opt,name=Changes,proto3" json:"Changes,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *ListResponseItems) Reset() {
	*x = ListResponseItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_items_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponseItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponseItems) ProtoMessage() {}

func (x *ListResponseItems) ProtoReflect() protoreflect.Message {
	mi := &file_items_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponseItems.ProtoReflect.Descriptor instead.
func (*ListResponseItems) Descriptor() ([]byte, []int) {
	return file_items_proto_rawDescGZIP(), []int{5}
}

func (x *ListResponseItems) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListResponseItems) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ListResponseItems) GetChanges() int64 {
	if x != nil {
		return x.Changes
	}
	return 0
}

func (x *ListResponseItems) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_items_proto protoreflect.FileDescriptor

var file_items_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x0c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x55, 0x72, 0x6c, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x89, 0x01, 0x0a, 0x09, 0x53, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xb3,
	0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x73, 0x63, 0x12, 0x20, 0x0a, 0x0b,
	0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34,
	0x0a, 0x0c, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x53, 0x6f, 0x72,
	0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0c, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x22, 0x37, 0x0a, 0x0d, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x22, 0x3e, 0x0a,
	0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x91, 0x01,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x32, 0x77, 0x0a, 0x0c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x34, 0x0a, 0x05, 0x46, 0x65, 0x74, 0x63, 0x68, 0x12, 0x13, 0x2e, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x12, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x72, 0x6b, 0x6f, 0x73, 0x68, 0x37,
	0x34, 0x34, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2f, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_items_proto_rawDescOnce sync.Once
	file_items_proto_rawDescData = file_items_proto_rawDesc
)

func file_items_proto_rawDescGZIP() []byte {
	file_items_proto_rawDescOnce.Do(func() {
		file_items_proto_rawDescData = protoimpl.X.CompressGZIP(file_items_proto_rawDescData)
	})
	return file_items_proto_rawDescData
}

var file_items_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_items_proto_goTypes = []interface{}{
	(*FetchRequest)(nil),          // 0: items.FetchRequest
	(*SortField)(nil),             // 1: items.SortField
	(*ListRequest)(nil),           // 2: items.ListRequest
	(*FetchResponse)(nil),         // 3: items.FetchResponse
	(*ListResponse)(nil),          // 4: items.ListResponse
	(*ListResponseItems)(nil),     // 5: items.ListResponseItems
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_items_proto_depIdxs = []int32{
	6, // 0: items.FetchRequest.created_at:type_name -> google.protobuf.Timestamp
	6, // 1: items.SortField.UpdatedAt:type_name -> google.protobuf.Timestamp
	1, // 2: items.ListRequest.SortingField:type_name -> items.SortField
	5, // 3: items.ListResponse.items:type_name -> items.ListResponseItems
	6, // 4: items.ListResponseItems.UpdatedAt:type_name -> google.protobuf.Timestamp
	0, // 5: items.ItemsService.Fetch:input_type -> items.FetchRequest
	2, // 6: items.ItemsService.List:input_type -> items.ListRequest
	3, // 7: items.ItemsService.Fetch:output_type -> items.FetchResponse
	4, // 8: items.ItemsService.List:output_type -> items.ListResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_items_proto_init() }
func file_items_proto_init() {
	if File_items_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_items_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchRequest); i {
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
		file_items_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SortField); i {
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
		file_items_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_items_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchResponse); i {
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
		file_items_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_items_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponseItems); i {
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
			RawDescriptor: file_items_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_items_proto_goTypes,
		DependencyIndexes: file_items_proto_depIdxs,
		MessageInfos:      file_items_proto_msgTypes,
	}.Build()
	File_items_proto = out.File
	file_items_proto_rawDesc = nil
	file_items_proto_goTypes = nil
	file_items_proto_depIdxs = nil
}
