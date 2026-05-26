package token

// 刷新令牌仓储接口
type Repository interface {
	Create(token *RefreshToken) error
	FindByToken(token string) (*RefreshToken, error) // 通过令牌值寻找
	DeleteByUserID(userID string) error              // 删除某用户的所有令牌
	DeleteExpired() error                            // 删除所有过期令牌
}
