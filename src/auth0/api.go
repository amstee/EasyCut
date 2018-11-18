package auth0

import (
	"github.com/amstee/easy-cut/src/config"
	"github.com/amstee/easy-cut/src/types"
	"net/http"
	"github.com/amstee/easy-cut/src/common"
	"fmt"
	"encoding/json"
	"errors"
	"github.com/amstee/easy-cut/src/request"
)

func NewOauth() types.OauthToken {
	return types.OauthToken{
		GrantType: "client_credentials",
		Audience: config.Content.Oauth.Audience,
		ClientSecret: config.Content.ClientSecret,
		ClientId: config.Content.ClientId,
	}
}

func LoadToken() (error) {
	url := config.GetOauth()
	data, err := common.JsonToReader(NewOauth()); if err != nil {
		fmt.Println(err)
		return err
	}
	req, err := http.NewRequest("POST", url, data); if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("content-type", "application/json")
	request.DisplayRequest(req)
	res, _ := http.DefaultClient.Do(req)
	if res == nil {
		return errors.New("unable to execute http request, check your internet connection")
	}
	defer res.Body.Close()

	request.DisplayResponse(res)
	if res.StatusCode != 200 {
		return errors.New("unable to retrieve api token")
	}
	if err := json.NewDecoder(res.Body).Decode(token); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetToken() (*types.TokenType, error) {
	if token == nil {
		if err := LoadToken(); err != nil {
			return nil, err
		}
	}
	return token, nil
}

var token = new(types.TokenType)
