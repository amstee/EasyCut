package vars

type Salon struct {
	UserId string 		`json:"user_id"`
	Address string 		`json:"address"`
	EmployeeNumber int 	`json:"employee_number"`
	Barber []string 	`json:"barbers"`
	Website string 		`json:"website"`
	Created string 		`json:"created"`
	Updated string 		`json:"updated"`
}