package main

import (
	"encoding/json"
	"github.com/Mohllal/go-challenge/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/ping", nil)
	router.ServeHTTP(w, req)

	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)

	response := utils.PingResponse{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	// grab the value & whether or not it exists
	value := response.Message

	// make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.Equal(t, "pong", value)
}
