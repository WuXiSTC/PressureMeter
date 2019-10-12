package Service

import (
	"../Model"
	"github.com/kataras/iris/core/errors"
	"mime/multipart"
)

//新建任务
//
//先检查任务是否存在再新建任务
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
