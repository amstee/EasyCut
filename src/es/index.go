package es

import (
	"github.com/olivere/elastic"
	"fmt"
	"github.com/amstee/easy-cut/src/logger"
	"github.com/pkg/errors"
)

func Index(index string, itype string, data interface{}) (*elastic.IndexResponse, error) {
	if up := CheckESService(); !up {
		return nil, errors.New("can't connect to database")
	}
	resp, err := Client.Index().Index(index).Type(itype).BodyJson(data).Do(Ctx); if err == nil {
		logger.Info.Printf("Indexed %s %s to index %s, type %s\n", itype, resp.Id, resp.Index, resp.Type)
	} else {
		return resp, err
	}
	return resp, Flush(index)
}

func IndexById(index string, itype string, id string, data interface{}) (*elastic.IndexResponse, error) {
	if up := CheckESService(); !up {
		return nil, errors.New("can't connect to database")
	}
	resp, err := Client.Index().Index(index).Type(itype).Id(id).BodyJson(data).Do(Ctx); if err == nil {
		logger.Info.Printf("Indexed %s %s to index %s, type %s\n", itype, resp.Id, resp.Index, resp.Type)
	} else {
		return resp, err
	}
	return resp, Flush(index)
}

func IndexExist(name string) (bool, error) {
	if up := CheckESService(); !up {
		return false, errors.New("can't connect to database")
	}
	return Client.IndexExists(name).Do(Ctx)
}

func RegisterIndex(name string, mapping string) error {
	if up := CheckESService(); !up {
		return errors.New("can't connect to database")
	}
	exist, err := IndexExist(name); if err != nil {
		return err
	}
	if exist {
		return nil
	}
	createIndex, err := Client.CreateIndex(name).BodyString(mapping).Do(Ctx)
	if err != nil {
		return err
	}
	if !createIndex.Acknowledged {
		fmt.Println("index not acknowledged")
	}
	return nil
}
