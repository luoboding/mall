package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	authentication "github.com/luoboding/mall/api/authentication"
	user "github.com/luoboding/mall/api/user"
	"github.com/luoboding/mall/middleware/authorization"
)

func main() {
	r := gin.Default()

	r.Use(authorization.Auth)

	r.POST("/user", user.Signup)
	r.POST("/authentication", authentication.Signin)
	r.GET("/product", func(c *gin.Context) {
		fmt.Println("product", c.Request.Header.Get("x-consumer-id"), c.Request.Header.Get("x-consumer-kind"))
	})
	r.Run(":3000") // 监听并在 0.0.0.0:8080 上启动服务
}
