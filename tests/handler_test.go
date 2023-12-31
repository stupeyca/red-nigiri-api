package tests

import (
	"net/http"
	"net/http/httptest"
	"red-nigiri-api/internal/server"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestIndexHandler(t *testing.T) {

	s := &server.Server{}

	router := gin.New()
	router.GET("/", s.IndexHandler)
	router.LoadHTMLGlob("../web/template/index.tmpl")

	// Create a test HTTP request
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	// Serve the HTTP request
	router.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}
