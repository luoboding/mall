package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luoboding/mall/user"
)

func main() {
	r := gin.Default()
	r.POST("/user", user.Signup)
	r.Run(":3000") // 监听并在 0.0.0.0:8080 上启动服务
}
