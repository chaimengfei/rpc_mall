package test

import (
	"fmt"
	"testing"
)

func TestCode( t *testing.T)  {
	//encPwd:=fmt.Sprintf("%x", md5.Sum([]byte("123456")))
	//fmt.Print(encPwd)
	//b4bdc1bc7155e855f808e40aadd85474
	sqlStr := fmt.Sprintf("insert into tab_order (id,goods_id,num,price,user_id) VALUES ('%s',%d,%d,%d,%d)", "order1243q432", 1, 1,100, 6)
	fmt.Println(sqlStr)
}
