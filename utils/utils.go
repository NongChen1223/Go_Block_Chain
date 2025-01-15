package utils

import (
	"golang.org/x/crypto/bcrypt"
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
