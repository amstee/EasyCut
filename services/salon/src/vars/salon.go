package vars

import (
	"net/url"
	"github.com/olivere/elastic"
	"strconv"
)

type Salon struct {
	Id string 			`json:"_id,omitempty"`
	UserId string 		`json:"user_id"`
	Name string 		`json:"name"`
	Address string 		`json:"address,omitempty"`
	EmployeeNumber int 	`json:"employee_number,omitempty"`
	Barber []string 	`json:"barbers,omitempty"`
	Website string 		`json:"website,omitempty"`
	Created string 		`json:"created,omitempty"`
	Updated string 		`json:"updated,omitempty"`
}

type ExtractQuery struct {
	Address string 			`json:"address,omitempty"`
	ExactAddress bool 		`json:"exact,omitempty"`
	EmployeeMinNumber int 	`json:"min_emp,omitempty"`
	EmployeeMaxNumber int 	`json:"max_emp,omitempty"`
}

func (e *ExtractQuery) ConstructQuery() (*elastic.BoolQuery) {
	query := elastic.NewBoolQuery()
	rangeQuery := elastic.NewRangeQuery("employee_number")
	checked := false

	if e.Address != "" {
		if e.ExactAddress {
			query = query.Must(elastic.NewTermQuery("address", e.Address))
		} else {
			query = query.Must(elastic.NewWildcardQuery("address", e.Address))
		}
	}
	if e.EmployeeMinNumber >= 1 {
		rangeQuery.Gte(e.EmployeeMinNumber)
		checked = true
	}
	if e.EmployeeMaxNumber >= 1 {
		rangeQuery.Lte(e.EmployeeMaxNumber)
		checked = true
	}
	if checked {
		query = query.Must(rangeQuery)
	}
	return query
}

func (e *ExtractQuery) Load(values url.Values) error {
	v, ok := values["address"]; if ok {
		e.Address = v[0]
	}
	v, ok = values["max_emp"]; if ok {
		i, err := strconv.Atoi(v[0]); if err != nil {
			return err
		}
		e.EmployeeMaxNumber = i
	}
	v, ok = values["min_emp"]; if ok {
		i, err := strconv.Atoi(v[0]); if err != nil {
			return err
		}
		e.EmployeeMinNumber = i
	}
	v, ok = values["exact"]; if ok {
		e.ExactAddress = v[0] == "true"
	}
	return nil
}