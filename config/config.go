package config

import (
	"database/sql"
	"fmt"
	_ `github.com/go-sql-driver/mysql`
)

const (
	username = `user1`
	password = `1234567890`
	database = `db_apigo`
)

var dsn = fmt.Sprintf(`%s:%s@/%s`,username,password,database)

func Mysql() (db *sql.DB,err error) {
	db, err = sql.Open(`mysql`,dsn)
	if err != nil {
		fmt.Println(err)
	}
	return 
}
