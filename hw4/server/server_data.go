package server

import (
	"database/sql"
	"fmt"
	"go/build"
	"html/template"
	"os"
)

var (
	staticDir string // directory of static html files
)

// template files for sending to client
// html template files path
const (
	registerTemplate   = "register.html"
	addItemTemplate    = "add.html"
	deleteItemTemplate = "delete.html"
	ShowItemsTemplate  = "show.html"
)

var htmlPosition string = build.Default.GOPATH + "/src/github.com/freakkid/Service-Computing/hw3/html/"

var htmlFilesNames = []string{
	htmlPosition + registerTemplate,
	htmlPosition + addItemTemplate,
	htmlPosition + deleteItemTemplate,
	htmlPosition + ShowItemsTemplate,
}

var templates = template.Must(template.ParseFiles(htmlFilesNames...))

// -------------------------------------------------

// database names and sqlstatements
var (
	dbName       = "todos"         // database name
	dbTableName  = "secondversion" // a table name
	dbPara       string            // database open parameter
	createDBPara string            // database open parameter (to create database)
)

// database execute statements
var dbStatements = map[string]string{
	"CREATEDB": "CREATE DATABASE IF NOT EXISTS " + dbName,
	"USEDB":    "USE " + dbName,
	"CREATETABLE": "CREATE TABLE IF NOT EXISTS " + dbTableName +
		" (id varchar(255) PRIMARY KEY, username varchar(255) NOT NULL UNIQUE, password varchar(255) NOT NULL, todos Text)",
	"REGISTER":  "INSERT INTO " + dbTableName + " (id, username, password, todos) values (?, ?, ?, ?)",
	"QUERYUSER": "SELECT id FROM " + dbTableName + " WHERE username=? AND password=?",
	"EDITTODOS": "UPDATE " + dbTableName + " set todos=? WHERE id=?",
	"SHOWTODOS": "SELECT todos FROM " + dbTableName + " WHERE id=?",
}

// -------------------------------------------------------------

// message to client

const (
	// the [result] of user's request
	SUCCESS = "success"
	FAIL    = "fail"
	// the [operation] of user's request
	REGISTER = "register"
	ADD      = "add"
	DELETE   = "delete"
	SHOW     = "show"
)

// the details of user request results
var messages = map[string]string{
	// empty username and password
	"EmptyUsernameOrPassword": "username and password should be non-empty",
	// register
	"RegisterSuccess": "register success",
	"RegisterFail":    "register fail",
	// add item
	"AddSuccess": "add success",
	"EmptyItem":  "add fail: the item should be non-empty",
	"AddFail":    "add fail: please check username and password",
	// delete item
	"DeleteSuccess": "delete success",
	"DeleteFail":    "delete fail: please check username and password and the item index should be valid",
	// show items
	"ShowSuccess": "show success： you have %d todo items",
	"ShowFail":    "show fail: please check username and password",
}

// exec some simple sql statement
func dbExec(db *sql.DB, DBStatement string) {
	if _, err := db.Exec(DBStatement); err != nil {
		db.Close()
		panic(err)
	}
}

func init() {
	// ----------- to create database when program starts run -------------
	// The paraments should reset when runs on a new platform
	var (
		username = "root"          // the username of mysql database
		password = "pincushion147" // the password of the username
		addrs    = "127.0.0.1"     // the tcp address
		port     = "3306"          // the port
	)
	dbPara = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&timeout=30s", username, password, addrs, port, dbName)
	createDBPara = fmt.Sprintf("%s:%s@tcp(%s:%s)/", username, password, addrs, port)

	// create database and table when init package
	db, err := openDB(createDBPara)
	if err != nil {
		panic(err)
	}

	dbExec(db, dbStatements["CREATEDB"])    // create database
	dbExec(db, dbStatements["USEDB"])       // use database
	dbExec(db, dbStatements["CREATETABLE"]) // create table in database

	db.Close()
	// ---------------------------------------------------------------------

	// ---------------------------- get paths ------------------------------
	// get current path
	var currentPath string
	if currentPath, err = os.Getwd(); err != nil {
		panic(err)
	}
	// directory of static html files
	staticDir = currentPath + "/assets/"
	// ---------------------------------------------------------------------

}
