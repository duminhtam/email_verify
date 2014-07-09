package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"database/sql"
)

func main(){
	db, err := sql.Open("postgres", "user=tam host=/dev/shm/regress-tam/pgsql0/data dbname=blocketdb")

	if err != nil {
		fmt.Print(err)
	}
	rows, err := db.Query("SELECT * FROM users")

	for rows.Next(){
		var email string


		if err := rows.Scan(&email); err != nil {
			fmt.Printf("email: %s\n", email)
		}

	}
//
//	fmt.Println(rows)

//	fmt.Print(VerifyEmail("tamdu@chotot.vn"))
}


