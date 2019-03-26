package actions

import (
	"net/http"

	"github.com/azwartin/go-request-receiver/server"
	"github.com/azwartin/go-request-receiver/server/helpers"
)

//ListTasksAction is action, that returns list of tasks in service
type ListTasksAction struct {
	BaseAction
}

//RunListTasksAction handle fetch task request
func RunListTasksAction(w http.ResponseWriter, r *http.Request) {
	action := ListTasksAction{
		BaseAction{
			Request:  r,
			Response: w,
		},
	}

	action.Run()
}

//Run http handler
func (a *ListTasksAction) Run() {
	pagination := helpers.Pagination{}
	pagination.Load(a.Request)
	storage := server.GetStorage()
	items := storage.LoadRange(pagination.GetLimit(), pagination.GetOffset())
	a.Output(items)
}
