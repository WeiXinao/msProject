package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/WeiXinao/msProject/api/proto/gen/task/v1"
	lx "github.com/WeiXinao/msProject/pkg/logx"
	"github.com/WeiXinao/msProject/task/internal/config"
	"github.com/WeiXinao/msProject/task/internal/events"
	"github.com/WeiXinao/msProject/task/internal/server"
	"github.com/WeiXinao/msProject/task/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/task.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		v1.RegisterTaskServiceServer(grpcServer, server.NewTaskServiceServer(svcCtx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	// defer s.Stop()

	ctx := context.Background()
	sg := service.NewServiceGroup()
	defer sg.Stop()
	for _, mq := range events.Consumers(c, ctx, svcCtx) {
		sg.Add(mq)	
	}

	sg.Add(s)

	// 设置 log 的 writer
	//从配置中读取日志配置，初始化日志
	writer, err := lx.NewZapWriter(&c.LogConfig)
	logx.Must(err)
	logx.SetWriter(writer)


	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	// s.Start()
	sg.Start()
}
