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
		w.Write([]byte(t))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"datetime": t})
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
