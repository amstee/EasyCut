package vars

type UserMetadata struct {
	Username string 	`json:"username"`
	Address string 		`json:"address"`
	Phone string 		`json:"phone"`
	Description string 	`json:"description"`
}

type User struct {
	UserId string 				`json:"user_id"`
	Email string				`json:"email"`
	Created string 				`json:"created_at"`
	Picture string 				`json:"picture"`
	LastLogin string 			`json:"last_login"`
	Metadata AppMetadata		`json:"app_metadata"`
	UserMetadata UserMetadata	`json:"user_metadata"`
}

type Barber struct {
	UserId string 			`json:"user_id"`
	Experience string 		`json:"experience,omitempty"`
	Style string 			`json:"style,omitempty"`
}

type BarberResponse struct {
	User User 				`json:"user"`
	Barber Barber			`json:"barber"`
}