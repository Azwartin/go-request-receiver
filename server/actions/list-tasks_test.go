package actions

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/azwartin/go-request-receiver/server"
	"github.com/azwartin/go-request-receiver/tasks"
)

func TestListTasksSuccess(t *testing.T) {
	tasks := []tasks.URLFetchTask{
		tasks.URLFetchTask{
			ID:     "0",
			Method: http.MethodGet,
			URL:    "https://google.com",
			Header: http.Header{"Accept-Language": []string{"ru_RU"}},
			Body:   "Hello",
		},
		tasks.URLFetchTask{
			ID:     "1",
			Method: http.MethodGet,
			URL:    "https://ya.ru",
			Header: http.Header{"Accept-Language": []string{"ru_RU"}},
		},
	}

	storage := server.GetStorage()
	for _, task := range tasks {
		storage.Store(task.ID, task)
	}

	req := httptest.NewRequest(http.MethodGet, "/tasks?page=1&page-size=3", nil)
	rec := httptest.NewRecorder()
	RunListTasksAction(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("Expect status 200, got %v", rec.Code)
	}

	if !strings.Contains(rec.Body.String(), tasks[0].URL) {
		t.Errorf("Expect response contains task url")
	}
}
