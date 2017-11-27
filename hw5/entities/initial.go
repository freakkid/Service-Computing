package entities

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-xorm/xorm"
)

// ORM engine
var xormEngine *xorm.Engine

func init() {
	var (
		username = "root"          // the username of mysql database
		password = "pincushion147" // the password of the username
		addrs    = "127.0.0.1"     // the tcp address
		port     = "3306"          // the port
		dbName   = "test"
		err      error
	)
	var dataSourceName string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", username, password, addrs, port, dbName)

	// create engine
	xormEngine, err = xorm.NewEngine("mysql", dataSourceName)
	checkErr(err)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
