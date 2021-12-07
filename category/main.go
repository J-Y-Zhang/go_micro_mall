package main

import (
    "github.com/J-Y-Zhang/mall/category/common"
    "github.com/J-Y-Zhang/mall/category/domain/dao"
    "github.com/J-Y-Zhang/mall/category/domain/service"
    "github.com/J-Y-Zhang/mall/category/handler"
    "github.com/J-Y-Zhang/mall/category/proto/category"
    "github.com/asim/go-micro/plugins/registry/consul/v4"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
    "go-micro.dev/v4"
    "go-micro.dev/v4/registry"
    "log"
    "strconv"
)

const (
    SERVICENAME = "category"
    VERSION = "latest"
    ADDRESS = "0.0.0.0:10001"
)


func main() {
    //配置中心
    consulConf, err := common.GetConsulConfig("82.156.19.233", 8500, "/micro/config")
    if err != nil {
        log.Fatal(err)
        return
    }

    //注册中心
    consulRegistry := consul.NewRegistry(func(options *registry.Options) {
        options.Addrs = []string{
            "82.156.19.233:8500",
        }
    })

    // Create service
    srv := micro.NewService(
        micro.Name(SERVICENAME),
        micro.Version(VERSION),
        micro.Address(ADDRESS),
        micro.Config(consulConf),
        micro.Registry(consulRegistry),
    )
    srv.Init()

    //获取mysql配置
    mysqlConf := common.GetMysqlConfigFromConsul(consulConf, "mysql")
    log.Println("获取mysql配置成功")

    //创建数据库连接
    mysqlStr := mysqlConf.User + ":" + mysqlConf.Pwd + "@tcp(" + mysqlConf.Host +
        ":" + strconv.Itoa(int(mysqlConf.Port)) + ")/" + mysqlConf.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open("mysql", mysqlStr)
    if err != nil {
        log.Fatal(err)
        return
    }
    log.Println("连接mysql成功")
    defer db.Close()
    db.SingularTable(true)

    manager := &dao.CategoryDBManager{db}

    categoryDataSrv := service.NewCategoryDataService(manager)
    err = category.RegisterCategoryHandler(srv.Server(), &handler.Category{categoryDataSrv})
    if err != nil {
        log.Fatal(err)
        return
    }

    if err := srv.Run(); err != nil{
        log.Fatal(err)
    }
}
