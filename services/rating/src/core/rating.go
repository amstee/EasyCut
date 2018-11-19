package core

import (
	"github.com/amstee/easy-cut/services/rating/src/vars"
	"github.com/amstee/easy-cut/src/es"
	"errors"
	"encoding/json"
)

func CreateRating(rate *vars.Rating) error {
	resp, err := es.Index("rating", "rating", rate); if err != nil {
		return err
	}
	rate.Id = resp.Id
	return nil
}

func FindRating(rate *vars.Rating, ratingId string) error {
	resp, err := es.GetById("rating", "rating", ratingId); if err != nil {
		return err
	}
	if resp.Found {
		if rate != nil {
			rate.Id = resp.Id
			return json.Unmarshal(*resp.Source, rate)
		}
		return nil
	}
	return errors.New("rating not found")
}

func FindRatings(extract vars.ExtractQuery) (*[]vars.Rating, error) {
	return nil, nil
}

func UpdateRating(rate vars.Rating, ratingId string) error {
	err := FindRating(nil, ratingId); if err != nil {
		return errors.New("this rating does not exist")
	}
	_, err = es.UpdateDoc("rating", "rating", ratingId, rate); if err != nil {
		return err
	}
	return nil
}

func DeleteRating(ratingId string, check bool) error {
	if check {
		err := FindRating(nil, ratingId); if err != nil {
			return errors.New("this rating does not exist")
		}
	}
	_, err := es.DeleteDoc("rating", "rating", ratingId); if err != nil {
		return err
	}
	return nil
}