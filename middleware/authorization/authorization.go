package authorization

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luoboding/mall/models/token"
)

func Auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "请登录",
		})
		return
	}
	claim, e := token.Parse(tokenString)
	if e != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "请登录",
		})
		return
	}
	// 将参数写回到header
	c.Request.Header.Set("x-consumer-id", fmt.Sprintf("%d", claim.Id))
	c.Request.Header.Set("x-consumer-kind", fmt.Sprintf("%d", claim.Kind))
}
