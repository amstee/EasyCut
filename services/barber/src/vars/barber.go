package vars

import "github.com/amstee/easy-cut/src/common"

type Barber struct {
	Name string 			`json:"name,omitempty"`
	Address string 			`json:"name,omitempty"`
	Price int				`json:"price,omitempty"`
	Experience string 		`json:"experience,omitempty"`
	Style string 			`json:"style,omitempty"`
	Created string 			`json:"created,omitempty"`
	Updated string 			`json:"updated,omitempty"`
}

type BarberResponse struct {
	User common.User 		`json:"user"`
	Barber Barber			`json:"barber"`
}