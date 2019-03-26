package tasks

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Task interface - base task interface.
// It has one method `Execute` which execute task and returns TaskResult
type Task interface {
	Execute() TaskResult
}

// TaskResult interface - base task result interface
type TaskResult interface{}

//URLFetchTask is task to request url with method and params and returns UrlFetchResult
type URLFetchTask struct {
	ID     string          `json:"id"`
	Method string          `json:"method"`
	URL    string          `json:"url"`
	Header http.Header     `json:"headers"`
	Body   string          `json:"body"`
	Result *URLFetchResult `json:"result"`
}

//prepareRequest prepares http request for task and returns it or error
func (task *URLFetchTask) prepareRequest() (*http.Request, error) {
	request, err := http.NewRequest(task.Method, task.URL, strings.NewReader(task.Body))
	if err != nil {
		return nil, err
	}

	request.Header = task.Header
	return request, nil
}

//Execute fetch response with given params, save result at field result and returns it as URLFetchResult struct
func (task *URLFetchTask) Execute() {
	result := URLFetchResult{
		ID:         task.ID,
		StatusCode: http.StatusInternalServerError,
	}
	task.Result = &result

	request, err := task.prepareRequest()
	if err != nil {
		result.Error = fmt.Sprintf("Can't create request: %v", err)
		return
	}

	client := http.Client{
		Timeout: time.Second * 10,
	}

	response, err := client.Do(request)
	if err != nil {
		result.Error = fmt.Sprintf("Can't perform request: %v", err)
		return
	}

	defer response.Body.Close()
	result.StatusCode = response.StatusCode
	result.Header = response.Header
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		result.Error = fmt.Sprintf("Can't perform request: %v", err)
		return
	}

	result.Body = string(body)
}

//URLFetchResult struct to represent service response at URLFetchTask
type URLFetchResult struct {
	ID         string      `json:"id"`
	StatusCode int         `json:"statusCode"`
	Header     http.Header `json:"headers"`
	Body       string      `json:"body"`
	Error      string      `json:"error"`
}
