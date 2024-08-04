package handlers

import (
	"matcha/components"
	"matcha/components/lib"
	"net/http"

	"goji.io"
	"goji.io/pat"
)

func GetHome(res http.ResponseWriter, req *http.Request) {
	lib.Page("Home", components.Home()).Render(req.Context(), res)
}

func initHome(root *goji.Mux) {
	root.HandleFunc(pat.Get("/home"), GetHome)
}
