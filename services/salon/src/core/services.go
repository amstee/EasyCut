package core

import (
	"github.com/amstee/easy-cut/services/salon/src/vars"
	"github.com/amstee/easy-cut/src/es"
	"encoding/json"
	"errors"
	"github.com/amstee/easy-cut/src/logger"
)

func CreateSalon(salon *vars.Salon) error {
	resp, err := es.Index("easy_cut", "salon", salon); if err != nil {
		return err
	}
	salon.Id = resp.Id
	return nil
}

func FindSalon(salon *vars.Salon, salonId string) error {
	resp, err := es.GetById("easy_cut", "salon", salonId); if err != nil {
		return err
	}
	if resp.Found {
		if salon != nil {
			return json.Unmarshal(*resp.Source, salon)
		}
		return nil
	}
	return errors.New("salon not found")
}

func FindSalons(extract vars.ExtractQuery) (*[]vars.Salon, error) {
	var salons []vars.Salon
	query := extract.ConstructQuery()
	result, err := es.Search("easy_cut", query, "", false, -1)
	if err != nil {
		return nil, err
	}
	if result.Hits.TotalHits > 0 {
		for _, hit := range result.Hits.Hits {
			var salon vars.Salon
			err := json.Unmarshal(*hit.Source, salon)
			if err != nil {
				logger.Error.Println("unable to unmarshal salon id = ", hit.Id, ", raw data = ", hit)
			} else {
				salons = append(salons, salon)
			}
		}
		return &salons, nil
	}
	return nil, errors.New("0 results")
}