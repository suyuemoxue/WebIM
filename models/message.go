package models

import (
	"CopyQQ/global"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	SendID    string `json:"sendID" gorm:"size:10"`          // 发送者
	ReceiveID string `json:"receiveID" gorm:"size:10"`       // 接受者
	Content   []byte `json:"content" gorm:"charset=utf8mb4"` // 消息
	MsgType   string `json:"msgType" gorm:"size:10"`         // 消息类型 群聊 私聊
	MediaType string `json:"mediaType" gorm:"size:10"`       // 消息类型 文字 图片 音视频
}

type MessageList []Message

var message Message

func (msg *Message) TableName() string {
	return "message"
}

// SaveMessage 保存消息
func (msg *Message) SaveMessage() bool {
	err := global.DB.Create(&msg).Error
	if err != nil {
		return false
	}
	return true
}
