package vars

import "github.com/pkg/errors"

type Rate struct {
	Barber string 		`json:"barber,omitempty"`
	Salon string 		`json:"salon,omitempty"`
	Appointment string 	`json:"appointment,omitempty"`
	Comment string 		`json:"comment,omitempty"`
	Stars int 			`json:"stars"`
	Created string 		`json:"created"`
	Updated string 		`json:"updated"`
}

type Rating struct {
	Id string 			`json:"_id,omitempty"`
	Comment string 		`json:"comment,omitempty"`
	Stars int 			`json:"stars"`
	Created string 		`json:"created"`
	Updated string 		`json:"updated"`
	UserId string 		`json:"user_id"`
	TargetId string 	`json:"target_id"`
	TargetType string 	`json:"target_type"`
}

func (r *Rate) GetTarget() (string, error) {
	if r.Barber != "" {
		return r.Barber, nil
	}
	if r.Salon != "" {
		return r.Salon, nil
	}
	if r.Appointment != "" {
		return r.Appointment, nil
	}
	return "", errors.New("invalid rating")
}

func (r *Rate) GetTargetType() (string, error) {
	if r.Barber != "" {
		return "barber", nil
	}
	if r.Salon != "" {
		return "salon", nil
	}
	if r.Appointment != "" {
		return "appointment", nil
	}
	return "", errors.New("invalid rating")
}