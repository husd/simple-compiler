package parser

import (
	"husd.com/v0/code"
)

/**
 *
 * 可以把源代码里的 String name = "xiaoming 你好" 分解成:
 * String
 * name
 * =
 * "xiaoming 你好"
 *
 * 可以看出，并不是简单的按某个字符分割，还需要过滤出注释 包括单行注释和多行注释。
 *
 */

type JavaTokenizer struct {
	reader    *UnicodeReader // reader
	source    code.JVersion  // jdk版本
	tokenKind tokenKind      // 当前token的类型

}

func NewJavaTokenizer(path string) *JavaTokenizer {

	javaTokenizer := JavaTokenizer{}
	javaTokenizer.reader = NewUnicodeReaderFromFile(path)
	javaTokenizer.source = code.JDK8
	javaTokenizer.tokenKind = TOKEN_KIND_ERROR

	return &javaTokenizer
}

func (jt *JavaTokenizer) ReadToken() *Token {

	//TODO husd start 核心方法 解析出来下一个token是什么
	//endPos := int(0)
	reader := jt.reader
	for {
		//pos := reader.CurrentPos()
		switch reader.ch {
		case '\t':
		case ' ':
		case Layout_char_ff:
			for {
				reader.ReadRune()
				if reader.ch == '\t' || reader.ch == ' ' || reader.ch == Layout_char_ff {
					break
				}
			}
			//
			break
		case Layout_char_lf:
			reader.ReadRune()
			break
		}

	}

	t := Token{}
	return &t
}
