package user

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/luoboding/mall/db"
	"github.com/luoboding/mall/utils"
	"github.com/luoboding/mall/utils/errors"
)

type UserKind int

const (
	UserKindNormal  UserKind = 1                   // 普通用户
	UserKindManager UserKind = UserKindNormal << 1 // 管理员
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Nickname  sql.NullString `json:"nickname"`
	Username  string         `json:"username" gorm:"index"`
	Password  string         `json:"password" gorm:"index"`
	Phone     string         `json:"phone" gorm:"index"`
	Gender    uint8          `json:"gender"`
	Avatar    sql.NullString `json:"avatar"`
	Status    uint           `json:"status" gorm:"index"`
	CreatedAt time.Time      `json:"created_at"`
}

func (user *User) Encrypt_password() {
	user.Password = utils.Encrypt(user.Password)
}

func (user *User) Validate() *errors.Error {
	if user.Username == "" || user.Password == "" {
		e := errors.New("参数错误")
		e.Code = 400
		return e
	}
	return nil
}

func (user *User) DoesSameUserNameExist() bool {
	var count int64
	db := db.Get_DB()
	db.Table("users").Where("username = ?", user.Username).Count(&count)
	return count > 0
}

func (user *User) Create() error {
	db := db.Get_DB()
	if err := user.Validate(); err != nil {
		err.Code = 400
		return err
	}
	if user.DoesSameUserNameExist() {
		e := errors.New("用户已经存在")
		e.Code = http.StatusConflict
		return e
	}
	user.Encrypt_password()
	result := db.Create(user)
	if result.Error != nil {
		e := errors.New(result.Error.Error())
		e.Code = http.StatusBadRequest
		return e
	}
	return nil
}
