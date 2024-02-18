package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/barbarosaffan/zort-backend/config"
	"github.com/barbarosaffan/zort-backend/repository/postgresql"
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

	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	db, err := postgresql.NewDB(cfg)

	if err != nil {
		log.Fatal("Failed to connect to DB: ", err)
	}

	defer db.Close()
}
