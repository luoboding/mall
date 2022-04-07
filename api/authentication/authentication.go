package authentication

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luoboding/mall/models/authentication"
)

func Signin(context *gin.Context) {
	var auth authentication.Authentication
	err := context.ShouldBindJSON(&auth)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
	}
	token, err := auth.Signin()
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"token": token,
	})
}
