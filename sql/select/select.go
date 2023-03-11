package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	name string
}

func main() {
	db, err := sql.Open("mysql", "root:123456789@/cursogo")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("select * from usuarios where id > ?", 5)
	defer rows.Close()

	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.name)
		fmt.Println(u)
	}
}
