package Task

import (
	"errors"
	"fmt"
	"sync"
)

//用于设置Jmeter分布式测试指令中输入的从机IP列表
var hosts [][]string
var hostsMu = new(sync.RWMutex)

func getHosts(i uint16) (list []string) {
	hostsMu.RLock()
	defer hostsMu.RUnlock()
	defer func() { recover() }()
	if len(hosts) <= int(i) {
		return nil
	}
	return hosts[i]
}

//修改Jmeter服务IP列表
//
//list[1]给线程1用、list[2]给线程2用，依此类推
func SetHosts(list [][]string) error {
	hostsMu.Lock()
	defer hostsMu.Unlock()
	if len(list) < int(conf.taskAccN) {
		return errors.New(fmt.Sprintf("请至少指定%d组IP", conf.taskAccN))
	}
	hosts = list
	return nil
}
