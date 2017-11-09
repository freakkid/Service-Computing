package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/freakkid/Service-Computing/hw3/tools"
)

/* Server handles some requests from client and log some neccessary information.
 * The correct URLs as follows:
 * user/registe?username=XXX&password=XXX
 * todo/add?username=XXX&&password=XXX&item=XXX
 * todo/delete?username=XXX&password=XXX&itemIndex=XXX
 * todo/show?username=XXX&&password=XXX
 * Some incorrect URLs may receive error message, some may get 404 page.
 */

// handle "user/register"
// correct URL: user/registe?username=XXX&password=XXX
// return infomation of registering fail/successfully
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // parsing the parameters
	var (
		username = r.FormValue("username")
		password = r.FormValue("password")
		message  string
	)
	if username == "" || password == "" { // if one of important parameters is empty
		message = messages["EmptyUsernameOrPassword"]
	} else if dealAddUserIntoDBFn(username, password) { // succeed to execute
		message = messages["RegisterSuccess"]
	} else {
		message = messages["RegisterFail"]
	}
	// render html template
	renderTemplate(w, registerTemplate, &TodoList{Username: username, Todos: []string{}, Message: message})
	// log infomation on server
	tools.LogOKInfo(r.Method, "register")
}

// handle "todo/add"
// correct URL: todo/add?username=XXX&&password=XXX&item=XXX
// return infomation of adding item fail/successfully
func AddItemHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var (
		username = r.FormValue("username")
		password = r.FormValue("password")
		item     = r.FormValue("item")
		message  string
	)
	if username == "" || password == "" {
		message = messages["EmptyUsernameOrPassword"]
	} else if dealAddItemIntoDBFn(username, password, item) {
		message = messages["AddSuccess"]
	} else {
		message = messages["AddFail"]
	}
	renderTemplate(w, addItemTemplate, &TodoList{Username: username, Todos: []string{}, Message: message})
	tools.LogOKInfo(r.Method, "add")
}

// handle "todo/delete"
// correct URL: todo/delete?username=XXX&password=XXX&itemIndex=XXX
// return infomation of deleting item fail/successfully
func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var (
		username        = r.FormValue("username")
		password        = r.FormValue("password")
		itemIndexString = r.FormValue("itemIndex")
		message         string
	)

	if username == "" || password == "" {
		message = messages["EmptyUsernameOrPassword"]
	} else {
		itemIdex, err := strconv.Atoi(itemIndexString)
		if err != nil || !dealDeleteItemIntoDBFn(username, password, itemIdex) {
			message = messages["DeleteFail"]
		} else {
			message = messages["DeleteSuccess"]
		}
	}
	renderTemplate(w, deleteItemTemplate, &TodoList{Username: username, Todos: []string{}, Message: message})
	tools.LogOKInfo(r.Method, "delete")
}

// handle "todo/show"
// correct URL: todo/show?username=XXX&&password=XXX
// return infomation of showing items fail/successfully
func ShowListHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var (
		username = r.FormValue("username")
		password = r.FormValue("password")
		message  string
		todoList []string
		result   bool
	)
	if username == "" || password == "" {
		message = messages["EmptyUsernameOrPassword"]
	} else {
		todoList, result = dealShowItemsFromDBFn(username, password)
		if result {
			message = fmt.Sprintf(messages["ShowSuccess"], len(todoList)-1)
		} else {
			message = messages["ShowFail"]
		}
	}
	renderTemplate(w, ShowItemsTemplate, &TodoList{Username: username, Todos: todoList, Message: message})
	tools.LogOKInfo(r.Method, "show")
}

// handle URL that could not be know the purpose of client
// return 404 page
func OtherHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
	tools.LogNoFound(r.Method)
}

// server listens requests from client
// add "" tools.LogPortListening " to log when server begins to work
func ListenAndServe(addr string, handler http.Handler) error {
	tools.LogPortListening(addr[1:])
	server := &http.Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}
