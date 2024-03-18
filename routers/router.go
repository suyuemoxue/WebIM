package routers

import (
	"CopyQQ/global"
	"CopyQQ/middlewares"
	"CopyQQ/routers/users"
	"CopyQQ/routers/websocket"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	Users users.RouterGroup
	WS    websocket.RouterGroup
}

var RouterGroupAPP RouterGroup

func InitRouter() {
	gin.SetMode(global.Config.System.Env) //设置了 Gin 框架的运行模式，根据全局配置中的环境参数来设置
	router := gin.Default()
	router.Use(middlewares.LogMiddleWare())

	//docs.SwaggerInfo.BasePath = "" //注册了一个用于 Swagger 文档展示的路由，可以通过该路由访问 Swagger 自动生成的 API 文档。
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//
	//router.LoadHTMLGlob("templates/*")
	//router.GET("/index", service.GetIndex)

	routerGroup := router.Group("")
	{
		RouterGroupAPP.Users.LoginRouter.InitLoginRouter(routerGroup) // loginGroup 用于提供登录注册等相关功能的路由与实现
		RouterGroupAPP.Users.UserRouter.InitUserRouter(routerGroup)   // userGroup 用于提供管理用户等功能的路由与实现
		RouterGroupAPP.WS.InitWebSocketRouter(routerGroup)
	}

	addr := global.Config.System.Address() //获取系统地址
	err := router.Run(addr)                //启动
	if err != nil {
		return
	}
}
