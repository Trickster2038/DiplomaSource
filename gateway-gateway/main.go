package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/gorilla/mux"
)

type ResponseFrame struct {
	StatusStr  string      `json:"status_str"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

type InRequsestCRUD struct {
	UserID   int `json:"user_id"`
	MetaInfo struct {
		ObjType string `json:"obj_type"`
		Action  string `json:"action"`
	} `json:"metainfo"`
	Data interface{} `json: "data"`
}

func crud_user(w http.ResponseWriter, req *http.Request) {
	var response ResponseFrame

	defer func() {
		if panicInfo := recover(); panicInfo != nil {
			var response ResponseFrame
			response.StatusStr = "error"
			response.StatusCode = 400
			response.Message = fmt.Sprintf("Top-level panic: %v", panicInfo)
			w.WriteHeader(response.StatusCode)
			json.NewEncoder(w).Encode(response)
		}
	}()

	reqBody, _ := ioutil.ReadAll(req.Body)
	var dataFrame InRequsestCRUD

	err := json.Unmarshal(reqBody, &dataFrame)

	if err != nil {
		defer func() {
			if r := recover(); r != nil {
				response.StatusStr = "error"
				response.StatusCode = 400
				response.Message = err.Error()
				w.WriteHeader(response.StatusCode)
				json.NewEncoder(w).Encode(response)
			}
		}()
		panic("JSON parsing error")
	}

	if dataFrame.MetaInfo.ObjType != "user" {
		panic(fmt.Sprintf("Wrong CRUD-Object Type: %s", dataFrame.MetaInfo.ObjType))
	}

	payload, err := json.Marshal(dataFrame)
	resp, err_post := http.Post("http://crud-microservice:8082/crud", "application/json", bytes.NewBuffer(payload))

	if err_post != nil {
		panic(fmt.Sprintf("Accesing CRUD-microservice error: %v", err.Error()))
	}

	var res ResponseFrame
	json.NewDecoder(resp.Body).Decode(&res)

	response.StatusStr = res.StatusStr
	response.StatusCode = res.StatusCode
	response.Data = res.Data
	if res.StatusStr == "ok" {
		response.Message = "CRUD operation with user done"
	} else {
		response.Message = "CRUD-microservice.User error: " + res.Message
	}
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		fmt.Println("Server start")
		r := mux.NewRouter().StrictSlash(true)
		r.HandleFunc("/cruduser", crud_user).Methods("POST")

		exec.Command("fuser", "-k", "8084/tcp").Output()
		http.ListenAndServe(":8084", r)
		fmt.Println("Server stop")
	}
}
