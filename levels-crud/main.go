package main

// TODO: All microservice =)

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func init_db() {
	fmt.Println("Connecting")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(0.0.0.0:3306)/levels", username, password))

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected")

	// defer the close till after the main function has finished
	// executing
	defer db.Close()
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		fmt.Println("Server start")
		r := mux.NewRouter().StrictSlash(true)
		// r.HandleFunc("/wavedrom", wavedrom).Methods("POST")

		init_db()

		exec.Command("fuser", "-k", "8082/tcp").Output()
		http.ListenAndServe(":8082", r)
		fmt.Println("Server stop")
	}
}
