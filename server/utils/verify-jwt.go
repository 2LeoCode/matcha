package utils

import (
	"fmt"
	"matcha/database"
	"net/http"
	"os"
	"slices"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyJwt(req *http.Request) *database.User {
	cookies := req.Cookies()
	tokenIdx := slices.IndexFunc(cookies, func(c *http.Cookie) bool {
		return c.Valid() == nil && c.Name == "matcha-auth-token"
	})
	if tokenIdx == -1 {
		return nil
	}
	value, err := jwt.Parse(cookies[tokenIdx].Value, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SERVER_JWT_SECRET")), nil
	})
	if err != nil {
		return nil
	}
	_, err = value.Claims.GetSubject()
	if err != nil {
		return nil
	}
	// TODO: retrieve user in database
	return &database.User{}
}
