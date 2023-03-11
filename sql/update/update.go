package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456789@/cursogo")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")

	stmt.Exec("Caio", 1)
	stmt.Exec("Julia", 2)

	stmt2, _ := db.Prepare("delete from usuarios where id = ?")

	stmt2.Exec(2)
}
