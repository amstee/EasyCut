package core

import (
	"github.com/amstee/easy-cut/services/barber/src/vars"
	"github.com/amstee/easy-cut/src/es"
	"github.com/pkg/errors"
	"encoding/json"
	"github.com/amstee/easy-cut/src/logger"
	"github.com/olivere/elastic"
)

func CreateBarber(barber *vars.Barber, userId string) error {
	err := FindBarber(nil, userId); if err == nil {
		return errors.New("this barber already exist")
	}
	_, err = es.IndexById("easy_cut", "barber", userId, barber); if err != nil {
		return err
	}
	return nil
}

func FindBarber(barber *vars.Barber, usedId string) error {
	resp, err := es.GetById("easy_cut", "barber", usedId); if err != nil {
		return err
	}
	logger.Info.Println(resp)
	if resp.Found {
		if barber != nil {
			return json.Unmarshal(*resp.Source, barber)
		}
		return nil
	}
	return errors.New("barber not found")
}

func FindBarbers(barbers map[string]*vars.BarberResponse) error {
	var users []string

	for k := range barbers {
		users = append(users, k)
	}
	query := elastic.NewTermQuery("_id", users)
	result, err := es.Search("easy_cut", query, "", false, -1)
	if err != nil {
		return err
	}
	if result.Hits.TotalHits > 0 {
		for _, hit := range result.Hits.Hits {
			err := json.Unmarshal(*hit.Source, &barbers[hit.Id].Barber)
			if err != nil {
				logger.Error.Println("unable to unmarshal barber id = ", hit.Id, ", raw data = ", hit)
			}
		}
	}
	return errors.New("no barber found")
}