package ws

import (
	"CopyQQ/global"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"log"
)

// SendMsg 发送消息的结构体
type SendMsg struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
}

// ReplyMsg 回复消息的结构体
type ReplyMsg struct {
	From    string `json:"from"`
	Code    int    `json:"code"`
	Content string `json:"content"`
}

// Client 用户结构体
type Client struct {
	SendID    string
	ReceiveID string
	Socket    *websocket.Conn
	SendMsg   chan []byte
}

func (c *Client) Read() {
	defer func() {
		Manager.Unregister <- c
		_ = c.Socket.Close()
	}()
	for {
		c.Socket.PongHandler()
		sendMsg := SendMsg{}
		err := c.Socket.ReadJSON(&sendMsg)
		if err != nil {
			logrus.Error(err)
			//log.Println("数据格式不正确")
			Manager.Unregister <- c
			_ = c.Socket.Close()
			break
		}
		if sendMsg.Type == 1 { // 发送消息,从redis获取消息发送者和接收者ID
			//r1, _ := global.RDB.Get(global.Ctx, c.SendID).Result()
			//r2, _ := global.RDB.Get(global.Ctx, c.ReceiveID).Result()
			global.RDB.Incr(global.Ctx, c.SendID)
			Manager.Broadcast <- &Broadcast{
				Client:  c,
				Message: []byte(sendMsg.Content),
				Type:    1,
			}
		}
	}
}

func (c *Client) Write() {
	defer func() {
		_ = c.Socket.Close()
	}()
	for {
		select {
		case message, ok := <-c.SendMsg:
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			replyMsg := ReplyMsg{
				From:    "50001",
				Code:    0,
				Content: fmt.Sprintf("%s", string(message)),
			}
			msg, _ := json.Marshal(replyMsg)
			_ = c.Socket.WriteMessage(1, msg)
		}
	}
}

// Broadcast 广播类（包括广播内容和源用户）
type Broadcast struct {
	Client  *Client
	Message []byte
	Type    int
}

// ClientManager 用户管理
type ClientManager struct {
	Clients    map[string]*Client
	Broadcast  chan *Broadcast
	Reply      chan *Client
	Register   chan *Client
	Unregister chan *Client
}

func (cm *ClientManager) Start() {
	for {
		log.Println("----------监听管道通信----------")
		select {
		case conn := <-Manager.Register:
			fmt.Printf("有新的连接: %v\n", conn.SendID)
			Manager.Clients[conn.SendID] = conn
			replyMsg := ReplyMsg{
				Code:    50002,
				Content: "已经连接到服务器",
			}
			msg, _ := json.Marshal(replyMsg)
			_ = conn.Socket.WriteMessage(1, msg)
		}
	}
}

// Message 信息转JSON (包括：发送者、接收者、内容)
type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

// Manager 定义一个管理Manager
var Manager = ClientManager{
	Clients:    make(map[string]*Client), // 参与连接的用户，出于性能的考虑，需要设置最大连接数
	Broadcast:  make(chan *Broadcast),
	Register:   make(chan *Client),
	Reply:      make(chan *Client),
	Unregister: make(chan *Client),
}
