package users

import (
	"CopyQQ/global"
	"CopyQQ/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfoService struct {
}

// GetUserList
// @Tags 查询所有用户
// @Success 200 {string} json{"code","message"}
// @Router /users/getUserList [get]
func (us *UserInfoService) GetUserList(context *gin.Context) {
	user := models.User{}
	userList := user.GetUserList()
	context.JSON(http.StatusOK, global.Response{
		Code: http.StatusOK,
		Data: userList,
		Msg:  "查询成功",
	})
}

// GetUserInfo
// @Tags 根据条件查询用户信息
// @Success 200 {string} json{"code","message"}
// @Router /users/getUserInfo [get]
// GetUserInfo 获取用户信息
func (us *UserInfoService) GetUserInfo(context *gin.Context) {
	user := models.User{}
	if !user.CheckUserExists("name", context.PostForm("name")) {
		context.JSON(-1, global.Response{
			Code: -1,
			Data: nil,
			Msg:  "用户名不存在",
		})
		return
	}
	if !user.GetUserInfo("name", context.PostForm("name")) {
		context.JSON(-1, global.Response{
			Code: -1,
			Data: nil,
			Msg:  "查询失败",
		})
		return
	}
	context.JSON(http.StatusOK, global.Response{
		Code: http.StatusOK,
		Data: user,
		Msg:  "查询成功",
	})
}
