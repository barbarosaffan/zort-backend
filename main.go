package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Zort struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var zorts = []Zort{
	{ID: "1", Title: "Zort 1", Body: "This is the body of Zort 1"},
	{ID: "2", Title: "Zort 2", Body: "This is the body of Zort 2"},
}

func getZorts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(zorts)
}

func main() {
	http.HandleFunc("/zorts", getZorts)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
