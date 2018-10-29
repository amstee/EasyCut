package es

import "github.com/olivere/elastic"

func GetById(index string, itype string, id string) (*elastic.GetResult, error) {
	return Client.Get().Index(index).Type(itype).Id(id).Do(Ctx)
}

func Search(index string, query elastic.Query, sortField string, asc bool, rang int) (*elastic.SearchResult, error) {
	return Client.Search().Index(index).Query(query).Sort(sortField, asc).From(0).Size(rang).Do(Ctx)
}