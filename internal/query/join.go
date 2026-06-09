package query

import "gorm.io/gorm"

func applyJoins(db *gorm.DB, joins []Join) *gorm.DB {

	for _, j := range joins {
		// 注意：这里是安全版本（只能写 ON，不允许自由 SQL）
		db = db.Joins(j.Cond)
	}

	return db
}

func applyPreload(db *gorm.DB, preload []string) *gorm.DB {

	for _, p := range preload {
		db = db.Preload(p)
	}

	return db
}
