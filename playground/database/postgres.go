package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	connStr := "postgres://postgres:wawawa@localhost:5432/django_postgres_docker?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Connect Error!")
	}
	log.Println("连接数据库成功!")
	rows, err := db.Query(`select name  from test where id=1`)
	//rows, err := db.Query("select version()")

	if err != nil {
		log.Fatal(err)

	}
	//panic("宕机")
	log.Println("执行查询成功!")

	//var results = make([]string, 0)
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal("Error!")
		}
		log.Println("当前版本:", name)

	}
}
