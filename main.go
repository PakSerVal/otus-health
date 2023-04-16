package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type healthStatus struct {
	Status string `json:"status"`
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		// Создаем ответ в формате JSON
		response := healthStatus{"OK"}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	})

	http.HandleFunc("/readiness", func(w http.ResponseWriter, r *http.Request) {
		// Создаем ответ в формате JSON
		response := healthStatus{"OK"}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	})

	log.Printf("Server started on port %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
