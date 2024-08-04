package main

import (
	"log"
	"matcha/database"
	"matcha/handlers"
	"net/http"
)

func main() {
	// http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
	// 	http.Redirect(res, req, "/counter", http.StatusSeeOther)
	// })
	// http.HandleFunc("/counter", handlers.Counter)

	if err := database.Database.Init(); err != nil {
		log.Fatalln(err)
	}
	defer database.Database.Close()
	log.Println("Successfully connected to database")

	root := handlers.InitRoot()

	http.ListenAndServe(":80", root)
}
