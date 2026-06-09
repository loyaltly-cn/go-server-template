package query

import "gorm.io/gorm"

func Execute(
	db *gorm.DB,
	model any,
	req Request,
) (Result, error) {

	var (
		list  []map[string]any
		total int64
	)

	// ====== 只构建过滤条件（不含分页） ======
	base := Apply(db, model, req)

	// ====== COUNT ======
	countDB := base.Session(&gorm.Session{})

	if err := countDB.Count(&total).Error; err != nil {
		return Result{}, err
	}

	// ====== LIST ======
	listDB := paginate(
		base.Session(&gorm.Session{}),
		req.Page,
		req.Size,
	)

	// 用 Scan 保证只返回 select 字段
	if err := listDB.
		Find(&list).
		Error; err != nil {
		return Result{}, err
	}

	return Result{
		List:  list,
		Total: total,
	}, nil
}
