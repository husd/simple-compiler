package util

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/**
 * 这个类是文件的工具类
 * @author hushengdong
 */

// ReadLine 按行读取文件
func ReadLine(filePth string, callback func([]byte)) error {
	f, err := os.Open(filePth)
	if err != nil {
		return err
	}
	defer f.Close()
	bfRd := bufio.NewReader(f)
	for {
		line, err := bfRd.ReadBytes('\n')
		callback(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}

// 是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// 是否目录
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 是否文件
func IsFile(path string) bool {
	return !IsDir(path)
}

//递归读取目录下的所有文件
func Walk(name string) *list.List {
	res := list.New()
	exist := Exists(name)
	if !exist {
		panic("文件或者目录不存在，请检查 " + name)
	}
	stackString := DefaultStack{}
	stackString.Init()
	stackString.Push(name)
	pathSep := string(os.PathSeparator)
	for !stackString.IsEmpty() {
		e := stackString.Pop()
		path := fmt.Sprintf("%v", e)
		//1 判断是文件或者目录
		isDir := IsDir(path)
		if isDir {
			//目录,列出目录下的所有的文件
			items, err := ioutil.ReadDir(path)
			if err != nil {
				panic("读取目录错误，请检查 " + path)
			}
			for _, f := range items {
				stackString.Push(path + pathSep + f.Name())
			}
		} else {
			fmt.Println("file is :" + path)
			res.PushBack(path)
		}
	}
	return res
}
