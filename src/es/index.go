package es

import "github.com/olivere/elastic"

func Index(index string, itype string, data interface{}) (*elastic.IndexResponse, error) {
	return Client.Index().Index(index).Type(itype).BodyJson(data).Do(Ctx)
}
