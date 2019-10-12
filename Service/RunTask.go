package Service

import (
	"../Model"
	"errors"
)

func StartTask(id string) error {
	task, exists := Model.TaskList.GetTask(id)
	if !exists {
		return errors.New("任务不存在")
	}
	return task.Start()
}

func StopTask(id string) error {

	task, exists := Model.TaskList.GetTask(id)
	if !exists {
		return errors.New("任务不存在")
	}
	return task.Stop()
}
