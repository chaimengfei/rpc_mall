package main

import (
	"cmf_mall/order/conf"
	"cmf_mall/order/model"
	"cmf_mall/order/rpc/impl"
	"cmf_mall/order/rpc/protos"
	"fmt"
	"github.com/yakaa/grpcx"
	"google.golang.org/grpc"
)

func main() {
	orderImpl := impl.NewOrderRpcImpl(model.NewOrderModel())
	rpcServer, err := grpcx.MustNewGrpcxServer(conf.OrderConf.RpcServerConfig, func(server *grpc.Server) {
		order.RegisterOrderRpcServer(server, orderImpl)
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	rpcServer.Run()
}
