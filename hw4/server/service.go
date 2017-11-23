package server

import (
	"fmt"
	"strconv"

	"github.com/freakkid/Service-Computing/hw4/tools"
)

type TodoList struct {
	Operation string
	Username  string
	Todos     []string
	Result    string
	Message   string
}

// judge if username or password is empty
func isEmptyUsernameOrPassword(username string, password string) bool {
	return username == "" || password == ""
}

// check if username or password is empty
// add user into datebase and check if it succeeded
// return result and details
func registerByUsernameAndPassword(username string, password string) TodoList {
	if isEmptyUsernameOrPassword(username, password) { // if one of important parameters is empty
		return TodoList{Operation: REGISTER, Username: username, Result: FAIL, Message: messages["EmptyUsernameOrPassword"]}
	} else if _, ok := operateOnDB(addUserIntoDB, username, tools.MD5Encryption(password), nil); ok {
		return TodoList{Operation: REGISTER, Username: username, Result: SUCCESS, Message: messages["RegisterSuccess"]}
	}
	return TodoList{Operation: REGISTER, Username: username, Result: FAIL, Message: messages["RegisterFail"]}
}

// check if item is empty
func isEmptyItem(item string) bool {
	return item == ""
}

// check if username or password is empty
// add item into datebase, check if item is empty and check if succeeded
// return result and details
func addItemByUsernamePasswordItem(username string, password string, item string) TodoList {
	if isEmptyUsernameOrPassword(username, password) { // if one of important parameters is empty
		return TodoList{Operation: ADD, Username: username, Result: FAIL, Message: messages["EmptyUsernameOrPassword"]}
	} else if isEmptyItem(item) {
		return TodoList{Operation: ADD, Username: username, Result: FAIL, Message: messages["EmptyItem"]}
	} else if _, ok := operateOnDB(addItemIntoDB, username, tools.MD5Encryption(password), item); ok {
		return TodoList{Operation: ADD, Username: username, Result: SUCCESS, Message: messages["AddSuccess"]}
	}
	return TodoList{Operation: ADD, Username: username, Result: FAIL, Message: messages["AddFail"]}
}

// check if username or password is empty
// delete item from datebase and check if succeeded
// return result and details
func deleteItemByUsernameAndPasswordItemindex(
	username string, password string, itemIndexString string) TodoList {
	if isEmptyUsernameOrPassword(username, password) { // if one of important parameters is empty
		return TodoList{Operation: DELETE, Username: username, Result: FAIL, Message: messages["EmptyUsernameOrPassword"]}
	} else if itemIndex, err := strconv.Atoi(itemIndexString); err == nil {

		if _, ok := operateOnDB(deleteItemIntoDB, username, tools.MD5Encryption(password), itemIndex); ok {
			return TodoList{Operation: DELETE, Username: username, Result: SUCCESS, Message: messages["DeleteSuccess"]}
		}
	}
	return TodoList{Operation: DELETE, Username: username, Result: FAIL, Message: messages["DeleteFail"]}
}

// check if username or password is empty
// show items from datebase and check if succeeded
// return todolist, result and details
func showItemsByUsernameAndPassword(username string, password string) TodoList {
	if isEmptyUsernameOrPassword(username, password) { // if one of important parameters is empty
		return TodoList{Operation: SHOW, Username: username, Result: FAIL, Message: messages["EmptyUsernameOrPassword"]}
	} else if todoList, ok := operateOnDB(showItemsFromDB, username, tools.MD5Encryption(password), nil); ok {
		return TodoList{
			Operation: SHOW, Username: username, Todos: todoList, Result: SUCCESS,
			Message: fmt.Sprintf(messages["ShowSuccess"], len(todoList)-1)}
	}
	return TodoList{Operation: SHOW, Username: username, Result: FAIL, Message: messages["ShowFail"]}
}
