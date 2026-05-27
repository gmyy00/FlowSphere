// Package dto 定义数据传输对象
// 用于 HTTP 请求和响应的数据结构
package dto

// RegisterRequest 是注册请求
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20,alphanum"` // 用户名，3-20字符
	Email    string `json:"email" binding:"required,email"`                    // 邮箱
	Password string `json:"password" binding:"required,min=6"`                 // 密码，至少6字符
}

// LoginRequest 是登录请求
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"` // 邮箱
	Password string `json:"password" binding:"required"`    // 密码
}

// RefreshRequest 是刷新令牌请求
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"` // 刷新令牌
}

// UserResponse 是用户信息响应
type UserResponse struct {
	ID        string `json:"id"`         // 用户 ID
	Username  string `json:"username"`   // 用户名
	Email     string `json:"email"`      // 邮箱
	CreatedAt string `json:"created_at"` // 创建时间
}

// LoginResponse 是登录响应
type LoginResponse struct {
	AccessToken  string       `json:"access_token"`  // 访问令牌
	RefreshToken string       `json:"refresh_token"` // 刷新令牌
	ExpiresIn    int64        `json:"expires_in"`    // 过期时间（秒）
	User         UserResponse `json:"user"`          // 用户信息
}

// TokenResponse 是令牌响应
type TokenResponse struct {
	AccessToken string `json:"access_token"` // 访问令牌
	ExpiresIn   int64  `json:"expires_in"`   // 过期时间（秒）
}
