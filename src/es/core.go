package es

import (
	"github.com/olivere/elastic"
	"github.com/amstee/easy-cut/src/config"
	"context"
	"fmt"
)

func IndexExist(name string) (bool, error) {
	return Client.IndexExists(name).Do(Ctx)
}

func RegisterIndex(name string, mapping string) error {
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

func GetVersion() (string, error) {
	return Client.ElasticsearchVersion(config.GetServiceURL("elasticsearch"))
}

func InitClient() error {
	var err error

	Client, err = elastic.NewClient(elastic.SetURL(config.GetServiceURL("elasticsearch")))
	if err != nil {
		return err
	}
	return nil
}

var Ctx = context.Background()
var Client *elastic.Client = nil
