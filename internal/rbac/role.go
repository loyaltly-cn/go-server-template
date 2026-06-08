package rbac

type Role string

const (
	RoleAnonymous Role = "anonymous"
	RoleUser      Role = "user"
	RoleAdmin     Role = "admin"
)
