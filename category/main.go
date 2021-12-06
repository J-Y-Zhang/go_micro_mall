package main

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
    "go-micro.dev/v4"
    "log"
)

const (
    SERVICENAME = "category"
    VERSION = "latest"
)


func main() {
    // Create service
    srv := micro.NewService(
        micro.Name(SERVICENAME),
        micro.Version(VERSION),
    )
    srv.Init()

    //创建数据库连接
    db, err := gorm.Open("mysql", "root:zjy20020508@tcp(82.156.19.233:3306)/mall_category?charset=utf8mb4&parseTime=True&loc=Local")
    if err != nil {
        log.Fatal(err)
        return
    }
    log.Println("连接mysql成功")
    defer db.Close()
    db.SingularTable(true)

}
