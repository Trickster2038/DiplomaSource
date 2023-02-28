package connection

import (
	"database/sql"
	"fmt"
	"os"
)

func Connect_db() *sql.DB {
	password := os.Getenv("MYSQL_PASS")
	db, err := sql.Open("mysql", fmt.Sprintf("db_user:%s@tcp(mysqldb:3306)/levels", password))

	if err != nil {
		panic(err.Error())
	}

	return db
}
