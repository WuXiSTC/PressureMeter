package Daemon

import "sync"

//一个线程安全的取消值记录量，用于取消列表记录中在当前队列中每个各有多少次取消
type count struct {
	n  uint64
	mu *sync.RWMutex
}

//加一次
func (c *count) more() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

//减一次
func (c *count) less() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.n > 0 {
		c.n--
	}
}
func (c *count) get() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.n
}

//取消列表
//
//队列中按照id存储每个任务的取消次数
var cancelQ = make(map[string]*count)

//取消一个在daemon中的任务
//
//在取消队列中的对应项取消次数加一
func CancelTask(id string) {
	canceln, exists := cancelQ[id]
	if !exists {
		cancelQ[id] = &count{1, new(sync.RWMutex)}
	} else {
		canceln.more() //取消次数加一
	}
}
