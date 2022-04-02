package authorization

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	token "github.com/luoboding/mall/models/token"
)

func Auth(c *gin.Context) {
	fmt.Println("c", c.Request.Header.Get("Authorization"))
	uri := c.Request.RequestURI
	tokenString := c.Request.Header.Get("Authorization")
	if uri != "/user" && uri != "/authentication" {
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "请登录",
			})
			return
		}
		claim, e := token.Parse(tokenString)
		if e != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "请登录",
			})
			return
		}
		// 将参数写回到header
		c.Request.Header.Set("x-consumer-id", fmt.Sprintf("%d", claim.Id))
		c.Request.Header.Set("x-consumer-kind", claim.Kind)
	}
}
