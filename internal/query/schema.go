package query

import (
	"strings"

	"gorm.io/gorm"
)

func getSchemaMap(db *gorm.DB, model any) map[string]string {

	stmt := &gorm.Statement{DB: db}

	_ = stmt.Parse(model)

	m := make(map[string]string)

	for _, f := range stmt.Schema.Fields {
		m[strings.ToLower(f.Name)] = f.DBName
	}

	return m
}

// 白名单（防止乱查字段）
func allowField(field string, schema map[string]string) (string, bool) {
	col, ok := schema[strings.ToLower(field)]
	return col, ok
}
