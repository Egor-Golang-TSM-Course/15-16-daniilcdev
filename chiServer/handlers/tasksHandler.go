package handlers

import (
	"encoding/json"
	"net/http"
	"path"
	"server-context/chiServer/services"
)

type TasksHandler struct {
	ts *services.TasksService
}

func NewTasksHandler(ts *services.TasksService) *TasksHandler {
	return &TasksHandler{ts: ts}
}

func (th *TasksHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	type payload struct {
		CreatedBy   services.OwnerId `json:"createdBy"`
		Title       string           `json:"title"`
		Description string           `json:"description"`
	}

	var p payload
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := th.ts.CreateTask(p.CreatedBy, p.Title, p.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

func (th *TasksHandler) GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	all, _ := th.ts.GetAllTasks()

	p, err := json.Marshal(struct {
		Tasks []services.TaskDto `json:"tasks"`
	}{
		Tasks: all,
	})

	if err != nil {
		http.Error(w, "serialization error", http.StatusInternalServerError)
		return
	}

	w.Write(p)
}

func (th *TasksHandler) GetUsersTasksHandler(w http.ResponseWriter, r *http.Request) {
	ownerId := r.URL.Query().Get("createdBy")
	tasks, err := th.ts.GetTasksByOwner(services.OwnerId(ownerId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	type payload struct {
		Result []services.TaskDto `json:"result"`
	}

	p, err := json.Marshal(payload{Result: tasks})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(p)
}

func (th *TasksHandler) GetTaskByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)

	task, err := th.ts.GetTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	p, err := json.Marshal(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(p)
}
