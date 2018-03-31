package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/frodonLD/GoLangRESTAPIWithGin/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var toto string

func initTestsData() {
	model.Notifications = []model.Notification{
		model.Notification{
			ID:      "1",
			Message: "Something is wrong in this world",
			Level: &model.NotificationLevel{
				Name:     "Warning",
				Bloquant: false,
			},
		},
		model.Notification{
			ID:      "2",
			Message: "This is the end of the world",
			Level: &model.NotificationLevel{
				Name:     "Critical",
				Bloquant: true,
			},
		},
	}
}

func resetData() {
	model.Notifications = []model.Notification{}
}

func setupDataTestCase(t *testing.T) func(t *testing.T) {
	initTestsData()
	return func(t *testing.T) {
		t.Log("teardown test case")
		resetData()
	}
}

func TestGetAllNotificationsHandler(t *testing.T) {
	teardownTestCase := setupDataTestCase(t)
	defer teardownTestCase(t)

	logsJSON, err := json.Marshal(&model.Notifications)
	if err != nil {
		t.Log("error:", err)
	}

	testCases := []struct {
		description      string
		httpMethod       string
		expectedHTTPCode int
		expectedBody     interface{}
	}{
		{
			"Case 200 : GET Logs, returned",
			http.MethodGet,
			http.StatusOK,
			logsJSON,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf(tc.description), func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			req, err := http.NewRequest(tc.httpMethod, "/logs", nil)
			assert.Nil(t, err)
			w := httptest.NewRecorder()
			contextTest, _ := gin.CreateTestContext(w)
			contextTest.Request = req

			GetAllNotificationsHandler(contextTest)

			assert.Equal(t, tc.expectedHTTPCode, w.Code)
			assert.Equal(t, tc.expectedBody, []byte(w.Body.String()))
		})
	}
}

func BenchmarkGetLogsHandler(b *testing.B) {
	initTestsData()
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/logs", GetAllNotificationsHandler)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/logs", nil)
	for n := 0; n < b.N; n++ {
		r.ServeHTTP(w, req)
	}
}

func BenchmarkGetLogHandler(b *testing.B) {
	initTestsData()
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/logs/:id", GetAllNotificationsHandler)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/logs/1", nil)
	for n := 0; n < b.N; n++ {
		r.ServeHTTP(w, req)
	}
}
