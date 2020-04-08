package model

import (
	"cmf_mall/integral/conf"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type IntegralModel struct {
	engine *gorm.DB
}

type Integral struct {
	Id         int       `gorm:"column:id"  json:"id"` //自增id
	UserId     int       `gorm:"column:user_id"  json:"userId"`
	Integral   int       `gorm:"column:integral" json:"integral"`
	UpdateTime time.Time `gorm:"column:update_time"  json:"updateTime"`
}

func (Integral) TableName() string {
	return conf.IntegralConf.Mysql.Table.Integral
}

var db *gorm.DB
var err error
func init() {
	db, err = gorm.Open("mysql", conf.IntegralConf.Mysql.DataSource)
	if err != nil {
		fmt.Println(err)
		return
	}
	db.LogMode(true)

	NewIntegralModel()
}

func NewIntegralModel() *IntegralModel {
	return &IntegralModel{engine: db}
}
func (m *IntegralModel) FindByUserId(userId int) (*Integral, error) {
	res := new(Integral)
	if err := m.engine.Raw("select * from tab_integral where user_id = ?", userId).Scan(res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

/*func (m *IntegralModel) UpdateIntegralByUserId(userId , integral int) (*Integral, error) {
	str:=m.IntegralSql(userId,integral,false)
	if err := m.engine.Exec(str).Error; err != nil {
		return nil, err
	}
	return m.FindByUserId(userId)
}*/
func (m *IntegralModel) ExecIntegralSql(sql string) error {
	if err := m.engine.Exec(sql).Error; err != nil {
		return err
	}
	return nil
}

func (m *IntegralModel) IntegralSql(userId, integral int, add bool) string {
	s := fmt.Sprintf("update tab_integral set `integral`= %d,update_time = CURRENT_TIME where user_id = %d ", integral, userId)
	if add {
		s = fmt.Sprintf("insert into tab_integral (user_id,`integral`,update_time) values (%d,%d,CURRENT_TIME)", userId, integral)
	}
	return s
}
