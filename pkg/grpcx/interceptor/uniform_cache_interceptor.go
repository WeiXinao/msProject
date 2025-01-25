package interceptor

import (
	"context"
	"fmt"
	"time"

	"github.com/WeiXinao/msProject/pkg/cachex"
	"github.com/WeiXinao/msProject/pkg/encrypts"
	"github.com/bytedance/sonic"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"google.golang.org/grpc"
)

type UniformCacheInterceporBuilder struct {
	cache cachex.Cache
	patternToRsp map[string]any
	expire time.Duration
}

func NewUniformCacheInterceptorBuilder(cache cachex.Cache) *UniformCacheInterceporBuilder {
	return &UniformCacheInterceporBuilder{
		cache: cache,
		expire: 5 * time.Minute,
		patternToRsp: make(map[string]any),
	}	
}

func (u *UniformCacheInterceporBuilder) Expire(exp time.Duration)  {
	u.expire = exp	
}

func (u *UniformCacheInterceporBuilder) AddPatternRespMap(pattern string, resp any) *UniformCacheInterceporBuilder {
	u.patternToRsp[pattern] = resp
	return u	
}

func (u *UniformCacheInterceporBuilder) cacheKey(req any, fullMethod string) (string, error) {
		respBytes, err := sonic.Marshal(req)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s:%s", fullMethod, encrypts.Md5(string(respBytes))), nil
}

func (u *UniformCacheInterceporBuilder) Build() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		logx.Debug("go by UniformCacheInterceporBuilder")
		typeResp, ok := u.patternToRsp[info.FullMethod]
		if !ok {
			logx.Debug("resp type not exists")
			return handler(ctx, req)
		}
		key, _ := u.cacheKey(req, info.FullMethod)
		respJson, err := u.cache.Get(ctx, key)
		if err == nil && respJson != "" {
			err = sonic.Unmarshal([]byte(respJson), &typeResp)
			logx.Infof("[interceptor UniformCacheInterceporBuilder] %#v", typeResp)
			if err == nil {
				return typeResp, nil
			}
		}
		if err != nil {
				logx.Error("[intercepter UniformCacheInterceporBuilder]", err)		
		}

		resp, err = handler(ctx, req)
		if err == nil {
			threading.GoSafe(func ()  {
				logx.Info("[intercepter UniformCacheInterceporBuilder]", resp)		
				ctx,  cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				er := u.cache.Put(ctx, key, resp, u.expire)
				if er != nil {
					logx.Error("[intercepter UniformCacheInterceporBuilder]", err)		
				}
			})
		}

		return resp, err
	}	
}
