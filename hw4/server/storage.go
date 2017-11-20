package server

import (
	"database/sql"
	"errors"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

/* Do some operations on database directly */
// ---------------------------------------------------------------
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

// check if user exists in database by username
func dbQueryUser(db *sql.DB, username string, password string) bool {
	rows, err := db.Query(dbStatements["QUERYUSER"], username, password)
	if err != nil {
		return false
	}
	for rows.Next() { // if exists return true
		return true
	}
	return false
}

// get todolist string from database by username and password
func dbQueryTodos(db *sql.DB, username string, password string) string {
	rows, err := db.Query(dbStatements["SHOWTODOS"], username, password)
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
	_, err := db.Exec(dbStatements["REGISTER"], username, password, "")
	return err
}

// add item in database
func addItemIntoDB(db *sql.DB, username string, password string, item string) error {
	// check if item is empty
	if item == "" {
		return errors.New("Could not add empty item")
	}
	// check if the user exists
	if !dbQueryUser(db, username, password) {
		return errors.New("The account with this password was not found")
	}
	// add item in todos
	var todosList []string = append(strings.Split(dbQueryTodos(db, username, password), ","), item)
	_, err := db.Exec(dbStatements["EDITTODOS"], strings.Join(todosList, ","), username, password)
	return err
}

// delete item by index in database
func deleteItemIntoDB(db *sql.DB, username string, password string, itemIndex int) error {
	// check if the user exists
	if !dbQueryUser(db, username, password) {
		return errors.New("The account with this password was not found")
	}
	// get item in todos
	var todosList []string = strings.Split(dbQueryTodos(db, username, password), ",")
	// check if item is empty, 0 is invalid becase the 0th one is empty string after spliting by ","
	if itemIndex < 1 || itemIndex > len(todosList)-1 {
		return errors.New("Could not find the item")
	}
	// delete an item by index
	todosList = append(todosList[:itemIndex], todosList[itemIndex+1:]...)
	_, err := db.Exec(dbStatements["EDITTODOS"], strings.Join(todosList, ","), username, password)
	return err
}

// get all todo items from database
func showItemsFromDB(db *sql.DB, username string, password string) ([]string, error) {
	// check if the user exists
	if !dbQueryUser(db, username, password) {
		return nil, errors.New("The account with this password was not found")
	}
	// get item in todos
	return strings.Split(dbQueryTodos(db, username, password), ","), nil
}
