package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"database/sql"
)

func main(){
	_, err := sql.Open("postgres", "user=tam host=/dev/shm/regress-tam/pgsql0/dataa dbname=blocketdb")

	if err != nil {
		fmt.Print(err)
	}
//	rows, err := db.Query("SELECT * FROM users")
//
//	fmt.Println(rows)

//	fmt.Print(VerifyEmail("tamdu@chotot.vn"))
}


