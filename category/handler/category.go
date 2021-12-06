package handler

import (
    context "context"
    "github.com/J-Y-Zhang/mall/category/domain/model"
    "github.com/J-Y-Zhang/mall/category/domain/service"
    "github.com/J-Y-Zhang/mall/category/proto/category"
)

type Category struct {
    CategoryDataSrv service.CategoryDataService
}

func (c Category) CreateCategory(ctx context.Context, request *category.CategoryRequest, response *category.CreateCategoryResponse) error {
    return c.CategoryDataSrv.AddCategory(&model.Category{
        CategoryName:        request.CategoryName,
        CategoryLevel:       request.CategoryLevel,
        CategoryParent:      request.,
        CategoryImage:       "",
        CategoryDescription: "",
    })
}

func (c Category) UpdateCategory(ctx context.Context, request *category.CategoryRequest, response *category.UpdateCategoryResponse) error {
    panic("implement me")
}

func (c Category) DeleteCategory(ctx context.Context, request *category.DeleteCategoryRequest, response *category.DeleteCategoryResponse) error {
    panic("implement me")
}

func (c Category) FindCategoryByName(ctx context.Context, request *category.FindByNameRequest, response *category.CategoryResponse) error {
    panic("implement me")
}

func (c Category) FindCategoryByID(ctx context.Context, request *category.FindByIDRequest, response *category.CategoryResponse) error {
    panic("implement me")
}

func (c Category) FindCategoryByLevel(ctx context.Context, request *category.FindByLevelRequest, response *category.CategoryListResponse) error {
    panic("implement me")
}

func (c Category) FindCategoryByParent(ctx context.Context, request *category.FindByParentRequest, response *category.CategoryListResponse) error {
    panic("implement me")
}

func (c Category) FindAllCategory(ctx context.Context, request *category.FindAllRequest, response *category.CategoryListResponse) error {
    panic("implement me")
}
