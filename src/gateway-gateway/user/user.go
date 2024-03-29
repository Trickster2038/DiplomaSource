package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gateway/request"
	"io/ioutil"
	"net/http"
)

func Check_user_exists(user_id int) {
	var user_request request.IdRequestFrame
	user_request.Data.ID = user_id
	user_request.MetaInfo.ObjType = "user"
	user_request.MetaInfo.Action = "read"
	payload, _ := json.Marshal(user_request)

	resp, err_post := http.Post("http://crud-microservice:8082/crud", "application/json", bytes.NewBuffer(payload))

	if err_post != nil {
		panic(fmt.Sprintf("Accesing CRUD-microservice.Users error: %v", err_post.Error()))
	}

	var res request.ResponseFrame
	json.NewDecoder(resp.Body).Decode(&res)
	if res.StatusStr != "ok" {
		panic(fmt.Sprintf("User with ID=%d does NOT exist", user_id))
	}
}

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
	var dataFrame request.InRequestCRUD

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
		response.Message = "CRUD operation with User done"
	} else {
		response.Message = "CRUD-microservice.User error: " + res.Message
	}
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}
