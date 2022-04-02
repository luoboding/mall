package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luoboding/mall/db"
	models "github.com/luoboding/mall/models/user"
)

func Signin(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil || user.Check() != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
	}
	if !user.IsExist() {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "用户不存在",
		})
	}
	user.Encrypt_password()
	connection := db.Get_DB()

	r := connection.Table("users").Where("username = ? and password = ?", user.Username, user.Password).First(&user)
	if r.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": r.Error.Error(),
		})
		return
	}
	token, e := user.Create_JWT()
	if e != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": e.Error(),
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"token": token,
	})
}
