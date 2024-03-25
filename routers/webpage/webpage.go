package webpage

import (
	"WebIM/service/webpage"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
}

func (wpr *RouterGroup) InitWebPageRouter(router *gin.RouterGroup) {
	pageRouter := router.Group("")
	{
		wpService := webpage.WpService{}
		pageRouter.GET("//favicon.ico", wpService.GetIco)
		pageRouter.GET("/index", wpService.GetIndex)
		pageRouter.GET("/login", wpService.GetLogin)
		pageRouter.GET("/register", wpService.GetRegister)
		pageRouter.GET("/chat", wpService.GetChat)
	}
}
