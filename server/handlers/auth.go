package handlers

import (
	"net/http"

	"matcha/handlers/auth"

	"goji.io"
	"goji.io/pat"
)

func GetAuth(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, "/auth/login", http.StatusSeeOther)
}

func initAuth(root *goji.Mux) {
	mux := goji.SubMux()

	root.Handle(pat.New("/auth/*"), mux)
	root.HandleFunc(pat.Get("/auth"), func(res http.ResponseWriter, req *http.Request) {
		http.Redirect(res, req, "/auth/login", http.StatusSeeOther)
	})

	mux.HandleFunc(pat.Get("/login"), auth.GetLogin)
	mux.HandleFunc(pat.Post("/login"), auth.PostLogin)
	mux.HandleFunc(pat.Get("/signin"), auth.GetSignin)
	mux.HandleFunc(pat.Post("/signin"), auth.PostSignin)
}
