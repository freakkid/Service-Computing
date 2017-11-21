package server

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	negroniInstance := negroni.Classic()
	muxInstance := mux.NewRouter()

	initRoutes(muxInstance, formatter)
	negroniInstance.UseHandler(muxInstance)

	return negroniInstance
}

func initRoutes(muxInstance *mux.Router, formatter *render.Render) {
	muxInstance.HandleFunc("/api/test", apiTestHandler(formatter)).Methods("GET")
	muxInstance.HandleFunc("/api/unknow", apiTestHandler(formatter)).Methods("GET") // 501

	// muxInstance.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(staticDir))))
	muxInstance.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))
	muxInstance.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))
	muxInstance.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))
	muxInstance.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))
}
