package src

import (
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"errors"
	"fmt"
	"strings"
	"github.com/amstee/easy-cut/services/auth/src/types"
	"github.com/amstee/easy-cut/services/auth/src/utils"
	"net/http"
	"bytes"
	"encoding/json"
)

func GetJwtMiddleware() (* jwtmiddleware.JWTMiddleware, error) {
	middleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: CheckTokenValidity,
		SigningMethod: jwt.SigningMethodRS256,
	})
	return middleware, nil
}

func CheckTokenValidity(token *jwt.Token) (interface{}, error) {
	aud := "https://easy-cut.eu.auth0.com/api/v2/"
	iss := "https://easy-cut.eu.auth0.com/"

	if token.Header["alg"] != jwt.SigningMethodRS256.Alg() {
		return token, errors.New("invalid signature")
	}
	checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false); if !checkAud {
		return token, errors.New("invalid audience")
	}
	checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false); if !checkIss {
		return token, errors.New("invalid issuer")
	}

	cert, err := utils.GetCertificate(token)
	if err != nil {
		return nil, err
	}
	result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
	return result, nil
}

func CheckScope(scope string, tokenString string) (bool) {
	token, err := jwt.ParseWithClaims(tokenString, &types.PermissionClaims{}, func (token *jwt.Token) (interface{}, error) {
		cert, err := utils.GetCertificate(token)
		if err != nil {
			return nil, err
		}
		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
	}); if err != nil {
		fmt.Println("invalid token string for scope checking")
		return false
	}
	claims, ok := token.Claims.(*types.PermissionClaims); if ok && token.Valid {
		result := strings.Split(claims.Scope, " ")
		for _, s := range result {
			if s == scope {
				return true
			}
		}
	}
	return false
}

func GetUserGroups(tokenInfo *types.TokenInfo) ([]string, error) {
	var userGroup types.UserGroups
	jsonStr, err := json.Marshal(tokenInfo); if err != nil {
		return nil, err
	}
	resp, err := http.Post("https://easy-cut.eu.auth0.com/tokeninfo", "application/json", bytes.NewBuffer(jsonStr)); if err != nil {
		return nil, err
	}
	fmt.Println(resp.StatusCode)
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&userGroup); if err != nil {
		fmt.Println(resp.Body)
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println(userGroup)
	return userGroup.AppMetadata.Authorization.Groups, nil
}

func CheckGroups(groups []string, tokenString string) (*types.GroupsResponse, error) {
	var isInGroup bool
	tokenInfo := types.TokenInfo{Token: tokenString}
	resp := types.GroupsResponse{Groups: make(map[string]bool)}

	userGroups, err := GetUserGroups(&tokenInfo); if err != nil {
		return nil, err
	}
	for _, group := range groups {
		isInGroup = false
		for _, userGroup := range userGroups {
			if group == userGroup {
				isInGroup = true
			}
		}
		resp.Groups[group] = isInGroup
	}
	return &resp, nil
}
