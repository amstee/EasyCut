package vars

type GroupResponse struct {
	Name string 		`json:"name"`
	Success bool 		`json:"success"`
	Message string 		`json:"message,omitempty"`
}

type Groups struct {
	Groups []Group		`json:"groups"`
}

type Group struct {
	Id string 			`json:"_id"`
	Name string 		`json:"name"`
	Description string 	`json:"description"`
	Members []string 	`json:"members,omitempty"`
}