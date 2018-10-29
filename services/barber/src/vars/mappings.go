package vars

import "github.com/amstee/easy-cut/src/es"

const BarberMapping = `
	{
		"settings": {
			"number_of_shards": 2,
			"number_of_replicas": 1
		},
		"mappings": {
			"barber": {
				"properties": {
					"user_id": {
						"type": "keyword"
					},
					"experience": {
						"type": "keyword"
					},
					"Style": {
						"type": "keyword"
					},
					"created": {
						"type": "date"
					},
					"updated": {
						"type": "date"
					}
				}
			}
		}
	}
`

func Register() error {
	return es.RegisterIndex("EasyCut", BarberMapping)
}
