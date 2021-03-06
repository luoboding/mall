package user

import (
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

type UserProlfile struct {
	Nickname string `json:"nickname" gorm:"comment:昵称"`
	Phone    string `json:"phone" gorm:"index;comment:电话号码"`
	Gender   uint8  `json:"gender" gorm:"index;comment:性别"`
	Avatar   string `json:"avatar" gorm:"comment:头像"`
	Status   uint   `json:"status" gorm:"index;comment:状态0正常1禁用"`
}

// 实现tabler协议
func (UserProlfile) TableName() string {
	return "users"
}

type User struct {
	UserProlfile
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"index;comment:用户名"`
	Password  string    `json:"password" gorm:"index;comment:密码"`
	CreatedAt time.Time `json:"created_at"`
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
		e := errors.New("密码错误")
		e.Code = http.StatusBadRequest
		return e
	}
	return nil
}

func (u *User) Update() error {
	conn := db.Get_DB()
	var user User
	q := conn.Find(&user)
	if q.Error != nil {
		return q.Error
	}
	action := q.Updates(u)
	if action.Error != nil {
		return action.Error
	}
	return nil
}
