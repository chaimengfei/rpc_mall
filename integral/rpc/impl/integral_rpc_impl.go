package impl

import (
	"cmf_mall/common/utils"
	"cmf_mall/integral/model"
	"cmf_mall/integral/rpc/protos"
	"context"
	"fmt"
)

type IntegralRpcImpl struct {
	integralModel  *model.IntegralModel
	rabbitMqServer *utils.RabbitMqServer
}

func NewIntegralRpcImpl(m *model.IntegralModel, mq *utils.RabbitMqServer) *IntegralRpcImpl {
	return &IntegralRpcImpl{integralModel: m, rabbitMqServer: mq}
}

func (impl *IntegralRpcImpl) ConsumeMessage(){
	impl.rabbitMqServer.ConsumeMessage(func(msg string) error {
		fmt.Println("ConsumeMessage() msg -->>",msg)
		return impl.integralModel.ExecIntegralSql(msg)
	})
}
func (impl *IntegralRpcImpl) Close(){
	impl.rabbitMqServer.CloseRabbitmqConn()
}

func (impl *IntegralRpcImpl) AddIntegral(ctx context.Context, req *integral.IntegralRequest) (resp *integral.IntegralResponse, err error) {
	//now, err := impl.integralModel.UpdateIntegralByUserId(req.UserId, old.Integral+req.Integral)
	sql := impl.integralModel.IntegralSql(int(req.UserId),int(req.Integral),true)
	err=impl.rabbitMqServer.PushMessage(sql)
	return &integral.IntegralResponse{UserId:req.UserId,Integral:req.Integral}, err
}
func (impl *IntegralRpcImpl) ConsumeIntegral(ctx context.Context, req *integral.IntegralRequest) (resp *integral.IntegralResponse, err error) {
	old, err := impl.integralModel.FindByUserId(int(req.UserId))
	if err != nil {
		return nil, err
	}
	ups := old.Integral - int(req.Integral)
	if ups < 0 {
		ups = 0
	}
	sql := impl.integralModel.IntegralSql(int(req.UserId), ups,false)
	err = impl.rabbitMqServer.PushMessage(sql)
	return  &integral.IntegralResponse{UserId:req.UserId,Integral:int64(ups)}, err
}
func (impl *IntegralRpcImpl) FindIntegralByUserid(ctx context.Context, req *integral.FindIntegralRequest) (resp *integral.IntegralResponse, err error) {
	old, err := impl.integralModel.FindByUserId(int(req.UserId))
	if err != nil {
		return nil, err
	}
	return  &integral.IntegralResponse{UserId:req.UserId,Integral:int64(old.Integral)}, nil
}
