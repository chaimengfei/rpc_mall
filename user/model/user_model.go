package model

import (
	"cmf_mall/user/conf"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserModel struct {
	engine *gorm.DB
}
type User struct {
	Id       int  `gorm:"column:id" json:"id"`
	Mobile   string `gorm:"column:mobile" json:"mobile"`
	Name     string `gorm:"column:name" json:"name"`
	Password string `gorm:"column:password" json:"password"`
}

func (User) TableName() string {
	return conf.UserConf.Mysql.Table.User
}
var db *gorm.DB // user的db实体对象,封装在userModel中
var err error
func init() {
	db, err = gorm.Open("mysql", conf.UserConf.Mysql.DataSource)
	if err != nil {
		fmt.Println(err)
		return
	}
	db.LogMode(true)

	NewUserModel()
}

func NewUserModel() *UserModel {
	return &UserModel{engine: db}
}

func (m *UserModel) ExistByMobile(mobile string) bool {
	user:=m.FindByMobile(mobile)
	if user !=nil{
		return true
	}
	return false
}

func (m *UserModel) FindByMobile(mobile string) *User {
	var user =&User{}
	if err:=m.engine.Raw("select * from tab_user where mobile = ?",mobile).Scan(user).Error;err!=nil{
		//record not found
		return nil
	}
	return user
}
func (m *UserModel) FindById(id int32) *User {
	var user =&User{}
	if err:=m.engine.Raw("select * from tab_user where id = ?",id).Scan(user).Error;err!=nil{
		//record not found
		return nil
	}
	return user
}

func (m *UserModel) InsertOne(mobile,password string) (int,error) {
	if err:=m.engine.Exec("insert into tab_user (mobile,password) values (?,?)",mobile,password).Error;err!=nil{
		return 0,err
	}
	var u User
	if err:=m.engine.Raw("select max(id) id from tab_user").Scan(&u).Error;err!=nil{
		return 0,err
	}
	return u.Id,nil
}
