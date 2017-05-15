package main

import (
	"html"
	"fmt"
	"log"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	//router.Methods("GET").Path("/").Handler(Index)
	router.Path("/").Methods("GET").HandlerFunc(Index)
	router.Path("/api/oauth").Methods("POST").HandlerFunc(AuthCode)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func AuthCode(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	option := GoogleCode{}
	if err := decoder.Decode(&option); err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(option)
}