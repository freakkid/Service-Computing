package server

import (
	"net/http"

	"github.com/unrolled/render"
)

/* Server provides some api to client
 * The correct URLs as follows:
 * /register
 * /addItem
 * /deleteItem
 * /showItems
 */

// handle "api/register"
// correct URL: /register
// return infomation of registering fail/successfully
func htmlRegisterHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		formatter.HTML(w, http.StatusOK, "register", nil)
	}
}

func htmlAddItemHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		formatter.HTML(w, http.StatusOK, "add", nil)
	}
}

func htmlDeleteItemHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		formatter.HTML(w, http.StatusOK, "delete", nil)
	}
}

func htmlShowItemsHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			formatter.HTML(w, http.StatusOK, "show", nil)
		} else {
			req.ParseForm() // parsing the parameters
			formatter.HTML(w, http.StatusOK, "show",
				showItemsByUsernameAndPassword(req.FormValue("username"), req.FormValue("password")))
		}
	}
}
