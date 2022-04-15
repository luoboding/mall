package file

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// 图片上传
func Create(c *gin.Context) {
	f, e := c.FormFile("file")
	name := c.Request.Form.Get("name")
	if e != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": e.Error(),
		})
		return
	}
	path, _ := os.Getwd()
	folder := filepath.Join(path, "files")
	if name == "" {
		name = filepath.Base(f.Filename)
	}
	os.Mkdir(folder, 0777)
	err := c.SaveUploadedFile(f, filepath.Join(folder, name))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": e.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
	})

}
