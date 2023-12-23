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
	tasks map[OwnerId][]task
}

type task struct {
	id          string
	title       string
	description string
}

func NewTasksService() *TasksService {
	return &TasksService{
		tasks: make(map[OwnerId][]task),
	}
}

func (ts *TasksService) CreateTask(ownerId OwnerId,
	title string,
	description string) (TaskDto, error) {
	pool, ok := ts.tasks[ownerId]

	if !ok {
		pool = make([]task, 0, 4)
	}

	newTask := task{id: randomString(6), title: title, description: description}
	pool = append(pool, newTask)
	ts.tasks[ownerId] = pool

	return TaskDto{
		Id:          newTask.id,
		Title:       newTask.title,
		Description: newTask.description,
		CreatedBy:   ownerId,
	}, nil
}

func (ts *TasksService) GetAllTasks() ([]TaskDto, error) {
	r := make([]TaskDto, 0, 4)
	for ownerId, pool := range ts.tasks {
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
	for ownerId, pool := range ts.tasks {
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
	userTasks, ok := ts.tasks[ownerId]

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
