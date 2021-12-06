package handler

import (
    "context"
    "github.com/J-Y-Zhang/mall/user/domain/model"
    "github.com/J-Y-Zhang/mall/user/domain/service"
    "github.com/J-Y-Zhang/mall/user/proto/user"
    "strconv"
)

type User struct {
    UserDataSvc service.UserDataServiceInterface
}

func (u User) Register(ctx context.Context, req *user.UserRegisterRequest, resp *user.UserRegisterResponse) error {
    newUser := &model.User{
        UserName:     req.UserName,
        UserNickName: req.UserNickName,
        UserPassword: req.UserPassword,
    }

    id, err := u.UserDataSvc.AddUser(newUser)
    if err != nil {
        return err
    }
    resp.Message = "注册用户成功, id: " + strconv.Itoa(int(id))
    return nil
}

func (u User) Login(ctx context.Context, req *user.UserLoginRequest, resp *user.UserLoginResponse) error {
    isOk, err := u.UserDataSvc.CheckPassword(req.UserName, req.UserPassword)
    if err != nil {
        resp.IsSuccess = false
        return err
    }
    resp.IsSuccess = isOk
    return nil
}

func (u User) GetInfo(ctx context.Context, req *user.UserInfoRequest, resp *user.UserInfoResponse) error {
    user, err := u.UserDataSvc.FindUserByName(req.UserName)
    if err != nil {
        return err
    }
    *resp = User2UserInfoResponse(user)
    return nil
}

func User2UserInfoResponse(user *model.User) (resp user.UserInfoResponse){
    resp.UserName = user.UserName
    resp.UserId = user.ID
    resp.UserNickName = user.UserNickName
    return
}


