package connection

import (
	"database/sql"
	"fmt"
	"os"
)

func Connect_db() *sql.DB {
	fmt.Println("DB Connecting")
	password := os.Getenv("MYSQL_PASS")
	fmt.Println("Pass:", password)
	db, err := sql.Open("mysql", fmt.Sprintf("db_user:%s@tcp(0.0.0.0:8089)/levels", password))

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("DB Connected")
	return db
}
