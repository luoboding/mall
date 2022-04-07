package product

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luoboding/mall/models/product"
)

func List(c *gin.Context) {
	size := c.DefaultQuery("size", "10")
	current := c.DefaultQuery("current", "1")
	price_gt := c.Query("price_gt")
	price_lt := c.Query("price_lt")
	catalogue_id := c.Query("catalogue_id")
	title := c.Query("title")
	order_by := c.QueryArray("order")

	isize, e := strconv.Atoi(size)
	if e != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "size参数为number类型",
		})
		return
	}

	icurrent, e := strconv.Atoi(current)
	if e != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "size参数为number类型",
		})
		return
	}

	pagination := product.Pagination{
		Size:    isize,
		Current: icurrent,
	}

	order := product.Order{
		By: order_by,
	}

	query := &product.ProductSearchQuery{
		Order:      order,
		Pagination: pagination,
	}

	if e != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": e.Error(),
		})
	}
	if price_gt != "" {
		i_price_gt, e := strconv.Atoi(price_gt)
		if e != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "price_gt参数为number类型",
			})
			return
		}
		query.Price_gt = i_price_gt
	}

	if price_lt != "" {
		i_price_lt, e := strconv.Atoi(price_lt)
		if e != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "price_gt参数为number类型",
			})
			return
		}
		query.Price_lt = i_price_lt
	}

	if catalogue_id != "" {
		i_catalogue_id, e := strconv.Atoi(catalogue_id)
		if e != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "catalogue_id参数为number类型",
			})
			return
		}
		query.Catalogue_ID = uint64(i_catalogue_id)
	}

	if title != "" {
		query.Title = title
	}

	response, e := product.List(query)
	if e != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": e.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func Create(c *gin.Context) {
	var request product.Product
	if e := c.ShouldBindJSON(&request); e != nil {
		fmt.Println("request is ", request)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}

	error := request.Create()
	if error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
	})
}

func One(c *gin.Context) {

}

func Update(c *gin.Context) {}
