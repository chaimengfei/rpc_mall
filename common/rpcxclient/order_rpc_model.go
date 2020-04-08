package rpcxclient

import (
	"cmf_mall/order/rpc/protos"
	"context"
	"github.com/yakaa/grpcx"
)

type OrderRpcModel struct {
	cli *grpcx.GrpcxClient
}

func NewOrderRpcModel(cli *grpcx.GrpcxClient)  *OrderRpcModel{
	return &OrderRpcModel{cli:cli}
}

func (m *OrderRpcModel) BookingGoods(goodsId int32, goodsNum int32,price int32, userId int32)  (*order.BookingGoodsResponse, error){
  conn,err:= m.cli.GetConnection()
  if err!=nil{
		return nil,err
	}
   rpcClient:=order.NewOrderRpcClient(conn)
   resp,err := rpcClient.BookingGoods(context.Background(),&order.BookingGoodsRequest{GoodsId:goodsId,Price:price,Num:goodsNum,UserId:userId})
   return resp,err
}

func (m *OrderRpcModel) FindByOrderId(id string) ( *order.FindByOrderIdResponse, error) {
  conn,err:=m.cli.GetConnection()
	if err!=nil{
		return nil,err
	}
  rpcClient:=order.NewOrderRpcClient(conn)
  resp,err:=rpcClient.FindByOrderId(context.Background(),&order.FindByOrderIdRequest{Id:id})
  return resp,err
}
