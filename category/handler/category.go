package handler

import (
    context "context"
    "github.com/J-Y-Zhang/mall/category/common"
    "github.com/J-Y-Zhang/mall/category/domain/model"
    "github.com/J-Y-Zhang/mall/category/domain/service"
    "github.com/J-Y-Zhang/mall/category/proto/category"
    "go-micro.dev/v4/util/log"
)

type Category struct {
    CategoryDataSrv service.CategoryDataServiceInterface
}

func (c Category) CreateCategory(ctx context.Context, request *category.CategoryRequest, response *category.CreateCategoryResponse) error {
    category := &model.Category{}
    err := common.SwapToByJson(request, category)
    if err != nil {
        return err
    }
    id, err := c.CategoryDataSrv.AddCategory(category)
    if err != nil {
        return err
    }
    response.CategoryId = id
    response.Message = "添加成功"
    return nil
}

func (c Category) UpdateCategory(ctx context.Context, request *category.CategoryRequest, response *category.UpdateCategoryResponse) error {
    category := &model.Category{}
    err := common.SwapToByJson(request, category)
    if err != nil {
        return err
    }
    err = c.CategoryDataSrv.UpdateCategory(category)
    if err != nil {
        return err
    }
    response.Message = "更新成功"
    return nil
}

func (c Category) DeleteCategory(ctx context.Context, request *category.DeleteCategoryRequest, response *category.DeleteCategoryResponse) error {
    err := c.CategoryDataSrv.DeleteCategory(request.CategoryId)
    if err != nil {
        return err
    }
    response.Message = "删除成功"
    return nil
}

func (c Category) FindCategoryByName(ctx context.Context, request *category.FindByNameRequest, response *category.CategoryResponse) error {
    category, err := c.CategoryDataSrv.FindCategoryByName(request.CategoryName)
    if err != nil {
        return err
    }
    return common.SwapToByJson(category, response)
}

func (c Category) FindCategoryByID(ctx context.Context, request *category.FindByIDRequest, response *category.CategoryResponse) error {
    category, err := c.CategoryDataSrv.FindCategoryByID(request.CategoryId)
    if err != nil {
        return err
    }
    return common.SwapToByJson(category, response)
}

func CategoryList2Response(list []model.Category, response *category.CategoryListResponse){
    for _, cg := range list{
        obj := &category.CategoryResponse{}
        err := common.SwapToByJson(cg, obj)
        if err != nil {
            log.Error(err)
            break
        }
        response.CategoryList = append(response.CategoryList, obj)
    }
}

func (c Category) FindCategoryByLevel(ctx context.Context, request *category.FindByLevelRequest, response *category.CategoryListResponse) error {
    categorySlice, err := c.CategoryDataSrv.FindCategoryByLevel(request.CategoryLevel)
    if err != nil {
        return err
    }
    CategoryList2Response(categorySlice, response)
    return nil
}

func (c Category) FindCategoryByParent(ctx context.Context, request *category.FindByParentRequest, response *category.CategoryListResponse) error {
    categorySlice, err := c.CategoryDataSrv.FindCategoryByParent(request.CategoryParent)
    if err != nil {
        return err
    }
    CategoryList2Response(categorySlice, response)
    return nil
}

func (c Category) FindAllCategory(ctx context.Context, request *category.FindAllRequest, response *category.CategoryListResponse) error {
    categorySlice, err := c.CategoryDataSrv.FindAllCategory()
    if err != nil {
        return err
    }
    CategoryList2Response(categorySlice, response)
    return nil
}
