package src

import (
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type CustomClaims struct {
	Scope string `json:"scope"`
	jwt.StandardClaims
}

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	Use string `json:"use"`
	N string `json:"n"`
	E string `json:"e"`
	X5c []string `json:"x5c"`
}

func getPermCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get("https://easy-cut.eu.auth0.com/.well-known/jwks.json"); if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks); if err != nil {
		return cert, err
	}

	for _, key := range jwks.Keys {
		if token.Header["kid"] == key.Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + key.X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}
	if cert == "" {
		err := errors.New("unable to find appropriate key")
		return cert, err
	}
	return cert, nil
}

func GetJwtMiddleware() (* jwtmiddleware.JWTMiddleware, error) {
	middleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: CheckTokenValidity,
		SigningMethod: jwt.SigningMethodES256,
	})
	return middleware, nil
}

func CheckTokenValidity(token *jwt.Token) (interface{}, error) {
	aud := "https://easy-cut.eu.auth0.com/api/v2/"
	checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
	if !checkAud {
		return token, errors.New("invalid issuer")
	}
	cert, err := getPermCert(token)
	if err != nil {
		return nil, err
	}
	result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
	return result, nil
}

func CheckScope(scope string, tokenString string) (bool) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, nil); if err != nil {
		fmt.Println("invalid token string for scope checking")
		return false
	}
	claims, _ := token.Claims.(*CustomClaims)
	result := strings.Split(claims.Scope, " ")
	for _, s := range result {
		if s == scope {
			return true
		}
	}
	return false
}

