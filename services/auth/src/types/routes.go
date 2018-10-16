package types

type StatusResponse struct {
	Status string 			`json:"status"`
	Service string			`json:"service"`
	Version string			`json:"version"`
}

type PermissionsParam struct {
	Scopes []string			`json:"scopes"`
}

type TokenInfo struct {
	Token string			`json:"id_token"`
}

type PermissionsResponse struct {
	Scopes map[string]bool	`json:"scopes"`
}

type GroupsParam struct {
	Groups []string			`json:"groups"`
}

type GroupsResponse struct {
	Groups map[string]bool	`json:"groups"`
}