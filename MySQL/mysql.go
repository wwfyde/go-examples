// 快速使用mysql
// 参考链接: https://pkg.go.dev/github.com/go-sql-driver/mysql
package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	//_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var cfg = mysql.Config{
	User:   os.Getenv("DBUSER"),
	Passwd: os.Getenv("DBPASS"),
	Net:    "tcp",
	Addr:   "127.0.0.1:3306",
	DBName: "go_examples",
}

func main() {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	check(err)
	err = db.Ping()
	check(err)
	log.Println("Connected!")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
