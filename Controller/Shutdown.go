package Controller

import "PressureMeter/Model"

func Shutdown() error {
	if err := Model.TaskList.StopAll(); err != nil {
		return err
	}
	if err := Model.TaskList.DelAll(); err != nil {
		return err
	}
	return nil
}
