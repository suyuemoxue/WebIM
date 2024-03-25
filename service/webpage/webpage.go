package webpage

import (
	"WebIM/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WpService struct {
}

// GetIndex
// @Tags 首页
// @Success 200 {string} welcome
// @Router /index [get]
func (wps *WpService) GetIndex(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", global.Response{})
}

func (wps *WpService) GetLogin(context *gin.Context) {
	context.HTML(http.StatusOK, "login.html", global.Response{})
}

func (wps *WpService) GetRegister(context *gin.Context) {
	context.HTML(http.StatusOK, "register.html", global.Response{})
}

func (wps *WpService) GetChat(context *gin.Context) {
	context.HTML(http.StatusOK, "chat.html", global.Response{})
}

func (wps *WpService) GetIco(context *gin.Context) {
	context.File("favicon.ico")
}
