package service

import (
	"github.com/J-Y-Zhang/mall/product/domain/dao"
	"github.com/J-Y-Zhang/mall/product/domain/model"
)

type ProductDataServiceInterface interface {
	AddProduct(*model.Product) (int64, error)
	DeleteProduct(int64) error
	UpdateProduct(*model.Product) error
	FindProductByID(int64) (*model.Product, error)
	FindAllProduct() ([]model.Product, error)
}

type ProductDataService struct {
	ProductDBManager dao.ProductDBManagerInterface
}

//创建
func NewProductDataService(productDBManager dao.ProductDBManagerInterface) ProductDataServiceInterface {
	return &ProductDataService{productDBManager}
}

//插入
func (u *ProductDataService) AddProduct(product *model.Product) (int64, error) {
	return u.ProductDBManager.CreateProduct(product)
}

//删除
func (u *ProductDataService) DeleteProduct(productID int64) error {
	return u.ProductDBManager.DeleteProductByID(productID)
}

//更新
func (u *ProductDataService) UpdateProduct(product *model.Product) error {
	return u.ProductDBManager.UpdateProduct(product)
}

//查找
func (u *ProductDataService) FindProductByID(productID int64) (*model.Product, error) {
	return u.ProductDBManager.FindProductByID(productID)
}

//查找
func (u *ProductDataService) FindAllProduct() ([]model.Product, error) {
	return u.ProductDBManager.FindAll()
}
