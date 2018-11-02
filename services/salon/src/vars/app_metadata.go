package vars

type Authorization struct {
	Groups []string `json:"groups"`
}

type AppMetadata struct {
	Authorization Authorization `json:"authorization"`
}

func GetSalonRole() AppMetadata {
	return AppMetadata{Authorization: Authorization{Groups: []string{"Salon"}}}
}