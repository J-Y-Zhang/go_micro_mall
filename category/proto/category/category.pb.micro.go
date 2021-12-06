// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/category/category.proto

package category

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Category service

func NewCategoryEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Category service

type CategoryService interface {
	CreateCategory(ctx context.Context, in *CategoryRequest, opts ...client.CallOption) (*CreateCategoryResponse, error)
	UpdateCategory(ctx context.Context, in *CategoryRequest, opts ...client.CallOption) (*UpdateCategoryResponse, error)
	DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...client.CallOption) (*DeleteCategoryResponse, error)
	FindCategoryByName(ctx context.Context, in *FindByNameRequest, opts ...client.CallOption) (*CategoryResponse, error)
	FindCategoryByID(ctx context.Context, in *FindByIDRequest, opts ...client.CallOption) (*CategoryResponse, error)
	FindCategoryByLevel(ctx context.Context, in *FindByLevelRequest, opts ...client.CallOption) (*CategoryListResponse, error)
	FindCategoryByParent(ctx context.Context, in *FindByParentRequest, opts ...client.CallOption) (*CategoryListResponse, error)
	FindAllCategory(ctx context.Context, in *FindAllRequest, opts ...client.CallOption) (*CategoryListResponse, error)
}

type categoryService struct {
	c    client.Client
	name string
}

func NewCategoryService(name string, c client.Client) CategoryService {
	return &categoryService{
		c:    c,
		name: name,
	}
}

func (c *categoryService) CreateCategory(ctx context.Context, in *CategoryRequest, opts ...client.CallOption) (*CreateCategoryResponse, error) {
	req := c.c.NewRequest(c.name, "Category.CreateCategory", in)
	out := new(CreateCategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) UpdateCategory(ctx context.Context, in *CategoryRequest, opts ...client.CallOption) (*UpdateCategoryResponse, error) {
	req := c.c.NewRequest(c.name, "Category.UpdateCategory", in)
	out := new(UpdateCategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...client.CallOption) (*DeleteCategoryResponse, error) {
	req := c.c.NewRequest(c.name, "Category.DeleteCategory", in)
	out := new(DeleteCategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) FindCategoryByName(ctx context.Context, in *FindByNameRequest, opts ...client.CallOption) (*CategoryResponse, error) {
	req := c.c.NewRequest(c.name, "Category.FindCategoryByName", in)
	out := new(CategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) FindCategoryByID(ctx context.Context, in *FindByIDRequest, opts ...client.CallOption) (*CategoryResponse, error) {
	req := c.c.NewRequest(c.name, "Category.FindCategoryByID", in)
	out := new(CategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) FindCategoryByLevel(ctx context.Context, in *FindByLevelRequest, opts ...client.CallOption) (*CategoryListResponse, error) {
	req := c.c.NewRequest(c.name, "Category.FindCategoryByLevel", in)
	out := new(CategoryListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) FindCategoryByParent(ctx context.Context, in *FindByParentRequest, opts ...client.CallOption) (*CategoryListResponse, error) {
	req := c.c.NewRequest(c.name, "Category.FindCategoryByParent", in)
	out := new(CategoryListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) FindAllCategory(ctx context.Context, in *FindAllRequest, opts ...client.CallOption) (*CategoryListResponse, error) {
	req := c.c.NewRequest(c.name, "Category.FindAllCategory", in)
	out := new(CategoryListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Category service

type CategoryHandler interface {
	CreateCategory(context.Context, *CategoryRequest, *CreateCategoryResponse) error
	UpdateCategory(context.Context, *CategoryRequest, *UpdateCategoryResponse) error
	DeleteCategory(context.Context, *DeleteCategoryRequest, *DeleteCategoryResponse) error
	FindCategoryByName(context.Context, *FindByNameRequest, *CategoryResponse) error
	FindCategoryByID(context.Context, *FindByIDRequest, *CategoryResponse) error
	FindCategoryByLevel(context.Context, *FindByLevelRequest, *CategoryListResponse) error
	FindCategoryByParent(context.Context, *FindByParentRequest, *CategoryListResponse) error
	FindAllCategory(context.Context, *FindAllRequest, *CategoryListResponse) error
}

func RegisterCategoryHandler(s server.Server, hdlr CategoryHandler, opts ...server.HandlerOption) error {
	type category interface {
		CreateCategory(ctx context.Context, in *CategoryRequest, out *CreateCategoryResponse) error
		UpdateCategory(ctx context.Context, in *CategoryRequest, out *UpdateCategoryResponse) error
		DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, out *DeleteCategoryResponse) error
		FindCategoryByName(ctx context.Context, in *FindByNameRequest, out *CategoryResponse) error
		FindCategoryByID(ctx context.Context, in *FindByIDRequest, out *CategoryResponse) error
		FindCategoryByLevel(ctx context.Context, in *FindByLevelRequest, out *CategoryListResponse) error
		FindCategoryByParent(ctx context.Context, in *FindByParentRequest, out *CategoryListResponse) error
		FindAllCategory(ctx context.Context, in *FindAllRequest, out *CategoryListResponse) error
	}
	type Category struct {
		category
	}
	h := &categoryHandler{hdlr}
	return s.Handle(s.NewHandler(&Category{h}, opts...))
}

type categoryHandler struct {
	CategoryHandler
}

func (h *categoryHandler) CreateCategory(ctx context.Context, in *CategoryRequest, out *CreateCategoryResponse) error {
	return h.CategoryHandler.CreateCategory(ctx, in, out)
}

func (h *categoryHandler) UpdateCategory(ctx context.Context, in *CategoryRequest, out *UpdateCategoryResponse) error {
	return h.CategoryHandler.UpdateCategory(ctx, in, out)
}

func (h *categoryHandler) DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, out *DeleteCategoryResponse) error {
	return h.CategoryHandler.DeleteCategory(ctx, in, out)
}

func (h *categoryHandler) FindCategoryByName(ctx context.Context, in *FindByNameRequest, out *CategoryResponse) error {
	return h.CategoryHandler.FindCategoryByName(ctx, in, out)
}

func (h *categoryHandler) FindCategoryByID(ctx context.Context, in *FindByIDRequest, out *CategoryResponse) error {
	return h.CategoryHandler.FindCategoryByID(ctx, in, out)
}

func (h *categoryHandler) FindCategoryByLevel(ctx context.Context, in *FindByLevelRequest, out *CategoryListResponse) error {
	return h.CategoryHandler.FindCategoryByLevel(ctx, in, out)
}

func (h *categoryHandler) FindCategoryByParent(ctx context.Context, in *FindByParentRequest, out *CategoryListResponse) error {
	return h.CategoryHandler.FindCategoryByParent(ctx, in, out)
}

func (h *categoryHandler) FindAllCategory(ctx context.Context, in *FindAllRequest, out *CategoryListResponse) error {
	return h.CategoryHandler.FindAllCategory(ctx, in, out)
}