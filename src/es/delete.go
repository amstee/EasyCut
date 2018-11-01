package es

import (
	"github.com/olivere/elastic"
	"github.com/amstee/easy-cut/src/logger"
)

func DeleteDoc(index string, itype string, id string) (*elastic.DeleteResponse, error) {
	resp, err := Client.Delete().Index(index).Type(itype).Id(id).Do(Ctx); if err == nil {
		logger.Info.Printf("deleted %s %s  to index %s", itype, id, index)
	} else {
		return resp, err
	}
	return resp, nil
}