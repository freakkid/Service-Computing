package server

import (
	"net/http"

	"github.com/unrolled/render"
)

/* Server provides some api to client
 * The correct URLs as follows:
 * /api/register?username=XXX&password=XXX
 * /api/add?username=XXX&&password=XXX&item=XXX
 * /api/delete?username=XXX&password=XXX&itemIndex=XXX
 * /api/show?username=XXX&&password=XXX
 * Some incorrect URLs may receive error message, some may get 404 page.
 */

// handle "api/register"
// correct URL: /api/register?username=XXX&password=XXX
// return infomation of registering fail/successfully
func apiRegisterHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm() // parsing the parameters

		formatter.JSON(w, http.StatusOK,
			registerByUsernameAndPassword(req.FormValue("username"), req.FormValue("password")))
	}
}

func apiAddItemHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm() // parsing the parameters

		formatter.JSON(w, http.StatusOK,
			addItemByUsernamePasswordItem(
				req.FormValue("username"), req.FormValue("password"), req.FormValue("item")))
	}
}

func apiDeleteItemHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm() // parsing the parameters

		formatter.JSON(w, http.StatusOK,
			deleteItemByUsernameAndPasswordItemindex(
				req.FormValue("username"), req.FormValue("password"), req.FormValue("itemIndex")))
	}
}

func apiShowItemsHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm() // parsing the parameters

		formatter.JSON(w, http.StatusOK,
			showItemsByUsernameAndPassword(req.FormValue("username"), req.FormValue("password")))
	}
}
