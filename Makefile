.PHONY: bff user
bff:
	@goctl api plugin -p goctl-go-compact go -api .\api\http\bff.api -dir .\bff --style go_zero

user:
	@goctl rpc protoc ./api/proto/user/v1/user.proto --go_out=. --go-grpc_out=. --zrpc_out=./user --style go_zero
