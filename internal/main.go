package main

import (
	. "husd.com/v0/lexical"
	"husd.com/v0/util"
)

func main() {

	//wordLex2 := WordLexicalV2{}
	//wordLex2.LexicalAnalysis("D:\\gitee-source\\go-framework")

	util.FindRepeatFileAndDelete("F:\\照片")
	//c := util.CountFile("F:\\照片\\20210815刘伟娟手机照片")
	//fmt.Println("总文件：" , c)
}

func main2() {

	wordLex := WordLexical{}

	str := "   this ois abc ! \n  second line \n third line"
	_list := wordLex.LexicalAnalysis(str)
	for item := _list.Front(); item != nil; item = item.Next() {
		//fmt.Println(item.Value)
	}
}
