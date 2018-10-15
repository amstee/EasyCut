package types

type PermissionsParam struct {
	Scopes []string			`json:"scopes"`
}

type PermissionsResponse struct {
	Scopes map[string]bool	`json:"scopes"`
}
