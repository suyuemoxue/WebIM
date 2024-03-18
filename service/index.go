package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetIndex
// @Tags 首页
// @Success 200 {string} welcome
// @Router /index [get]
func GetIndex(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{})
}
