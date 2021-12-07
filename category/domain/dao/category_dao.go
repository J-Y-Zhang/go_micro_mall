package dao

import (
	"github.com/J-Y-Zhang/mall/category/domain/model"
    "github.com/jinzhu/gorm"
)

type CategoryDBManagerInterface interface {
    //初始化数据库表
    InitTable() error
    FindCategoryByID(int64) (*model.Category, error)
    CreateCategory(*model.Category) (int64, error)
    DeleteCategoryByID(int64) error
    UpdateCategory(*model.Category) error
    FindAll() ([]model.Category, error)
    FindCategoryByLevel(uint32) ([]model.Category, error)
    FindCategoryByParent(int64) ([]model.Category, error)
    FindCategoryByName(string) (*model.Category, error)
}

type CategoryDBManager struct {
    MysqlDB *gorm.DB
}

func (u CategoryDBManager) UpdateCategory(category *model.Category) error {
    return u.MysqlDB.Model(category).Update(category).Error
}

func (u CategoryDBManager) FindCategoryByLevel(lv uint32) (res []model.Category, err error) {
    return res, u.MysqlDB.Where("category_level = ?", lv).Find(res).Error
}

func (u CategoryDBManager) FindCategoryByParent(parent_id int64) (res []model.Category, err error) {
    return res, u.MysqlDB.Where("category_parent = ?", parent_id).Find(res).Error
}


func (u CategoryDBManager) FindAll() (res []model.Category, err error) {
    return res, u.MysqlDB.Find(&res).Error
}

func (u CategoryDBManager) InitTable() error {
    return u.MysqlDB.CreateTable(&model.Category{}).Error
}

func (u CategoryDBManager) FindCategoryByName(name string) (*model.Category, error) {
    category := &model.Category{}
    return category, u.MysqlDB.Where("category_name = ?", name).First(category).Error
}

func (u CategoryDBManager) FindCategoryByID(id int64) (*model.Category, error) {
    category := &model.Category{}
    return category, u.MysqlDB.First(category, id).Error
}

func (u CategoryDBManager) CreateCategory(category *model.Category) (int64, error) {
    return category.ID, u.MysqlDB.Create(category).Error
}

func (u CategoryDBManager) DeleteCategoryByID(id int64) error {
    return u.MysqlDB.Delete(&model.Category{}, id).Error
}

