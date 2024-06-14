package manager

import (
	"fmt"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/pankaj0812/go-taskmaster/task"
)

type Manager struct {
	Pending       queue.Queue
	TaskDb        map[string][]*task.Task
	EventDb       map[string][]*task.TaskEvent
	Workers       []string
	WorkerTaskMap map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
}

func (m *Manager) SelectWorker() {
	fmt.Println("I will select a worker to run the task")
}

func (m *Manager) UpdateTasks() {
	fmt.Println("I will update the tasks")
}

func (m *Manager) SendWork() {
	fmt.Println("I will send work to the worker")
}
