package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetDatetime(t *testing.T) {
	router := gin.Default()
	router.GET("/datetime", GetDatetime)
	t.Run("testing a GET request", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/datetime", nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("testing a not-GET request", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/datetime", nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}

func TestStartServer(t *testing.T) {
	go func() {
		if err := StartServer(); err != nil && err != http.ErrServerClosed {
			t.Fatalf("error starting the server: %v", err)
		}
	}()
}
