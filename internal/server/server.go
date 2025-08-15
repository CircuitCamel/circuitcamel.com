package server

import (
	"fmt"
	"log"
	"net/http"

	"circuitcamel.com/internal/cache"
	"circuitcamel.com/internal/models"
	"circuitcamel.com/internal/site/blog"
	"circuitcamel.com/internal/site/index"
	"circuitcamel.com/internal/site/now"
	"circuitcamel.com/internal/site/teapot"
	"circuitcamel.com/internal/utils"
	"github.com/gorilla/mux"
)

func StartServer(conf models.Config) {
	go cache.LoadAll()

	server := mux.NewRouter().StrictSlash(true)

	server.HandleFunc("/", index.Index)
	server.HandleFunc("/now", now.Now)
	server.HandleFunc("/teapot", teapot.Teapot)
	server.HandleFunc("/blog", blog.BlogList)
	server.HandleFunc("/blog/{slug}", blog.Blog)

	server.PathPrefix("/static").Handler(
		http.StripPrefix("/static", http.FileServer(http.Dir("./static"))),
	)
	server.PathPrefix("/content").Handler(
		http.StripPrefix("/content", http.FileServer(http.Dir("./content"))),
	)
	server.NotFoundHandler = http.HandlerFunc(notfound)

	fmt.Printf("Server running on port: %s", conf.PORT)
	if conf.ENV == "production" {
		http.ListenAndServeTLS(":"+conf.PORT, conf.CRT, conf.KEY, server)
	} else if conf.ENV == "staging" {
		log.Fatal(http.ListenAndServeTLS(":"+conf.PORT, conf.CRT, conf.KEY, server))
	} else {
		log.Fatal(http.ListenAndServe(":"+conf.PORT, server))
	}
}

func notfound(w http.ResponseWriter, r *http.Request) {
	utils.ErrPage(w, r, 404)
}
