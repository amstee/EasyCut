package core

import "github.com/amstee/easy-cut/services/barber/src/vars"

func MtoL(barbers map[string]*vars.BarberResponse) []*vars.BarberResponse {
	var response []*vars.BarberResponse

	for _, v := range barbers {
		response = append(response, v)
	}
	return response
}