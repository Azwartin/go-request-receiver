package actions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Action is controller action - http request handler
type Action interface {
	Run()
}

//BaseAction is base controller action
type BaseAction struct {
	Request  *http.Request
	Response http.ResponseWriter
}

//GetBody - get body from http request as slice of byte
func (a *BaseAction) GetBody() ([]byte, error) {
	body, err := ioutil.ReadAll(a.Request.Body)
	defer a.Request.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("Error reading body: %v", err)
	}

	return body, nil
}

//OutputError - write error as plain text to writer
func (a *BaseAction) OutputError(err string, status int) {
	http.Error(a.Response, err, status)
}

//Output - write output as json to writer
func (a *BaseAction) Output(output interface{}) {
	json, err := json.Marshal(output)
	if err != nil {
		log.Printf("Error marshal output: %v", err)
		a.OutputError("Can't marshal output", http.StatusInternalServerError)
		return
	}

	a.Response.Header().Set("Content-Type", "application/json")
	a.Response.Write(json)
}
