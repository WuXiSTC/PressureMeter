package Service

import (
	"../Model"
)

func StartTask(id string) error {
	return Model.TaskList.Start(id)
}

func StopTask(id string) error {
	return Model.TaskList.Stop(id)
}
