package task

import (
	"context"

	"github.com/WeiXinao/msProject/project/internal/domain"
	"github.com/bytedance/sonic"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
)

type SaveTaskStagesEvent struct {
	TaskStagesList []*domain.TaskStages
}

type Producer interface {
	ProduceSaveTaskStagesEvent(ctx context.Context, evt SaveTaskStagesEvent) error
}

func (kq *KqSyncProducer) ProduceSaveTaskStagesEvent(ctx context.Context, evt SaveTaskStagesEvent) error {
	data, err := sonic.Marshal(&evt)
	if err != nil {
		return err
	}
	if err = kq.Producer.Push(ctx, string(data)); err != nil {
		logx.WithContext(ctx).Errorf("[event ProduceSaveTaskStagesEvent] %#v", err)
	}
	return err
}

type KqSyncProducer struct {
	Producer *kq.Pusher
}

func NewKqSyncProducer(producer *kq.Pusher) Producer {
	return &KqSyncProducer{
		Producer: producer,
	} 
}

