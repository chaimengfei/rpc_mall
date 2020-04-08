package controller

import (
	"cmf_mall/order/service"
	"github.com/gin-gonic/gin"
)

type (
	OrderController struct {
		orderService *service.OrderService
	}
)

func NewOrderController(orderService *service.OrderService) *OrderController {
	return &OrderController{orderService:orderService}
}
func (o *OrderController) OrderList(c *gin.Context) {
	orders:=o.orderService.FindOrders()
	c.JSON(200,orders)

}
func (o *OrderController) OrderBuy(c *gin.Context) {
	var orderBuy service.OrderBuyRequest
	err:=c.BindJSON(&orderBuy)
	if err!=nil{
		c.JSON(400,"OrderBuy Parse Error")
		return
	}
	uId,b:=c.Get("userId")
	if b{
		orderBuy.UserId=uId.(int)
	}
	err = o.orderService.BookingGoods(&orderBuy)
	if err!=nil{
		c.JSON(500,"BookingGoods Handler Error")
		return
	}
	c.JSON(200,"BookingGoods Success")
}

