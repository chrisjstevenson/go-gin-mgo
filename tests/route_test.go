package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/chrisjstevenson/go-gin-mgo/handlers/things"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestList(t *testing.T) {
	router := gin.New()
	router.GET("/testList", things.List)

	req, _ := http.NewRequest("GET", "/stuff", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.NotNil(t, resp.Body.String())
}

func TestCreate(t *testing.T) {
	router := gin.New()
	router.GET("/testCreate", things.Create)

	json := []byte(`{"animal":"furby"}`)


	req, _ := http.NewRequest("POST", "/stuff", bytes.NewBuffer(json))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.NotNil(t, resp.Body.String())
}