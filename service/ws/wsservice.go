package ws

import (
	"WebIM/global"
	"WebIM/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"log"
)

func (cm *ClientManager) Start() {
	for {
		log.Println("----------监听管道通信----------")
		select {
		case conn := <-Manager.Register: // 监听用户连接
			fmt.Printf("有新的连接: %v\n", conn.SendID)
			Manager.Clients[conn.SendID] = conn
			replyMsg := ReplyMsg{
				From:    "系统消息",
				Code:    50002,
				Content: "已经连接到服务器",
			}
			msg, _ := json.Marshal(replyMsg)
			_ = conn.Socket.WriteMessage(1, msg)
		case conn := <-Manager.Unregister: // 监听用户断开连接
			fmt.Printf("连接断开%s\n", conn.SendID)
			if _, ok := Manager.Clients[conn.SendID]; ok {
				replyMsg := &ReplyMsg{
					From:    "系统消息",
					Code:    50003,
					Content: "连接中断",
				}
				msg, _ := json.Marshal(replyMsg)
				_ = conn.Socket.WriteMessage(8, msg)
				close(conn.SendMsg)
				delete(Manager.Clients, conn.SendID)
			}
		case broadcast := <-Manager.Broadcast: // 遍历已连接的客户端，并向接收者发送消息
			switch broadcast.MsgType {
			case "private":
				message := broadcast.Message
				receiveID := broadcast.Client.ReceiveID

				msg := &models.Message{
					SendID:    broadcast.Client.SendID,
					ReceiveID: broadcast.Client.ReceiveID,
					Content:   message,
					MsgType:   broadcast.MsgType,
					MediaType: broadcast.MediaType,
				}

				for id, client := range Manager.Clients {
					if id != receiveID { // 寻找接收方ID对应的ws连接，找不到就开下一次循环，找到就继续执行
						continue
					}
					select {
					case client.SendMsg <- message:
					default:
						close(client.SendMsg)
						delete(Manager.Clients, client.SendID)
					}
				}
				res := msg.SaveMessage()
				if !res {
					fmt.Println("保存失败")
					return
				}
				fmt.Println("保存成功")
			case "group":
				message := broadcast.Message
				receiveID := broadcast.Client.GroupID
				fmt.Println(message)
				fmt.Println(receiveID)
				msg := &models.Message{
					SendID:    broadcast.Client.SendID,
					ReceiveID: broadcast.Client.GroupID,
					Content:   message,
					MsgType:   broadcast.MsgType,
					MediaType: broadcast.MediaType,
				}
				fmt.Println(msg)
				res := msg.SaveMessage()
				if !res {
					fmt.Println("保存失败")
					return
				}
				fmt.Println("保存成功")
			}
		}
	}
}

func (c *Client) Read() { // 从客户端ws不断读取消息(json格式)，并将消息放进SendMsg{}
	defer func() {
		Manager.Unregister <- c
		_ = c.Socket.Close()
	}()
	for {
		//c.Socket.PongHandler()
		//_, _ = global.RDB.Expire(global.Ctx, c.SendID, time.Hour*24*30*3).Result() // 建立连接3个月过期
		sendMsg := SendMsg{}
		var data map[string]any
		err := c.Socket.ReadJSON(&data) // 从客户端中读取一个JSON格式的消息到data中
		if err != nil {
			logrus.Error(err)
			Manager.Unregister <- c
			_ = c.Socket.Close()
			break
		}
		// 解析消息类型
		sendMsg.Message = []byte(data["message"].(string))
		sendMsg.MsgType = data["msgType"].(string)
		sendMsg.MediaType = data["mediaType"].(string)
		Manager.Broadcast <- &Broadcast{
			Client:    c,
			Message:   sendMsg.Message,
			MsgType:   sendMsg.MsgType,
			MediaType: sendMsg.MediaType,
		}
		global.RDB.Publish(global.Ctx, c.SendID, Manager.Broadcast)
	}
}

func (c *Client) Write() { // 将消息发送给指定用户
	defer func() {
		_ = c.Socket.Close()
	}()
	for {
		select {
		// 上述start方法中，第三个case会将消息放进对应的client中，这里监听client.SendMsg通道并将消息显示出来
		case message, ok := <-c.SendMsg:
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			msg, _ := json.Marshal(string(message))
			_ = c.Socket.WriteMessage(1, msg)
		}
	}
}
