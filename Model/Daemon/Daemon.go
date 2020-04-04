package Daemon

import (
	"fmt"
	"gitee.com/WuXiSTC/PressureMeter/util"
	"github.com/yindaheng98/go-utility/SortedSet"
	"net"
	"sync"
	"time"
)

//任务基础接口
type TaskInterface interface {
	GetID() string                                                                  //获取任务ID
	Start(shutdownPort uint16, duration time.Duration, ipList *[]net.TCPAddr) error //启动
	Wait()                                                                          //等待
	Stop() error                                                                    //停止
}

type task struct {
	TaskInterface
	duration time.Duration
}

var taskQ chan *task                //任务队列，用以存储要执行的任务的地址
var Qn = &count{0, new(sync.Mutex)} //任务队列当前长度

//运行一个任务
func run1task(i uint16) {
	tsk := <-taskQ
	if tsk == nil {
		return
	}
	util.Log(fmt.Sprintf("Daemon %d: get task %s", i, tsk.GetID()))
	cancelQMu.RLock()
	Qn.less() //队列中任务数量-1
	cancel, exists := cancelQ[tsk.GetID()]
	cancelQMu.RUnlock()
	if exists && cancel.less() >= 0 { //如果已取消就不运行
		return
	}
	if err := tsk.Start(conf.BasePort-i, tsk.duration, getIPList(i)); err == nil { //否则就运行
		util.Log(fmt.Sprintf("Daemon %d: started task %s", i, tsk.GetID()))
		tsk.Wait()
		util.Log(fmt.Sprintf("Daemon %d: stopped task %s", i, tsk.GetID()))
	} else { //运行出错则停止
		util.LogE(err)
		err = tsk.Stop()
		util.LogE(err)
	}
	taskDurationMu.Lock()
	defer taskDurationMu.Unlock()
	durations, exists := taskDuration[tsk.GetID()]
	if exists && len(durations) >= 1 {
		taskDuration[tsk.GetID()] = durations[1:]
	}
	durationCached = false
}

var toStop = false
var stopped = make(chan uint16)
var taskDuration = make(map[string][]time.Duration)
var taskDurationMu = new(sync.RWMutex)

//将一个任务交给daemon运行
//
//入任务队列
func AddTask(tsk TaskInterface, duration time.Duration) {
	taskQ <- &task{
		TaskInterface: tsk,
		duration:      duration,
	} //入队列
	taskDurationMu.Lock()
	defer taskDurationMu.Unlock()
	durations := taskDuration[tsk.GetID()]
	durations = append(durations, duration)
	taskDuration[tsk.GetID()] = durations
	Qn.more() //队列中任务数量+1
	durationCached = false
}

//停止Daemon运行
func Stop() {
	toStop = true
	close(taskQ)
	for i := uint16(0); i < conf.TaskAccN; i++ {
		goi := <-stopped
		util.Log(fmt.Sprintf("Daemon %d stopped", goi))
	}
}

var durationCache = time.Duration(0)
var durationCached = false

func ExpectDuration() time.Duration {
	if durationCached {
		return durationCache
	}
	taskDurationMu.RLock()
	defer taskDurationMu.RUnlock()
	set := SortedSet.New(uint64(conf.TaskAccN))
	for i := uint16(0); i < conf.TaskAccN; i++ {
		set.Update(element(i), 0)
	}
	for _, durations := range taskDuration {
		for _, duration := range durations {
			set.DeltaUpdate(set.Sorted(1)[0], float64(duration))
		}
	}
	d, ds := float64(0), set.SortedAll()
	if len(ds) >= 1 {
		d, _ = set.GetWeight(ds[len(ds)-1])
	}
	durationCache = time.Duration(d)
	durationCached = true
	return durationCache
}

type element uint16

func (e element) GetName() string {
	return fmt.Sprintf("%d", e)
}
