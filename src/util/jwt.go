package util

import (
	"net/http"
	"strings"
	"errors"
	"entity"

	jwt "github.com/dgrijalva/jwt-go"
)

// FromAuthHeader is a "TokenExtractor" that takes a give request and extracts
// the JWT token from the Authorization header.
func FromAuthHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", nil // No error, just no token
	}

	// TODO: Make this a bit more robust, parsing-wise
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", errors.New("Authorization header format must be Bearer {token}")
	}

	return authHeaderParts[1], nil
}

func ParseTokenWithClaims(r *http.Request) (*entity.TokenClaims, error){
	tokenString, err := FromAuthHeader(r)
	if err != nil || tokenString == "" {
		return nil, err
	}
	token, err := jwt.ParseWithClaims(tokenString, &entity.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(Secret), nil
    })
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*entity.TokenClaims)
    if ok && token.Valid {
        return claims, nil
    } 
	
    return nil, err
}

func GetClaimsFromRequest(r *http.Request) jwt.MapClaims{
	user := r.Context().Value("user")
	return user.(*jwt.Token).Claims.(jwt.MapClaims)
}

