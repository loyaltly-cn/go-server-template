package query

type Result struct {
	List  []map[string]any `json:"list"`
	Total int64            `json:"total"`
}
