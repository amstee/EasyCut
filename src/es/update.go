package es

import (
	"github.com/amstee/easy-cut/src/logger"
	"github.com/olivere/elastic"
)

func UpdateDoc(index string, itype string, id string, update interface{}) (*elastic.UpdateResponse, error) {
	resp, err := Client.Update().Index(index).Type(itype).Id(id).Doc(update).Do(Ctx); if err == nil {
		logger.Info.Printf("updated %s %s  to index %s", itype, id, index)
	} else {
		return resp, err
	}
	return resp, Flush(index)
}