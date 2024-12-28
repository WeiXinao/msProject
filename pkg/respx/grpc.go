package respx

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ToStatusErr(err *Error) error {
	return status.Error(codes.Code(err.Code), err.Msg)
}

func FromStatusErr(err error) *Error {
	e, _ := status.FromError(err)
	return NewError(int(e.Code()), e.Message())
}
