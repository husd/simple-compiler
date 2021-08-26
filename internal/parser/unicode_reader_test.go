package parser

import (
	"fmt"
	"testing"
)

func TestRune(t *testing.T) {

	var str = "\uFF41"
	//str = "a"
	//str = "a"
	buf := []byte(str)
	reader := NewUnicodeReader(buf)
	for {
		succ, ch, _ := reader.ReadRune()
		if !succ {
			break
		}
		fmt.Printf("ch is: %c\n", ch)
	}
}
