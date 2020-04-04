package Daemon

import "sync"

//一个线程安全的取消值记录量，用于取消列表记录中在当前队列中每个各有多少次取消
type count struct {
	n  int64 //这个值最小只会到-1
	mu *sync.Mutex
}

//加一次
func (c *count) more() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.n < 0 {
		c.n = 1
	}
	c.n++
	return c.n
}

//减一次
func (c *count) less() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.n >= 0 {
		c.n--
	}
	return c.n
}

func (c *count) Get() int64 {
	return c.n
}

//取消列表
//
//队列中按照id存储每个任务的取消次数
var cancelQ = make(map[string]*count)
var cancelQMu = new(sync.RWMutex)

//取消一个在daemon中的任务
//
//在取消队列中的对应项取消次数加一
func CancelTask(id string) {
	cancelQMu.Lock()
	defer cancelQMu.Unlock()

	taskDurationMu.Lock()
	defer taskDurationMu.Unlock()

	canceln, exists := cancelQ[id]
	if !exists {
		cancelQ[id] = &count{1, new(sync.Mutex)}
	} else {
		canceln.more() //取消次数加一
	}

	durations, exists := taskDuration[id]
	if exists && len(durations) >= 1 {
		taskDuration[id] = durations[1:]
	}
	durationCached = false
}
