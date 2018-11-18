package core

import (
	"github.com/amstee/easy-cut/services/barber/src/vars"
	"github.com/amstee/easy-cut/src/es"
	"github.com/pkg/errors"
	"encoding/json"
	"github.com/amstee/easy-cut/src/logger"
	"github.com/olivere/elastic"
)

func DeleteBarber(id string) error {
	err := FindBarber(nil, id); if err != nil {
		return errors.New("this barber does not exist")
	}
	_, err = es.DeleteDoc("barber", "barber", id); if err != nil {
		return err
	}
	return nil
}

func UpdateBarber(barber *vars.Barber, id string) error {
	err := FindBarber(nil, id); if err != nil {
		return errors.New("this barber does not exist")
	}
	_, err = es.UpdateDoc("barber", "barber", id, barber); if err != nil {
		return err
	}
	return nil
}

func CreateBarber(barber *vars.Barber, userId string) error {
	err := FindBarber(nil, userId); if err == nil {
		return errors.New("this barber already exist")
	}
	_, err = es.IndexById("barber", "barber", userId, barber); if err != nil {
		return err
	}
	return nil
}

func FindBarber(barber *vars.Barber, usedId string) error {
	resp, err := es.GetById("barber", "barber", usedId); if err != nil {
		return err
	}
	if resp.Found {
		if barber != nil {
			return json.Unmarshal(*resp.Source, barber)
		}
		return nil
	}
	return errors.New("barber not found")
}

func FindBarbers(barbers map[string]*vars.BarberResponse) error {
	var query elastic.Query
	var users []string

	for k := range barbers {
		users = append(users, k)
	}
	if len(users) == 1 {
		query = elastic.NewTermQuery("_id", users[0])
	} else {
		query = elastic.NewTermsQuery("_id", users)
	}
	result, err := es.Search("barber", query, "", false, -1, 200)
	if err != nil {
		return err
	}
	if result.Hits.TotalHits > 0 {
		for _, hit := range result.Hits.Hits {
			if barbers[hit.Id] != nil {
				err := json.Unmarshal(*hit.Source, &barbers[hit.Id].Barber)
				if err != nil {
					logger.Error.Println("unable to unmarshal barber id = ", hit.Id, ", raw data = ", hit)
				}
			} else {
				return errors.New("results incoherent")
			}
		}
		return nil
	}
	return errors.New("no barber found")
}