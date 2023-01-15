package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gateway/request"
	"io/ioutil"
	"net/http"
)

func Crud_user(w http.ResponseWriter, req *http.Request) {
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
	var dataFrame request.InRequsestCRUD

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

	var res request.ResponseFrame
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
