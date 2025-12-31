package user

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s service) CreateUser(user *User) (User, error) {
	createdUser, err := s.repo.CreateUser(user)
	if err != nil {
		return User{}, err
	}
	return *createdUser, nil

}

func (s service) DeleteUser() {
	s.repo.DeleteUser("1")
}
