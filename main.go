package main

import (
	"log"
	"net/http"
	"todolist/task"
)

//Main function, creates routes and launchs server
func main() {
	router := task.NewRouter()
	log.Fatal(http.ListenAndServe(":9000", router))
}
