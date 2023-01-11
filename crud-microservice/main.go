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

type User struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
}

type LevelsBrief struct {
	ID         int    `json:"id"`
	Level_type int    `json:"level_type"`
	Seqnum     int    `json:"seqnum"`
	Cost       int    `json:"cost"`
	Is_active  bool   `json:"is_active"`
	Name       string `json:"name"`
	Brief      string `json:"brief, omitempty"`
}

func (level_brief LevelsBrief) create() {
	db := connect_db()
	_, err := db.Query("UPDATE LevelsBrief SET seqnum = seqnum + 1 WHERE seqnum >= ?",
		level_brief.Seqnum)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	_, err = db.Query("INSERT INTO LevelsBrief (level_type, seqnum, cost, is_active, name, brief) VALUES (?, ?, ?, ?, ?, ?)",
		level_brief.Level_type,
		level_brief.Seqnum,
		level_brief.Cost,
		level_brief.Is_active,
		level_brief.Name,
		level_brief.Brief)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	defer db.Close()
}

func (level_brief *LevelsBrief) read(id int) {
	db := connect_db()
	err := db.QueryRow("SELECT id, level_type, seqnum, cost, is_active, name, brief FROM LevelsBrief where id = ?",
		id).
		Scan(&level_brief.ID,
			&level_brief.Level_type,
			&level_brief.Seqnum,
			&level_brief.Cost,
			&level_brief.Is_active,
			&level_brief.Name,
			&level_brief.Brief)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	defer db.Close()
}

func connect_db() *sql.DB {
	fmt.Println("DB Connecting")
	password := os.Getenv("MYSQL_PASS")
	fmt.Println("Pass:", password)
	db, err := sql.Open("mysql", fmt.Sprintf("db_user:%s@tcp(0.0.0.0:8089)/levels", password))

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("DB Connected")

	// defer the close till after the main function has finished
	// executing
	return db
}

func init_db() {
	fmt.Println("DB Connecting")
	password := os.Getenv("MYSQL_PASS")
	fmt.Println("Pass:", password)
	db, err := sql.Open("mysql", fmt.Sprintf("db_user:%s@tcp(0.0.0.0:8089)/levels", password))

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("DB Connected")

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT id, nickname FROM Users")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var user User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&user.ID, &user.Nickname)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		fmt.Println(user.ID, user.Nickname)
	}
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		fmt.Println("Server start")
		r := mux.NewRouter().StrictSlash(true)

		// level := LevelsBrief{0, 1, 2, 3, true, "Level1", "Test level"}
		// level2 := LevelsBrief{0, 1, 2, 3, true, "Level1", ""}
		// level2.create()

		var lvl_br LevelsBrief
		lvl_br.read(2)

		fmt.Println("Lvl name: ", lvl_br.Name)

		exec.Command("fuser", "-k", "8082/tcp").Output()
		http.ListenAndServe(":8082", r)
		fmt.Println("Server stop")
	}
}
