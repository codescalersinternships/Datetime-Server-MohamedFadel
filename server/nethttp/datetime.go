package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

func GetDatetime(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	t := time.Now().Format(time.RFC3339)

	acceptHeader := r.Header.Get("Accept")
	if !strings.Contains(acceptHeader, "application/json") {
		w.Header().Set("Content-Type", "text/plain")
		_, err := w.Write([]byte(t))
		if err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"datetime": t}); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}
}

func StartServer() error {
	http.HandleFunc("/datetime", GetDatetime)
	return http.ListenAndServe(":8000", nil)
}

func main() {
	if err := StartServer(); err != nil {
		log.Fatal(err)
	}
}
