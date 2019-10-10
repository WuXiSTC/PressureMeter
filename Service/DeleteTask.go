package Service

import "../Model"

func DeleteTask(id string) error {
	return Model.TaskList.DelTask(id)
}
