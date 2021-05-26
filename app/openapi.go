package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
)

// expose generated swagger spec of the api docs
func initAPISpec(router *httprouter.Router) {
	path := "/tmp/docs"
	fileServer := http.FileServer(http.Dir(path))

	router.GET("/docs/*filepath", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		filepath := p.ByName("filepath")
		if filepath != "/swagger.json" && filepath != "/swagger.yaml" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		r.URL.Path = filepath
		fileServer.ServeHTTP(w, r)
	})
}

// create endpoint for swagger UI
// will fetch generated swagger spec
func initSwagger(router *httprouter.Router) {
	urlConfig := httpSwagger.URL("http://localhost:8080/docs/swagger.json")
	swagHandler := httpSwagger.Handler(urlConfig)
	router.GET("/swagger/*filepath", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		swagHandler(w, r)
	})
}
