package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"runtime"
	"stats/personal"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type PersonalStatsRequest struct {
	UserID   int    `json:"user_id"`
	StatType string `json:"stat_type"`
}

type ResponseFrame struct {
	StatusStr  string      `json:"status_str"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func personal_stats(w http.ResponseWriter, req *http.Request) {
	var reqFrame PersonalStatsRequest
	var response ResponseFrame

	defer func() {
		if panicInfo := recover(); panicInfo != nil {
			response.StatusStr = "error"
			response.StatusCode = 400
			response.Message = fmt.Sprintf("Top-level panic: %v", panicInfo)
			w.WriteHeader(response.StatusCode)
			json.NewEncoder(w).Encode(response)
		}
	}()

	reqBody, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(reqBody, &reqFrame)

	if reqFrame.StatType == "general_progress" {
		response.Data = personal.General_progress(reqFrame.UserID)
	} else if reqFrame.StatType == "each_level_status" {
		response.Data = personal.Levels_statuses(reqFrame.UserID)
	} else if reqFrame.StatType == "avg_efforts" {
		response.Data = personal.Average_efforts_per_level(reqFrame.UserID)
	} else if reqFrame.StatType == "monthly_activity" {
		response.Data = personal.Monthly_activity(reqFrame.UserID)
	} else {
		panic("Unknown personal stat type")
	}

	response.StatusStr = "ok"
	response.StatusCode = 200
	response.Message = ""
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		fmt.Println("Server start")
		r := mux.NewRouter().StrictSlash(true)
		r.HandleFunc("/personalstats", personal_stats).Methods("POST")

		exec.Command("fuser", "-k", "8085/tcp").Output()
		http.ListenAndServe(":8085", r)
		fmt.Println("Server stop")
	}
}
