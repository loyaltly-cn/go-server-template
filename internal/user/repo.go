package user

import (
	"gorm.io/gorm"
)

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

func (r *Repo) GetByOpenID(openid string) (*User, error) {

	var user User

	err := r.db.
		Where("open_id = ?", openid).
		First(&user).
		Error

	return &user, err
}

func (r *Repo) FindByID(id int64) (*User, error) {

	var u User

	err := r.db.First(&u, id).Error
	if err != nil {
		return nil, err
	}

	return &u, nil
}
