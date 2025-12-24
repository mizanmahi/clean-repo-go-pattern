package user

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s service) CreateUser() {
	s.repo.CreateUser(&User{})
}

func (s service) DeleteUser() {
	s.repo.DeleteUser("1")
}
