package Daemon

import (
	"errors"
	"fmt"
	"net"
)

//用于设置Jmeter分布式测试指令中输入的从机IP列表
var IPList []*[]net.TCPAddr

func getIPList(i uint16) (list *[]net.TCPAddr) {
	defer func() { recover() }()
	list = IPList[i]
	return
}

//修改Jmeter服务IP列表
func SetIPList(list [][]net.TCPAddr) error {
	if len(list) < int(conf.TaskAccN) {
		return errors.New(fmt.Sprintf("请至少指定%d组IP", conf.TaskAccN))
	}
	for i := 0; i < int(conf.TaskAccN); i++ {
		IPList[i] = &list[i]
	}
	return nil
}
