package parser

/**
 * 所有的符号的定义
 * @author hushengdong
 */

type Symbol struct {
}

var empty_symbol *Symbol

func init() {

	empty_symbol = &Symbol{}
}

func NewEmptySymbol() *Symbol {

	return empty_symbol
}
