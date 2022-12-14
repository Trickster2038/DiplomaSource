package main

import (
    "fmt"
    "os/exec"
    "runtime"
    //"strings"
    "os"
)

func execute() {
    out, err := exec.Command("screenfetch").Output()
    if err != nil {
        fmt.Printf("%s", err)
    }
    fmt.Println("Command Successfully Executed")
    output := string(out[:])
    fmt.Println(output)
}

func execute2() {
    out, err := exec.Command("bash", "-c", "iverilog -o adder adder_tb.v adder.v").Output()
    if err != nil {
        fmt.Printf("%s", err)
    }
    fmt.Println("Command Successfully Executed")
    output := string(out[:])
    fmt.Println(output)

    // =====

    out, err = exec.Command("bash", "-c", "vvp adder").Output()
    if err != nil {
        fmt.Printf("%s", err)
    }
    fmt.Println("Command Successfully Executed")
    output = string(out[:])
    fmt.Println(output)
}

func create_or_update(user_id string, level_id string, device_src  string, tb_src string) {
    os.MkdirAll((user_id + "/" + level_id), os.ModePerm)
    f, _ := os.Create((user_id + "/" + level_id + "/device.v"))
    f.WriteString(device_src)
    f, _ = os.Create((user_id + "/" + level_id + "/tb.v"))
    f.WriteString(tb_src)
}

func main() {
    if runtime.GOOS == "windows" {
        fmt.Println("Can't Execute this on a windows machine")
    } else {
        execute2()
        create_or_update("a", "b", "g", "h")
    }
}
