package product

import "time"

type Product struct {
	Id           uint64   `gorm:"primaryKey"`
	Title        string   // 产品标题
	Sub_Title    string   // 产品简介
	Thumbnail    string   // 标题缩略图
	Pictures     []string // 产品图片
	Catalogue_ID uint64   // 分类id
	Description  string   // 产品详细描述
	Sort         uint8    // 顺序
	Status       uint     `gorm:"index"` // 状态 0 开启 1 下架 2 售罄
	CreatedAt    time.Time
}
