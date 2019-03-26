package actions

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/azwartin/go-request-receiver/tasks"
)

func TestFetchTasksSuccess(t *testing.T) {
	tasks := []tasks.URLFetchTask{
		tasks.URLFetchTask{
			Method: http.MethodGet,
			URL:    "https://google.com",
			Header: http.Header{"Accept-Language": []string{"ru_RU"}},
			Body:   "Hello",
		},
		tasks.URLFetchTask{
			Method: http.MethodGet,
			URL:    "https://ya.ru",
			Header: http.Header{"Accept-Language": []string{"ru_RU"}},
		},
	}
	json, _ := json.Marshal(tasks)
	req := httptest.NewRequest(http.MethodPost, "/fetch-tasks", bytes.NewReader(json))
	rec := httptest.NewRecorder()
	RunFetchTaskAction(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("Expect status 200, got %v", rec.Code)
	}
}

func TestFetchTasksError(t *testing.T) {
	body := tasks.URLFetchResult{
		StatusCode: 300,
	}
	json, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/fetch-tasks", bytes.NewReader(json))
	rec := httptest.NewRecorder()
	RunFetchTaskAction(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expect status 400, got %v", rec.Code)
	}
}
