package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupHealthTestCase(t *testing.T) func(t *testing.T) {
	return func(t *testing.T) {
		t.Log("teardown test case")
	}
}

func TestHealthCheckHandler(t *testing.T) {
	teardownTestCase := setupHealthTestCase(t)
	defer teardownTestCase(t)

	testCases := []struct {
		description      string
		httpMethod       string
		expectedHTTPCode int
		expectedHeaders  map[string][]string
		expectedBody     interface{}
	}{
		{
			"Case 200 : GET Health, returned",
			http.MethodGet,
			http.StatusOK,
			nil,
			`{"alive":true}`,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf(tc.description), func(t *testing.T) {

			gin.SetMode(gin.TestMode)
			req, err := http.NewRequest(tc.httpMethod, "/_health", nil)
			assert.Nil(t, err)
			w := httptest.NewRecorder()
			contextTest, _ := gin.CreateTestContext(w)
			contextTest.Request = req

			HealthCheck(contextTest)

			assert.Equal(t, tc.expectedHTTPCode, w.Code)
			assert.Equal(t, tc.expectedBody, w.Body.String())
		})
	}
}

func BenchmarkHealthCheckHandler(b *testing.B) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	//r.Use(gin.Recovery())
	r.GET("/_health", HealthCheck)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/_health", nil)

	for n := 0; n < b.N; n++ {
		r.ServeHTTP(w, req)
	}
}
