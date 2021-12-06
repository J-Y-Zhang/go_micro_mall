package service

import (
    "github.com/J-Y-Zhang/mall/category/domain/dao"
    "github.com/J-Y-Zhang/mall/category/domain/model"
)

type CategoryDataServiceInterface interface {
    AddCategory(*model.Category) (int64, error)
    DeleteCategory(int64) error
    UpdateCategory(*model.Category) error
    FindCategoryByName(string) (*model.Category, error)
    FindCategoryByID(int64) (*model.Category, error)
    FindAllCategory() ([]*model.Category, error)
    FindCategoryByLevel(uint32) ([]*model.Category, error)
    FindCategoryByParent(int64) ([]*model.Category, error)
}

type CategoryDataService struct {
    categoryDBManager dao.CategoryDBManagerInterface
}

func NewCategoryDataService(manager dao.CategoryDBManagerInterface) CategoryDataServiceInterface {
    return &CategoryDataService{
        categoryDBManager: manager,
    }
}

func (u CategoryDataService) FindCategoryByID(id int64) (*model.Category, error) {
    return u.categoryDBManager.FindCategoryByID(id)
}

func (u CategoryDataService) FindAllCategory() ([]*model.Category, error) {
    return u.categoryDBManager.FindAll()
}

func (u CategoryDataService) FindCategoryByLevel(lv uint32) ([]*model.Category, error) {
    return u.categoryDBManager.FindCategoryByLevel(lv)
}

func (u CategoryDataService) FindCategoryByParent(parent_id int64) ([]*model.Category, error) {
    return u.categoryDBManager.FindCategoryByParent(parent_id)
}


func (u CategoryDataService) AddCategory(category *model.Category) (int64, error) {
    return u.categoryDBManager.CreateCategory(category)
}

func (u CategoryDataService) DeleteCategory(id int64) error {
    return u.categoryDBManager.DeleteCategoryByID(id)
}

func (u CategoryDataService) UpdateCategory(category *model.Category) error {
    return u.categoryDBManager.UpdateCategory(category)
}

func (u CategoryDataService) FindCategoryByName(name string) (*model.Category, error) {
    return u.categoryDBManager.FindCategoryByName(name)
}


