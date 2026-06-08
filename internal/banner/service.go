package banner

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Create(req CreateBannerRequest) error {
	return s.repo.Create(&Banner{
		Image: req.Image,
	})
}

func (s *Service) List() ([]Banner, error) {
	return s.repo.GetAll()
}

func (s *Service) Get(id int64) (Banner, error) {
	return s.repo.GetByID(id)
}

func (s *Service) Update(id int64, req UpdateBannerRequest) error {
	return s.repo.Update(&Banner{
		ID:    id,
		Image: req.Image,
	})
}

func (s *Service) Delete(id int64) error {
	return s.repo.Delete(id)
}
