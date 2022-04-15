package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/luoboding/mall/api/authentication"
	"github.com/luoboding/mall/api/catalogue"
	"github.com/luoboding/mall/api/file"
	"github.com/luoboding/mall/api/product"
	"github.com/luoboding/mall/api/user"
	"github.com/luoboding/mall/middleware/authorization"
	"github.com/luoboding/mall/sql/migrations"
	"github.com/luoboding/mall/utils"
)

type Test struct {
	Name string `example:"title"`
}

func main() {
	// 运行migration
	migrations.Migrate()

	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/user", user.Signup)
	r.POST("/authentication", authentication.Signin)

	product_group := r.Group("/product")
	product_group.Use(authorization.Auth)
	product_group.POST("", product.Create)
	product_group.GET("", product.List)
	product_group.GET("/:id", product.One)
	product_group.PATCH("/:id", product.Update)

	catalogue_group := r.Group("/catalogue")
	catalogue_group.Use(authorization.Auth)
	catalogue_group.POST("", catalogue.Create)
	catalogue_group.PATCH("/:id", catalogue.Update)
	catalogue_group.GET("/:id", catalogue.One)

	file_group := r.Group("/file")
	file_group.Use(authorization.Auth)
	file_group.POST("", file.Create)

	// 获取环境变量
	port := utils.If(os.Getenv("PORT") != "", os.Getenv("PORT"), "3000")
	r.Run(fmt.Sprintf(":%s", port))
}
