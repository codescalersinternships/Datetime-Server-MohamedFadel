package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestGetDatetime(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/datetime", GetDatetime)

	t.Run("testing a GET request with JSON Accept header", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/datetime", nil)
		req.Header.Set("Accept", "application/json")
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		if contentType := rr.Header().Get("Content-Type"); !strings.Contains(contentType, "application/json") {
			t.Errorf("handler returned wrong content type: got %v want application/json", contentType)
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

		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		if contentType := rr.Header().Get("Content-Type"); !strings.Contains(contentType, "text/plain") {
			t.Errorf("handler returned wrong content type: got %v want text/plain", contentType)
		}

		_, err := time.Parse(time.RFC3339, strings.TrimSpace(rr.Body.String()))
		if err != nil {
			t.Errorf("Response body is not in RFC3339 format: %v", err)
		}
	})

	t.Run("testing a not-GET request", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/datetime", nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
		}
	})
}

func TestStartServer(t *testing.T) {
	errChan := make(chan error, 1)

	go func() {
		if err := StartServer(); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	select {
	case err := <-errChan:
		t.Fatalf("error starting the server: %v", err)
	case <-time.After(2 * time.Second):
	}
}
