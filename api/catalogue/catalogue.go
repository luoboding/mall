package catalogue

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luoboding/mall/models/catalogue"
)

func Create(c *gin.Context) {
	var catalogue catalogue.Catalogue
	err := c.ShouldBindJSON(&catalogue)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	e := catalogue.Create()
	if e != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": e.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "创建成功",
	})

}

func Update(c *gin.Context) {
	id := c.Param("id")
	var data catalogue.Catalogue
	e := c.ShouldBindJSON(&data)
	if e != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	i, e := strconv.Atoi(id)
	if e != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	data.ID = uint(i)
	update_error := data.Update()
	if update_error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": e.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "修改成功",
	})
}
