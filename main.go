package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

type chartData struct {
	Values []int `json:"value"`
}
type address struct {
	Id       int     `json:"id"`
	Lat      float64 `json:"lat"`
	Long     float64 `json:"long"`
	Location string  `json:"location"`
}

type allAddress []address

func Index(w http.ResponseWriter, r *http.Request) {
	nums, err := readFile("datachart.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(nums)
	resProcess := &chartData{Values: nums}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resProcess)
}
func Points(w http.ResponseWriter, r *http.Request) {
	allAddress := readFileLatLong("Point_Of_Interest.txt")
	//resProcess := &chartData{Values: nums}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(allAddress)
}
func Data(w http.ResponseWriter, r *http.Request) {
	nums, err := readFile("datamodemchart2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(nums)
	resProcess := &chartData{Values: nums}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resProcess)
}
func readFile(fname string) (nums []int, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	// Assign cap to avoid resize on every append.
	nums = make([]int, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		// Atoi better suits the job when we know exactly what we're dealing
		// with. Scanf is the more general option.
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}
func readFileLatLong(fname string) allAddress {
	var addressSelected allAddress
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil
	}

	lines := strings.Split(string(b), "\n")

	// Assign cap to avoid resize on every append.
	for idx, l := range lines {
		var addressFinal address
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		geoElem := strings.Split(string(l), ",")
		addressFinal.Id = idx
		latfloat, _ := strconv.ParseFloat(geoElem[0], 64)
		longfloat, _ := strconv.ParseFloat(geoElem[1], 64)
		addressFinal.Lat = latfloat
		addressFinal.Long = longfloat
		addressFinal.Location = strings.TrimRight(geoElem[2], "\r")
		// Atoi better suits the job when we know exactly what we're dealing
		// with. Scanf is the more general option.
		//n, err := strconv.Atoi(l)
		//if err != nil { return nil, err }
		addressSelected = append(addressSelected, addressFinal)
	}

	return addressSelected
}
func createRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entre")
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		fmt.Println("id is missing in parameters")
	}
	fmt.Println(`id := `, id)
	fmt.Println(`r.Body := `, r.Body)

	w.WriteHeader(http.StatusOK)
	jsonResponse, jsonError := json.Marshal(id)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}
	w.Write(jsonResponse)
}
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api", Index).Methods("GET")
	router.HandleFunc("/points", Points).Methods("GET")
	router.HandleFunc("/data", Data).Methods("GET")
	router.HandleFunc("/createRoute/{id}", createRouter).Methods("POST")
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://ornate-shortbread-20015a.netlify.app", "http://localhost:8080", "http://127.0.0.1:5500", "https://wondrous-dango-4bd51e.netlify.app", "https://magenta-dusk-9af42a.netlify.app"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)
	port := os.Getenv(("PORT"))
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
