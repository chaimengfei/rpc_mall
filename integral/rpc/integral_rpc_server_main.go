package main

import (
	"cmf_mall/common/utils"
	"cmf_mall/integral/conf"
	"cmf_mall/integral/model"
	"cmf_mall/integral/rpc/impl"
	"cmf_mall/integral/rpc/protos"
	"fmt"
	"github.com/yakaa/grpcx"
	"google.golang.org/grpc"
)

func main() {
	mqHost := conf.IntegralConf.RabbitMq.DataSource + conf.IntegralConf.RabbitMq.VirtualHost
	mqQueue := conf.IntegralConf.RabbitMq.QueueName
	mq, err := utils.NewRabbitMqServer(mqHost, mqQueue)
	if err != nil {
		fmt.Println(err)
	}
	integralImpl := impl.NewIntegralRpcImpl(model.NewIntegralModel(), mq)
	rpcServer, err := grpcx.MustNewGrpcxServer(conf.IntegralConf.RpcServerConfig, func(server *grpc.Server) {
		integral.RegisterIntegralRpcServer(server, integralImpl)
	})
	integralImpl.ConsumeMessage()
	defer integralImpl.Close()
	rpcServer.Run()
}

