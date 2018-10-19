package types

type OauthToken struct {
	GrantType string 		`json:"grant_type"`
	Audience string			`json:"audience"`
	ClientId string			`json:"client_id"`
	ClientSecret string		`json:"client_secret"`
}

type TokenType struct {
	Type string		`json:"token_type"`
	Access string 	`json:"access_token"`
	Scope string	`json:"scope"`
	Expires int 	`json:"expires_in"`
}

func (t *TokenType) Format() string {
	return t.Type + " " + t.Access
}

