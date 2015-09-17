package server

import (
    "encoding/json"
    "net/http"

    "github.com/julienschmidt/httprouter"
    "github.com/chreble/todo/task"
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

// ListTask handles GET requests on /tasks
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
    if err = json.NewDecoder(req.Body).Decode(&tsk); err != nil {
        w.WriteHeader(400)
        return
    }
    t := tasks.Create(tsk)
    if err != nil {
        w.WriteHeader(400)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)

    json.NewEncoder(w).Encode(t)
}