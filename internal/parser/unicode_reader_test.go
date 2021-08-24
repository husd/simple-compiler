package parser

import (
	"fmt"
	"testing"
)

func TestRune(t *testing.T) {

	var str = "北京"
	buf := []byte(str)
	reader := NewUnicodeReader(buf)
	fmt.Printf("str[%d]=%c\n", 0, reader.ReadRune())
	fmt.Printf("str[%d]=%c\n", 1, reader.ReadRune())

	for inx, ch := range str {

		fmt.Printf("str[%d]=%c\n", inx, ch)
	}
}
