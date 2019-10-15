package util

import "os"

//创建一个新的空文件，文件已存在就清空内容
func EmptyFile(path string) error {
	if err := DeleteFile(path); err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm) //打开文件流
	if err != nil {
		return err
	}
	LogE(f.Close())
	return nil
}

//创建一个新的空文件，文件已存在就什么都不做
func MakeFile(path string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer func() { LogE(f.Close()) }()
	return nil
}

//删除文件，文件已经没了就什么都不做返回nil
func DeleteFile(path string) error {
	if err := os.Remove(path); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return nil
}
