package query

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func mapFields(input []string, schema map[string]string) []string {

	out := make([]string, 0, len(input))

	for _, f := range input {
		if col, ok := schema[strings.ToLower(f)]; ok {
			out = append(out, col)
		}
	}

	return out
}

func Apply(db *gorm.DB, model any, req Request) *gorm.DB {

	schema := getSchemaMap(db, model)

	// SELECT
	if len(req.Select) > 0 {
		cols := mapFields(req.Select, schema)
		if len(cols) > 0 {
			db = db.Select(cols)
		}
	}

	// WHERE
	if req.Where != nil {
		sql, args := buildNode(req.Where, schema)
		if sql != "" {
			db = db.Where(sql, args...)
		}
	}

	// JOIN
	db = applyJoins(db, req.Joins)

	// PRELOAD
	db = applyPreload(db, req.Preload)

	// ORDER
	if req.OrderBy != "" {
		if col, ok := allowField(req.OrderBy, schema); ok {

			order := "asc"
			if strings.ToLower(req.Order) == "desc" {
				order = "desc"
			}

			db = db.Order(col + " " + order)
		}
	}

	return db
}

func buildNode(n *Node, schema map[string]string) (string, []any) {

	if len(n.Children) == 0 {

		col, ok := allowField(n.Field, schema)
		if !ok {
			return "", nil
		}

		switch n.Op {

		case "=":
			return fmt.Sprintf("%s = ?", col), []any{n.Value}

		case ">":
			return fmt.Sprintf("%s > ?", col), []any{n.Value}

		case "<":
			return fmt.Sprintf("%s < ?", col), []any{n.Value}

		case "like":
			return fmt.Sprintf("%s LIKE ?", col), []any{"%" + fmt.Sprint(n.Value) + "%"}
		}
	}

	if n.Op == "and" {
		sqls := []string{}
		args := []any{}

		for _, c := range n.Children {
			s, a := buildNode(c, schema)
			if s != "" {
				sqls = append(sqls, s)
				args = append(args, a...)
			}
		}

		return "(" + strings.Join(sqls, " AND ") + ")", args
	}

	if n.Op == "or" {
		sqls := []string{}
		args := []any{}

		for _, c := range n.Children {
			s, a := buildNode(c, schema)
			if s != "" {
				sqls = append(sqls, s)
				args = append(args, a...)
			}
		}

		return "(" + strings.Join(sqls, " OR ") + ")", args
	}

	return "", nil
}
