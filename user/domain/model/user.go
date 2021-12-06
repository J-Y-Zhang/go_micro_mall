package model

type User struct {
    //主键
    ID int64 `gorm:"primary_key;not_null;auto_increment"`

    //用户名
    UserName string `gorm:"unique_index;not_null"`

    //昵称
    UserNickName string

    //密码
    UserPassword string
}
