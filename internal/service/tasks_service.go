package service

import (
	"errors"
	"fmt"
	"log"
	"tasks/internal/core"
	r "tasks/internal/repository"
)

type TaskService struct {
	Repo r.TaskReporitory
}

func NewTaskService() TaskService {
	repository := r.NewTaskRepository()
	return TaskService{Repo: repository}
}

func (s *TaskService) Add(task_message string) error {
	if task_message == "" {
		return errors.New("la tarea debe contener un mensaje")
	}
	return s.Repo.Add(task_message)
}

func (s *TaskService) Remove(id string) error {
	if id == "" {
		return errors.New("el id no es valido")
	}
	return s.Repo.Remove(id)
}

func (s *TaskService) Get(id string) (core.Task, error) {
	if id == "" {
		return core.Task{}, errors.New("el id no es valido")
	}
	return s.Repo.Get(id)
}

func (s *TaskService) Update(id, task_message string) error {
	if id == "" || task_message == "" {
		return errors.New("la tarea debe contener un mensaje y id")
	}
	return s.Repo.Update(id, task_message)
}

func (s *TaskService) List(n int) []core.Task {
	if n <= 0 {
		return []core.Task{}
	}
	tasks, err := s.Repo.List(n)
	if err != nil {
		log.Fatal(err)
		return []core.Task{}
	}
	if len(tasks) <= 0 {
		fmt.Println("no existen tareas registradas")
	}
	return tasks
}
