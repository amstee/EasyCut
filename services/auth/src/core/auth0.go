package core

import (
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"errors"
	"fmt"
	"strings"
	"github.com/amstee/easy-cut/services/auth/src/vars"
	"net/http"
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

	cert, err := GetCertificate(token)
	if err != nil {
		return nil, err
	}
	result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
	return result, nil
}

func GetPermissionClaims(tokenString string) (*vars.PermissionClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &vars.PermissionClaims{}, func (token *jwt.Token) (interface{}, error) {
		cert, err := GetCertificate(token)
		if err != nil {
			return nil, err
		}
		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
	}); if err != nil {
		return nil, err
	}
	if token.Valid {
		claims, ok := token.Claims.(*vars.PermissionClaims)
		if ok {
			return claims, nil
		}
		return nil, errors.New("unable to extract claims")
	}
	return nil, errors.New("invalid token")
}

func CheckScope(scope string, tokenString string) (bool) {
	claims, err := GetPermissionClaims(tokenString); if err != nil {
		fmt.Println(err)
		return false
	}
	result := strings.Split(claims.Scope, " ")
	for _, s := range result {
		if s == scope {
			return true
		}
	}
	return false
}

func GetUser(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		cert, err := GetCertificate(token); if err != nil {
			return nil, err
		}
		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
	}); if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*jwt.StandardClaims); if ok && token.Valid {
		return claims.Subject, nil
	}
	return "", errors.New("unable to extract claims")
}

func GetUserGroups(tokenInfo *vars.TokenInfo) ([]string, error) {
	var userGroup vars.UserGroups
	userId, err := GetUser(tokenInfo.Token); if err != nil {
		return nil, err
	}
	url := "https://easy-cut.eu.auth0.com/api/v2/users/" + userId
	req, err := http.NewRequest("GET", url, nil); if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer " + tokenInfo.Token)
	client := &http.Client{}
	resp, err := client.Do(req); if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&userGroup); if err != nil {
		return nil, err
	}
	return userGroup.AppMetadata.Authorization.Groups, nil
}

func CheckGroups(groups []string, tokenString string) (*vars.GroupsResponse, error) {
	var isInGroup bool
	tokenInfo := vars.TokenInfo{Token: tokenString}
	resp := vars.GroupsResponse{Groups: make(map[string]bool)}

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
