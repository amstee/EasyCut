package vars

import "github.com/amstee/easy-cut/src/es"

const AppointmentMapping = `
	{
		"settings": {
			"number_of_shards": 2,
			"number_of_replicas": 1
		},
		"mappings": {
			"appointment": {
				"properties": {
					"user_id": {
						"type": "keyword"
					},
					"barber_id": {
						"type": "keyword"
					},
					"date": {
						"type": "date"
					},
					"description": {
						"type": "keyword"
					},
					"duration": {
						"type": "integer"
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
	return es.RegisterIndex("appointment", AppointmentMapping)
}