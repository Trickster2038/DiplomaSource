package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os/exec"
	"runtime"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
)

/*
TODO:
- web
- docker
- errors
- checks
- UI align of signals labels
- HEX labels for signals?
*/

type VCD_Struct struct {
	Signal []struct {
		Name string `json:"name"`
		Type struct {
			Name  string `json:"name"`
			Width int    `json:"width"`
		} `json:"type"`
		Data [][]interface{} `json:"data"`
	} `json:"vcd_parsed"`
}

type WD_Signal struct {
	Name string   `json:"name"`
	Wave string   `json:"wave"`
	Data []string `json:"data"`
}

type WD_Struct struct {
	Signal []WD_Signal `json:"signal"`
}

type SingleValueDump struct {
	Time  int
	Value string
}

type ResponseFrame struct {
	Status_str  string `json:"status_str"`
	Status_code int    `json:"status_code"`
	Message     string `json:"message,omitempty"`
	Wavedrom    WD_Struct `json:"wavedrom,omitempty"`
}

// FIXME: more effective algorithm
func findGCD(nums []int) int {
	gcd := 0
	sort.Ints(nums)
	l := nums[0]
	for i := 1; i <= l; i++ {
		flag := true
		for _, x := range nums {
			if x%i != 0 {
				flag = false
				break
			}

		}
		if flag {
			gcd = i
		}
	}
	return gcd
}

func (vcdFrame VCD_Struct) getSortedTimings() []int {
	var timeArray []int
	for _, x := range vcdFrame.Signal {
		for _, y := range x.Data {
			num := int(math.Round(y[0].(float64)))
			if num != 0 {
				timeArray = append(timeArray, num)
			}
		}
	}

	sort.Ints(timeArray)
	return timeArray
}

func (vcdFrame VCD_Struct) getMaxValueWidth() int {
	var WireWidth []int
	for _, x := range vcdFrame.Signal {
		WireWidth = append(WireWidth, x.Type.Width)
	}
	sort.Ints(WireWidth)
	return WireWidth[len(WireWidth)-1]
}

/*
returns (waveform, datalabels)
*/
func (vcdFrame VCD_Struct) parseValues(end_scale int, width_scale int) (map[string]string, map[string][]string) {
	var parsedData = map[string][]string{}
	var parsedWaves = map[string]string{}

	timings := vcdFrame.getSortedTimings()
	tick_amount := findGCD(timings)
	end_time := timings[len(timings)-1] + tick_amount*end_scale
	for i := 0; i < end_time/tick_amount; i++ {
		for _, x := range vcdFrame.Signal {
			flag := false
			name := x.Name
			single_wire := true
			if x.Type.Width > 1 {
				name += "[0:" + strconv.Itoa(x.Type.Width-1) + "]"
				single_wire = false
			}
			for _, y := range x.Data {
				if int(math.Round(y[0].(float64))) == i*tick_amount {
					if single_wire {
						parsedWaves[name] += y[1].(string)
					} else {
						parsedData[name] = append(parsedData[name], y[1].(string))
						parsedWaves[name] += "="
					}
					flag = true
					break
				}
			}
			if !flag {
				parsedWaves[name] += "."
			}
			for j := 1; j < (vcdFrame.getMaxValueWidth()*width_scale)/2; j++ {
				parsedWaves[name] += "."
			}
		}
	}
	return parsedWaves, parsedData
}

func (vcdFrame VCD_Struct) encodeWD(end_scale int, width_scale int) WD_Struct {
	var wavedrom WD_Struct
	waves, data := vcdFrame.parseValues(end_scale, width_scale)
	for name := range waves {
		wd_data := []string{}
		if len(data[name]) != 0 {
			wd_data = data[name]
		}
		wavedrom.Signal = append(wavedrom.Signal, WD_Signal{name, waves[name], wd_data})
	}
	return wavedrom
}

func wavedrom(w http.ResponseWriter, req *http.Request) {
	var response ResponseFrame

	defer func() {
		if panicInfo := recover(); panicInfo != nil {
			var response ResponseFrame
			response.Status_str = "error"
			response.Status_code = 400
			response.Message = fmt.Sprintf("Top-level panic: %v", panicInfo)
			w.WriteHeader(response.Status_code)
			json.NewEncoder(w).Encode(response)
		}
	}()

	reqBody, _ := ioutil.ReadAll(req.Body)
	var reqVCD VCD_Struct

	err := json.Unmarshal(reqBody, &reqVCD)

	if err != nil {
		defer func() {
			if r := recover(); r != nil {
				response.Status_str = "error"
				response.Status_code = 400
				response.Message = err.Error()
				w.WriteHeader(response.Status_code)
				json.NewEncoder(w).Encode(response)
			}
		}()
		panic("JSON parsing error")
	}

	response.Status_str = "ok"
	response.Status_code = 200
	response.Wavedrom = reqVCD.encodeWD(3, 1)
	w.WriteHeader(response.Status_code)
	json.NewEncoder(w).Encode(response)
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		fmt.Println("Server start")
		r := mux.NewRouter().StrictSlash(true)
		r.HandleFunc("/wavedrom", wavedrom).Methods("POST")

		exec.Command("fuser", "-k", "8081/tcp").Output()
		http.ListenAndServe(":8081", r)
		fmt.Println("Server stop")
	}
}
