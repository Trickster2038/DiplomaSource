package levels

import (
	"encoding/json"
	"fmt"
	"gateway/request"
	"io/ioutil"
	"net/http"
)

func Crud_levels(w http.ResponseWriter, req *http.Request) {
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

	if (dataFrame.MetaInfo.ObjType != "levels_brief") &&
		(dataFrame.MetaInfo.ObjType != "levels_data") {
		panic(fmt.Sprintf("Wrong CRUD-Object Type: %s", dataFrame.MetaInfo.ObjType))
	}

	if (dataFrame.MetaInfo.ObjType == "levels_data") &&
		(dataFrame.MetaInfo.Action == "delete") {
		panic("LevelsData can be only ARCHIVED by archivating LevelsBrief")
	}

	if (dataFrame.MetaInfo.Action == "create") ||
		(dataFrame.MetaInfo.Action == "update") ||
		(dataFrame.MetaInfo.Action == "delete") {

	}
}
