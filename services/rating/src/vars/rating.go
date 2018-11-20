package vars

import (
	"github.com/pkg/errors"
	"net/url"
	"strconv"
	"github.com/olivere/elastic"
)

type UpdateRating struct {
	Comment string 		`json:"comment,omitempty"`
	Stars int 			`json:"stars,omitempty"`
	Updated string 		`json:"updated,omitempty"`
}

type Rate struct {
	Barber string 		`json:"barber,omitempty"`
	Salon string 		`json:"salon,omitempty"`
	Appointment string 	`json:"appointment,omitempty"`
	Comment string 		`json:"comment,omitempty"`
	Stars int 			`json:"stars"`
	Created string 		`json:"created,omitempty"`
	Updated string 		`json:"updated,omitempty"`
}

type Rating struct {
	Id string 			`json:"_id,omitempty"`
	Comment string 		`json:"comment,omitempty"`
	Stars int 			`json:"stars"`
	Created string 		`json:"created,omitempty"`
	Updated string 		`json:"updated,omitempty"`
	UserId string 		`json:"user_id"`
	TargetId string 	`json:"target_id"`
	TargetType string 	`json:"target_type"`
}

type ExtractQuery struct {
	MinStars int 		`json:"min_stars,omitempty"`
	MaxStars int		`json:"max_stars,omitempty"`
	TargetId string 	`json:"target_id"`
	TargetType string	`json:"target_type,omitempty"`
}

func (e *ExtractQuery) ConstructQuery() (*elastic.BoolQuery) {
	query := elastic.NewBoolQuery()
	rangeQuery := elastic.NewRangeQuery("stars")
	checked := false
	query.Must(elastic.NewTermQuery("target_id", e.TargetId))
	query.Must(elastic.NewTermQuery("target_type", e.TargetType))
	if e.MinStars >= 1 {
		rangeQuery.Gte(e.MinStars)
		checked = true
	}
	if e.MaxStars >= 1 {
		rangeQuery.Lte(e.MaxStars)
		checked = true
	}
	if checked {
		query = query.Must(rangeQuery)
	}
	return query
}

func (e *ExtractQuery) Load(values url.Values) error {
	e.MinStars = 0
	e.MaxStars = 5
	e.TargetType = "barber"
	v, ok := values["target_id"]; if ok {
		e.TargetId = v[0]
	}
	v, ok = values["target_type"]; if ok {
		e.TargetType = v[0]
	}
	v, ok = values["min_stars"]; if ok {
		i, err := strconv.Atoi(v[0]); if err != nil {
			return err
		}
		e.MinStars = i
	}
	v, ok = values["max_stars"]; if ok {
		i, err := strconv.Atoi(v[0]); if err != nil {
			return err
		}
		e.MinStars = i
	}
	return nil
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