// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/product/product.proto

package product

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

// Api Endpoints for Product service

func NewProductEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Product service

type ProductService interface {
	AddProduct(ctx context.Context, in *ProductInfo, opts ...client.CallOption) (*AddProductResponse, error)
	FindProductByID(ctx context.Context, in *IDRequest, opts ...client.CallOption) (*ProductInfo, error)
	UpdateProduct(ctx context.Context, in *ProductInfo, opts ...client.CallOption) (*CommonResponse, error)
	DeleteProductByID(ctx context.Context, in *IDRequest, opts ...client.CallOption) (*CommonResponse, error)
	FindAllProduct(ctx context.Context, in *FindAllRequest, opts ...client.CallOption) (*ProductListResponse, error)
}

type productService struct {
	c    client.Client
	name string
}

func NewProductService(name string, c client.Client) ProductService {
	return &productService{
		c:    c,
		name: name,
	}
}

func (c *productService) AddProduct(ctx context.Context, in *ProductInfo, opts ...client.CallOption) (*AddProductResponse, error) {
	req := c.c.NewRequest(c.name, "Product.AddProduct", in)
	out := new(AddProductResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) FindProductByID(ctx context.Context, in *IDRequest, opts ...client.CallOption) (*ProductInfo, error) {
	req := c.c.NewRequest(c.name, "Product.FindProductByID", in)
	out := new(ProductInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) UpdateProduct(ctx context.Context, in *ProductInfo, opts ...client.CallOption) (*CommonResponse, error) {
	req := c.c.NewRequest(c.name, "Product.UpdateProduct", in)
	out := new(CommonResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) DeleteProductByID(ctx context.Context, in *IDRequest, opts ...client.CallOption) (*CommonResponse, error) {
	req := c.c.NewRequest(c.name, "Product.DeleteProductByID", in)
	out := new(CommonResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) FindAllProduct(ctx context.Context, in *FindAllRequest, opts ...client.CallOption) (*ProductListResponse, error) {
	req := c.c.NewRequest(c.name, "Product.FindAllProduct", in)
	out := new(ProductListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Product service

type ProductHandler interface {
	AddProduct(context.Context, *ProductInfo, *AddProductResponse) error
	FindProductByID(context.Context, *IDRequest, *ProductInfo) error
	UpdateProduct(context.Context, *ProductInfo, *CommonResponse) error
	DeleteProductByID(context.Context, *IDRequest, *CommonResponse) error
	FindAllProduct(context.Context, *FindAllRequest, *ProductListResponse) error
}

func RegisterProductHandler(s server.Server, hdlr ProductHandler, opts ...server.HandlerOption) error {
	type product interface {
		AddProduct(ctx context.Context, in *ProductInfo, out *AddProductResponse) error
		FindProductByID(ctx context.Context, in *IDRequest, out *ProductInfo) error
		UpdateProduct(ctx context.Context, in *ProductInfo, out *CommonResponse) error
		DeleteProductByID(ctx context.Context, in *IDRequest, out *CommonResponse) error
		FindAllProduct(ctx context.Context, in *FindAllRequest, out *ProductListResponse) error
	}
	type Product struct {
		product
	}
	h := &productHandler{hdlr}
	return s.Handle(s.NewHandler(&Product{h}, opts...))
}

type productHandler struct {
	ProductHandler
}

func (h *productHandler) AddProduct(ctx context.Context, in *ProductInfo, out *AddProductResponse) error {
	return h.ProductHandler.AddProduct(ctx, in, out)
}

func (h *productHandler) FindProductByID(ctx context.Context, in *IDRequest, out *ProductInfo) error {
	return h.ProductHandler.FindProductByID(ctx, in, out)
}

func (h *productHandler) UpdateProduct(ctx context.Context, in *ProductInfo, out *CommonResponse) error {
	return h.ProductHandler.UpdateProduct(ctx, in, out)
}

func (h *productHandler) DeleteProductByID(ctx context.Context, in *IDRequest, out *CommonResponse) error {
	return h.ProductHandler.DeleteProductByID(ctx, in, out)
}

func (h *productHandler) FindAllProduct(ctx context.Context, in *FindAllRequest, out *ProductListResponse) error {
	return h.ProductHandler.FindAllProduct(ctx, in, out)
}
