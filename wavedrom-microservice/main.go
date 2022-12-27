package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type VCD_obj struct {
	Children []struct {
		Name string `json:"name"`
		// Data []struct {

		// } `json:"data"`
	} `json:"children"`
	// Children []byte `json:"children"`
	Gg string `json:"gg"`
}

func main() {
	b, err := os.ReadFile("data.json") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	// fmt.Println(b) // print the content as 'bytes'

	// s := string(b) // convert content to a 'string'

	// fmt.Println(s) // print the content as a 'string'

	var VcdFrame VCD_obj
	json.Unmarshal(b, &VcdFrame)

	fmt.Println(string(VcdFrame.Children[0].Name))
}
