package banner

import "server/internal/query"

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

func (s *Service) Update(id int64, req UpdateBannerRequest) error {
	return s.repo.Update(&Banner{
		ID:    id,
		Image: req.Image,
	})
}

func (s *Service) Delete(id int64) error {
	return s.repo.Delete(id)
}

func (s *Service) Patch(id int64, req PatchBannerRequest) error {
	updates := map[string]interface{}{}

	if req.Image != nil {
		updates["image"] = *req.Image
	}

	if len(updates) == 0 {
		return nil
	}

	return s.repo.Patch(id, updates)
}

func (s *Service) Query(
	req query.Request,
) (query.Result, error) {

	return s.repo.Query(req)
}
