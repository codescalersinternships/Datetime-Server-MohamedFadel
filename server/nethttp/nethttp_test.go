package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDatetime(t *testing.T) {
	req := httptest.NewRequest("GET", "/datetime", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetDatetime)

	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
