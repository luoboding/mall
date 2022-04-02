package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	user "github.com/luoboding/mall/models/user"
)

func Signup(context *gin.Context) {
	var user user.User
	err := context.ShouldBindJSON(&user)
	if err != nil || user.Check() != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	if user.IsExist() {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "用户已存在",
		})
		return
	}
	user.Encrypt_password()
	e := user.Create()
	if e != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": e.Error(),
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "注册成功",
	})
}
