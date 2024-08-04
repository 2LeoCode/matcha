package handlers

import (
	"matcha/middlewares"
	"net/http"

	"goji.io"
	"goji.io/pat"
)

func InitRoot() *goji.Mux {
	root := goji.NewMux()
	defer root.HandleFunc(pat.Get("/*"), func(res http.ResponseWriter, req *http.Request) {
		http.Redirect(res, req, "/not-found", http.StatusNotFound)
	})

	root.Use(middlewares.Auth)
	initAuth(root)
	initHome(root)
	initNotFound(root)
	return root
}
