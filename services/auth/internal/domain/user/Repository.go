package user

// 用户数据仓储接口，用于持久化数据
type Repository interface {
	Create(user *User) error
	FindByID(id string) (*User, error)
	FindByEmain(email string) (*User, error)
	FindByUsername(name string) (*User, error)
	Delete(id string) error
	Update(user *User) error
}
