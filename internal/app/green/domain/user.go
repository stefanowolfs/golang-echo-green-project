package domain

type User struct {
	ID   uint
	Name string
}

type UserRepository interface {
	GetByID(id uint) (*User, error)
	NewUser(user *User) (*User, error)
}
