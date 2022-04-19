package user

import (
	"fmt"
	"net/http"
	"strconv"

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

func Update(c *gin.Context) {
	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	var u user.User
	fmt.Println("ok")
	if error := c.ShouldBindJSON(&u); error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": error.Error(),
		})
		return
	}
	u.ID = uint(id)
	err := u.Update()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
	})
}
