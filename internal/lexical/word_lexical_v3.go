package lexical

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
)

// WordLexicalV3
/**
 * 要实现词法分析器功能
 */
type WordLexicalV3 struct{}

// LexicalAnalysis 解析单个文件 按行读取单个文件，把解析到的数据，存入到链表里。
// 这个方法对于v1版本的方法来说，增加了1个功能，会过滤多行注释 /* */ 这样的注释
// 该方法实现了以下功能：
// 1 过滤单行注释
// 2 过滤多行注释
// 3 分割次素
// 4 自动处理换行和空格 过滤多余的空格
// 5 识别出当前token，并转换为 TokenTag
// 基本实现了词法分析器的功能
func (a *WordLexicalV3) LexicalAnalysis(path string) *list.List {

	res := list.New()

	fd, err := os.Open(path)
	defer fd.Close()
	if err != nil {
		panic("读取文件失败 " + path)
	}
	buf := bufio.NewReader(fd)
	lineNum := 0

	comment02Flag := false //多行注释的标志，如果为true，就表示扫描到了多行注释的开始 /*
	for {
	NEWLINE:
		data, _, eof := buf.ReadLine()
		if eof == io.EOF {
			break
		}
		lineNum++
		line := string(data)
		// ----------- 解析单行内容 -----------
		length := len(line)
		start := 0
		//去掉了开头的空格
		for ; start < length && blankChar(line[start]); start++ {
		}
		//idStart 是为了记录当前token从哪个索引开始的
		idStart := false
		for i := start; i < length; i++ {
			flag := judgeCondition(line, i, length)
			//忽略多行注释的一切内容，直到找到了多行注释的结束
			if comment02Flag && flag != comment02End {
				goto NEXTLOOP
			}
			switch flag {
			case blank: //由于之前已经去过空格了，再有空格，就是拿到token了
				goto TOKEN
			case id:
				if !idStart {
					idStart = true
					start = i
				}
				//如果是最后一个字符，直接结束
				if i == length-1 {
					i++ //加1的原因是，切片是左开右闭的 str[1:2]
					goto TOKEN
				} else {
					goto NEXTLOOP
				}
			case eofFlag:
				i++
				goto TOKEN
			case comment01:
				//单行注释，就需要先拿到token，然后跳过这一行
				if i > start && idStart {
					w := line[start:i]
					fmt.Println("token is :", w)
					idStart = false
				}
				goto NEWLINE
			case comment02Start:
				//多行注释的开始，就需要一直找多行注释的结束
				if i > start && idStart {
					w := line[start:i]
					fmt.Println("token is :", w)
					idStart = false
				}
				comment02Flag = true
				i++ // 一定要注意，因为我们读取了2个字符，所以要跳过下一个字符
				goto NEXTLOOP
			case comment02End:
				comment02Flag = false
				i++ // 一定要注意，因为我们读取了2个字符，所以要跳过下一个字符
				goto NEXTLOOP
			}
		TOKEN:
			//这里就是要声明token了
			if i > start && idStart {
				w := line[start:i]
				fmt.Println("token is :", w)
				idStart = false
			}
		NEXTLOOP:
		}
		// ----------- 解析单行内容 -----------
	}
	if comment02Flag {
		//多行注释没有结束 是否要提示？ // TODO husd
	}
	return res
}

// 1 空格 直接忽略
// 2 字母或者数字等 ，直接 i++ 看下一个字符是什么
// 3 // 2个斜线标识单行注释，这个时候要忽略这一样的内容，但是要注意，只能忽略//后面的内容，前面的内容不能忽略
// 4 /* 表示多行注释的开头
// 5 */ 表示多行注释的结尾
// 6 \n 表示这一行结束了
const (
	blank          = 1
	id             = 2
	comment01      = 3
	comment02Start = 4
	comment02End   = 5
	eofFlag        = 6
)

func judgeCondition(line string, index int, max int) int {

	if index >= max {
		return eofFlag
	}
	ch := line[index]
	if blankChar(ch) {
		return blank
	}
	if eofChar(ch) {
		return eofFlag
	}
	if index+1 < max {
		nextChar := line[index+1]
		if ch == '/' && nextChar == '/' {
			return comment01
		}
		if ch == '/' && nextChar == '*' {
			return comment02Start
		}
		if ch == '*' && nextChar == '/' {
			return comment02End
		}
	}
	return id
}
