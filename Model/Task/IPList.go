package Task

import (
	"errors"
	"fmt"
	"net"
	"sync"
)

//用于设置Jmeter分布式测试指令中输入的从机IP列表
var ipList [][]net.TCPAddr
var ipListMu = new(sync.RWMutex)

func getIPList(i uint16) (list []net.TCPAddr) {
	ipListMu.RLock()
	defer ipListMu.RUnlock()
	defer func() { recover() }()
	if len(ipList) <= int(i) {
		return nil
	}
	return ipList[i]
}

//修改Jmeter服务IP列表
//
//list[1]给线程1用、list[2]给线程2用，依此类推
func SetIPList(list [][]net.TCPAddr) error {
	ipListMu.Lock()
	defer ipListMu.Unlock()
	if len(list) < int(conf.taskAccN) {
		return errors.New(fmt.Sprintf("请至少指定%d组IP", conf.taskAccN))
	}
	ipList = list
	return nil
}
