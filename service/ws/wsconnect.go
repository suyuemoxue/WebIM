package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

type WebsocketService struct {
}

var userLock sync.RWMutex

func (ws *WebsocketService) HandleWebSocket(c *gin.Context) {
	//获取当前用户信息
	query := c.Request.URL.Query()
	userID := query.Get("uid")
	toUserID := query.Get("touid")
	toGroupID := query.Get("tgid")

	// 处理客户端发起的 WebSocket 连接升级请求的逻辑，如果升级过程顺利完成，那么后续就可以通过连接对象（conn）进行 WebSocket 消息的收发操作
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("ws connect fail")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("用户:", userID, "ws connect success")
	client := &Client{
		SendID:    userID,
		ReceiveID: toUserID,
		GroupID:   toGroupID,
		Socket:    conn,
		SendMsg:   make(chan []byte),
	}
	Manager.Register <- client
	fmt.Println("发送方:", client.SendID, "接收方:", client.ReceiveID, "群号:", client.GroupID)
	go client.Read()
	go client.Write()
}
