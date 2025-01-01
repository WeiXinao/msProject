package main

import (
	"flag"
	"fmt"
	lx "github.com/WeiXinao/msProject/pkg/logx"
	"github.com/zeromicro/go-zero/core/logx"

	projectv1 "github.com/WeiXinao/msProject/api/proto/gen/project/v1"
	"github.com/WeiXinao/msProject/project/internal/config"
	"github.com/WeiXinao/msProject/project/internal/server"
	"github.com/WeiXinao/msProject/project/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/project.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		projectv1.RegisterProjectServiceServer(grpcServer, server.NewProjectServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	// 设置 log 的 writer
	//从配置中读取日志配置，初始化日志
	writer, err := lx.NewZapWriter(&c.LogConfig)
	logx.Must(err)
	logx.SetWriter(writer)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
