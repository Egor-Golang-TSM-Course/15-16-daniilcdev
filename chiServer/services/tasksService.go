package services

import (
	"errors"
	"fmt"
)

type OwnerId string

type TaskDto struct {
	Id          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	CreatedBy   OwnerId `json:"createdBy"`
}

type TasksService struct {
	storage KeyValueStorage[OwnerId, []Task]
}

type Task struct {
	id          string
	title       string
	description string
}

func NewTasksService(storage KeyValueStorage[OwnerId, []Task]) *TasksService {
	return &TasksService{
		storage: storage,
	}
}

func (ts *TasksService) CreateTask(ownerId OwnerId,
	title string,
	description string) (TaskDto, error) {
	pool, ok := ts.storage.GetValue(ownerId)

	if !ok {
		pool = make([]Task, 0, 4)
	}

	newTask := Task{id: randomString(6), title: title, description: description}
	pool = append(pool, newTask)
	ts.storage.SetValue(ownerId, pool)

	return TaskDto{
		Id:          newTask.id,
		Title:       newTask.title,
		Description: newTask.description,
		CreatedBy:   ownerId,
	}, nil
}

func (ts *TasksService) GetAllTasks() ([]TaskDto, error) {
	all := ts.storage.All()

	r := make([]TaskDto, 0, 4)
	for ownerId, pool := range all {
		for _, v := range pool {
			r = append(r, TaskDto{
				Id:          v.id,
				Title:       v.title,
				Description: v.description,
				CreatedBy:   ownerId,
			})
		}
	}

	return r, nil
}

func (ts *TasksService) GetTask(id string) (TaskDto, error) {
	all := ts.storage.All()

	for ownerId, pool := range all {
		for _, task := range pool {
			if task.id != id {
				continue
			}

			return TaskDto{
				Id:          task.id,
				Title:       task.title,
				Description: task.description,
				CreatedBy:   ownerId,
			}, nil
		}
	}

	return TaskDto{}, fmt.Errorf("task with id %s not found", id)
}

func (ts *TasksService) GetTasksByOwner(ownerId OwnerId) ([]TaskDto, error) {
	userTasks, ok := ts.storage.GetValue(ownerId)

	if !ok {
		return nil, errors.New("owner not found")
	}

	r := make([]TaskDto, len(userTasks))
	for i, v := range userTasks {
		r[i] = TaskDto{
			Id:          v.id,
			Title:       v.title,
			Description: v.description,
			CreatedBy:   ownerId,
		}
	}

	return r, nil
}
