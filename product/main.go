package main

import (
	"fmt"
	"github.com/J-Y-Zhang/mall/product/common"
	"github.com/J-Y-Zhang/mall/product/domain/dao"
	service2 "github.com/J-Y-Zhang/mall/product/domain/service"
	"github.com/J-Y-Zhang/mall/product/handler"
	"github.com/J-Y-Zhang/mall/product/proto/product"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	opentracing2 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"log"
	"strconv"
)

const (
	SERVICENAME = "product"
	VERSION     = "latest"
	ADDRESS     = "0.0.0.0:10002"
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

	//链路追踪
	tracer, closer, err := common.NewTracer(SERVICENAME, "82.156.19.233:6831")
	if err != nil {
		log.Println(err)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

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

	// New Service
	service := micro.NewService(
		micro.Name(SERVICENAME),
		micro.Version(VERSION),
		micro.Address(ADDRESS),
		micro.Config(consulConf),
		micro.Registry(consulRegistry),
		//绑定链路追踪(Handler)
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	// Initialise service
	service.Init()

	manager := &dao.ProductDBManager{
		MysqlDB: db,
	}

	//建表
	manager.InitTable()
	fmt.Println("建表成功")

	productDataSrv := service2.NewProductDataService(manager)

	err = product.RegisterProductHandler(service.Server(), handler.Product{productDataSrv})
	if err != nil {
		log.Println(err)
		return
	}
	if err := service.Run(); err != nil {
		log.Println(err)
		return
	}
}
