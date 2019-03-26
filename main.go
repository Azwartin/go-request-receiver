package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/azwartin/go-request-receiver/server"
	"github.com/azwartin/go-request-receiver/server/actions"
)

func main() {
	config := parseServerConfig()
	initHandlers()
	err := http.ListenAndServe(config.Address+":"+strconv.FormatUint(uint64(config.Port), 10), nil)
	log.Fatalf("Server error: %v", err)
}

//Parse server config
func parseServerConfig() server.Config {
	address := flag.String("address", "", "Sets the address on which server will accept request")
	port := flag.Uint("port", 1048, "Sets the port on which server will accept request")
	flag.Parse()
	return server.Config{Address: *address, Port: *port}
}

//Initialization of http request handlers
func initHandlers() {
	//fetch tasks
	http.HandleFunc("/fetch-tasks", actions.RunFetchTaskAction)
	//delete task
	http.HandleFunc("/delete-task", actions.RunDeleteTaskAction)
	//list of tasks
	http.HandleFunc("/tasks", actions.RunListTasksAction)
}
