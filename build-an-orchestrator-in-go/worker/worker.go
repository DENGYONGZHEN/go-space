package worker

import (
	"build-an-orchestrator-in-go/task"
	"fmt"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]*task.Task
	TaskCount int
}

func (w *Worker) CollectStates() {
	fmt.Println("I will collect status")
}

func (w *Worker) RunTask() {
	fmt.Println("I will start or stop a task")
}

func (w *Worker) StartTask() {
	fmt.Println("I will start a task")

}

func (w *Worker) StopTask() {
	fmt.Println("I will stop a task")

}
