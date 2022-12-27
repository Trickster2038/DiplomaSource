package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type VCD_obj struct {
	Children []struct {
		Name string `json:"name"`
		Type struct {
			Name  string `json:"name"`
			Width int    `json:"width"`
		} `json:"type"`
		Data []interface{} `json:"data"`
	} `json:"children"`
}

type SingleValueDump struct {
	Time  int
	Value string
}

func main() {
	b, err := os.ReadFile("data.json") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	var VcdFrame VCD_obj
	json.Unmarshal(b, &VcdFrame)

	fmt.Println("Name: ", string(VcdFrame.Children[0].Name))
	fmt.Println("Type: ", string(VcdFrame.Children[0].Type.Name))
	fmt.Println("Width: ", int(VcdFrame.Children[0].Type.Width))
	s := VcdFrame.Children[0].Data[1].([]interface{})
	fmt.Println("Time: ", s[0])
	fmt.Println("Value: ", s[1])
}
