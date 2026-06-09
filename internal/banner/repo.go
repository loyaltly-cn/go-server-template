package banner

import (
	"server/internal/query"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(b *Banner) error {
	return r.db.Create(b).Error
}

func (r *Repository) Update(b *Banner) error {
	return r.db.Save(b).Error
}

func (r *Repository) Delete(id int64) error {
	return r.db.Delete(&Banner{}, id).Error
}

func (r *Repository) Patch(id int64, updates map[string]interface{}) error {
	return r.db.Model(&Banner{}).
		Where("id = ?", id).
		Updates(updates).
		Error
}

func (r *Repository) Query(
	req query.Request,
) (query.Result, error) {

	return query.Execute(
		r.db.Model(&Banner{}),
		&Banner{},
		req,
	)
}
