package tool

import (
	"errors"
	"github.com/TskFok/GinApi/app/utils/conf"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	Phone                string `json:"phone,omitempty"`
	Uid                  uint32 `json:"uid,omitempty"`
	jwt.RegisteredClaims `json:"jwt.RegisteredClaims"`
}

func JwtToken(id uint32) string {
	secret := []byte(conf.JwtSecret)

	newClaims := &Claims{
		Phone: "1881121211221",
		Uid:   id,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "user_system",
			//三小时超时
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour * time.Duration(1))),
			//生效时间
			NotBefore: jwt.NewNumericDate(time.Now()),
			//签发时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	token, err := t.SignedString(secret)

	if nil != err {
		panic(err)
	}

	return token
}

func TokenInfo(token string) (*Claims, error) {
	tokens, error := jwt.ParseWithClaims(token, &Claims{}, secret())

	if nil != error {
		return nil, error
	}

	if claims, ok := tokens.Claims.(*Claims); ok && tokens.Valid {
		return claims, nil
	}

	return nil, errors.New("unknown error")
}

func secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.JwtSecret), nil
	}
}
