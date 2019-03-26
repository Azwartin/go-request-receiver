package actions

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"

	"github.com/azwartin/go-request-receiver/server"
	"github.com/azwartin/go-request-receiver/tasks"
	t "github.com/azwartin/go-request-receiver/tasks"
)

//FetchTaskAction is action, that performs request to another service, based on user params
type FetchTaskAction struct {
	BaseAction
}

//RunFetchTaskAction handle fetch task request
func RunFetchTaskAction(w http.ResponseWriter, r *http.Request) {
	action := FetchTaskAction{
		BaseAction{
			Request:  r,
			Response: w,
		},
	}

	action.Run()
}

//Run http handler
func (a *FetchTaskAction) Run() {
	tasks, err := a.getTaskFromRequest()
	if err != nil {
		log.Print(err)
		a.OutputError("Can't parse body", http.StatusBadRequest)
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(len(tasks))
	for _, task := range tasks {
		go func(task *t.URLFetchTask) {
			defer wg.Done()
			task.Execute()
			storage := server.GetStorage()
			id := ""
			//generate uniq id for task
			for {
				id = strconv.Itoa(rand.Int())
				if _, ok := storage.Load(id); !ok {
					break
				}
			}

			task.ID = id
			task.Result.ID = id
			storage.Store(id, task)
		}(task)
	}

	wg.Wait()
	a.Output(tasks)
}

func (a *FetchTaskAction) getTaskFromRequest() ([]*tasks.URLFetchTask, error) {
	body, err := a.GetBody()
	if err != nil {
		return nil, err
	}

	tasks := []*tasks.URLFetchTask{}
	err = json.Unmarshal(body, &tasks)
	if err != nil {
		return nil, fmt.Errorf("Error parsing body: %v", err)
	}

	return tasks, nil
}
