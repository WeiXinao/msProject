package pprofx

import (
	"net/http"
	"net/http/pprof"

	"github.com/zeromicro/go-zero/rest"
)

const (
	DefaultPrefix = "/debug"
)

func getPrefix(prefixOptions ...string) string {
	prefix := DefaultPrefix
	if len(prefixOptions) > 0 {
		prefix = prefixOptions[0]
	}
	return prefix
}

func Register(server *rest.Server, prefixOptions ...string) {
	routeRegister(server, prefixOptions...)
}

func routeRegister(server *rest.Server, prefixOptions ...string) {
	server.AddRoutes([]rest.Route{
		{
			Method: http.MethodGet,
			Path: "/pprof",
			Handler: pprof.Index,
		},
		{
			Method: http.MethodGet,
			Path: "/cmdline",
			Handler: pprof.Cmdline,
		},
		{
			Method: http.MethodGet,
			Path: "/profile",
			Handler: pprof.Profile,
		},
		{
			Method: http.MethodPost,
			Path: "/symbol",
			Handler: pprof.Symbol,
		},
		{
			Method: http.MethodGet,
			Path: "/symbol",
			Handler: pprof.Symbol,
		},
		{
			Method: http.MethodGet,
			Path: "/trace",
			Handler: pprof.Trace,
		},
		{
			Method: http.MethodGet,
			Path: "/allocs",
			Handler: pprof.Handler("allocs").ServeHTTP,
		},
		{
			Method: http.MethodGet,
			Path: "/block",
			Handler: pprof.Handler("block").ServeHTTP,
		},
		{
			Method: http.MethodGet,
			Path: "/goroutine",
			Handler: pprof.Handler("goroutine").ServeHTTP,
		},
		{
			Method: http.MethodGet,
			Path: "/heap",
			Handler: pprof.Handler("heap").ServeHTTP,
		},
		{
			Method: http.MethodGet,
			Path: "/mutex",
			Handler: pprof.Handler("mutex").ServeHTTP,
		},
		{
			Method: http.MethodGet,
			Path: "/threadcreate",
			Handler: pprof.Handler("threadcreate").ServeHTTP,
		},
	}, 
	rest.WithPrefix(getPrefix(prefixOptions...)))
}