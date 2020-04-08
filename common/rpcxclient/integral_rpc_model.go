package rpcxclient

import (
	"cmf_mall/integral/rpc/protos"
	"context"
	"github.com/yakaa/grpcx"
)

type IntegralRpcModel struct {
	cli *grpcx.GrpcxClient
}

func NewIntegralRpcModel(cli *grpcx.GrpcxClient) *IntegralRpcModel {
	return &IntegralRpcModel{cli: cli}
}

func (m *IntegralRpcModel) AddIntegral(userId, integralNum int) (*integral.IntegralResponse, error) {
	conn, err := m.cli.GetConnection()
	if err != nil {
		return nil, err
	}
	rpcClient := integral.NewIntegralRpcClient(conn)
	resp, err := rpcClient.AddIntegral(context.Background(), &integral.IntegralRequest{UserId: int32(userId), Integral: int64(integralNum)})
	return resp, err
}

func (m *IntegralRpcModel) ConsumeIntegral(userId, conintegral int) (*integral.IntegralResponse, error) {
	conn, err := m.cli.GetConnection()
	if err != nil {
		return nil, err
	}
	rpcClient := integral.NewIntegralRpcClient(conn)
	resp, err := rpcClient.ConsumeIntegral(context.Background(), &integral.IntegralRequest{UserId: int32(userId), Integral: int64(conintegral)})
	return resp, err
}

func (m *IntegralRpcModel) FindIntegralByUserid(userId int) (*integral.IntegralResponse, error) {
	conn, err := m.cli.GetConnection()
	if err != nil {
		return nil, err
	}
	rpcClient := integral.NewIntegralRpcClient(conn)
	resp, err := rpcClient.FindIntegralByUserid(context.Background(), &integral.FindIntegralRequest{UserId: int32(userId)})
	return resp, err
}

