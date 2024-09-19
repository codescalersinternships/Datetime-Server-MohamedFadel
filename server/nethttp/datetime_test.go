package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetDatetime(t *testing.T) {
	t.Run("testing a GET request with JSON Accept header", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/datetime", nil)
		req.Header.Set("Accept", "application/json")
		rr := httptest.NewRecorder()

		GetDatetime(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		var response map[string]string
		err := json.Unmarshal(rr.Body.Bytes(), &response)
		if err != nil {
			t.Errorf("Failed to parse JSON response: %v", err)
		}

		_, exist := response["datetime"]
		if !exist {
			t.Errorf("JSON response does not contain 'datetime' key")
		}
	})

	t.Run("testing a GET request with plain text Accept header", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/datetime", nil)
		req.Header.Set("Accept", "text/plain")
		rr := httptest.NewRecorder()

		GetDatetime(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		if contentType := rr.Header().Get("Content-Type"); contentType != "text/plain" {
			t.Errorf("handler returned wrong content type: got %v want %v", contentType, "text/plain")
		}
	})

	t.Run("testing a not-GET request", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/datetime", nil)
		rr := httptest.NewRecorder()

		GetDatetime(rr, req)

		if status := rr.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
		}
	})
}

func TestStartServer(t *testing.T) {
	go func() {
		if err := StartServer(); err != nil && err != http.ErrServerClosed {
			t.Fatalf("error starting the server: %v", err)
		}
	}()
}
