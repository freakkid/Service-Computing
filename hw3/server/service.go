package server

import (
	"net/http"

	"github.com/freakkid/Service-Computing/hw3/tools"
)

type TodoList struct {
	Username string
	Todos    []string
	Message  string
}

// to render html template to return to client
// choose html template acconding to templateName
func renderTemplate(w http.ResponseWriter, templateName string, todoList *TodoList) {
	if err := templates.ExecuteTemplate(w, templateName, todoList); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

/* open database and make sure to enable to be pinged
 * to encrypt password by MD5
 * do some operation on database and close connect after that
 * return the result of operation (true/false)
 */

func dealAddUserIntoDBFn(username string, password string) bool {
	db, err := openDB(dbPara) // open database and ensure that be able to ping the database
	if err != nil {
		return false
	}

	defer db.Close()
	if addUserIntoDB(db, username, tools.MD5Encryption(password)) != nil {
		return false
	}
	return true
}

func dealAddItemIntoDBFn(username string, password string, item string) bool {
	db, err := openDB(dbPara)
	if err != nil {
		return false
	}

	defer db.Close()
	if err = addItemIntoDB(db, username, tools.MD5Encryption(password), item); err != nil {
		return false
	}
	return true
}

func dealDeleteItemIntoDBFn(username string, password string, itemIndex int) bool {
	db, err := openDB(dbPara)
	if err != nil {
		return false
	}

	defer db.Close()
	if deleteItemIntoDB(db, username, tools.MD5Encryption(password), itemIndex) != nil {
		return false
	}
	return true
}

func dealShowItemsFromDBFn(username string, password string) ([]string, bool) {
	db, err := openDB(dbPara)
	if err != nil {
		return nil, false
	}

	defer db.Close()
	todoList, err := showItemsFromDB(db, username, tools.MD5Encryption(password))
	if err != nil {
		return nil, false
	}
	return todoList, true
}
