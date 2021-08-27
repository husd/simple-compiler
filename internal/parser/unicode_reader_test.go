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
	reader := NewUnicodeReader(buf)
	for {
		succ := reader.scanRune()
		if !succ {
			break
		}
		fmt.Printf("ch is: %c\n", reader.ch)
	}
}
