package server

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		Directory:  "templates",
		Extensions: []string{".html"},
		IndentJSON: true,
	})

	negroniInstance := negroni.Classic()
	muxInstance := mux.NewRouter()

	initRoutes(muxInstance, formatter)
	negroniInstance.UseHandler(muxInstance)

	return negroniInstance
}

func initRoutes(muxInstance *mux.Router, formatter *render.Render) {

	muxInstance.HandleFunc("/api/register", apiRegisterHandler(formatter))
	muxInstance.HandleFunc("/api/add", apiAddItemHandler(formatter))
	muxInstance.HandleFunc("/api/delete", apiDeleteItemHandler(formatter))
	muxInstance.HandleFunc("/api/show", apiShowItemsHandler(formatter))

	muxInstance.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	muxInstance.HandleFunc("/register", htmlRegisterHandler(formatter))
	muxInstance.HandleFunc("/addItem", htmlAddItemHandler(formatter))
	muxInstance.HandleFunc("/deleteItem", htmlDeleteItemHandler(formatter))
	muxInstance.HandleFunc("/showItems", htmlShowItemsHandler(formatter))

	muxInstance.PathPrefix("/register").Handler(http.StripPrefix("/register/", http.FileServer(http.Dir(staticDir))))
	muxInstance.PathPrefix("/addItem").Handler(http.StripPrefix("/addItem/", http.FileServer(http.Dir(staticDir))))
	muxInstance.PathPrefix("/deleteItem").Handler(http.StripPrefix("/deleteItem/", http.FileServer(http.Dir(staticDir))))
	muxInstance.PathPrefix("/showItems").Handler(http.StripPrefix("/showItems/", http.FileServer(http.Dir(staticDir))))

	muxInstance.PathPrefix("/js").Handler(http.FileServer(http.Dir(staticDir)))
	muxInstance.PathPrefix("/css").Handler(http.FileServer(http.Dir(staticDir)))
	muxInstance.PathPrefix("/images").Handler(http.FileServer(http.Dir(staticDir)))

	muxInstance.NotFoundHandler = notImplementedHandler() // redirect 404 to 501
}

// handle /unknown  -- 501
func notImplementedHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("501 - the request method is not supported by the server and cannot be handled!"))
	}
}
