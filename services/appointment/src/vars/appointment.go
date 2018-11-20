package vars

import (
	"github.com/olivere/elastic"
	"net/url"
	"github.com/pkg/errors"
	"github.com/amstee/easy-cut/src/request"
	"github.com/amstee/easy-cut/src/config"
	"net/http"
)

type Appointment struct {
	Id string 				`json:"_id,omitempty"`
	UserId string 			`json:"user_id"`
	BarberId string 		`json:"barber_id"`
	Date string 			`json:"date"`
	Duration int 			`json:"duration"`
	Description string 		`json:"description,omitempty"`
	Created string 			`json:"created,omitempty"`
	Updated string 			`json:"updated,omitempty"`
}

type UpdateAppointment struct {
	Date string 		`json:"date,omitempty"`
	Duration int 		`json:"duration,omitempty"`
	Description string 	`json:"description,omitempty"`
}

type ExtractQuery struct {
	UserId string 			`json:"user_id"`
	BarberId string 		`json:"barber_id"`
	MinDate string 			`json:"min_date"`
	MaxDate string 			`json:"max_date"`
}

func (a *Appointment) CheckBarber(token string) error {
	resp, err := request.ExpectJson(config.GetServiceURL("barber") + "/get/" + a.UserId, http.MethodGet,
									"Bearer " + token, nil, nil)
	if err == nil && request.IsValid(resp.StatusCode) {
		return nil
	}
	return err
}

func (a *Appointment) CheckUser(token string) error {
	resp, err := request.ExpectJson(config.GetServiceURL("user") + "/get/" + a.BarberId, http.MethodGet,
		"Bearer " + token, nil, nil)
	if err == nil && request.IsValid(resp.StatusCode) {
		return nil
	}
	return err
}

func (a *Appointment) Verify(token string, id string) error {
	if a.UserId == "" && a.BarberId == "" {
		return errors.New("missing data")
	}
	if a.UserId == "" && a.BarberId != "" {
		a.UserId = id
		err := a.CheckBarber(token); if err != nil {
			return err
		}
	} else if a.UserId != "" && a.BarberId == "" {
		a.BarberId = id
		err := a.CheckBarber(token); if err != nil {
			return err
		}
		err = a.CheckUser(token); if err != nil {
			return err
		}
	} else {
		if a.UserId != id && a.BarberId != id {
			return errors.New("invalid user or barber")
		}
		err := a.CheckBarber(token); if err != nil {
			return err
		}
		err = a.CheckUser(token); if err != nil {
			return err
		}
	}
	return nil
}

func (e *ExtractQuery) ConstructQuery() (*elastic.BoolQuery) {
	query := elastic.NewBoolQuery()
	rangeQuery := elastic.NewRangeQuery("date") // .Format("dd/MM/yyyy||yyyy")
	checked := false

	if e.UserId != "" {
		query.Must(elastic.NewTermQuery("user_id", e.UserId))
	}
	if e.BarberId != "" {
		query.Must(elastic.NewTermQuery("barber_id", e.BarberId))
	}
	if e.MinDate != "" {
		rangeQuery.Gte(e.MinDate)
		checked = true
	}
	if e.MaxDate != "" {
		rangeQuery.Lte(e.MaxDate)
		checked = true
	}
	if checked {
		query = query.Must(rangeQuery)
	}
	return query
}

func (e *ExtractQuery) Load(values url.Values) error {
	v, ok := values["user_id"]; if ok {
		e.UserId = v[0]
	}
	v, ok = values["barber_id"]; if ok {
		e.BarberId = v[0]
	}
	v, ok = values["min_date"]; if ok {
		e.MinDate = v[0]
	}
	v, ok = values["max_date"]; if ok {
		e.MaxDate = v[0]
	}
	if e.UserId == "" && e.BarberId == "" {
		return errors.New("A user or barber must be specified")
	}
	return nil
}