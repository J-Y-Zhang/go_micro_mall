package dao

import (
	"github.com/J-Y-Zhang/mall/user/domain/model"
    "github.com/jinzhu/gorm"
)

type UserDBManagerInterface interface {
    //初始化数据库表
    InitTable() error

    //根据用户名称查找用户信息
    FindUserByName(string) (*model.User, error)
    //根据用户ID查找用户信息
    FindUserByID(int64) (*model.User, error)

    //创建用户
    CreateUser(user *model.User) (int64, error)

    //根据ID删除用户
    DeleteUserByID(int64) error

    //更新用户信息
    UpdateUser(*model.User) error

    //查找全部用户
    FindAll() ([]*model.User, error)
}

type UserDBManager struct {
    mysqlDB *gorm.DB
}

func (u UserDBManager) FindAll() (res []*model.User, err error) {
    return res, u.mysqlDB.Find(&res).Error
}

func (u UserDBManager) InitTable() error {
    return u.mysqlDB.CreateTable(&model.User{}).Error
}

func (u UserDBManager) FindUserByName(name string) (*model.User, error) {
    user := &model.User{}
    return user, u.mysqlDB.Where("user_name = ?", name).First(user).Error
}

func (u UserDBManager) FindUserByID(id int64) (*model.User, error) {
    user := &model.User{}
    return user, u.mysqlDB.First(user, id).Error
}

func (u UserDBManager) CreateUser(user *model.User) (int64, error) {
    return user.ID, u.mysqlDB.Create(user).Error
}

func (u UserDBManager) DeleteUserByID(id int64) error {
    return u.mysqlDB.Delete(&model.User{}, id).Error
}

func (u UserDBManager) UpdateUser(user *model.User) error {
    return u.mysqlDB.Model(user).Update(user).Error
}

func NewUserDBManager(db *gorm.DB) UserDBManagerInterface {
    return &UserDBManager{
        mysqlDB: db,
    }
}
