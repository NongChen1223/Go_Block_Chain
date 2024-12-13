package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pwd string) (string, error) {
	//bcrypt.GenerateFromPassword接收两个参数，第一个是转换为字节byte类型的切片，第二个是加密强度
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 12)
	return string(hash), err
}
