package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/chrisjstevenson/go-gin-mgo/handlers/things"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestList(t *testing.T) {
	router := gin.New()
	router.GET("/testList", things.List)

	req, _ := http.NewRequest("GET", "/foo", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.NotNil(t, resp.Body.String())
}
