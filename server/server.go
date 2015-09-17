package server

import (
	"encoding/json"
	"net/http"

	"github.com/chreble/todo/task"
	"github.com/julienschmidt/httprouter"
)

var tasks task.Tasks

func init() {
	tasks = task.Tasks{}
}

// RegisterHandlers registers httprouter handlers
func RegisterHandlers() *httprouter.Router {
	rt := httprouter.New()
	rt.GET("/tasks", ListTasks)
	rt.POST("/tasks", CreateTask)
	return rt
}

// ListTasks handles GET requests on /tasks
func ListTasks(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	res := tasks.All()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	json.NewEncoder(w).Encode(res)
}

// CreateTask handles POST request
func CreateTask(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	tsk := task.Task{}
	var err error

	// Unmarshal JSON to Task Object
	if err = json.NewDecoder(req.Body).Decode(&tsk); err != nil {
		w.WriteHeader(400)
		return
	}

	// Create task
	t := tasks.Create(tsk)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	json.NewEncoder(w).Encode(t)
}
