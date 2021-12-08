package dao

import (
	"github.com/J-Y-Zhang/mall/product/domain/model"
	"github.com/jinzhu/gorm"
)

type ProductDBManagerInterface interface {
	//初始化数据库表
	InitTable() error
	FindProductByID(int64) (*model.Product, error)
	CreateProduct(*model.Product) (int64, error)
	DeleteProductByID(int64) error
	UpdateProduct(*model.Product) error
	FindAll() ([]model.Product, error)
}

type ProductDBManager struct {
	MysqlDB *gorm.DB
}

//初始化表
func (u *ProductDBManager) InitTable() error {
	return u.MysqlDB.CreateTable(&model.ProductSeo{}, &model.ProductImage{}, &model.ProductSize{}, &model.Product{}).Error
}

//根据ID查找Product信息
func (u *ProductDBManager) FindProductByID(productID int64) (product *model.Product, err error) {
	product = &model.Product{}
	return product, u.MysqlDB.First(product, productID).Error
}

//创建Product信息
func (u *ProductDBManager) CreateProduct(product *model.Product) (int64, error) {
	return product.ID, u.MysqlDB.Create(product).Error
}

//根据ID删除Product信息
func (u *ProductDBManager) DeleteProductByID(productID int64) error {
	//开启事务
	tx := u.MysqlDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	//删除
	if err := tx.Unscoped().Where("id = ?", productID).Delete(&model.Product{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("image_product_id = ?", productID).Delete(&model.ProductImage{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("size_product_id = ?", productID).Delete(&model.ProductSize{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("seo_product_id = ?", productID).Delete(&model.ProductSeo{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

//更新Product信息
func (u *ProductDBManager) UpdateProduct(product *model.Product) error {
	return u.MysqlDB.Model(product).Update(product).Error
}

//获取结果集
func (u *ProductDBManager) FindAll() (productAll []model.Product, err error) {
	return productAll, u.MysqlDB.Preload("ProductImages").Preload("ProductSizes").Preload("ProductSeo").Find(&productAll).Error
}
