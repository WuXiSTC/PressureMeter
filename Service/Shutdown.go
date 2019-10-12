package Service

import "../Model"

func Shutdown() {
	Model.TaskList.StopAll()
	Model.TaskList.DelAll()
}
