package Daemon

//任务基础接口
type TaskInterface interface {
	GetID() string       //获取任务ID
	Start(thread uint16) //指定任务运行线程编号，启动
	Wait()               //等待
	Stop(thread uint16)  //停止
}

type TaskState int

const (
	NOTEXISTS TaskState = -1
	QUEUEING  TaskState = 0
	RUNNING   TaskState = 1
)

//获取任务状态
func GetState(id string) TaskState {
	runningsMu.RLock()
	defer runningsMu.RUnlock()
	for _, task := range runnings {
		if task != nil && task.GetID() == id {
			return RUNNING
		}
	}
	if queue.Exists(id) {
		return QUEUEING
	}
	return NOTEXISTS
}
