package impl

import (
	"cmf_mall/user/model"
	"cmf_mall/user/rpc/protos"
	"context"
)

type UserRpcImpl struct {
	userModel *model.UserModel
}

func NewUserRpcImpl(m *model.UserModel) *UserRpcImpl {
	return &UserRpcImpl{userModel:m}
}

func (impl *UserRpcImpl) FindUserByMobile(ctx context.Context, req *user.FindUserByMobileRequest) (resp *user.UserResponse, err error){
	u:=impl.userModel.FindByMobile(req.Mobile)
	return &user.UserResponse{Id:int32(u.Id),Name:u.Name},nil
}
func (impl *UserRpcImpl) FindUserById(ctx context.Context, req *user.FindUserByIdRequest) (resp *user.UserResponse, err error){
	u:=impl.userModel.FindById(req.Id)
	return &user.UserResponse{Id:int32(u.Id),Name:u.Name},nil
}
