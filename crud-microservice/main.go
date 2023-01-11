package main

// TODO:
// - Web-CRUD handlers
// - SolutionEfforts ORM

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
	"crud/solutionefforts"
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

		level_data = levelsdata.LevelsData{ID: 2, WideDescription: "1tttt", Code: "2t", Question: "3t", Answer: "4t"}
		// level_data.CreateOrUpdate()

		level_data.Update()

		var eff solutionefforts.SolutionEffort
		eff.UserID = 1
		eff.LevelID = 7
		eff.IsSuccessful = true

		eff.Create()

		exec.Command("fuser", "-k", "8082/tcp").Output()
		http.ListenAndServe(":8082", r)
		fmt.Println("Server stop")
	}
}
