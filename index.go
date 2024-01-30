package main

import (
	"log"
	"fmt"
	"net/http"
	"os"
	"io"
	"github.com/rs/cors"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", getOneEvent).Methods("GET")
	router.HandleFunc("/createRoute/{id}", createRouter).Methods("POST")
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "https://ornate-shortbread-20015a.netlify.app"},
		AllowCredentials: true,
	})
	io.WriteString(w, "Hello, getOneEvent!\n")
	//os.Setenv("PORT", "3000")
	handler := c.Handler(router)
	port := os.Getenv(("PORT"))
	fmt.Println("port:", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
func createRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entre")
}
func getOneEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entre")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	/* json.NewEncoder(w).Encode(allAddress) */
}