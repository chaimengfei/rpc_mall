package impl

import (
	"cmf_mall/common/utils"
	"cmf_mall/order/model"
	"cmf_mall/order/rpc/protos"
	"context"
	"errors"
	"time"
)

type OrderRpcImpl struct {
	orderModel *model.OrderModel
	mqServer *utils.RabbitMqServer
}

func NewOrderRpcImpl(m *model.OrderModel) *OrderRpcImpl  {
	return &OrderRpcImpl{orderModel:m}
}

func (impl *OrderRpcImpl) BookingGoods(ctx context.Context, req *order.BookingGoodsRequest) (resp *order.BookingGoodsResponse, err error) {
	now := time.Now()
	orderId := string(now.Unix()) + "--" + string(req.UserId)
	//--------------------rabbitmq publish sql-----------
	sql := impl.orderModel.InsertOrderSql(orderId, int(req.GoodsId), int(req.Num),int(req.Price), int(req.UserId))
	err = impl.mqServer.PushMessage(sql)
	return &order.BookingGoodsResponse{OrderId:orderId,CreateTime:now.Unix()}, err
}

func (impl *OrderRpcImpl) FindByOrderId(ctx context.Context, req *order.FindByOrderIdRequest) (resp *order.FindByOrderIdResponse, err error) {
	o:=impl.orderModel.FindByOrderId(req.Id)
	if o==nil{
		return nil,errors.New("Order Not Found")
	}
	return &order.FindByOrderIdResponse{Id:o.Id,GoodsId:int32(o.GoodsId),Num:int32(o.Num),UserId:int32(o.UserId),CreateTime:o.CreateTime.Unix()},nil
}
