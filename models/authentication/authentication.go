package authentication

import (
	"errors"

	"github.com/luoboding/mall/db"
	"github.com/luoboding/mall/models/token"
	"github.com/luoboding/mall/models/user"
)

type Authentication struct {
	Username string `gorm:"index"`
	Password string
	ID       uint `gorm:"primaryKey"`
}

func (a *Authentication) Create_JWT() (string, error) {
	return token.New(uint64(a.ID), user.UserKindNormal)
}

func (a *Authentication) Signin() (string, error) {
	u := &user.User{
		ID:       a.ID,
		Password: a.Password,
		Username: a.Username,
	}
	if err := u.Validate(); err != nil {
		return "", err
	}
	if !u.DoesSameUserNameExist() {
		return "", errors.New("用户不存在")
	}
	u.Encrypt_password()
	connection := db.Get_DB()
	r := connection.Table("users").Where("username = ? and password = ?", u.Username, u.Password).First(a)
	if r.Error != nil {
		return "", r.Error
	}
	return a.Create_JWT()
}
