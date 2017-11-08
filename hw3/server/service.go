package server

import (
	"github.com/freakkid/Service-Computing/hw3/tools"
)

/* open database and make sure to enable to be pinged
 * to encrypt password by MD5
 * do some operation on database and close connect after that
 * return the result of operation (true/false)
 */

func DealAddUserIntoDBFn(username string, password string) bool {
	db, err := OpenDB(DBPara) // open database and ensure that be able to ping the database
	if err != nil {
		return false
	}

	defer db.Close()
	if AddUserIntoDB(db, username, tools.MD5Encryption(password)) != nil {
		return false
	}
	return true
}

func DealAddItemIntoDBFn(username string, password string, item string) bool {
	db, err := OpenDB(DBPara)
	if err != nil {
		return false
	}

	defer db.Close()
	if err = AddItemIntoDB(db, username, tools.MD5Encryption(password), item); err != nil {
		return false
	}
	return true
}

func DealDeleteItemIntoDBFn(username string, password string, itemIndex int) bool {
	db, err := OpenDB(DBPara)
	if err != nil {
		return false
	}

	defer db.Close()
	if DeleteItemIntoDB(db, username, tools.MD5Encryption(password), itemIndex) != nil {
		return false
	}
	return true
}

func DealShowItemsFromDBFn(username string, password string) ([]string, bool) {
	db, err := OpenDB(DBPara)
	if err != nil {
		return nil, false
	}

	defer db.Close()
	todoList, err := ShowItemsFromDB(db, username, tools.MD5Encryption(password))
	if err != nil {
		return nil, false
	}
	return todoList, true
}
