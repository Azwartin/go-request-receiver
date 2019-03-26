package actions

import (
	"net/http"

	"github.com/azwartin/go-request-receiver/server"
)

//DeleteTasksAction is action, that delete task by id
type DeleteTasksAction struct {
	BaseAction
}

//RunDeleteTaskAction handle delete task by id request
func RunDeleteTaskAction(w http.ResponseWriter, r *http.Request) {
	action := DeleteTasksAction{
		BaseAction{
			Request:  r,
			Response: w,
		},
	}

	action.Run()
}

//Run http handler
func (a *DeleteTasksAction) Run() {
	id := a.Request.FormValue("id")
	if id == "" {
		a.OutputError("ID required", http.StatusBadRequest)
		return
	}

	storage := server.GetStorage()
	if _, ok := storage.Load(id); !ok {
		a.OutputError("Task not found", http.StatusNotFound)
		return
	}

	storage.Delete(id)
	a.Output("OK")
}
