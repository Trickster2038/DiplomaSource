package main

// TODO:
// - All microservice =)
// - read_all() for levels_brief
// - read levels_brief & levels_data (or just call .read(id) twice?)

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
		panic(err.Error())
	}

	_, err = db.Query("INSERT INTO LevelsBrief (level_type, seqnum, cost, is_active, name, brief) VALUES (?, ?, ?, ?, ?, ?)",
		level_brief.Level_type,
		level_brief.Seqnum,
		level_brief.Cost,
		level_brief.Is_active,
		level_brief.Name,
		level_brief.Brief)
	if err != nil {
		panic(err.Error())
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
		panic(err.Error())
	}

	defer db.Close()
}

func (level_brief LevelsBrief) update() {
	db := connect_db()

	var old_level_brief LevelsBrief

	old_level_brief.read(level_brief.ID) // can call panic

	if level_brief.Seqnum > old_level_brief.Seqnum {
		_, err := db.Query("UPDATE LevelsBrief SET seqnum = seqnum - 1 WHERE seqnum > ? AND seqnum < ?",
			old_level_brief.Seqnum, level_brief.Seqnum)
		if err != nil {
			panic(err.Error())
		}
	} else if level_brief.Seqnum < old_level_brief.Seqnum {
		_, err := db.Query("UPDATE LevelsBrief SET seqnum = seqnum + 1 WHERE seqnum > ? AND seqnum < ?",
			level_brief.Seqnum, old_level_brief.Seqnum)
		if err != nil {
			panic(err.Error())
		}
	}

	_, err := db.Query("UPDATE LevelsBrief SET "+
		"level_type = ?, "+
		"seqnum = ?, "+
		"cost = ?, "+
		"is_active = ?, "+
		"name = ?, "+
		"brief = ? "+
		"WHERE id = ?",
		level_brief.Level_type,
		level_brief.Seqnum,
		level_brief.Cost,
		level_brief.Is_active,
		level_brief.Name,
		level_brief.Brief,
		level_brief.ID)
	if err != nil {
		panic(err.Error())
	}

}

// setting is_active = FALSE
func (level_brief LevelsBrief) delete() {
	db := connect_db()

	level_brief.read(level_brief.ID) // can call panic
	if level_brief.Is_active {
		_, err := db.Query("UPDATE LevelsBrief SET seqnum = seqnum - 1 WHERE seqnum > ?",
			level_brief.Seqnum)
		if err != nil {
			panic(err.Error())
		}

		_, err = db.Query("UPDATE LevelsBrief SET is_active = false WHERE id = ?",
			level_brief.ID)
		if err != nil {
			panic(err.Error())
		}
	} else {
		panic("already deleted (archived)")
	}

}

func connect_db() *sql.DB {
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

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		fmt.Println("Server start")
		r := mux.NewRouter().StrictSlash(true)

		// level := LevelsBrief{0, 1, 2, 3, true, "Level1", "Test level"}
		level2 := LevelsBrief{0, 1, 5, 3, true, "Level221", "Test descr"}
		level2.create()

		// var lvl_br LevelsBrief
		// lvl_br.read(2)

		// fmt.Println("Lvl name: ", lvl_br.Name)

		level := LevelsBrief{5, 1, 0, 3, true, "Level1111", "Test level"}
		level.delete()

		exec.Command("fuser", "-k", "8082/tcp").Output()
		http.ListenAndServe(":8082", r)
		fmt.Println("Server stop")
	}
}
