package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"go_server/config"
	"golang.org/x/crypto/bcrypt"
	"time"
)

/**
 *  HashPassword
 *  @Description: 加密密码
 *  @param pwd
 *  @return string
 *  @return error
 */
func HashPassword(pwd string) (string, error) {
	//bcrypt.GenerateFromPassword接收两个参数，第一个是转换为字节byte类型的切片，第二个是加密强度
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 12)
	return string(hash), err
}

/**
 *	CompareHashAndPassword
 *  @Description: 比较加密密码和明文密码是否相等
 *  @param e 数据库密码
 *  @param p 明文密码
 *  @return bool
 *  @return error
 */
func CompareHashAndPassword(e string, p string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(e), []byte(p))
	if err != nil {
		return false, err
	}
	return true, nil
}

/**
 *  GenerateJWT
 *  @Description: 用用户名生成一个JWT
 *  @param username
 *  @return string
 *  @return error
 */
func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, //加密方式 可选
		jwt.MapClaims{
			"username": username,                                                                              //用户名
			"exp":      time.Now().Add(time.Hour * time.Duration(config.AppConfig.JWT.ExpirationHour)).Unix(), //过期时间时间戳 获取当前时间 + 72小时 Unix()转时间戳方法
		})
	signedToken, err := token.SignedString([]byte(config.AppConfig.JWT.PrivateKey))
	return "Bearer" + signedToken, err
}
