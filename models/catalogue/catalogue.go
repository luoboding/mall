package catalogue

import (
	"time"

	"github.com/luoboding/mall/db"
)

type Catalogue struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string // 分类标题
	Thumbnail string // 分类图标 参考美团外卖
	Sort      uint8  // 顺序
	Status    uint   `gorm:"index"` // 状态 0 开启 1 禁用
	CreatedAt time.Time
}

func (c *Catalogue) Validate() bool {
	return c.Title != "" && c.Thumbnail != ""
}

func (c *Catalogue) Exist() bool {
	connection := db.Get_DB()
	var count int64
	connection.Table("catalogues").Where("title = ?", c.Title).Count(&count)
	return count > 0
}

func (c *Catalogue) Create() error {
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
