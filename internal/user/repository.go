package user

import "gorm.io/gorm"

type Repository interface {
	CreateUser(user *User) (*User, error)
	DeleteUser(id string) error
}

type repository struct {
	db any
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateUser(user *User) (*User, error) {

	return &User{
		ID:    "1",
		Name:  "mizan",
		Email: "mizan@mail.com",
	}, nil

}

func (r *repository) DeleteUser(id string) error {
	return nil
}
