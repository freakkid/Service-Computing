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

type TodoList struct {
	Username string
	Todos    []string
	Message  string
}

// to render html template to return to client
// choose html template acconding to templateName
func RenderTemplate(w http.ResponseWriter, templateName string, todoList *TodoList) {
	if err := Templates.ExecuteTemplate(w, templateName[5:], todoList); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

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
		message = Messages["EmptyUsernameOrPassword"]
	} else if DealAddUserIntoDBFn(username, password) { // succeed to execute
		message = Messages["RegisterSuccess"]
	} else {
		message = Messages["RegisterFail"]
	}
	// render html template
	RenderTemplate(w, RegisterTemplate, &TodoList{Username: username, Todos: []string{}, Message: message})
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
		message = Messages["EmptyUsernameOrPassword"]
	} else if DealAddItemIntoDBFn(username, password, item) {
		message = Messages["AddSuccess"]
	} else {
		message = Messages["AddFail"]
	}
	RenderTemplate(w, AddItemTemplate, &TodoList{Username: username, Todos: []string{}, Message: message})
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
		message = Messages["EmptyUsernameOrPassword"]
	} else {
		itemIdex, err := strconv.Atoi(itemIndexString)
		if err != nil || !DealDeleteItemIntoDBFn(username, password, itemIdex) {
			message = Messages["DeleteFail"]
		} else {
			message = Messages["DeleteSuccess"]
		}
	}
	RenderTemplate(w, DeleteItemTemplate, &TodoList{Username: username, Todos: []string{}, Message: message})
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
		message = Messages["EmptyUsernameOrPassword"]
	} else {
		todoList, result = DealShowItemsFromDBFn(username, password)
		if result {
			message = fmt.Sprintf(Messages["ShowSuccess"], len(todoList)-1)
		} else {
			message = Messages["ShowFail"]
		}
	}
	RenderTemplate(w, ShowItemsTemplate, &TodoList{Username: username, Todos: todoList, Message: message})
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
