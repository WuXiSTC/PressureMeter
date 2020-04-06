package Controller

import "gitee.com/WuXiSTC/PressureMeter/Model"

func Shutdown() error {
	Model.TaskList.StopAll()
	if err := Model.TaskList.DelAll(); err != nil {
		return err
	}
	return nil
}
