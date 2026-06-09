package query

type Request struct {
	Select []string `json:"select"`
	Where  *Node    `json:"where"`

	Joins   []Join   `json:"joins"`
	Preload []string `json:"preload"`

	OrderBy string `json:"order_by"`
	Order   string `json:"order"`

	Page int `json:"page"`
	Size int `json:"size"`
}

type Node struct {
	Op       string  `json:"op"`
	Field    string  `json:"field,omitempty"`
	Value    any     `json:"value,omitempty"`
	Children []*Node `json:"children,omitempty"`
}

type Join struct {
	Table string `json:"table"`
	Cond  string `json:"cond"` // users.id = orders.user_id
}
