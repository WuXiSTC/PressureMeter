package Service

import (
	"../Model"
	"errors"
	"mime/multipart"
)

//新建任务
//
//先检查任务是否存在再新建任务，如果存在则报错
func CreateTask(jmx multipart.File, taskId string) error {
	if Model.TaskList.Exists(taskId) {
		return errors.New("任务已存在")
	}
	task, err := Model.Task(taskId, jmx)
	if err != nil {
		return err
	}
	return Model.TaskList.AddTask(task)
}
