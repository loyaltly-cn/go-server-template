package user

type Service struct {
	repo *Repo
}

func NewService(r *Repo) *Service {
	return &Service{repo: r}
}

func (s *Service) Create(req CreateUserRequest) (User, error) {

	user := User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	err := s.repo.Create(&user)
	return user, err
}

func (s *Service) GetByID(id int64) (User, error) {
	return s.repo.GetByID(id)
}
