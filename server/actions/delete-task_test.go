package actions

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/azwartin/go-request-receiver/server"
	"github.com/azwartin/go-request-receiver/tasks"
)

func TestDeleteTaskSuccess(t *testing.T) {
	task := tasks.URLFetchTask{
		ID:     "0",
		Method: http.MethodGet,
		URL:    "https://google.com",
		Header: http.Header{"Accept-Language": []string{"ru_RU"}},
		Body:   "Hello",
	}

	storage := server.GetStorage()
	storage.Store(task.ID, task)
	req := httptest.NewRequest(http.MethodGet, "/delete-task?id=0", nil)
	rec := httptest.NewRecorder()
	RunDeleteTaskAction(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("Expect status 200, got %v", rec.Code)
	}

	if _, ok := storage.Load(task.ID); ok {
		t.Errorf("Expect that task removed, but found")
	}
}

func TestDeleteTaskNotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/delete-task?id=0", nil)
	rec := httptest.NewRecorder()
	RunDeleteTaskAction(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Errorf("Expect status 404, got %v", rec.Code)
	}
}
