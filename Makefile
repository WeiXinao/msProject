.PHONY: bff user project task modelgen modelgen
bff:
	@goctl api go -api ./api/http/bff.api -dir ./bff --style go_zero

user:
	@goctl rpc protoc ./api/proto/user/v1/user.proto --go_out=. --go-grpc_out=. --zrpc_out=./user --style go_zero

project:
	@goctl rpc protoc ./api/proto/project/v1/project.proto --go_out=. --go-grpc_out=. --zrpc_out=./project --style go_zero

task:
	@goctl rpc protoc ./api/proto/task/v1/task.proto --go_out=. --go-grpc_out=. --zrpc_out=./task --style go_zero

modelgen:
	@./pkg/model_generator/bin/modelgen model --dsn 'root:123456@tcp(127.0.0.1:3307)/ms_project?charset=utf8' --table ms_task_work_time --dst ./task/internal/repo/dao/types.go

msggen:
	@./pkg/model_generator/bin/modelgen msg --dsn 'root:123456@tcp(127.0.0.1:3307)/ms_project?charset=utf8' --table ms_project_member --dst ./api/proto/project/v1/project.proto
	
	