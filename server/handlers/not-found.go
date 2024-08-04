package handlers

import (
	"matcha/components"
	"matcha/components/lib"
	"net/http"

	"goji.io"
	"goji.io/pat"
)

func GetNotFound(res http.ResponseWriter, req *http.Request) {
	lib.Page("Not found", components.NotFound()).Render(req.Context(), res)
}

func initNotFound(root *goji.Mux) {
	root.HandleFunc(pat.Get("/not-found"), GetNotFound)

}
