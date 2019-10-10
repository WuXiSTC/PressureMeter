package Model

type taskList struct {
	tasks map[string]task
}

var TaskList = taskList{make(map[string]task)}

//插入一个任务
//
//先删除ID对应的任务再插入
func (tasklist *taskList) AddTask(tsk task) error {
	err := tasklist.DelTask(tsk.id)
	if err != nil {
		return err
	}
	tasklist.tasks[tsk.id] = tsk
	return nil
}

//按照ID获取任务
func (tasklist *taskList) GetTask(id string) task {
	return tasklist.tasks[id]
}

//按照ID删除任务
func (tasklist *taskList) DelTask(id string) error {
	tsk, exists := tasklist.tasks[id]
	if exists {
		err := tsk.Delete()
		if err != nil {
			return err
		}
		delete(tasklist.tasks, id)
	}
	return nil
}

func (tasklist *taskList) Exists(id string) bool {
	_, exists := tasklist.tasks[id]
	return exists
}
