package middlewares

import (
	"context"
	"matcha/utils"
	"net/http"
	"regexp"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		user := utils.VerifyJwt(req)
		urlPath := req.URL.EscapedPath()
		isAuthPage := regexp.MustCompile(`^(?:/auth)|(?:/auth/.*)$`).MatchString(urlPath)
		isNotFound := urlPath == "/not-found"
		switch {
		case isNotFound:
			fallthrough
		default:
			next.ServeHTTP(res, req.WithContext(context.WithValue(req.Context(), "session-user", user)))
		case isAuthPage && user != nil:
			http.Redirect(res, req, "/home", http.StatusSeeOther)
		case !isAuthPage && user == nil:
			http.Redirect(res, req, "/auth", http.StatusNetworkAuthenticationRequired)
		}
	})
}
