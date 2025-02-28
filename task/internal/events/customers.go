package events

import (
	"context"

	"github.com/WeiXinao/msProject/task/internal/config"
	"github.com/WeiXinao/msProject/task/internal/events/task"
	"github.com/WeiXinao/msProject/task/internal/svc"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

func Consumers(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {
    return []service.Service{
        //Listening for changes in consumption flow status
        kq.MustNewQueue(c.KqConsumerConf, task.NewSaveTaskStagesEventCustomer(ctx, svcContext)),
        //.....
    }
}