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
