package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func GetDatetime(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC3339)
	fmt.Fprint(w, t)
}

func main() {
	http.HandleFunc("/datetime", GetDatetime)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
