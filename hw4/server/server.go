package service

import (
	"net/http"
	"os"

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

	//initRou

}

func initRoutes(muxInstance *mux.Router, formatter *render.Render) {
	if root, err := os.Getwd(); err != nil {
		panic("")
	} else {

		//muxInstance.HandleFunc("")
		muxInstance.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir)))
	}
}
