package service

import (
    "errors"
    "github.com/J-Y-Zhang/mall/user/domain/dao"
    "github.com/J-Y-Zhang/mall/user/domain/model"
    "golang.org/x/crypto/bcrypt"
)

type UserDataServiceInterface interface {
    AddUser(*model.User) (int64, error)
    DeleteUser(int64) error
    UpdateUser(*model.User, bool) error
    FindUserByName(string) (*model.User, error)
    CheckPassword(string, string) (bool, error)
}

type UserDataService struct {
    userDBManager dao.UserDBManagerInterface
}

func NewUserDataService(manager dao.UserDBManagerInterface) UserDataServiceInterface {
    return &UserDataService{
        userDBManager: manager,
    }
}

//加密用户密码
func GeneratePassword(pwd string) ([]byte, error){
    return bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
}

func ValidatePassword(pwd, hashed string) (bool, error){
    if bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pwd)) != nil{
        return false, errors.New("密码比对错误")
    }
    return true, nil
}

func (u UserDataService) AddUser(user *model.User) (int64, error) {
    pwdByte, err := GeneratePassword(user.UserPassword)
    if err != nil {
        return user.ID, err
    }
    user.UserPassword = string(pwdByte)
    return u.userDBManager.CreateUser(user)
}

func (u UserDataService) DeleteUser(id int64) error {
    return u.DeleteUser(id)
}

func (u UserDataService) UpdateUser(user *model.User, isPwdChanged bool) error {
    if isPwdChanged{
        pwdByte, err := GeneratePassword(user.UserPassword)
        if err != nil {
            return err
        }
        user.UserPassword = string(pwdByte)
    }

    return u.userDBManager.UpdateUser(user)
}

func (u UserDataService) FindUserByName(name string) (*model.User, error) {
    return u.userDBManager.FindUserByName(name)
}

//传入用户名和密码
func (u UserDataService) CheckPassword(name string, pwd string) (bool, error) {
    user, err := u.userDBManager.FindUserByName(name)
    if err != nil {
        return false, err
    }

    return ValidatePassword(pwd, user.UserPassword)
}

