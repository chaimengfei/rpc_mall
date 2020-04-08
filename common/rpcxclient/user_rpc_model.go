package rpcxclient

import (
	"cmf_mall/user/rpc/protos"
	"context"
	"github.com/yakaa/grpcx"
)

type UserRpcModel struct {
   cli *grpcx.GrpcxClient
}

func NewUserRpcModel(cli *grpcx.GrpcxClient) *UserRpcModel {
	return &UserRpcModel{cli:cli}
}

func (m *UserRpcModel)  FindUserByMobile (mobile string) (*user.UserResponse,error){
	conn, err := m.cli.GetConnection()
	if err != nil {
		return nil, err
	}
	rpcClient:=user.NewUserRpcClient(conn)
	resp,err:=rpcClient.FindUserByMobile(context.Background(),&user.FindUserByMobileRequest{Mobile:mobile})
	if err!=nil{
		return nil,err
	}
	return resp,nil
}
func (m *UserRpcModel) FindUserById(id int32) (*user.UserResponse, error) {
	conn, err := m.cli.GetConnection()
	if err != nil {
		return nil, err
	}
	rpcClient := user.NewUserRpcClient(conn)
	resp, err := rpcClient.FindUserById(context.Background(), &user.FindUserByIdRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return resp, nil
}


