package service

import (
	"cmf_mall/common/rpcxclient"
	"cmf_mall/order/model"
	"fmt"
	"github.com/google/uuid"
)

type OrderService struct {
  orderModel *model.OrderModel
	integralRpcModel *rpcxclient.IntegralRpcModel
	userRpcModel *rpcxclient.UserRpcModel
}
type OrderBuyRequest struct {
	UserId    int `json:"userId"`
	GoodsId   int `json:"goodsId"  binding:"required"`
	Num       int `json:"num" binding:"required"`
	Price     int `json:"price" binding:"required"`
}

func NewOrderService(orderModel *model.OrderModel, integralRpcModel *rpcxclient.IntegralRpcModel) *OrderService {
	return &OrderService{orderModel: orderModel, integralRpcModel: integralRpcModel}
}
func (s *OrderService) BookingGoods(buy *OrderBuyRequest) error {
	orderId := uuid.New().String()
	consumePrice := int(buy.Price)
	err := s.orderModel.InsertOrder(orderId, buy.GoodsId, buy.Num, buy.Price, buy.UserId)
	integralResp, err := s.integralRpcModel.ConsumeIntegral(int(buy.UserId), consumePrice)
	if err != nil {
		return err
	} else {
		fmt.Println("order_service BookingGoods integralResp -->>", integralResp.UserId, integralResp.Integral)
		return nil
	}
}
func (s *OrderService) FindByOrderId(id string) *model.Order {
	order:=s.orderModel.FindByOrderId(id)
	return order

}
func (s *OrderService) FindOrders() []*model.Order {
	orders:=s.orderModel.FindOrders()
	return orders
}
