package main

import (
	"fmt"
	"gateway/check"
	"gateway/levels"
	"gateway/stats"
	"gateway/user"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/gorilla/mux"
)

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		fmt.Println("Server start")
		r := mux.NewRouter().StrictSlash(true)
		r.HandleFunc("/user", user.Crud_user).Methods("POST")
		r.HandleFunc("/levels", levels.Crud_levels).Methods("POST")
		r.HandleFunc("/check", check.Check).Methods("POST")
		r.HandleFunc("/stats", stats.Stats).Methods("POST")

		exec.Command("fuser", "-k", "8084/tcp").Output()
		http.ListenAndServe(":8084", r)
		fmt.Println("Server stop")
	}
}
