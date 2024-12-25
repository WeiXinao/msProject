.PHONY: bff user
bff:
	@goctl api go --api ..\..\api\http\user.api --dir .

user:
	@goctl rpc protoc ./api/proto/user/v1/user.proto --go_out=. --go-grpc_out=. --zrpc_out=./user
