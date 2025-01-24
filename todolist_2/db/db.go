package db

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"os"
)

func ConnectDB() *sql.DB {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "todolist_2",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic("db open err | err : " + err.Error())
	}

	err2 := db.Ping()
	if err2 != nil {
		panic("db Ping err | err : " + err2.Error())
	}

	fmt.Println("db연결 성공")

	return db
}
