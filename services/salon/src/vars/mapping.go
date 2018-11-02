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
						"type": "string"
					},
					"address": {
						"type": "string"
					},
					"employee_number": {
						"type": "integer"
					},
					"barbers": {
						"type": "array"
					},
					"website": {
						"type": "string"
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
	return es.RegisterIndex("easy_cut", SalonMapping)
}