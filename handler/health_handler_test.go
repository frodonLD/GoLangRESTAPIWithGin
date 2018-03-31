package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func setupHealthTestCase(t *testing.T) func(t *testing.T) {
	return func(t *testing.T) {
		t.Log("teardown test case")
	}
}

func TestHealthCheckHandler(t *testing.T) {
	teardownTestCase := setupHealthTestCase(t)
	defer teardownTestCase(t)

	expectedBody := `{"alive":true}`
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
			expectedBody,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf(tc.description), func(t *testing.T) {

			req, err := http.NewRequest(tc.httpMethod, "/api/v1/_health", nil)
			assert.Nil(t, err)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			respRes := w.Result()
			returnedHTTPCode := respRes.StatusCode

			assert.Equal(t, tc.expectedHTTPCode, returnedHTTPCode)

			for expectedHeaderKey, expectedHeaderValue := range tc.expectedHeaders {
				got := respRes.Header[expectedHeaderKey]
				assert.ElementsMatch(t, got, expectedHeaderValue, "testing header %s", expectedHeaderKey)
			}
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
