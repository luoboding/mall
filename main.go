package main

import (
	"github.com/gin-gonic/gin"
	authentication "github.com/luoboding/mall/api/authentication"
	user "github.com/luoboding/mall/api/user"
)

func main() {
	r := gin.Default()
	r.POST("/user", user.Signup)
	r.POST("/authentication", authentication.Signin)
	r.Run(":3000") // 监听并在 0.0.0.0:8080 上启动服务
}
