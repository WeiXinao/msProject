package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"

	lx "github.com/WeiXinao/msProject/pkg/logx"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/WeiXinao/msProject/bff/internal/config"
	"github.com/WeiXinao/msProject/bff/internal/handler"
	"github.com/WeiXinao/msProject/bff/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/bff.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	
	httpx.SetErrorHandler(func(err error) (int, any) {
		var e *respx.Error
		switch {
		case errors.As(err, &e):
			return http.StatusOK, respx.Fail(e)
		default:
			return http.StatusInternalServerError, nil
		}
	})

	// 设置 log 的 writer
	//从配置中读取日志配置，初始化日志
	writer, err := lx.NewZapWriter(&c.LogConfig)
	logx.Must(err)
	logx.SetWriter(writer)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
