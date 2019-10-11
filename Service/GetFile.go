package Service

import "../Model"

//通过ID获取配置文件路径
//
//ID存在则返回路径，否则返回空字符串
func GetConfigFilePath(id string) string {
	task, exists := Model.TaskList.GetTask(id)
	if !exists {
		return ""
	}
	return task.GetConfigFilePath()
}

//通过ID获取结果文件路径
//
//ID存在则返回路径，否则返回空字符串
func GetResultFilePath(id string) string {
	task, exists := Model.TaskList.GetTask(id)
	if !exists {
		return ""
	}
	return task.GetResultFilePath()
}

func GetLogFilePath(id string) string {
	task, exists := Model.TaskList.GetTask(id)
	if !exists {
		return ""
	}
	return task.GetLogFilePath()
}
