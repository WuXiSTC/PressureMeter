package Service

import (
	"../Model"
	"github.com/kataras/iris/core/errors"
	"mime/multipart"
)

//新建任务
func CreateTask(jmx multipart.File, taskId string) error {
	if Model.TaskList.Exists(taskId) {
		return errors.New("此任务已存在")
	}
	return UpdateTask(jmx, taskId)
}
