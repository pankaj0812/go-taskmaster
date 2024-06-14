package worker

import (
	"fmt"

	"github.com/pankaj0812/go-taskmaster/task"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]*task.Task
	TaskCount int
}

func (w *Worker) CollectStats() {
	fmt.Println("I am collecting stats")
}

func (w *Worker) RunTask() {
	fmt.Println("I will start or stop a task, after identifying the state")
}

func (w *Worker) StopTask() {
	fmt.Println("I will stop a task")
}

func (w *Worker) StartTask() {
	fmt.Println("I will start a task")
}
