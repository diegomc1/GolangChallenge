package task

import (
	"net/http"
	"todolist/logger"

	"github.com/gorilla/mux"
)

//Pointer to controller struct
var controller = &Controller{Repository: Repository{}}

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes defines the list of routes of our API
type Routes []Route

var routes = Routes{
	Route{
		"AddTask",
		"POST",
		"/AddTask",
		controller.AddTask,
	},
	Route{
		"UpdateTask",
		"PUT",
		"/UpdateTask",
		controller.UpdateTask,
	},
	Route{
		"DeleteTask",
		"DELETE",
		"/DeleteTask",
		controller.DeleteTask,
	},
	Route{
		"DropDatabase",
		"POST",
		"/dropDatabase",
		controller.DropDatabase,
	},
	Route{
		"DropCollection",
		"POST",
		"/dropCollection",
		controller.DropCollection,
	},
	Route{
		"GetEcho",
		"POST",
		"/Echo",
		controller.Echo,
	},
}

//NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	//This for loop gives the logger function args
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
