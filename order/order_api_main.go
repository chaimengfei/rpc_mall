package main

import (
	"cmf_mall/common/middleware"
	"cmf_mall/common/rpcxclient"
	"cmf_mall/order/conf"
	"cmf_mall/order/controller"
	"cmf_mall/order/model"
	"cmf_mall/order/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/yakaa/grpcx"
)

func main() {
	var orderConf = conf.OrderConf
	rpcClient, err := grpcx.MustNewGrpcxClient(orderConf.IntegralRpc)
	if err != nil {
		fmt.Println("user_api_main.go rpcClient err ", err.Error())
		return
	}
	integralRpcModel := rpcxclient.NewIntegralRpcModel(rpcClient)
	orderService := service.NewOrderService(model.NewOrderModel(), integralRpcModel)
	orderController := controller.NewOrderController(orderService)
	redisClient:=redis.NewClient(&redis.Options{Addr:orderConf.Redis.DataSource,Password:orderConf.Redis.Auth})
    auth:=middleware.NewAuthorization(redisClient)
	r := gin.Default()
	orderR := r.Group("/order")
	orderR.Use(auth.Auth)
	{
		orderR.GET("/list", orderController.OrderList)
		orderR.POST("/buy", orderController.OrderBuy)
	}
	r.Run(orderConf.Port)

}
