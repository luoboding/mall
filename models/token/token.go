package token

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/luoboding/mall/models/user"
)

const (
	VERIFY_KEY = "mall"
)

type UserAuthorization struct {
	Id   uint64
	Kind user.UserKind
	Name string
}

type UserClaims struct {
	*jwt.RegisteredClaims
	*UserAuthorization
}

func (c *UserAuthorization) Create_JWT() (string, error) {
	claim := &UserClaims{
		&jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		c,
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return t.SignedString([]byte(VERIFY_KEY))
}

// 生成令牌
func New(id uint64, kind user.UserKind) (string, error) {
	user := UserAuthorization{
		Id:   id,
		Kind: kind,
		Name: "张三",
	}
	return user.Create_JWT()
}

// 解析令牌
func Parse(tokenString string) (*UserAuthorization, error) {
	token, e := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(VERIFY_KEY), nil
	})
	if !token.Valid {
		return nil, e
	}
	claim := token.Claims.(*UserClaims)
	return claim.UserAuthorization, e
}
