package vars

import (
	"net/url"
)

type Search struct {
	Group string		`json:"group"`
	Nickname string 	`json:"nickname"`
	Email string 		`json:"email"`
	Username string 	`json:"username"`
}

func (search *Search) Build(vals url.Values) {
	if v, ok := vals["group"]; ok {
		search.Group = v[0]
	}
	if v, ok := vals["nickname"]; ok {
		search.Nickname = v[0]
	}
	if v, ok := vals["username"]; ok {
		search.Username = v[0]
	}
	if v, ok := vals["email"]; ok {
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
	if search.Username != "" {
		if len(s) > 3 {
			s += "%20AND%20"
			s += "user_metadata.username%3A" + search.Username
		}
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
