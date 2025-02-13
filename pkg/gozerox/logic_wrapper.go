package gozerox

import (
	"context"
	"reflect"

	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/zeromicro/go-zero/core/logx"
)

func RpcLogicWrapper[LogicType any, ReqType any, RespType any](ctx context.Context, logic LogicType, req ReqType, serviceFunc func (methodName string, logic LogicType, req ReqType) (RespType, error)) (RespType, error) {
	methodName := reflect.TypeOf(logic).Method(0).Name
	resp, err := serviceFunc(methodName, logic, req)
	if err != nil {
		logx.Errorf("[logic %s] %#v", methodName, err)
		return resp, respx.ToStatusErr(respx.ErrInternalServer)
	}
	return resp, nil
}