package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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

		if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
			t.Errorf("handler returned wrong content type: got %v want %v", contentType, "application/json")
		}

		var response map[string]string
		err := json.Unmarshal(rr.Body.Bytes(), &response)
		if err != nil {
			t.Errorf("Failed to parse JSON response: %v", err)
		}

		datetime, exist := response["datetime"]
		if !exist {
			t.Errorf("JSON response does not contain 'datetime' key")
		}

		_, err = time.Parse(time.RFC3339, datetime)
		if err != nil {
			t.Errorf("Datetime is not in RFC3339 format: %v", err)
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

		_, err := time.Parse(time.RFC3339, rr.Body.String())
		if err != nil {
			t.Errorf("Response body is not in RFC3339 format: %v", err)
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
	errChan := make(chan error, 1)

	go func() {
		if err := StartServer(); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
		close(errChan)
	}()

	select {
	case err := <-errChan:
		if err != nil {
			t.Fatalf("error starting the server: %v", err)
		}
	case <-time.After(2 * time.Second):
	}
}
