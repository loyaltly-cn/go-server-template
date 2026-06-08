package user

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(name, email string) (*User, error) {

	user := &User{
		Name:  name,
		Email: email,
	}

	err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) GetByID(id int64) (*User, error) {
	return s.repo.GetByID(id)
}
