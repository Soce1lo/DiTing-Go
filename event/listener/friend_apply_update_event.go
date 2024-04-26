package listener

import (
	"DiTing-Go/dal/model"
	"DiTing-Go/domain/enum"
	"DiTing-Go/global"
	common_enum "DiTing-Go/pkg/enum"
	"DiTing-Go/utils/jsonUtils"
	"context"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/spf13/viper"
)

func init() {
	host := viper.GetString("rocketmq.host")
	// 设置推送消费者
	rocketConsumer, _ := rocketmq.NewPushConsumer(
		//消费组
		consumer.WithGroupName(enum.UpdataFriendApplyTopic),
		// namesrv地址
		consumer.WithNameServer([]string{host}),
	)
	err := rocketConsumer.Subscribe(enum.UpdataFriendApplyTopic, consumer.MessageSelector{}, UpdateFriendApplyEvent)
	if err != nil {
		global.Logger.Panicf("subscribe error: %s", err.Error())
	}
	err = rocketConsumer.Start()
	if err != nil {
		global.Logger.Panicf("start consumer error: %s", err.Error())
	}
}

// UpdateFriendApplyEvent 更新已读好友申请事件
func UpdateFriendApplyEvent(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	for i := range ext {
		// 解码
		msg := model.UserApply{}
		if err := jsonUtils.UnmarshalMsg(&msg, ext[i]); err != nil {
			global.Logger.Errorf("jsonUtils unmarshal error: %s", err.Error())
			return consumer.ConsumeRetryLater, nil
		}

		// 执行更新操作
		userApply := global.Query.UserApply
		userApplyQ := userApply.WithContext(ctx)
		_, err := userApplyQ.Where(userApply.TargetID.Eq(msg.UID), userApply.ReadStatus.Eq(common_enum.NO)).Update(userApply.ReadStatus, msg.ReadStatus)
		if err != nil {
			global.Logger.Errorf("更新好友申请表失败 %s", err)
			return consumer.ConsumeRetryLater, nil
		}
	}
	return consumer.ConsumeSuccess, nil
}
