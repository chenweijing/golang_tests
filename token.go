package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

/*
iss: 签发者
sub: 面向的用户
aud: 接收方
exp: 过期时间
nbf: 生效时间
iat: 签发时间
jti: 唯一身份标识
*/

var (
	key []byte = []byte("Hello World！This is secret!")
)

// 标准产生方式
// 产生json web token
func GenToken() string {
	claims := &jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix()),
		ExpiresAt: int64(time.Now().Unix() + 1000),
		Issuer:    "Bitch",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		fmt.Println("err:", err)
		return ""
	}

	return ss
}

// 校验token是否有效
func CheckToken(token string) bool {
	_, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		fmt.Println("parase with claims failed.", err)
		return false
	}
	return true
}

// token 加入自定义模块
type UserInfo struct {
	ID       uint64
	Username string
}

// 产生token，里面包含自定义模块
func CreateToken(user *UserInfo) (tokenss string, err error) {
	//自定义claim
	claim := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Unix() + 1000,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenss, err = token.SignedString(key)
	return tokenss, err
}

// 方法，返回一个keyfunc的接口。
func secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return key, nil
	}
}

// 解析
func ParseToken(tokenss string) (user *UserInfo, err error) {
	user = &UserInfo{}
	token, err := jwt.Parse(tokenss, secret())

	if err != nil {
		return
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("cannot convert claim to mapclaim")
		return
	}
	//验证token，如果token被修改过则为false
	if !token.Valid {
		err = errors.New("token is invalid")
		return
	}

	user.ID = uint64(claim["id"].(float64))
	user.Username = claim["username"].(string)
	return
}

func main() {
	// 标准
	token := GenToken()
	fmt.Println("token:", token)
	fmt.Println(CheckToken(token))

	fmt.Println("-------------- new token ------------------")

	// 自定义
	user := &UserInfo{
		ID:       1234567,
		Username: "tom",
	}

	// create token
	tokenss, err := CreateToken(user)
	fmt.Println("token:", tokenss, " err:", err)

	// check token
	fmt.Println(CheckToken(tokenss))

	// Parse
	user2, err2 := ParseToken(tokenss)
	fmt.Println("user:", user2, " err2:", err2)
}
