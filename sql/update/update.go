package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:my_password@/learngo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("update users set name = ? where id = ?")

	stmt.Exec("Shion", 1)
	stmt.Exec("Dhoko", 2)

	stmt2, _ := db.Prepare("delete from users where id = ?")
	stmt2.Exec(3)
}
