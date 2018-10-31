package common

type AppMetadata struct {
	Authorization struct {
		Groups []string	`json:"groups"`
	}	`json:"authorization"`
}

type UserMetadata struct {
	Username string 	`json:"username"`
	Address string 		`json:"address"`
	Phone string 		`json:"phone"`
	Description string 	`json:"description"`
}

type Identity struct {
	Connection string 		`json:"connection"`
	UserId string 			`json:"user_id"`
	Provider string 		`json:"provider"`
	Social bool 			`json:"isSocial"`
}

type User struct {
	Email string				`json:"email"`
	EmailVerified bool			`json:"email_verified"`
	Updated string				`json:"updated_at"`
	Created string 				`json:"created_at"`
	Picture string 				`json:"picture"`
	UserId string 				`json:"user_id"`
	LastLogin string 			`json:"last_login"`
	Identities []Identity		`json:"identities"`
	Metadata AppMetadata		`json:"app_metadata"`
	UserMetadata UserMetadata	`json:"user_metadata"`
}