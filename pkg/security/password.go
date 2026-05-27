package security

import (
	"golang.org/x/crypto/bcrypt"
)

// bcryptCost 是 bcrypt 哈希的计算成本
// 值越大计算越慢，安全性越高
const bcryptCost = 12

// HashPassword 对密码进行 bcrypt 哈希
// 返回哈希后的字符串
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	return string(bytes), err
}

// CheckPassword 验证密码是否匹配哈希值
// 返回 true 表示密码正确
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
