package core

import (
	"net/http"
	"encoding/json"
	"github.com/amstee/easy-cut/services/auth/src/cache"
	"github.com/dgrijalva/jwt-go"
	"errors"
	"github.com/amstee/easy-cut/services/auth/src/types"
)

func FindKey(token *jwt.Token) (string, bool) {
	for _, key := range cache.JWKS.Keys {
		if token.Header["kid"] == key.Kid {
			cert := "-----BEGIN CERTIFICATE-----\n" + key.X5c[0] + "\n-----END CERTIFICATE-----"
			return cert, true
		}
	}
	return "", false
}

func GetCertificate(token *jwt.Token) (string, error) {
	cert, ok := FindKey(token); if !ok {
		err := SetJwks(); if err != nil {
			return cert, err
		}
		cert, ok = FindKey(token); if !ok {
			return cert, errors.New("unable to find appropriate key")
		}
	}
	return cert, nil
}

func SetJwks() error {
	resp, err := http.Get("https://easy-cut.eu.auth0.com/.well-known/jwks.json"); if err != nil {
		return err
	}
	defer resp.Body.Close()

	if cache.JWKS == nil {
		cache.JWKS = new(types.Jwks)
	}
	err = json.NewDecoder(resp.Body).Decode(cache.JWKS); if err != nil {
		return err
	}
	return nil
}

