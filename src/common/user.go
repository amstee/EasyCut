package common

type AppMetadata struct {
	Authorization struct {
		Groups []string	`json:"groups"`
	}	`json:"authorization"`
}

type UserMetadata struct {
	Username string 	`json:"username,omitempty"`
	Address string 		`json:"address,omitempty"`
	Phone string 		`json:"phone,omitempty"`
	Description string 	`json:"description,omitempty"`
	FavoriteBarbers []string 	`json:"favorite_barbers,omitempty"`
	FavoriteSalons []string 	`json:"favorite_salons,omitempty"`
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