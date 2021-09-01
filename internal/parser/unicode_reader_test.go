package parser

import (
	"fmt"
	"testing"
)

func TestRune(t *testing.T) {

	var str = "中华人民123\u4E25"
	//str = "a"
	//str = "a"
	buf := []byte(str)
	reader := NewUnicodeReader(&buf)
	for {
		succ := reader.scanRune()
		if !succ {
			break
		}
		fmt.Printf("ch is: %c\n", reader.ch)
	}
}

//测试把一个char放进去
func TestPutRune(t *testing.T) {

	var str = "中"
	//str = "a"
	//str = "a"
	buf := []byte(str)
	reader := NewUnicodeReader(&buf)

	reader.putChar('a')
	reader.putChar('a')
	reader.putChar('a')
	reader.putChar('中')
	reader.putChar('中')
}

func TestSlice(t *testing.T) {

	const SBUF_MAX = 8
	sbuf := [1]int{}

	//sbuf = [2]int{}
	sbuf[0] = 123
	//sbuf[1] = 124
	fmt.Println(sbuf)
}

func TestSwitch(t *testing.T) {

	a := 2
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
