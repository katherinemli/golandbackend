package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
	"os"
)

func check(e error) {
    if e != nil {
        fmt.Println(e)
        panic(e)
    }
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Index functoin")
    indexHTML, err := ioutil.ReadFile("datachart.txt")
    check(err)
    fmt.Println(indexHTML)
    w.Write(indexHTML)
}

func main() {
    http.HandleFunc("/", Index)

    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    check(err)
}