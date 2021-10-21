package parser

/**
 * 括号的模式
 * @author hushengdong
 */
type parenthesesResult int

const (
	CAST            parenthesesResult = 0 // 转换
	EXPLICIT_LAMBDA parenthesesResult = 1 // 明确的lambda表达式
	IMPLICIT_LAMBDA parenthesesResult = 2 // 含蓄的lambda表达式
	PARENS          parenthesesResult = 3 // 单纯括号
)
