package vars

import "github.com/amstee/easy-cut/src/common"

type Barber struct {
	Experience string 		`json:"experience,omitempty"`
	Style string 			`json:"style,omitempty"`
	Created string 			`json:"created,omitempty"`
	Updated string 			`json:"updated,omitempty"`
}

type BarberResponse struct {
	User common.User 		`json:"user"`
	Barber Barber			`json:"barber"`
}