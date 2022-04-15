package catalogue

import (
	"errors"
	"fmt"
	"time"

	"github.com/luoboding/mall/db"
)

type Catalogue struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"comment:标题"`              // 分类标题
	Thumbnail string    `json:"thumbnail" gorm:"comment:缩略图"`         // 分类图标 参考美团外卖
	Sort      uint8     `json:"sort" gorm:"index;comment:排序"`         // 顺序
	Status    uint      `json:"status" gorm:"index;comment:状态0正常1禁用"` // 状态 0 开启 1 禁用
	CreatedAt time.Time `json:"created_at"`
}

func (c *Catalogue) validate() bool {
	fmt.Println("c", c.Title, c.Thumbnail)
	return c.Title != "" && c.Thumbnail != ""
}

func (c *Catalogue) doseNameExist() bool {
	connection := db.Get_DB()
	var count int64
	connection.Table("catalogues").Where("title = ?", c.Title).Count(&count)
	return count > 0
}

func (c *Catalogue) Create() error {
	if !c.validate() {
		return errors.New("参数错误")
	}
	if c.doseNameExist() {
		return errors.New("分类已存在")
	}
	connection := db.Get_DB()
	result := connection.Create(c)
	return result.Error
}

func (c *Catalogue) Update() error {
	connection := db.Get_DB()
	var instance Catalogue
	result := connection.Table("catalogues").Where("id = ?", c.ID).First(&instance)
	if result.Error != nil {
		return result.Error
	}
	action := connection.Model(instance).Updates(c)
	return action.Error
}

// 包方法
func One(id int) (*Catalogue, error) {
	var one Catalogue
	connection := db.Get_DB()
	r := connection.First(&one, id)
	return &one, r.Error
}
