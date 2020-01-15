package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     string = "172.36.21.200"
	port     int    = 5432
	user     string = "gpadmin"
	passworf string = "gpadmin"
	dbname   string = "db01"
)

func gpconnect() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, passworf, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("链接成功")
}

func main() {
	gpconnect()
}
