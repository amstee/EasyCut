package es

import (
	"github.com/olivere/elastic"
	"github.com/amstee/easy-cut/src/config"
	"context"
	"github.com/amstee/easy-cut/src/logger"
)

type setter func() error

func Flush(index string) error {
	_, err := Client.Flush().Index(index).Do(Ctx)
	if err != nil {
		return err
	}
	return nil
}

func CheckESService() bool {
	if connected == false {
		InitClient(nil)
	}
	return connected
}

func GetVersion() (string, error) {
	return Client.ElasticsearchVersion(config.GetServiceURL("elasticsearch"))
}

func InitClient(callable func() error) error {
	var err error

	if callable != nil {
		call = callable
	}
	Client, err = elastic.NewClient(elastic.SetURL(config.GetServiceURL("elasticsearch")))
	if err != nil {
		logger.Error.Println("Elastic search connection failed")
		return err
	}
	logger.Error.Println("Elastic search connection success")
	connected = true
	if call == nil {
		return nil
	}
	return call()
}

var call setter = nil
var connected = false
var Ctx = context.Background()
var Client *elastic.Client = nil
