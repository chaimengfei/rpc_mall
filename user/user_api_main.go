package main

import (
	"cmf_mall/common/rpcxclient"
	"cmf_mall/user/conf"
	"cmf_mall/user/controller"
	"cmf_mall/user/model"
	"cmf_mall/user/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/yakaa/grpcx"
)

func main() {
	var conf = conf.UserConf
	rpcClient,err:=grpcx.MustNewGrpcxClient(conf.IntegralRpc)
	if err != nil {
	fmt.Println("user_api_main.go rpcClient err ",err.Error())
		return
	}
	integralRpcModel := rpcxclient.NewIntegralRpcModel(rpcClient)
	redisClient:=redis.NewClient(&redis.Options{Addr:conf.Redis.DataSource,Password:conf.Redis.Auth})
	userService := service.NewUserService(model.NewUserModel(), integralRpcModel,redisClient)
	userController := controller.NewUserController(userService)
	r := gin.Default()
	userR := r.Group("/user")
	userR.Use()
	{
		userR.POST("/register", userController.Register)
		userR.POST("/login", userController.Login)
	}
	r.Run(conf.Port)

}
