package server

import (
	"database/sql"
	"fmt"
	"html/template"
)

// template files for sending to client
// html template files path
const (
	RegisterTemplate   = "html/register.html"
	AddItemTemplate    = "html/add.html"
	DeleteItemTemplate = "html/delete.html"
	ShowItemsTemplate  = "html/show.html"
)

var htmlFilesNames = []string{
	RegisterTemplate,
	AddItemTemplate,
	DeleteItemTemplate,
	ShowItemsTemplate,
}

var Templates = template.Must(template.ParseFiles(htmlFilesNames...))

// -------------------------------------------------

// database names and sqlstatements
var (
	DBName       = "todos"          // database name
	DBTableName  = "primaryversion" // a table name
	DBPara       string             // database open parameter
	CreateDBPara string             // database open parameter (to create database)
)

// database execute statements
var DBStatements = map[string]string{
	"CREATEDB": "CREATE DATABASE IF NOT EXISTS " + DBName,
	"USEDB":    "USE " + DBName,
	"CREATETABLE": "CREATE TABLE IF NOT EXISTS " + DBTableName +
		" (username varchar(255) PRIMARY KEY, password varchar(255) NOT NULL, todos Text)",
	"REGISTER":  "INSERT INTO " + DBTableName + " (username, password, todos) values (?, ?, ?)",
	"EDITTODOS": "UPDATE " + DBTableName + " set todos=? WHERE username=? AND password=?",
	"SHOWTODOS": "SELECT todos FROM " + DBTableName + " WHERE username=? AND password=?",
	"QUERYUSER": "SELECT username FROM " + DBTableName + " WHERE username=? AND password=?",
}

// -------------------------------------------------------------

// message to client
var Messages = map[string]string{
	"EmptyUsernameOrPassword": "username and password should be non-empty",
	"RegisterSuccess":         "register success",
	"RegisterFail":            "register fail: the username may have been used",
	"AddSuccess":              "add success",
	"AddFail":                 "add fail: please check username and password and the item should be non-empty",
	"DeleteSuccess":           "delete success",
	"DeleteFail":              "delete fail: please check username and password and the item index should be valid",
	"ShowSuccess":             "show success： you have %d todo items",
	"ShowFail":                "show fail: please check username and password",
}

// exec some simple sql statement
func DBExec(db *sql.DB, DBStatement string) {
	if _, err := db.Exec(DBStatement); err != nil {
		db.Close()
		panic(err)
	}
}

// to create database when program starts run
func init() {
	// The paraments should reset when runs on a new platform
	var (
		username = "root"          // the username of mysql database
		password = "pincushion147" // the password of the username
		addrs    = "127.0.0.1"     // the tcp address
		port     = "3306"          // the port
	)
	DBPara = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&timeout=30s", username, password, addrs, port, DBName)
	CreateDBPara = fmt.Sprintf("%s:%s@tcp(%s:%s)/", username, password, addrs, port)

	// // create database and table when init package
	db, err := OpenDB(CreateDBPara)
	if err != nil {
		panic(err)
	}

	DBExec(db, DBStatements["CREATEDB"])    // create database
	DBExec(db, DBStatements["USEDB"])       // use database
	DBExec(db, DBStatements["CREATETABLE"]) // create table in database

	db.Close()
}
