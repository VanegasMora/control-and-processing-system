package server

import (
	"sync"
)

type TaskQueue struct {
	mu    sync.Mutex
	tasks []Task
}

type Task struct {
	Type    string
	Payload map[string]interface{}
}

func NewTaskQueue() *TaskQueue {
	return &TaskQueue{
		tasks: make([]Task, 0),
	}
}

func (tq *TaskQueue) Enqueue(taskType string, payload map[string]interface{}) {
	tq.mu.Lock()
	defer tq.mu.Unlock()
	tq.tasks = append(tq.tasks, Task{
		Type:    taskType,
		Payload: payload,
	})
}

func (tq *TaskQueue) Dequeue() *Task {
	tq.mu.Lock()
	defer tq.mu.Unlock()
	if len(tq.tasks) == 0 {
		return nil
	}
	task := tq.tasks[0]
	tq.tasks = tq.tasks[1:]
	return &task
}
