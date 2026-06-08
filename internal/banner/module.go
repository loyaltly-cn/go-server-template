package banner

import "gorm.io/gorm"

type Module struct {
	Handler *Handler
	Service *Service
	Repo    *Repository
}

func NewModule(db *gorm.DB) *Module {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	return &Module{
		Repo:    repo,
		Service: service,
		Handler: handler,
	}
}
