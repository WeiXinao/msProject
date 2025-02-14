package gozerox

import (
	"context"
	"reflect"

	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
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

func HttpLogicWrapper[LogicType any, ReqType any, RespType any](ctx context.Context, logic LogicType, req ReqType, serviceFunc func (methodName string, logic LogicType, req ReqType) (RespType, error)) (RespType, error) {
	methodName := reflect.TypeOf(logic).Method(0).Name
	resp, err := serviceFunc(methodName, logic, req)
	if err != nil {
		actualErr, ok := status.FromError(err)
		if !ok {
			logx.Errorf("[logic %s] %#v", methodName, err)
		} else {
			logx.Errorf("[logic %s] %#v", methodName, actualErr)
		}
		return resp, respx.ErrInternalServer
	}
	return resp, nil
}

func HttpLogicWrapperWithoutReq[LogicType any, RespType any](ctx context.Context, logic LogicType, serviceFunc func (methodName string, logic LogicType) (RespType, error)) (RespType, error) {
	methodName := reflect.TypeOf(logic).Method(0).Name
	resp, err := serviceFunc(methodName, logic)
	if err != nil {
		actualErr, ok := status.FromError(err)
		if !ok {
			logx.Errorf("[logic %s] %#v", methodName, err)
		} else {
			logx.Errorf("[logic %s] %#v", methodName, actualErr)
		}
		return resp, respx.ErrInternalServer
	}
	return resp, nil
}