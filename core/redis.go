package core

import (
	"CopyQQ/global"
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

func InitRedis() {
	global.RDB = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	// 设置超时时间
	timeoutCtx, cancel := context.WithTimeout(global.Ctx, 100*time.Millisecond)
	defer cancel()
	pong, err := global.RDB.Ping(timeoutCtx).Result()
	if err != nil {
		log.Println("Error:", err)
		return
	}
	log.Println("redis connect success,", pong)
}

// Publish 发布消息到Redis
func Publish(ctx context.Context, channel string, msg string) error {
	err := global.RDB.Publish(ctx, channel, msg).Err()
	return err
}

// Subscribe 订阅Redis
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := global.RDB.PSubscribe(ctx, channel)
	msg, err := sub.ReceiveMessage(ctx)
	return msg.Payload, err
}
