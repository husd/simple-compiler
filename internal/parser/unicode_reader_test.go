package parser

import (
	"fmt"
	"husd.com/v0/util"
	"testing"
)

/**
 * 主要测试了读取中文的问题，数字和字母的问题都不大
 * @author hushengdong
 */
func TestRune(t *testing.T) {

	var str = "中华人民123\u4E25"
	//str = "a"
	//str = "a"
	buf := []byte(str)
	reader := NewUnicodeReader(&buf)
	for {
		succ := reader.ScanRune()
		if !succ {
			break
		}
		fmt.Printf("ch is: %c\n", reader.ch)
	}
}

//测试把一个char放进去
func TestPutRune(t *testing.T) {

	var str = "abc123&^%$!~中文<>?"
	buf := []byte(str)
	reader := NewUnicodeReader(&buf)
	for _, v := range str {
		reader.putChar(v)
	}
	fmt.Println("sbuf:", string(reader.sbuf[:reader.spos]))
	util.AssertEquals(t, "测试putChar 主要是测试中文", str, string(reader.sbuf[:reader.spos]))
}

func TestSlice(t *testing.T) {

	const SBUF_MAX = 8
	sbuf := make([]int, 10, 10)
	newSbuf := sbuf[0:]
	fmt.Println("sbuf len:", len(newSbuf), " cap:", cap(newSbuf))
	sbuf[0] = 123
	sbuf[1] = 123
	sbuf[2] = 123
	sbuf[3] = 123
	fmt.Println(sbuf)
}

// go语言的switch 不需要break 空的case也没问题 以下代码什么都不会输出
func TestSwitch(t *testing.T) {

	a := 1
	switch a {
	case 1:
	case 2:
		fmt.Println("a is 2")
	case 5:
		fmt.Println("a is 5")
	default:
		fmt.Println("a is default")
	}
}
