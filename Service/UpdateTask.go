package Service

import (
	"../Model"
	"mime/multipart"
)

//更新任务
func UpdateTask(jmx multipart.File, taskId string) error {
	task, err := Model.Task(taskId, jmx)
	if err != nil {
		return err
	}
	err = Model.TaskList.AddTask(*task)
	if err != nil {
		return err
	}
	return nil
}
