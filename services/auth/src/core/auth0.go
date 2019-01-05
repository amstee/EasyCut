package core

import (
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"errors"
	"strings"
	"github.com/amstee/easy-cut/services/auth/src/vars"
	"net/http"
	"encoding/json"
	"github.com/amstee/easy-cut/services/auth/src/config"
	"github.com/amstee/easy-cut/src/logger"
)

func GetJwtMiddleware() (* jwtmiddleware.JWTMiddleware, error) {
	middleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: CheckTokenValidity,
		SigningMethod: jwt.SigningMethodRS256,
	})
	return middleware, nil
}

func CheckTokenValidity(token *jwt.Token) (interface{}, error) {
	b := false

	if token.Header["alg"] != jwt.SigningMethodRS256.Alg() {
		return token, errors.New("invalid signature")
	}
	for _, elem := range config.Content.Audiences {
		checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(elem, false); if checkAud {
			b = true
		}
	}
	if !b {
		return token, errors.New("invalid audience")
	}
	checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(config.Content.Issuer, false); if !checkIss {
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
		logger.Error.Println(err)
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

func GetClaims(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		cert, err := GetCertificate(token); if err != nil {
			return nil, err
		}
		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
	}); if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*jwt.StandardClaims); if ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("unable to extract claims")
}

func GetUserGroups(tokenInfo *vars.TokenInfo) ([]string, error) {
	var userGroups []vars.UserGroup
	var res []string

	claims, err := GetClaims(tokenInfo.Token); if err != nil {
		return nil, err
	}
	if config.Content.TPrefix == claims.Subject[:6] {
		url := config.Content.Perms + "/get/" + claims.Subject[6:]

		req, err := http.NewRequest("GET", url, nil); if err != nil {
			return nil, err
		}
		req.Header.Add("Authorization", "Bearer " + tokenInfo.Token)
		client := &http.Client{}
		resp, err := client.Do(req); if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&userGroups); if err != nil {
			return nil, err
		}
		for _, group := range userGroups {
			res = append(res, group.Name)
		}
		return res, nil
	} else {
		if claims.Issuer == config.Content.Issuer {
			return []string{"Service"}, nil
		} else {
			return nil, errors.New("token issuer unknown")
		}
	}
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
			if group == userGroup || userGroup == "Service" { // Service bypass every groups
				isInGroup = true
			}
		}
		resp.Groups[group] = isInGroup
	}
	return &resp, nil
}
