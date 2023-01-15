package main

import (
	"fmt"
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

		exec.Command("fuser", "-k", "8084/tcp").Output()
		http.ListenAndServe(":8084", r)
		fmt.Println("Server stop")
	}
}
