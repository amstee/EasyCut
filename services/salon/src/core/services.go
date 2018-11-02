package core

import (
	"github.com/amstee/easy-cut/services/salon/src/vars"
	"github.com/amstee/easy-cut/src/es"
	"encoding/json"
	"errors"
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