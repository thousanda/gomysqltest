package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	ID   int64
	Name string
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		log.Printf(err.Error())
		return
	}
	defer db.Close()
	q := "SELECT * FROM t_user;"
	rows, err := db.Query(q)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			log.Printf(err.Error())
		}
		fmt.Printf("ID: %v, Name: %v\n", user.ID, user.Name)
	}
	err = rows.Err()
	if err != nil {
		log.Printf(err.Error())
	}
}
