package vars

type AppMetadata struct {
	Authorization struct {
		Groups []string	`json:"groups"`
	}	`json:"authorization"`
}

type UserCreation struct {
	Connection	string 		`json:"connection"`
	Email		string		`json:"email"`
	Username	string		`json:"username"`
	Password	string		`json:"password"`
	Phone		string		`json:"phone_number"`
	EVerified	bool		`json:"email_verified"`
	VerifyEmail	bool		`json:"verify_email"`
	VerifyPhone	bool		`json:"phone_verified"`
	Metadata 	AppMetadata	`json:"app_metadata"`
}