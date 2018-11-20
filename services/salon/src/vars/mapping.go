package vars

import "github.com/amstee/easy-cut/src/es"

const SalonMapping = `
	{
		"settings": {
			"number_of_shards": 2,
			"number_of_replicas": 1
		},
		"mappings": {
			"salon": {
				"properties": {
					"user_id": {
						"type": "keyword"
					},
					"name": {
						"type": "keyword"
					},
					"address": {
						"type": "keyword"
					},
					"employee_number": {
						"type": "integer"
					},
					"barbers": {
						"type": "array"
					},
					"website": {
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
	return es.RegisterIndex("salon", SalonMapping)
}