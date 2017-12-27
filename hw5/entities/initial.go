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

	// creatse engine
	xormEngine, err = xorm.NewEngine(driverName, dataSourceName)
	checkErr(err)

	xormEngine.SetMapper(core.GonicMapper{})

	// sync the struct changes to database
	checkErr(xormEngine.Sync2(new(UserInfo)))
	checkErr(xormEngine.Sync2(new(User)))
	checkErr(xormEngine.Sync2(new(Meeting)))
	_, _ = xormEngine.Insert(User{UserName: "llleel", Password: "ssse", Email: " 111 234 ", Phone: "111"})
	_, _ = xormEngine.Insert(User{UserName: "llleel2", Password: "ssse", Email: " 111 123", Phone: "111"})
	_, _ = xormEngine.Insert(User{UserName: "llleel3", Password: "ssse", Email: "1111", Phone: "111"})
	_, _ = xormEngine.Insert(User{UserName: "llleel4", Password: "ssse", Email: "oo 111 dd", Phone: "111"})
	_, _ = xormEngine.Insert(User{UserName: "llleel5", Password: "ssse", Email: "11ee", Phone: "111"})
	var users []User
	xormEngine.Where("email LIKE ?", "% 111 %").Limit(2).Find(&users)
	var users1 []User
	xormEngine.In("password", "ssse").Find(&users1)
	fmt.Println(users)
	fmt.Println(users1)

	//checkErr(err)
}
