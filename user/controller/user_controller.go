package controller

import (
	"cmf_mall/user/service"
	"github.com/gin-gonic/gin"
)

type (
	UserController struct {
		userService *service.UserService
	}
)

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService:userService}
}
func ( u *UserController) Register(c *gin.Context) {
	var req service.RegisterReq
	var resp *service.Resp
	if err:=c.BindJSON(&req);err==nil{
		resp=u.userService.Register(&req)
	}
	c.JSON(resp.Code,resp.Msg)

}
func ( u *UserController) Login(c *gin.Context) {
	var req service.LoginReq
	var resp *service.Resp
	if err:=c.BindJSON(&req);err==nil{
		resp=u.userService.Login(&req)
	}
	c.JSON(resp.Code,resp.Msg)
}
