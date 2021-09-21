package parser

import (
	"husd.com/v0/util"
	"testing"
)

/**
 *
 * @author hushengdong
 */

func TestJavacParser_literal(t *testing.T) {

	var str string
	c := util.NewContext()
	str = " int a = 10 ; "
	p := NewJavacParserWithString(str, c)
	p.ParseStatement()
}
