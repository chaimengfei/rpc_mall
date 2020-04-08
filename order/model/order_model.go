package model

import (
	"cmf_mall/order/conf"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type (
	OrderModel struct {
		engine *gorm.DB
	}

	Order struct {
		Id         string    `gorm:"column:id"  json:"id"` //时间戳+userId+GoodsId
		GoodsId    int       `gorm:"column:goods_id" json:"goodsId"`
		Num        int       `gorm:"column:num"  json:"num"`
		Price      int       `gorm:"column:price"  json:"price"`
		UserId     int       `gorm:"column:user_id"  json:"userId"`
		CreateTime time.Time `gorm:"column:create_time"  json:"createTime"` //mysql 声明为timestamp
	}
)
func (Order) TableName() string {
	return conf.OrderConf.Mysql.Table.Order
}
var db *gorm.DB
var err error
func init() {
	db, err = gorm.Open("mysql", conf.OrderConf.Mysql.DataSource)
	if err != nil {
		fmt.Println(err)
		return
	}
	db.LogMode(true)

	NewOrderModel()
}
func NewOrderModel() *OrderModel {
	return &OrderModel{engine: db}
}

func (m *OrderModel) InsertOrder(orderId string, goodsId, goodsNum,price, userId int) error {
	sqlStr := m.InsertOrderSql(orderId, goodsId, goodsNum, price,userId)
	if err := m.engine.Exec(sqlStr).Error; err != nil {
		fmt.Println("InsertOrder Error %s", err.Error())
		return err
	}
	return nil
}

func (m *OrderModel) InsertOrderSql(orderId string, goodsId , goodsNum ,price, userId int) string {
	sqlStr := fmt.Sprintf("insert into tab_order (id,goods_id,num,price,user_id,create_time) VALUES ('%s',%d,%d,%d,%d,CURRENT_TIME)", orderId, goodsId, goodsNum,price, userId)
	return sqlStr
}

func (m *OrderModel) FindByOrderId(id string) *Order {
	var data = &Order{}
	if err := m.engine.Raw("select * from tab_order where id = ?", id).Scan(&data).Error; err != nil {
		fmt.Println("FindByOrderId Error ", err.Error())
		return nil
	}
	return data
}

func (m *OrderModel) FindOrders() []*Order {
	var datas []*Order
	if err := m.engine.Raw("select * from tab_order").Scan(&datas).Error; err != nil {
		fmt.Println("FindOrders Error ", err.Error())
		return nil
	}
	return datas
}
