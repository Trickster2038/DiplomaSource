package stats

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gateway/request"
	"io/ioutil"
	"net/http"
)

type StatsRequest struct {
	Scope    string `json:"scope"`
	StatType string `json:"stat_type"`
	UserId   int    `json:"user_id"`
}

func Stats(w http.ResponseWriter, req *http.Request) {
	var response request.ResponseFrame

	defer func() {
		if panicInfo := recover(); panicInfo != nil {
			var response request.ResponseFrame
			response.StatusStr = "error"
			response.StatusCode = 400
			response.Message = fmt.Sprintf("Top-level panic: %v", panicInfo)
			w.WriteHeader(response.StatusCode)
			json.NewEncoder(w).Encode(response)
		}
	}()

	reqBody, _ := ioutil.ReadAll(req.Body)
	var dataFrame StatsRequest

	err := json.Unmarshal(reqBody, &dataFrame)
	payload, _ := json.Marshal(dataFrame)

	var endpoint string

	if dataFrame.Scope == "personal" {
		endpoint = "personalstats"
	} else if dataFrame.Scope == "general" {
		endpoint = "generalstats"
	} else {
		panic(fmt.Sprintf("Unknown stats scope: %s", dataFrame.Scope))
	}

	resp, err_post := http.Post(fmt.Sprintf("http://stats-microservice:8085/%s", endpoint), "application/json", bytes.NewBuffer(payload))

	if err_post != nil {
		panic(fmt.Sprintf("Accesing stats-microservice error: %v", err.Error()))
	}

	var res request.ResponseFrame
	json.NewDecoder(resp.Body).Decode(&res)

	response.StatusStr = res.StatusStr
	response.StatusCode = res.StatusCode
	response.Data = res.Data
	if res.StatusStr == "ok" {
		response.Message = ""
	} else {
		response.Message = "Stats-microservice error: " + res.Message
	}
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}
