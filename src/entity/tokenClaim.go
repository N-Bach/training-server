package entity

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type TokenClaims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}