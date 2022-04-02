package catalogue

import (
	"database/sql"
	"time"
)

type Catalogue struct {
	ID        uint `gorm:"primaryKey"`
	Nickname  sql.NullString
	Username  string `gorm:"index"`
	Password  string
	Phone     string `gorm:"index"`
	Gender    uint8
	Avatar    sql.NullString
	Status    uint `gorm:"index"`
	CreatedAt time.Time
}
