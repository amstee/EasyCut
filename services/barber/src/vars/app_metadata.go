package vars

type Authorization struct {
	Groups []string `json:"groups"`
}

type AppMetadata struct {
	Authorization Authorization `json:"authorization"`
}

func GetBarberCreation() AppMetadata {
	return AppMetadata{Authorization: Authorization{Groups: []string{"Barber"}}}
}