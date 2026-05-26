package token

import (
	"time"
)

// 刷新令牌实体
type RefreshToken struct {
	ID       string
	User_id  string
	Token    string
	ExpireAt time.Time // 过期时间
	CreateAt time.Time
}
