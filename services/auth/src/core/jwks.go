package core

import (
	"net/http"
	"encoding/json"
	"github.com/amstee/easy-cut/services/auth/src/cache"
	"github.com/dgrijalva/jwt-go"
	"errors"
	"github.com/amstee/easy-cut/services/auth/src/vars"
	"github.com/amstee/easy-cut/services/auth/src/config"
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
	resp, err := http.Get(config.Content.Jwks); if err != nil {
		return err
	}
	defer resp.Body.Close()

	if cache.JWKS == nil {
		cache.JWKS = new(vars.Jwks)
	}
	err = json.NewDecoder(resp.Body).Decode(cache.JWKS); if err != nil {
		return err
	}
	return nil
}

