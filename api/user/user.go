package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luoboding/mall/models/user"
)

func Signup(context *gin.Context) {
	var user user.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	e := user.Create()
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": e.Error(),
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "注册成功",
	})
}
