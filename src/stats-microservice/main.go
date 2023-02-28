package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"runtime"
	"stats/general"
	"stats/personal"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type PersonalStatsRequest struct {
	UserID   int    `json:"user_id"`
	StatType string `json:"stat_type"`
}

type GeneralStatsRequest struct {
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
	} else if reqFrame.StatType == "activity_borders" {
		response.Data = personal.Activity_borders(reqFrame.UserID)
	} else {
		panic("Unknown personal stat type")
	}

	response.StatusStr = "ok"
	response.StatusCode = 200
	response.Message = ""
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}

func general_stats(w http.ResponseWriter, req *http.Request) {
	var reqFrame GeneralStatsRequest
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

	if reqFrame.StatType == "each_level_passed" {
		response.Data = general.Each_level_passed()
	} else if reqFrame.StatType == "each_level_avg_efforts" {
		response.Data = general.Each_level_avg_efforts()
	} else if reqFrame.StatType == "solutions_distribution" {
		response.Data = general.Solutions_count_distribution()
	} else if reqFrame.StatType == "activity_by_month" {
		response.Data = general.Activity_by_month()
	} else if reqFrame.StatType == "top_month_users" {
		response.Data = general.Top_last_month_active_users()
	} else {
		panic("Unknown general stat type")
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
		r.HandleFunc("/generalstats", general_stats).Methods("POST")

		exec.Command("fuser", "-k", "8085/tcp").Output()
		http.ListenAndServe(":8085", r)
		fmt.Println("Server stop")
	}
}
