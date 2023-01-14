package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"

	"github.com/gorilla/mux"
)

type SourceFiles struct {
	UserID    int    `json:"user_id"`
	LevelID   int    `json:"level_id"`
	DeviceSrc string `json:"device_src"`
	TbSrc     string `json:"tb_src"`
}

type ResponseFrame struct {
	StatusStr         string `json:"status_str"`
	StatusCode        int    `json:"status_code"`
	Message           string `json:"message, omitempty"`
	Value_change_dump string `json:"value_change_dump, omitempty"`
}

func add_dump_macros(user_id int, level_id int, tb_src string) string {
	var r = regexp.MustCompile(`\$dumpfile\(.*\);`)
	s := r.ReplaceAllString(tb_src, ``)
	r = regexp.MustCompile(`\$dumpvars;`)
	s = r.ReplaceAllString(s, "$$dumpfile(\""+"user"+
		strconv.Itoa(user_id)+"/level"+strconv.Itoa(level_id)+"/device"+
		".vcd\");\n$$dumpvars;")
	return s
}

func create_or_update(user_id int, level_id int, device_src string, tb_src string) int {
	os.MkdirAll(("user" + strconv.Itoa(user_id) + "/level" + strconv.Itoa(level_id)), 0777)
	f, err1 := os.Create(("user" + strconv.Itoa(user_id) + "/level" + strconv.Itoa(level_id) + "/device.v"))
	_, err2 := f.WriteString(device_src)
	f, err3 := os.Create(("user" + strconv.Itoa(user_id) + "/level" + strconv.Itoa(level_id) + "/tb.v"))
	_, err4 := f.WriteString(tb_src)
	if (err1 != nil) || (err2 != nil) || (err3 != nil) || (err4 != nil) {
		return 1
	}
	return 0
}

func compile_and_visualise(user_id int, level_id int) int {
	device_path := "user" + strconv.Itoa(user_id) + "/level" + strconv.Itoa(level_id) + "/device.v"
	tb_path := "user" + strconv.Itoa(user_id) + "/level" + strconv.Itoa(level_id) + "/tb.v"
	out_path := "user" + strconv.Itoa(user_id) + "/level" + strconv.Itoa(level_id) + "/device"

	_, err := exec.Command("/bin/sh", "-c", ("iverilog -o " + out_path + " " + device_path + " " + tb_path)).Output()
	if err != nil {
		fmt.Printf("CMD-ERROR: %s\n", err)
		return 1
	}

	_, err = exec.Command("/bin/sh", "-c", ("vvp " + out_path)).Output()
	if err != nil {
		return 2
	}

	return 0
}

func build(w http.ResponseWriter, req *http.Request) {
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
	var dataFrame SourceFiles

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

	r := regexp.MustCompile(`.*\$dumpvars;.*`)
	if !r.MatchString(dataFrame.TbSrc) {
		defer func() {
			if r := recover(); r != nil {
				response.StatusStr = "error"
				response.StatusCode = 400
				response.Message = "testbench without \"$dumpvars\""
				w.WriteHeader(response.StatusCode)
				json.NewEncoder(w).Encode(response)
			}
		}()
		panic("Testbench without \"$dumpvars\"")
	}

	defer func() {
		os.RemoveAll("user" + strconv.Itoa(dataFrame.UserID) + "/level" + strconv.Itoa(dataFrame.LevelID))
	}()

	dataFrame.TbSrc = add_dump_macros(dataFrame.UserID,
		dataFrame.LevelID, dataFrame.TbSrc)

	err_int := create_or_update(dataFrame.UserID, dataFrame.LevelID,
		dataFrame.DeviceSrc, dataFrame.TbSrc)
	if err_int != 0 {
		defer func() {
			if r := recover(); r != nil {
				response.StatusStr = "error"
				response.StatusCode = 400
				response.Message = "writing files error"
				w.WriteHeader(response.StatusCode)
				json.NewEncoder(w).Encode(response)
			}
		}()
		panic("Writing files error")
	}

	defer func() {
		err := os.RemoveAll("user" + strconv.Itoa(dataFrame.UserID) + "/level" + strconv.Itoa(dataFrame.LevelID))
		if err != nil {
			fmt.Printf("error: %s", err)
		}
	}()

	err_int = compile_and_visualise(dataFrame.UserID, dataFrame.LevelID)
	if err_int != 0 {
		defer func() {
			if r := recover(); r != nil {
				response.StatusStr = "error"
				response.StatusCode = 400
				if err_int == 1 {
					response.Message = "synthethis error"
				} else {
					response.Message = "simulation error"
				}
				w.WriteHeader(response.StatusCode)
				json.NewEncoder(w).Encode(response)
			}
		}()
		panic("Simulation error")
	}

	value_change_dump, _ := os.ReadFile("user" + strconv.Itoa(dataFrame.UserID) + "/level" +
		strconv.Itoa(dataFrame.LevelID) + "/device.vcd")

	response.StatusStr = "ok"
	response.StatusCode = 200
	response.Message = "compiled successfully"
	response.Value_change_dump = string(value_change_dump)
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		fmt.Println("Server start")
		r := mux.NewRouter().StrictSlash(true)
		r.HandleFunc("/build", build).Methods("POST")

		exec.Command("fuser", "-k", "8080/tcp").Output()
		http.ListenAndServe(":8080", r)
		fmt.Println("Server stop")
	}
}
