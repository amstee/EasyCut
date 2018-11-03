package es

import (
	"github.com/olivere/elastic"
	"github.com/pkg/errors"
)

func GetById(index string, itype string, id string) (*elastic.GetResult, error) {
	if up := CheckESService(); !up {
		return nil, errors.New("can't connect to database")
	}
	return Client.Get().Index(index).Type(itype).Id(id).Do(Ctx)
}

func Search(index string, query elastic.Query, sortField string, asc bool, rang int) (*elastic.SearchResult, error) {
	if up := CheckESService(); !up {
		return nil, errors.New("can't connect to database")
	}
	search := Client.Search().Index(index).Query(query)
	if sortField != "" {
		search.Sort(sortField, asc)
	}
	if rang != -1 {
		search.From(0).Size(rang)
	}
	return search.Do(Ctx)
}

