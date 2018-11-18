package core

import (
	"github.com/amstee/easy-cut/services/perms/src/vars"
	"github.com/amstee/easy-cut/src/request"
	"github.com/amstee/easy-cut/src/config"
	"github.com/amstee/easy-cut/src/auth0"
	"github.com/amstee/easy-cut/src/logger"
)

func GetUserGroups(data *[]vars.Group, userId string) error {
	token, err := auth0.GetToken(); if err != nil {
		return err
	}
	resp, err := request.ExpectJson(config.GetApi() + "users/" + config.Content.Api.Tprefix + userId + "/groups",
								"GET", token.Format(), nil, data)
	if err != nil {
		logger.Error.Printf("Failed to get groups %s status = %d", err.Error(), resp.StatusCode)
		return err
	}
	return nil
}

func UpdateGroup(userId string, group vars.Group, b bool) error {
	t := "PATCH"

	token, err := auth0.GetToken(); if err != nil {
		return err
	}
	if b {
		t = "DELETE"
	}
	resp, err := request.ExpectJson(config.GetApi() + "groups/" + group.Id + "/members",
									t, token.Format(), []string{config.Content.Api.Tprefix + userId}, nil)
	if err != nil {
		logger.Error.Printf("Failed to update group %s status = %d", err.Error(), resp.StatusCode)
		return err
	}
	return nil
}

func RetrieveGroups() (*vars.Groups ,error) {
	var groups vars.Groups

	token, err := auth0.GetToken(); if err != nil {
		return nil, err
	}
	resp, err := request.ExpectJson(config.GetApi() + "groups", "GET", token.Format(), nil, &groups)
	if err != nil {
		logger.Error.Printf("Failed to update group %s status = %d", err.Error(), resp.StatusCode)
		return nil, err
	}
	return &groups, nil
}