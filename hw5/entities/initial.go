package entities

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

// ORM engine
var xormEngine *xorm.Engine

// create database
func createDB(driverName string, createDBPara string, createDataBaseStmt string) {
	// open a database
	db, err := sql.Open(driverName, createDBPara)
	defer db.Close()
	checkErr(err)

	// create database if not exist
	_, err = db.Exec(createDataBaseStmt)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {

	const (
		username   = "root"          // the username of mysql database
		password   = "pincushion147" // the password of the username
		addrs      = "127.0.0.1"     // the tcp address
		port       = "3306"          // the port
		driverName = "mysql"         // name of sql driver
		dbName     = "test"          // database name
		//dbTableName = "user_info"     // table name
	)

	var (
		dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
			username, password, addrs, port, dbName)
		createDBPara       = fmt.Sprintf("%s:%s@tcp(%s:%s)/", username, password, addrs, port)
		createDataBaseStmt = "CREATE DATABASE IF NOT EXISTS " + dbName
		err                error
	)

	// create database before creating xorm engine
	createDB(driverName, createDBPara, createDataBaseStmt)

	// create engine
	xormEngine, err = xorm.NewEngine(driverName, dataSourceName)
	checkErr(err)

	xormEngine.SetMapper(core.GonicMapper{})

	// sync the struct changes to database
	err = xormEngine.Sync2(new(UserInfo))
	checkErr(err)
}
