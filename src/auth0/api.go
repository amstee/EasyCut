package auth0

import (
	"github.com/amstee/easy-cut/src/config"
	"github.com/amstee/easy-cut/src/types"
	"net/http"
	"github.com/amstee/easy-cut/src/common"
	"fmt"
	"encoding/json"
	"errors"
)

func NewOauth() types.OauthToken {
	return types.OauthToken{
		GrantType: "client_credentials",
		Audience: config.Content.Domain + config.Content.Api,
		ClientSecret: config.Content.ClientSecret,
		ClientId: config.Content.ClientId,
	}
}

func LoadToken() (error) {
	url := config.Content.Domain + config.Content.Oauth
	data, err := common.JsonToReader(NewOauth()); if err != nil {
		fmt.Println(err)
		return err
	}
	req, err := http.NewRequest("POST", url, data); if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("content-type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

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
