// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: proto/category/category.proto

package category

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

type CategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryName string `protobuf:"bytes,1,opt,name=category_name,json=categoryName,proto3" json:"category_name,omitempty"`
	//多级分类
	CategoryLevel       uint32 `protobuf:"varint,2,opt,name=category_level,json=categoryLevel,proto3" json:"category_level,omitempty"`
	CategoryParent      int64  `protobuf:"varint,3,opt,name=category_parent,json=categoryParent,proto3" json:"category_parent,omitempty"`
	CategoryImage       string `protobuf:"bytes,4,opt,name=category_image,json=categoryImage,proto3" json:"category_image,omitempty"`
	CategoryDescription string `protobuf:"bytes,5,opt,name=category_description,json=categoryDescription,proto3" json:"category_description,omitempty"`
}

func (x *CategoryRequest) Reset() {
	*x = CategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_category_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryRequest) ProtoMessage() {}

func (x *CategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_category_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryRequest.ProtoReflect.Descriptor instead.
func (*CategoryRequest) Descriptor() ([]byte, []int) {
	return file_proto_category_category_proto_rawDescGZIP(), []int{0}
}

func (x *CategoryRequest) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

func (x *CategoryRequest) GetCategoryLevel() uint32 {
	if x != nil {
		return x.CategoryLevel
	}
	return 0
}

func (x *CategoryRequest) GetCategoryParent() int64 {
	if x != nil {
		return x.CategoryParent
	}
	return 0
}

func (x *CategoryRequest) GetCategoryImage() string {
	if x != nil {
		return x.CategoryImage
	}
	return ""
}

func (x *CategoryRequest) GetCategoryDescription() string {
	if x != nil {
		return x.CategoryDescription
	}
	return ""
}

type CreateCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message    string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	CategoryId int64  `protobuf:"varint,2,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

func (x *CreateCategoryResponse) Reset() {
	*x = CreateCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_category_category_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCategoryResponse) ProtoMessage() {}

func (x *CreateCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_category_category_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCategoryResponse.ProtoReflect.Descriptor instead.
func (*CreateCategoryResponse) Descriptor() ([]byte, []int) {
	return file_proto_category_category_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCategoryResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateCategoryResponse) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

type UpdateCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateCategoryResponse) Reset() {
	*x = UpdateCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_category_category_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCategoryResponse) ProtoMessage() {}

func (x *UpdateCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_category_category_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCategoryResponse.ProtoReflect.Descriptor instead.
func (*UpdateCategoryResponse) Descriptor() ([]byte, []int) {
	return file_proto_category_category_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateCategoryResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId int64 `protobuf:"varint,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

func (x *DeleteCategoryRequest) Reset() {
	*x = DeleteCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_category_category_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCategoryRequest) ProtoMessage() {}

func (x *DeleteCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_category_category_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCategoryRequest.ProtoReflect.Descriptor instead.
func (*DeleteCategoryRequest) Descriptor() ([]byte, []int) {
	return file_proto_category_category_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteCategoryRequest) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

type DeleteCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteCategoryResponse) Reset() {
	*x = DeleteCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_category_category_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCategoryResponse) ProtoMessage() {}

func (x *DeleteCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_category_category_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCategoryResponse.ProtoReflect.Descriptor instead.
func (*DeleteCategoryResponse) Descriptor() ([]byte, []int) {
	return file_proto_category_category_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteCategoryResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type FindByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryName string `protobuf:"bytes,1,opt,name=category_name,json=categoryName,proto3" json:"category_name,omitempty"`
}

func (x *FindByNameRequest) Reset() {
	*x = FindByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_category_category_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByNameRequest) ProtoMessage() {}

func (x *FindByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_category_category_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByNameRequest.ProtoReflect.Descriptor instead.
func (*FindByNameRequest) Descriptor() ([]byte, []int) {
	return file_proto_category_category_proto_rawDescGZIP(), []int{5}
}

func (x *FindByNameRequest) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

type CategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId          int64  `protobuf:"varint,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	CategoryName        string `protobuf:"bytes,2,opt,name=category_name,json=categoryName,proto3" json:"category_name,omitempty"`
	CategoryLevel       uint32 `protobuf:"varint,3,opt,name=category_level,json=categoryLevel,proto3" json:"category_level,omitempty"`
	CategoryParent      int64  `protobuf:"varint,4,opt,name=category_parent,json=categoryParent,proto3" json:"category_parent,omitempty"`
	CategoryImage       string `protobuf:"bytes,5,opt,name=category_image,json=categoryImage,proto3" json:"category_image,omitempty"`
	CategoryDescription string `protobuf:"bytes,6,opt,name=category_description,json=categoryDescription,proto3" json:"category_description,omitempty"`
}

func (x *CategoryResponse) Reset() {
	*x = CategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_category_category_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryResponse) ProtoMessage() {}

func (x *CategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_category_category_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryResponse.ProtoReflect.Descriptor instead.
func (*CategoryResponse) Descriptor() ([]byte, []int) {
	return file_proto_category_category_proto_rawDescGZIP(), []int{6}
}

func (x *CategoryResponse) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *CategoryResponse) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

func (x *CategoryResponse) GetCategoryLevel() uint32 {
	if x != nil {
		return x.CategoryLevel
	}
	return 0
}

func (x *CategoryResponse) GetCategoryParent() int64 {
	if x != nil {
		return x.CategoryParent
	}
	return 0
}

func (x *CategoryResponse) GetCategoryImage() string {
	if x != nil {
		return x.CategoryImage
	}
	return ""
}

func (x *CategoryResponse) GetCategoryDescription() string {
	if x != nil {
		return x.CategoryDescription
	}
	return ""
}

type FindByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId string `protobuf:"bytes,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

func (x *FindByIDRequest) Reset() {
	*x = FindByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_category_category_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByIDRequest) ProtoMessage() {}

func (x *FindByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_category_category_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByIDRequest.ProtoReflect.Descriptor instead.
func (*FindByIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_category_category_proto_rawDescGZIP(), []int{7}
}

func (x *FindByIDRequest) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

type FindByLevelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryLevel uint32 `protobuf:"varint,1,opt,name=category_level,json=categoryLevel,proto3" json:"category_level,omitempty"`
}

func (x *FindByLevelRequest) Reset() {
	*x = FindByLevelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_category_category_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByLevelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByLevelRequest) ProtoMessage() {}

func (x *FindByLevelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_category_category_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByLevelRequest.ProtoReflect.Descriptor instead.
func (*FindByLevelRequest) Descriptor() ([]byte, []int) {
	return file_proto_category_category_proto_rawDescGZIP(), []int{8}
}

func (x *FindByLevelRequest) GetCategoryLevel() uint32 {
	if x != nil {
		return x.CategoryLevel
	}
	return 0
}

type FindByParentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryParent int64 `protobuf:"varint,1,opt,name=category_parent,json=categoryParent,proto3" json:"category_parent,omitempty"`
}

func (x *FindByParentRequest) Reset() {
	*x = FindByParentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_category_category_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByParentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByParentRequest) ProtoMessage() {}

func (x *FindByParentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_category_category_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByParentRequest.ProtoReflect.Descriptor instead.
func (*FindByParentRequest) Descriptor() ([]byte, []int) {
	return file_proto_category_category_proto_rawDescGZIP(), []int{9}
}

func (x *FindByParentRequest) GetCategoryParent() int64 {
	if x != nil {
		return x.CategoryParent
	}
	return 0
}

type FindAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindAllRequest) Reset() {
	*x = FindAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_category_category_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllRequest) ProtoMessage() {}

func (x *FindAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_category_category_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllRequest.ProtoReflect.Descriptor instead.
func (*FindAllRequest) Descriptor() ([]byte, []int) {
	return file_proto_category_category_proto_rawDescGZIP(), []int{10}
}

type CategoryListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryList []*CategoryResponse `protobuf:"bytes,1,rep,name=categoryList,proto3" json:"categoryList,omitempty"`
}

func (x *CategoryListResponse) Reset() {
	*x = CategoryListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_category_category_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryListResponse) ProtoMessage() {}

func (x *CategoryListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_category_category_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryListResponse.ProtoReflect.Descriptor instead.
func (*CategoryListResponse) Descriptor() ([]byte, []int) {
	return file_proto_category_category_proto_rawDescGZIP(), []int{11}
}

func (x *CategoryListResponse) GetCategoryList() []*CategoryResponse {
	if x != nil {
		return x.CategoryList
	}
	return nil
}

var File_proto_category_category_proto protoreflect.FileDescriptor

var file_proto_category_category_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0xe0, 0x01, 0x0a, 0x0f, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a,
	0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x50, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x31, 0x0a, 0x14, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x53, 0x0a, 0x16,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49,
	0x64, 0x22, 0x32, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x38, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22,
	0x32, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x38, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x82, 0x02,
	0x0a, 0x10, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x27, 0x0a, 0x0f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x31, 0x0a, 0x14, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x32, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0x3b, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x22, 0x3e, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x50, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x50, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x56, 0x0a, 0x14, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a,
	0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xa0, 0x05,
	0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x4f, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x19, 0x2e, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x19, 0x2e,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1f,
	0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x2e, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x49, 0x44, 0x12, 0x19, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x55, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x42, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1c, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x12, 0x1d, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x42, 0x79, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4d, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x1b, 0x5a, 0x19, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x3b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_category_category_proto_rawDescOnce sync.Once
	file_proto_category_category_proto_rawDescData = file_proto_category_category_proto_rawDesc
)

func file_proto_category_category_proto_rawDescGZIP() []byte {
	file_proto_category_category_proto_rawDescOnce.Do(func() {
		file_proto_category_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_category_category_proto_rawDescData)
	})
	return file_proto_category_category_proto_rawDescData
}

var file_proto_category_category_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_category_category_proto_goTypes = []interface{}{
	(*CategoryRequest)(nil),        // 0: category.CategoryRequest
	(*CreateCategoryResponse)(nil), // 1: category.CreateCategoryResponse
	(*UpdateCategoryResponse)(nil), // 2: category.UpdateCategoryResponse
	(*DeleteCategoryRequest)(nil),  // 3: category.DeleteCategoryRequest
	(*DeleteCategoryResponse)(nil), // 4: category.DeleteCategoryResponse
	(*FindByNameRequest)(nil),      // 5: category.FindByNameRequest
	(*CategoryResponse)(nil),       // 6: category.CategoryResponse
	(*FindByIDRequest)(nil),        // 7: category.FindByIDRequest
	(*FindByLevelRequest)(nil),     // 8: category.FindByLevelRequest
	(*FindByParentRequest)(nil),    // 9: category.FindByParentRequest
	(*FindAllRequest)(nil),         // 10: category.FindAllRequest
	(*CategoryListResponse)(nil),   // 11: category.CategoryListResponse
}
var file_proto_category_category_proto_depIdxs = []int32{
	6,  // 0: category.CategoryListResponse.categoryList:type_name -> category.CategoryResponse
	0,  // 1: category.Category.CreateCategory:input_type -> category.CategoryRequest
	0,  // 2: category.Category.UpdateCategory:input_type -> category.CategoryRequest
	3,  // 3: category.Category.DeleteCategory:input_type -> category.DeleteCategoryRequest
	5,  // 4: category.Category.FindCategoryByName:input_type -> category.FindByNameRequest
	7,  // 5: category.Category.FindCategoryByID:input_type -> category.FindByIDRequest
	8,  // 6: category.Category.FindCategoryByLevel:input_type -> category.FindByLevelRequest
	9,  // 7: category.Category.FindCategoryByParent:input_type -> category.FindByParentRequest
	10, // 8: category.Category.FindAllCategory:input_type -> category.FindAllRequest
	1,  // 9: category.Category.CreateCategory:output_type -> category.CreateCategoryResponse
	2,  // 10: category.Category.UpdateCategory:output_type -> category.UpdateCategoryResponse
	4,  // 11: category.Category.DeleteCategory:output_type -> category.DeleteCategoryResponse
	6,  // 12: category.Category.FindCategoryByName:output_type -> category.CategoryResponse
	6,  // 13: category.Category.FindCategoryByID:output_type -> category.CategoryResponse
	11, // 14: category.Category.FindCategoryByLevel:output_type -> category.CategoryListResponse
	11, // 15: category.Category.FindCategoryByParent:output_type -> category.CategoryListResponse
	11, // 16: category.Category.FindAllCategory:output_type -> category.CategoryListResponse
	9,  // [9:17] is the sub-list for method output_type
	1,  // [1:9] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_proto_category_category_proto_init() }
func file_proto_category_category_proto_init() {
	if File_proto_category_category_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_category_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryRequest); i {
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
		file_proto_category_category_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCategoryResponse); i {
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
		file_proto_category_category_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCategoryResponse); i {
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
		file_proto_category_category_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCategoryRequest); i {
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
		file_proto_category_category_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCategoryResponse); i {
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
		file_proto_category_category_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByNameRequest); i {
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
		file_proto_category_category_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryResponse); i {
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
		file_proto_category_category_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByIDRequest); i {
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
		file_proto_category_category_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByLevelRequest); i {
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
		file_proto_category_category_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByParentRequest); i {
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
		file_proto_category_category_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllRequest); i {
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
		file_proto_category_category_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryListResponse); i {
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
			RawDescriptor: file_proto_category_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_category_category_proto_goTypes,
		DependencyIndexes: file_proto_category_category_proto_depIdxs,
		MessageInfos:      file_proto_category_category_proto_msgTypes,
	}.Build()
	File_proto_category_category_proto = out.File
	file_proto_category_category_proto_rawDesc = nil
	file_proto_category_category_proto_goTypes = nil
	file_proto_category_category_proto_depIdxs = nil
}
