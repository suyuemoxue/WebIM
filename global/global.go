package global

import (
	"CopyQQ/config"
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 声明一系列全局常量
const (
	PulishKey = "ws"
)

// 声明一系列全局变量
var (
	Config *config.Config // 用于保存配置文件
	DB     *gorm.DB       // 连接mysql数据库
	RDB    *redis.Client
	Logger logger.Interface // 用于打印日志
	Ctx    = context.Background()
)

// Response 用于返回请求的的结构体
type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}
