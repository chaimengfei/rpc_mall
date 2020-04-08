package main

import (
	"cmf_mall/user/conf"
	"cmf_mall/user/model"
	"cmf_mall/user/rpc/impl"
	"cmf_mall/user/rpc/protos"
	"fmt"
	"github.com/yakaa/grpcx"
	"google.golang.org/grpc"
)

func main()  {
	rpcImpl := impl.NewUserRpcImpl(model.NewUserModel())
	rpcServer, err := grpcx.MustNewGrpcxServer(conf.UserConf.RpcServerConfig, func(server *grpc.Server) {
		user.RegisterUserRpcServer(server, rpcImpl)
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	rpcServer.Run()
}
