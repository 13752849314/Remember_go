package utils

import (
	"github.com/dgrijalva/jwt-go"
	"remember/config"
	"remember/entity"
	"strconv"
	"time"
)

var key []byte

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func init() {
	key = []byte(config.Configure.Remember.Key)
}

func CreatToken(user *entity.User) (string, error) {
	claim := new(MyClaims)
	claim.Username = user.Username
	claim.Id = strconv.Itoa(int(user.ID))
	claim.NotBefore = time.Now().Unix()
	claim.ExpiresAt = time.Now().Unix() + 60*60*3
	claim.Issuer = config.Configure.Remember.Developers
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return signedString, nil
}

func CheckToken(token string) (*MyClaims, error) {
	token1, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	claim := token1.Claims.(*MyClaims)
	return claim, nil
}
