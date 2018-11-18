package core

import (
	"github.com/amstee/easy-cut/services/salon/src/vars"
	"github.com/amstee/easy-cut/src/es"
	"encoding/json"
	"errors"
	"github.com/amstee/easy-cut/src/logger"
)

func UpdateSalon(salon vars.UpdateSalon, salonId string) error {
	err := FindSalon(nil, salonId); if err != nil {
		return errors.New("this salon does not exist")
	}
	_, err = es.UpdateDoc("salon", "salon", salonId, salon); if err != nil {
		return err
	}
	return nil
}

func CreateSalon(salon *vars.Salon) error {
	resp, err := es.Index("salon", "salon", salon); if err != nil {
		return err
	}
	salon.Id = resp.Id
	return nil
}

func FindSalon(salon *vars.Salon, salonId string) error {
	resp, err := es.GetById("salon", "salon", salonId); if err != nil {
		return err
	}
	if resp.Found {
		if salon != nil {
			salon.Id = resp.Id
			return json.Unmarshal(*resp.Source, salon)
		}
		return nil
	}
	return errors.New("salon not found")
}

func FindSalons(extract vars.ExtractQuery) (*[]vars.Salon, error) {
	var salons []vars.Salon
	query := extract.ConstructQuery()
	result, err := es.Search("salon", query, "", false, -1, 50)
	if err != nil {
		return nil, err
	}
	if result.Hits.TotalHits > 0 {
		for _, hit := range result.Hits.Hits {
			var salon vars.Salon
			salon.Id = hit.Id
			err := json.Unmarshal(*(hit.Source), &salon)
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

func DeleteSalon(salonId string, check bool) error {
	if check {
		err := FindSalon(nil, salonId); if err != nil {
			return errors.New("this salon does not exist")
		}
	}
	_, err := es.DeleteDoc("salon", "salon", salonId); if err != nil {
		return err
	}
	return nil
}