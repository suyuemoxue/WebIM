package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

type WebsocketService struct {
}

var userlock sync.RWMutex

// 创建一个 WebSocket 升级器，并设置了一个简单的请求来源检查函数，以确保连接的安全性。
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (ws *WebsocketService) HandleWebSocket(c *gin.Context) {
	//获取当前用户信息
	query := c.Request.URL.Query()
	userID := query.Get("uid")
	touserID := query.Get("touid")

	// 处理客户端发起的 WebSocket 连接升级请求的逻辑，如果升级过程顺利完成，那么后续就可以通过连接对象（conn）进行 WebSocket 消息的收发操作
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("ws connect fail")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("ws connect success")
	client := &Client{
		SendID:    userID,
		ReceiveID: touserID,
		Socket:    conn,
		SendMsg:   make(chan []byte),
	}
	Manager.Register <- client
	go func() {
		userlock.Lock()
		defer userlock.Unlock()

		client.Read()
		client.Write()

		err := conn.Close()
		if err != nil {
			log.Println("Error closing connection:", err)
		}
	}()
}
