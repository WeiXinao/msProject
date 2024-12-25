package respx

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %v, msg: %v", e.Code, e.Msg)
}

func (e *Error) AsStatusErr() error {
	return status.Error(codes.Code(e.Code), e.Msg)
}

func NewError(code int, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}
