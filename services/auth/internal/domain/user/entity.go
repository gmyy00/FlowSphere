package user

import "time"

// User 用户实体定义
type User struct {
	ID           string
	Username     string
	Email        string
	PasswordHash string
	CreateAt     time.Time
	UpdateAt     time.Time
}

// Role 为用户角色类型
type Role string

// 管理员角色与普通用户角色
const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)
