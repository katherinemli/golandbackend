package main

import (
    "fmt"
    "io/ioutil"
	"encoding/json"
    "net/http"
	"strconv"
    "strings"
	"os"
)
type chartData struct {
    Values []int `json:"value"`
}
func check(e error) {
    if e != nil {
        fmt.Println(e)
        panic(e)
    }
}

func Index(w http.ResponseWriter, r *http.Request) {
    nums, err := readFile("datachart.txt")
    if err != nil { panic(err) }
    fmt.Println(nums)
    resProcess := &chartData{Values: nums}
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(resProcess)
}
func readFile(fname string) (nums []int, err error) {
    b, err := ioutil.ReadFile(fname)
    if err != nil { return nil, err }

    lines := strings.Split(string(b), "\n")
    // Assign cap to avoid resize on every append.
    nums = make([]int, 0, len(lines))

    for _, l := range lines {
        // Empty line occurs at the end of the file when we use Split.
        if len(l) == 0 { continue }
        // Atoi better suits the job when we know exactly what we're dealing
        // with. Scanf is the more general option.
        n, err := strconv.Atoi(l)
        if err != nil { return nil, err }
        nums = append(nums, n)
    }

    return nums, nil
}
func main() {
    http.HandleFunc("/", Index)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    //err := http.ListenAndServe(":8080", nil)
    check(err)
	
}