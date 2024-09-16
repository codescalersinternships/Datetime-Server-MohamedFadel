package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func GetDatetime(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Invalid http method, only GET is allowed", http.StatusBadRequest)
		return
	}

	t := time.Now().Format(time.RFC3339)
	fmt.Fprint(w, t)
}

func StartServer() error {
	http.HandleFunc("/datetime", GetDatetime)
	return http.ListenAndServe(":8000", nil)
}

func main() {
	if err := StartServer(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
