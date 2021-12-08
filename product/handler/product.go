package handler

import (
	"context"
	"github.com/J-Y-Zhang/mall/product/common"
	"github.com/J-Y-Zhang/mall/product/domain/model"
	"github.com/J-Y-Zhang/mall/product/domain/service"
	"github.com/J-Y-Zhang/mall/product/proto/product"
)

type Product struct {
	ProductDataService service.ProductDataServiceInterface
}

func (p Product) AddProduct(ctx context.Context, req *product.ProductInfo, resp *product.AddProductResponse) error {
	newProduct := &model.Product{}
	err := common.SwapToByJson(req, newProduct)
	if err != nil {
		return err
	}
	id, err := p.ProductDataService.AddProduct(newProduct)
	if err != nil {
		return err
	}
	resp.ProductId = id

	return nil
}

func (p Product) FindProductByID(ctx context.Context, req *product.IDRequest, resp *product.ProductInfo) error {
	productData, err := p.ProductDataService.FindProductByID(req.ProductId)
	if err != nil {
		return err
	}
	return common.SwapToByJson(productData, resp)
}

func (p Product) UpdateProduct(ctx context.Context, req *product.ProductInfo, resp *product.CommonResponse) error {
	newProduct := &model.Product{}
	err := common.SwapToByJson(req, newProduct)
	if err != nil {
		return err
	}
	err = p.ProductDataService.UpdateProduct(newProduct)
	if err != nil {
		return err
	}
	resp.Msg = "更新成功"
	return nil
}

func (p Product) DeleteProductByID(ctx context.Context, req *product.IDRequest, resp *product.CommonResponse) error {
	err := p.ProductDataService.DeleteProduct(req.ProductId)
	if err != nil {
		return err
	}
	resp.Msg = "删除成功"
	return nil
}

func (p Product) FindAllProduct(ctx context.Context, req *product.FindAllRequest, resp *product.ProductListResponse) error {
	allProduct, err := p.ProductDataService.FindAllProduct()
	if err != nil {
		return err
	}
	for _, pro := range allProduct {
		proInfo := &product.ProductInfo{}
		err := common.SwapToByJson(pro, proInfo)
		if err != nil {
			return err
		}
		resp.ProductList = append(resp.ProductList, proInfo)
	}

	return nil
}
