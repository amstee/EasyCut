package core

import (
	"github.com/amstee/easy-cut/services/barber/src/vars"
	"github.com/amstee/easy-cut/src/es"
	"github.com/pkg/errors"
	"encoding/json"
)

func CreateBarber(barber *vars.Barber, userId string) error {
	_, err := es.IndexById("easy_cut", "barber", userId, barber); if err != nil {
		return err
	}
	return nil
}

func FindBarber(barber *vars.Barber, usedId string) error {
	resp, err := es.GetById("easy_cut", "barber", usedId); if err != nil {
		return err
	}
	if resp.Found {
		return json.Unmarshal(*resp.Source, barber)
	}
	return errors.New("barber not found")
}