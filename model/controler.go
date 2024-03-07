package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver       = "mysql"
	dataSourceName = "root:123@tcp(localhost:3306)/testdb"
)

func StartConn() {
	db, err := sql.Open(dbDriver, dataSourceName)
	// sql.Open(driverName , dataSourceName string)
	if err != nil {
		fmt.Println("error validating sql.open artumgents")
		panic(err.Error())
	}
	// fmt.Println(db.Query("show databases"))
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connection with db.ping")
		panic(err.Error())
	}

	// func (db *sql.DB) Query(query string, args ...interface{})(* Rows,error){

	// }
	insert, err := db.Query(`INSERT INTO testdb.students(id,firstname,lastname) VALUES (1,"andrew","doghry");`)
	// insert, err := db.Query("show databases;")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	sel, err := db.Query("select firstname from testdb.students as db;")
	if err != nil {
		panic(err.Error())
	}
	var data string
	for sel.Next() {
		sel.Scan(&data)
		fmt.Println("select => ", data)
	}
	if err != nil {
		panic(err.Error())
	}
	defer sel.Close()
}
