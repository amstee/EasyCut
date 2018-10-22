package vars

import (
	"net/url"
	"fmt"
)

type Search struct {
	Group string		`json:"group"`
	Nickname string 	`json:"nickname"`
	Email string 		`json:"email"`
}

func (search *Search) Build(vals url.Values) {
	if v, ok := vals["group"]; ok {
		fmt.Println(v)
		search.Group = v[0]
	}
	if v, ok := vals["nickname"]; ok {
		fmt.Println(v)
		search.Nickname = v[0]
	}
	if v, ok := vals["email"]; ok {
		fmt.Println(v)
		search.Email = v[0]
	}
}

func (search *Search) GetSearch() (string) {
	s := "?q="

	if search.Group != "" {
		s += "app_metadata.authorization.groups%3A%22" + search.Group + "%22"
	}
	if search.Nickname != "" {
		if len(s) > 3 {
			s += "%20AND%20"
		}
		s += "nickname%3A" + search.Nickname
	}
	if search.Email != "" {
		if len(s) > 3 {
			s += "%20AND%20"
		}
		s +="email%3A" + search.Email
	}
	s += "&search_engine=v3"
	return s
}
