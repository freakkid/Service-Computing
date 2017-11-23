package server

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/freakkid/Service-Computing/hw4/tools"
	_ "github.com/go-sql-driver/mysql"
)

/* Do some operations on database directly */

// open database and make sure to enable to be pinged
// to encrypt password by MD5
// do some operation on database and close connect after that
// return the result of operation (true/false)
func operateOnDB(function interface{}, username string, password string, parameter interface{}) ([]string, bool) {
	db, err := openDB(dbPara) // open database and ensure that be able to ping the database
	if err != nil {
		return nil, false
	}

	defer db.Close()

	// judge function type
	// if is func addUserIntoDB(db *sql.DB, username string, password string) error
	if addUserIntoDBFunc, ok := function.(func(*sql.DB, string, string) error); ok {
		if addUserIntoDBFunc(db, username, password) != nil {
			return nil, false
		}
		return nil, true

		// if is addItemIntoDB(db *sql.DB, username string, password string, item string) error
	} else if addItemIntoDBFunc, ok := function.(func(*sql.DB, string, string, string) error); ok {
		var item string
		if item, ok = parameter.(string); ok { // get string parameter item
			if addItemIntoDBFunc(db, username, password, item) == nil {
				return nil, true
			}
		}
		return nil, false

		// if is deleteItemIntoDB(db *sql.DB, username string, password string, itemIndex int) error
	} else if deleteItemIntoDBFunc, ok := function.(func(*sql.DB, string, string, int) error); ok {
		var itemIndex int
		if itemIndex, ok = parameter.(int); ok {
			if deleteItemIntoDBFunc(db, username, password, itemIndex) == nil {
				return nil, true
			}
		}
		return nil, false

		// if is showItemsFromDB(db *sql.DB, username string, password string) ([]string, error)
	} else if showItemsFromDBFunc, ok := function.(func(*sql.DB, string, string) ([]string, error)); ok {
		todoList, err := showItemsFromDBFunc(db, username, password)
		if err != nil {
			return nil, false
		}
		return todoList, true

		// must not exist other kinds of function
	} else {
		panic(errors.New("wrong parameters"))
	}

}

// some basic operation: open databases, check user exist or get todolist by username
// try to open database and return error if fails
func openDB(DBpara string) (*sql.DB, error) {
	db, err := sql.Open("mysql", DBpara)
	if err != nil || db.Ping() != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

// check if user exists in database by username, return bool result and user's id if exsits
func dbQueryUser(db *sql.DB, username string, password string) (bool, string) {
	rows, err := db.Query(dbStatements["QUERYUSER"], username, password)
	if err != nil {
		return false, ""
	}
	for rows.Next() { // if exists
		var id string
		if rows.Scan(&id) == nil { // if get id return true
			return true, id
		}
	}
	return false, ""
}

// get todolist string from database by username and password
func dbQueryTodos(db *sql.DB, id string) string {
	rows, err := db.Query(dbStatements["SHOWTODOS"], id)
	if err != nil {
		return ""
	}

	var todos string // get todos
	for rows.Next() {
		rows.Scan(&todos)
	}
	return todos
}

// ---------------------------------------------------------------
// only do some main operations: add user, update or select todolist
// add user in database
func addUserIntoDB(db *sql.DB, username string, password string) error {
	_, err := db.Exec(dbStatements["REGISTER"], tools.GetUUID(), username, password, "")
	return err
}

func getItemsById(db *sql.DB, id string) []string {
	return strings.Split(dbQueryTodos(db, id), ",")
}

// add item in database
func addItemIntoDB(db *sql.DB, username string, password string, item string) error {
	// check if the user exists
	if result, id := dbQueryUser(db, username, password); result {
		// add item in todos
		var todosList []string = append(getItemsById(db, id), item)
		_, err := db.Exec(dbStatements["EDITTODOS"], strings.Join(todosList, ","), id)
		return err
	}

	return errors.New("The account with this password was not found")
}

// parameter@todoListLength is len(strings.Split(dbQueryTodos(db, id), ","))-1
// to get actual todolistLength by "len-1" after spliting by "," because
// the first one string is empty after spliting by "," that should be ignored.
// check if item index is invalid; valid item index : from 0 to todolist actual length
func isInvalidItemIndex(itemIndex int, todoListLength int) bool {
	return itemIndex < 1 || itemIndex > todoListLength
}

// delete item by index in database
func deleteItemIntoDB(db *sql.DB, username string, password string, itemIndex int) error {
	// check if the user exists
	if result, id := dbQueryUser(db, username, password); result {
		// get item in todos
		var todosList []string = getItemsById(db, id)
		// check if itemIndex is valid, 0 is invalid becase the 0th one is empty string after spliting by ","
		if isInvalidItemIndex(itemIndex, len(todosList)-1) {
			return errors.New("Could not find the item")
		}

		// delete an item by index
		todosList = append(todosList[:itemIndex], todosList[itemIndex+1:]...)
		_, err := db.Exec(dbStatements["EDITTODOS"], strings.Join(todosList, ","), id)
		return err
	}

	return errors.New("The account with this password was not found")
}

// get all todo items from database
func showItemsFromDB(db *sql.DB, username string, password string) ([]string, error) {
	// check if the user exists
	if result, id := dbQueryUser(db, username, password); result {
		// get item in todos
		return getItemsById(db, id), nil
	}

	return nil, errors.New("The account with this password was not found")
}
