package repository

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"tasks/internal/core"
	"time"
)

const jsonFilePath = "/home/daniel/Escritorio/GITHUB/task_tracker_cli/tasks.json"

type TaskReporitory struct {
}

func NewTaskRepository() TaskReporitory {
	return TaskReporitory{}
}

func NewID() string {
	return fmt.Sprintf("%04d", rand.Intn(10000))
}

func (t *TaskReporitory) Add(task_message string) error {
	tasks, err := t.Read()
	if err != nil {
		return err
	}
	newTasks := core.Task{
		Id:        NewID(),
		Message:   task_message,
		Status:    "not-done",
		CreatedAt: time.Now().Format(time.RFC1123),
		UpdatedAt: time.Now().Format(time.RFC1123),
	}
	tasks = append(tasks, newTasks)
	out, err := json.MarshalIndent(tasks, " ", " ")
	if err != nil {
		return err
	}
	os.WriteFile("../../tasks.json", out, 0644)
	return nil

}

func (t *TaskReporitory) Remove(id string) error {
	tasks, err := t.Read()
	if err != nil {
		return err
	}
	tasksCopy := make([]core.Task, len(tasks)-1)

	for _, task := range tasks {
		if task.Id == id {
			continue
		}
		tasksCopy = append(tasksCopy, task)

	}
	out, err := json.MarshalIndent(tasksCopy, " ", " ")
	if err != nil {
		return err
	}
	os.WriteFile("../../tasks.json", out, 0644)
	return nil

}

func (t *TaskReporitory) Update(id, task_message string) error {
	tasks, err := t.Read()
	if err != nil {
		return err
	}
	var tasks_updated []core.Task

	for _, task := range tasks {
		if task.Id == id {
			task.Message = task_message
			task.UpdatedAt = time.Now().Format(time.RFC1123)
			tasks_updated = append(tasks_updated, task)
			continue
		}
		tasks_updated = append(tasks_updated, task)

	}
	out, err := json.MarshalIndent(tasks_updated, " ", " ")
	if err != nil {
		return err
	}
	os.WriteFile("../../tasks.json", out, 0644)
	return nil

}

// lee el archivo y converte a structuras de go el json
func (t *TaskReporitory) Read() ([]core.Task, error) {
	// load data
	var tasks []core.Task
	data, _ := os.ReadFile(jsonFilePath)
	if err := json.Unmarshal(data, &tasks); err != nil && err.Error() != "unexpected end of JSON input" {
		return []core.Task{}, err
	}
	return tasks, nil
}

// lee el archivo y retona el numero de tareas solicitado
func (t *TaskReporitory) List(status string) ([]core.Task, error) {
	// load data
	var tasks []core.Task
	data, _ := os.ReadFile(jsonFilePath)
	if err := json.Unmarshal(data, &tasks); err != nil && err.Error() != "unexpected end of JSON input" {
		return []core.Task{}, err
	}
	var tasks_by_status []core.Task
	for _, task := range tasks {
		if task.Status == status {
			tasks_by_status = append(tasks_by_status, task)
		}

	}
	return tasks_by_status, nil
}

// Marca una tarea como hecha o en progreso
func (t *TaskReporitory) Mark(id, status string) error {
	tasks, err := t.Read()
	if err != nil {
		return err
	}
	var tasks_updated []core.Task

	for _, task := range tasks {
		if task.Id == id {
			task.Status = status
			task.UpdatedAt = time.Now().Format(time.RFC1123)
			tasks_updated = append(tasks_updated, task)
			continue
		}
		tasks_updated = append(tasks_updated, task)

	}
	out, err := json.MarshalIndent(tasks_updated, " ", " ")
	if err != nil {
		return err
	}
	os.WriteFile("../../tasks.json", out, 0644)
	return nil
}
