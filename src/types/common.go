package types

type HttpMessage struct {
	Message string	`json:"message"`
	Success bool	`json:"success"`
}

type StatusResponse struct {
	Status string 			`json:"status"`
	Service string			`json:"service"`
	Version string			`json:"version"`
}

type UpdateGroup struct {
	Name string 		`json:"name"`
	Delete bool			`json:"delete"`
}

type UpdateGroups struct {
	Groups []UpdateGroup 	`json:"groups"`
}

func CreateGroup(group string, delete bool) UpdateGroups {
	var res UpdateGroups

	res.Groups = append(res.Groups, UpdateGroup{Name: group, Delete: delete})
	return res
}