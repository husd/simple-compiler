package util

import (
	"bufio"
	"container/list"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func FindRepeatFileAndDelete(name string) {
	fMap := make(map[string]string)
	exist := Exists(name)
	if !exist {
		panic("文件或者目录不存在，请检查 " + name)
	}
	stackString := DefaultStack{}
	stackString.Init()
	stackString.Push(name)
	pathSep := string(os.PathSeparator)
	total := 0

	toDelete := list.New()

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
				file := path + pathSep + f.Name()
				file2 := fMap[f.Name()]
				if file2 != "" {
					if !ignoreFile(file) {
						m1, _ := md5Str(file)
						m2, _ := md5Str(file2)
						if m1 == m2 {
							fmt.Println(file + "       |        " + file2)
							toDelete.PushBack(file)
						}
					}
				} else {
					fMap[f.Name()] = file
				}
				stackString.Push(file)
			}
		} else {
			total++
		}
	}
	fmt.Printf("\r%d", total)
	fmt.Println(" start to delete file")
	deleteFileCount := 0
	for i := toDelete.Front(); i != nil; i = i.Next() {
		deleteFileCount++
		p := fmt.Sprintf("%v", i.Value)
		fmt.Println("delete file: ", p)
		os.Remove(p)
	}
	fmt.Println("total delete file :", deleteFileCount)
}

func md5Str(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	_, _ = io.Copy(hash, file)
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func ignoreFile(name string) bool {

	l := len(name)
	sux := name[l-4 : l]
	return name[l-4] != '.' || sux == ".MOV" || sux == ".MP4" || sux == ".mov" || sux == ".mp4" || sux == ".AAE"
}

func vedioFile(name string) bool {

	l := len(name)
	sux := name[l-4 : l]
	return sux == ".MOV" || sux == ".MP4" || sux == ".mov" || sux == ".mp4"
}

func CountFile(name string) int {
	res := 0
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
			//fmt.Println("file is :" + path)
			res++
		}
	}
	return res
}

//按行读取文件
func ReadFileLines(path string) *list.List {

	res := list.List{}
	fd, err := os.Open(path)
	defer fd.Close()
	if err != nil {
		panic("读取文件失败 " + path)
	}
	buf := bufio.NewReader(fd)
	for {
		data, _, eof := buf.ReadLine()
		if eof == io.EOF {
			break
		}
		res.PushBack(data)
	}
	return &res
}
