package main

import (
	"context"
	"fmt"
	"github.com/J-Y-Zhang/mall/product/common"
	"github.com/J-Y-Zhang/mall/product/proto/product"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	opentracing2 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4"
	_ "github.com/go-sql-driver/mysql"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"log"
)

func main() {
	//注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"82.156.19.233:8500",
		}
	})

	//链路追踪
	//参一是微服务的服务名, 参二是jaeger的collecter地址
	tracer, closer, err := common.NewTracer("product.client", "82.156.19.233:6831")
	if err != nil {
		log.Println(err)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	// New Service
	service := micro.NewService(
		micro.Name("product.client"),
		micro.Version("latest"),
		micro.Registry(consulRegistry),
		//绑定链路追踪(Handler)
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
	)

	productService := product.NewProductService("product", service.Client())

	productAdd := &product.ProductInfo{
		ProductName:        "imooc",
		ProductSku:         "lwk",
		ProductPrice:       1.1,
		ProductDescription: "imooc-lwk",
		ProductCategoryId:  1,
		ProductImages: []*product.ProductImage{
			{
				ImageName: "lwk-image",
				ImageCode: "lwkimage01",
				ImageUrl:  "lwkimage01",
			},
			{
				ImageName: "lwk-image02",
				ImageCode: "lwkimage02",
				ImageUrl:  "lwkimage02",
			},
		},
		ProductSizes: []*product.ProductSize{
			{
				SizeName: "lwk-size",
				SizeCode: "lwk-size-code",
			},
		},
		ProductSeo: &product.ProductSeo{
			SeoTitle:       "lwk-seo",
			SeoKeywords:    "lwk-seo",
			SeoDescription: "lwk-seo",
			SeoCode:        "lwk-seo",
		},
	}
	response, err := productService.AddProduct(context.TODO(), productAdd)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}
