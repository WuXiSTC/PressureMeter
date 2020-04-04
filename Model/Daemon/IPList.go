package Daemon

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
	list = ipList[i]
	return
}

//修改Jmeter服务IP列表
func SetIPList(list [][]net.TCPAddr) error {
	ipListMu.Lock()
	defer ipListMu.Unlock()
	if len(list) < int(conf.TaskAccN) {
		return errors.New(fmt.Sprintf("请至少指定%d组IP", conf.TaskAccN))
	}
	ipList = list
	return nil
}
