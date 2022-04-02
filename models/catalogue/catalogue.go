package catalogue

import (
	"database/sql"
	"time"
)

type Catalogue struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Thumbnail sql.NullString // 分类图标 参考美团外卖
	Sort      uint8          // 顺序
	Status    uint           `gorm:"index"` // 状态 0 开启 1 禁用
	CreatedAt time.Time
}
