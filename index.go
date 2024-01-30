package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", createRouter)
    http.HandleFunc("/foo", getOneEvent)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)

    if err != nil {
        panic(err)
    }
}
func createRouter(res http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(res, "hello, world")

}
func getOneEvent(res http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(res, "{\"foo\":\"baz\"}")
}