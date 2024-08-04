package auth

import (
	"matcha/components/auth"
	"matcha/components/lib"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GetLogin(res http.ResponseWriter, req *http.Request) {
	lib.Page("Log-in", auth.Login(false)).Render(req.Context(), res)
}

func PostLogin(res http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		// Handle error
	} else {
		username := req.Form.Get("username")
		password := req.Form.Get("password")

		// TODO: check if user is in database

		// for now assume that the user is found and password matches
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": username,
		})
		if tokenStr, err := token.SignedString(os.Getenv("SERVER_JWT_SECRET")); err != nil {
			// handle error
		} else {
			http.SetCookie(res, &http.Cookie{
				Name:  "matcha-auth-token",
				Value: tokenStr,
			})
			http.Redirect(res, req, "/home", http.StatusSeeOther)
		}
	}
}
