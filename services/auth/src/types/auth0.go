package types

import "github.com/dgrijalva/jwt-go"

type UserGroups struct {
	AppMetadata struct{
		Authorization struct {
			Groups []string	`json:"groups"`
		}	`json:"authorization"`
	}	`json:"app_metadata"`
}

type GroupClaims struct {
	Group string `json:"group"`
	jwt.StandardClaims
}

type PermissionClaims struct {
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
