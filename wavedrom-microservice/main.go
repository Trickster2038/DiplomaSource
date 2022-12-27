package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"sort"
)

type VCD_obj struct {
	Children []struct {
		Name string `json:"name"`
		Type struct {
			Name  string `json:"name"`
			Width int    `json:"width"`
		} `json:"type"`
		Data [][]interface{} `json:"data"`
	} `json:"children"`
}

type WD_Signal struct {
	Name string   `json:"name"`
	Wave string   `json:"wave"`
	Data []string `json:"data"`
}

type WD_obj struct {
	Signal []WD_Signal `json:"signal"`
}

type SingleValueDump struct {
	Time  int
	Value string
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

func (vcdFrame VCD_obj) getTimings() []int {
	var timeArray []int
	for _, x := range vcdFrame.Children {
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

/*
returns (waveform, datalabels)
*/
func (vcdFrame VCD_obj) parseValues(end_scale int) (map[string]string, map[string][]string) {
	var parsedData = map[string][]string{}
	var parsedWaves = map[string]string{}

	timings := vcdFrame.getTimings()
	tick_amount := findGCD(timings)
	end_time := timings[len(timings)-1] + tick_amount*end_scale
	for i := 0; i <= end_time/tick_amount; i++ {
		for _, x := range vcdFrame.Children {
			for _, y := range x.Data {
				if int(math.Round(y[0].(float64))) == i*tick_amount {
					parsedData[x.Name] = append(parsedData[x.Name], y[1].(string))
					parsedWaves[x.Name] += "="
				} else {
					parsedWaves[x.Name] += "."
				}
			}
		}
	}
	return parsedWaves, parsedData
}

func (vcdFrame VCD_obj) encodeWD(end_scale int) WD_obj {
	var wavedrom WD_obj
	waves, data := vcdFrame.parseValues(end_scale)
	for name := range waves {
		wavedrom.Signal = append(wavedrom.Signal, WD_Signal{name, waves[name], data[name]})
	}
	return wavedrom
}

func main() {
	b, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Print(err)
	}

	var vcdFrame VCD_obj
	json.Unmarshal(b, &vcdFrame)

	fmt.Println("Name: ", string(vcdFrame.Children[0].Name))
	fmt.Println("Type: ", string(vcdFrame.Children[0].Type.Name))
	fmt.Println("Width: ", int(vcdFrame.Children[0].Type.Width))
	fmt.Println("Data: ", vcdFrame.Children[0].Data[1][0], vcdFrame.Children[0].Data[1][1])

	fmt.Println("GCD: ", findGCD(vcdFrame.getTimings()))

	wave, data := vcdFrame.parseValues(3)
	fmt.Println("Parsed Values, Data: ", wave, data)
	s, _ := json.MarshalIndent(vcdFrame.encodeWD(3), "", " ")
	fmt.Println("Wavedrom ", string(s))
}
