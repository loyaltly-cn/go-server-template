package banner

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(b *Banner) error {
	return r.db.Create(b).Error
}

func (r *Repository) GetAll() ([]Banner, error) {
	var list []Banner
	err := r.db.Find(&list).Error
	return list, err
}

func (r *Repository) GetByID(id int64) (Banner, error) {
	var b Banner
	err := r.db.First(&b, id).Error
	return b, err
}

func (r *Repository) Update(b *Banner) error {
	return r.db.Save(b).Error
}

func (r *Repository) Delete(id int64) error {
	return r.db.Delete(&Banner{}, id).Error
}
