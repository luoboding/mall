package main

import (
	"github.com/gin-gonic/gin"
	authentication "github.com/luoboding/mall/api/authentication"
	catalogue "github.com/luoboding/mall/api/catalogue"
	user "github.com/luoboding/mall/api/user"
	"github.com/luoboding/mall/middleware/authorization"
)

func main() {
	r := gin.Default()

	r.Use(authorization.Auth)

	r.POST("/user", user.Signup)
	r.POST("/authentication", authentication.Signin)
	catalogue_group := r.Group("/catalogue")
	{
		catalogue_group.POST("", catalogue.Create)
		catalogue_group.PATCH("/:id", catalogue.Update)
	}
	r.Run(":3000") // 监听并在 0.0.0.0:8080 上启动服务
}
