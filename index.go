package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", getDataChart)
    http.HandleFunc("/foo", getOneEvent)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)

    if err != nil {
        panic(err)
    }
}
func getDataChart(res http.ResponseWriter, r *http.Request) {
	var letter_goodness = [...]float32 {.0817, .0149, .0278, .0425, .1270, .0223, .0202, .0609, .0697, .0015, .0077, .0402, .0241, .0675, .0751, .0193, .0009, .0599, .0633, .0906, .0276, .0098, .0236, .0015, .0197, .0007 }
	fmt.Fprintln(res, letter_goodness)

}
func getOneEvent(res http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(res, "{\"foo\":\"baz\"}")
}