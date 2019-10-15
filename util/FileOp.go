package util

import "os"

//创建一个新的空文件，或者清空文件
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

//删除文件
func DeleteFile(path string) error {
	if err := os.Remove(path); err != nil {
		if !os.IsExist(err) {
			return nil
		}
		return err
	}
	return nil
}
