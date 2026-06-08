package user

import "gorm.io/gorm"

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) Create(u *User) error {
	return r.db.Create(u).Error
}

func (r *Repo) GetByID(id int64) (User, error) {
	var u User
	err := r.db.First(&u, id).Error
	return u, err
}
