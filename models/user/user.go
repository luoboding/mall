package user

import (
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/luoboding/mall/db"
	token "github.com/luoboding/mall/models/token"
)

const (
	SALT = "mall"
)

func encrypt(input string) string {
	sum := sha256.Sum256([]byte(input))
	result := sha256.Sum256([]byte(fmt.Sprintf("%x", sum) + SALT))
	return fmt.Sprintf("%x", result)
}

type User struct {
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

func (user *User) Encrypt_password() {
	user.Password = encrypt(user.Password)
}

func (user *User) Check() error {
	if user.Username == "" || user.Password == "" {
		return errors.New("参数错误")
	}
	return nil
}

func (user *User) Create() error {
	db := db.Get_DB()
	result := db.Create(&user)
	return result.Error
}

func (user *User) IsExist() bool {
	var count int64
	db := db.Get_DB()
	db.Table("users").Where("username = ?", user.Username).Count(&count)
	return count > 0
}

func (user *User) Create_JWT() (string, error) {
	return token.New(uint64(user.ID), "user")
}