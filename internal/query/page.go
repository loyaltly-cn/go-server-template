package query

import "gorm.io/gorm"

func paginate(db *gorm.DB, page, size int) *gorm.DB {

	if page <= 0 {
		page = 1
	}

	if size <= 0 {
		size = 20
	}

	offset := (page - 1) * size

	return db.Offset(offset).Limit(size)
}
