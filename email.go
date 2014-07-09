package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"database/sql"
)

func main(){
	db, err := sql.Open("postgres", "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full")

	if err != nil {
		fmt.Print(err)
	}
	rows, err := db.Query("SELECT * FROM users")

	for rows.Next(){
		rows.Scan("a")
	}
//	fmt.Print(VerifyEmail("tamdu@chotot.vn"))
}


