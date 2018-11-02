package vars

type Salon struct {
	Id string 			`json:"_id,omitempty"`
	UserId string 		`json:"user_id"`
	Address string 		`json:"address,omitempty"`
	EmployeeNumber int 	`json:"employee_number,omitempty"`
	Barber []string 	`json:"barbers,omitempty"`
	Website string 		`json:"website,omitempty"`
	Created string 		`json:"created,omitempty"`
	Updated string 		`json:"updated,omitempty"`
}