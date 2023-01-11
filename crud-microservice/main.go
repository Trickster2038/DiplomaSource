package main

// TODO:
// - All microservice =)
// - read_all() for levels_brief
// - read levels_brief & levels_data (or just call .read(id) twice?)

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"crud/levelsbrief"
	"crud/levelsdata"
)

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		fmt.Println("Server start")
		r := mux.NewRouter().StrictSlash(true)

		// level := LevelsBrief{0, 1, 2, 3, true, "Level1", "Test level"}

		var lvl_br levelsbrief.LevelsBrief
		lvl_br.Read(1)
		fmt.Println("Lvl name: ", lvl_br.Name)

		var level_data levelsdata.LevelsData
		level_data.Read(1)
		res, _ := json.Marshal(level_data)
		fmt.Println(string(res))

		level_data = levelsdata.LevelsData{2, "1tt", "2t", "3t", "4t"}
		level_data.CreateOrUpdate()

		exec.Command("fuser", "-k", "8082/tcp").Output()
		http.ListenAndServe(":8082", r)
		fmt.Println("Server stop")
	}
}
