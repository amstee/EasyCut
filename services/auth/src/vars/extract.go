package vars

import (
	"net/url"
)

type ExtractContent struct {
	User bool 			`json:"user"`
}

func (e *ExtractContent) Load(values url.Values) {
	_, ok := values["user"]
	e.User = ok
}