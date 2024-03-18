package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	FromID    string // 发送者
	TargetID  string // 接受者
	msgType   string // 消息类型 群聊 私聊 广播
	MediaType string // 消息类型 文字 图片 音视频
}

func (msg *Message) TableName() string {
	return "message"
}
