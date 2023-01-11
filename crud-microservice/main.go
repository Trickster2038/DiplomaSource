package main

// TODO:
// - All microservice =)
// - read_all() for levels_brief
// - read levels_brief & levels_data (or just call .read(id) twice?)

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"crud/levelsbrief"
)

type User struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		fmt.Println("Server start")

		// connection.PrintHello()
		r := mux.NewRouter().StrictSlash(true)

		// level := LevelsBrief{0, 1, 2, 3, true, "Level1", "Test level"}
		// level2 := LevelsBrief{0, 1, 5, 3, true, "Level221", "Test descr"}
		// level2.create()

		var lvl_br levelsbrief.LevelsBrief
		lvl_br.Read(6)
		fmt.Println("Lvl name: ", lvl_br.Name)

		// level := LevelsBrief{5, 1, 0, 3, true, "Level1111", "Test level"}
		// level.delete()

		exec.Command("fuser", "-k", "8082/tcp").Output()
		http.ListenAndServe(":8082", r)
		fmt.Println("Server stop")
	}
}
