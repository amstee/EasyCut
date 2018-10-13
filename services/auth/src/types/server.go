package types

type HttpMessage struct {
	Message string	`json:"message"`
	Success bool	`json:"success"`
}

type PermissionsParam struct {
	Scopes []string			`json:"scopes"`
}

type PermissionsResponse struct {
	Scopes map[string]bool	`json:"scopes"`
}
