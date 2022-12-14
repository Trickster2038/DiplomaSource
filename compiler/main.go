package main

import (
    "fmt"
    "os/exec"
    "runtime"
    //"strings"
    "os"
    //"net/http"
)

func create_or_update(user_id string, level_id string, device_src  string, tb_src string) {
    os.MkdirAll((user_id + "/" + level_id), os.ModePerm)
    f, _ := os.Create((user_id + "/" + level_id + "/device.v"))
    f.WriteString(device_src)
    f, _ = os.Create((user_id + "/" + level_id + "/tb.v"))
    f.WriteString(tb_src)
}

func compile_and_visualise(user_id string, level_id string) {
    device_path := user_id + "/" + level_id + "/device.v"
    tb_path := user_id + "/" + level_id + "/tb.v"
    out_path := user_id + "/" + level_id + "/device"

    fmt.Printf("%s", "\np1\n")
    _, err := exec.Command("bash", "-c", ("iverilog -o " + out_path + " " + device_path + " " + tb_path)).Output()
    //_, err := exec.Command("bash", "-c", ("iverilog -o  " + out_path + " " + device_path + " " + tb_path)).Output()
    if err != nil {
        fmt.Printf("%s", err)
    }
    fmt.Printf("%s", "\np2\n")

    _, err = exec.Command("bash", "-c", ("vvp " + out_path)).Output()
    if err != nil {
        fmt.Printf("%s", err)
    }
    
}

// func build(w http.ResponseWriter, req *http.Request) {
//     for name, headers := range req.Header {
//         for _, h := range headers {
//             fmt.Fprintf(w, "%v: %v\n", name, h)
//         }
//     }
//     //create_or_update(user_id, level_id, device_src, tb_src)
//     //compile_and_visualise(user, level)
// }

func main() {
    if runtime.GOOS == "windows" {
        fmt.Println("Can't Execute this on a windows machine")
    } else {
        //execute2()
        f, _ := os.ReadFile("adder.v")
        device_src := string(f)

        f, _ = os.ReadFile("adder_tb.v")
        tb_src := string(f)

        user := "user1"
        level := "level1"

        
        create_or_update(user, level, device_src, tb_src)
        compile_and_visualise(user, level)
        // fmt.Printf("%s", "gg")

        //http.HandleFunc("/build", build)

        //http.ListenAndServe(":8090", nil)
    }
}
