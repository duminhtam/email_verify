package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"database/sql"
)

func main(){
	db, err := sql.Open("postgres", "host='/dev/shm/regress-the/pgsql0/data' dbname=blocketdb user=the password=")

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


