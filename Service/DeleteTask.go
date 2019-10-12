package Service

import "../Model"

//删除任务
//
//删除成功返回nil，其他情况返回错误信息
func DeleteTask(id string) error {
	return Model.TaskList.DelTask(id)
}
