package parser

import (
	"fmt"
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
	str = " boolean a = false ; "
	p := NewJavacParserWithString(str, c)
	tree := p.ParseExpression()

	fmt.Println(tree)
}

func TestJavacParser_statement(t *testing.T) {

	var str string
	c := util.NewContext()
	str = "{ {} {} {} }"
	p := NewJavacParserWithString(str, c)
	tree := p.ParseStatement()
	fmt.Println(tree)
}

func TestJavacParser_if(t *testing.T) {

	var str string
	c := util.NewContext()
	str = "if ( a == 10) {} "
	p := NewJavacParserWithString(str, c)
	tree := p.ParseStatement()
	fmt.Println(tree)
}

func TestJavacParser_if_else(t *testing.T) {

	var str string
	c := util.NewContext()
	str = "if ( a == 10) {} else if ( a == 20) {} else {} "
	p := NewJavacParserWithString(str, c)
	tree := p.ParseStatement()
	fmt.Println(tree)
}
