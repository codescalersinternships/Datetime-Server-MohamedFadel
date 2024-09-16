package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDatetime(t *testing.T) {
	t.Run("testing a GET request", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/datetime", nil)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(GetDatetime)

		handler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)

	})

	t.Run("testing a not-GET request", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/datetime", nil)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(GetDatetime)

		handler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestStartServer(t *testing.T) {
	go func() {
		if err := StartServer(); err != nil && err != http.ErrServerClosed {
			t.Fatalf("error starting the server: %v", err)
		}
	}()
}
