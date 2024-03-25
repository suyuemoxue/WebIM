package websocket

import (
	"WebIM/service/ws"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
}

func (r *RouterGroup) InitWebSocketRouter(router *gin.RouterGroup) {
	wsRouter := router.Group("ws").Use()
	{
		wsService := ws.WebsocketService{}
		wsRouter.GET("/chat", wsService.HandleWebSocket)
	}
}
