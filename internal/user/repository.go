package user

import "gorm.io/gorm"

type Repository interface {
	CreateUser(user *User) (*User, error)
	DeleteUser(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateUser(user *User) (*User, error) {

	result := r.db.Create(user)

	return user, result.Error

}

func (r *repository) DeleteUser(id string) error {
	return nil
}
