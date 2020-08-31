package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:my_password@/learngo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("insert into users(name) values(?)")
	stmt.Exec("John")
	stmt.Exec("Penny")

	res, _ := stmt.Exec("Peter")

	id, _ := res.LastInsertId()
	fmt.Println(id)

	lines, err := res.RowsAffected()
	fmt.Println(lines)
}
