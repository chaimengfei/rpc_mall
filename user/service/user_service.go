package service

import (
	"cmf_mall/common/middleware"
	"cmf_mall/common/rpcxclient"
	"cmf_mall/user/model"
	"crypto/md5"
	"fmt"
	"strconv"
	"github.com/go-redis/redis"
)

// -----------------------------------
type UserService struct {
	userModel        *model.UserModel
	integralRpcModel *rpcxclient.IntegralRpcModel
	redisCache *redis.Client

}
type (
	RegisterReq struct {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
	}
	LoginReq struct {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
	}
	Resp struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
)

func NewUserService(userModel *model.UserModel,integralRpcModel *rpcxclient.IntegralRpcModel,redisCache *redis.Client) *UserService {
	return &UserService{userModel: userModel,integralRpcModel:integralRpcModel,redisCache:redisCache}
}

func (s *UserService) Register(req *RegisterReq) *Resp {
	var resp = &Resp{}
	exist := s.userModel.ExistByMobile(req.Mobile)
	if !exist {
		encPwd:=fmt.Sprintf("%x", md5.Sum([]byte(req.Password)))
		uId,err := s.userModel.InsertOne(req.Mobile, encPwd)
		if err != nil {
			resp.Code = 500
			resp.Msg = "Register Error " + err.Error()
			return resp
		}
		//------IntergralRpcModel.AddIntergral()------------------
		integralResp,err:=s.integralRpcModel.AddIntegral(uId,1000)
		if err != nil {
			resp.Code = 500
			resp.Msg = "Register Error " + err.Error()
			return resp
		}else{
			fmt.Println("Register integralResp -->>",integralResp.UserId,integralResp.Integral)
		}
	}
	return resp
}

func (s *UserService) Login(req *LoginReq) *Resp{
	var resp = &Resp{}
	user := s.userModel.FindByMobile(req.Mobile)
	if user == nil {
		resp.Code = 500
		resp.Msg = "User Not Found "
		return resp
	}
	if user.Password != fmt.Sprintf("%x", md5.Sum([]byte(req.Password))) {
		resp.Code = 500
		resp.Msg = "Password Error"
		return resp
	}
	//------RedisCache.set(userId)------------------
	authToken:=fmt.Sprintf("%x", md5.Sum([]byte(user.Mobile+strconv.Itoa(int(user.Id)))))
	s.redisCache.Set(authToken,user.Id,middleware.AuthorizationExpire)
	resp.Code=200
	resp.Msg = authToken
	return resp
}
