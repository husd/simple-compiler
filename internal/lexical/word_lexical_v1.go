package lexical

import (
	"container/list"
	"fmt"
	"husd.com/v0/util"
)

/**
 * WordLexical 非常简单的程序，把字符串分割成单词 这个文件的方法，独立，简单，是一个开胃菜。
 * 万事开头难，先从这里开始，可以增加一点信心。
 * @author hushengdong
 */
type WordLexical struct{}

/**
 * LexicalAnalysis 实现了把输入的字符串，解析成了单词 自动过滤空格 和 换行
 * 换行也是停止符号 如果是注释的话，会区分单行注释 还是双行注释
 * 这个方法只解析单个方法，因此没办法判断多行注释。 只实现了过滤单行注释
 */
func (a *WordLexical) LexicalAnalysis(str string) *list.List {

	res := list.List{}
	length := len(str)
	start := 0
	//去掉了开头的空格
	for ; start < length && util.EndChar(str[start]); start++ {
	}
	// 判断是不是单行注释 //
	//遍历这个字符串
	for i := start; i < length; {
		for i < length && !util.EndChar(str[i]) {
			if str[i] == '/' && i+1 < length && str[i+1] == '/' {
				//单行注释 // 直接结束，跳出循环
				w := str[start:i]
				res.PushBack(w)
				goto end
			}
			i++
		}
		w := str[start:i]
		fmt.Println(w)
		res.PushBack(w)
		//找到空格了，要继续跳过空格
		for i < length && util.EndChar(str[i]) {
			i++
		}
		start = i
	}
end:
	return &res
}
