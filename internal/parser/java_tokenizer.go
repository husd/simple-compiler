package parser

import (
	"fmt"
	"husd.com/v0/code"
	"husd.com/v0/compiler"
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
	reader *UnicodeReader // reader
	source code.JVersion  // jdk版本
	tk     tokenKind      // 当前token的类型
}

func NewJavaTokenizer(path string) *JavaTokenizer {

	javaTokenizer := JavaTokenizer{}
	javaTokenizer.reader = NewUnicodeReaderFromFile(path)
	javaTokenizer.source = code.JDK8
	javaTokenizer.tk = TOKEN_KIND_ERROR

	return &javaTokenizer
}

func (jt *JavaTokenizer) ReadToken() *Token {

	//TODO husd start 核心方法 解析出来下一个token是什么
	//endPos := int(0)
	reader := jt.reader
loop:
	for {
		//bp := reader.CurrentPos()
		switch reader.ch {
		case '\t':
		case ' ':
		case Layout_char_ff: // 空格 tab 换页等
			for {
				reader.ReadRune()
				if reader.ch == '\t' || reader.ch == ' ' || reader.ch == Layout_char_ff {
					break
				}
			}
			//
			break
		case Layout_char_lf: // 换行
			reader.ReadRune()
			break
		case Layout_char_cr: // 回车
			reader.ReadRune()
			if reader.ch == Layout_char_lf { //有的操作系统是： CRLF
				reader.ReadRune()
			}
			break
		case 'A':
		case 'B':
		case 'C':
		case 'D':
		case 'E':
		case 'F':
		case 'G':
		case 'H':
		case 'I':
		case 'J':
		case 'K':
		case 'L':
		case 'M':
		case 'N':
		case 'O':
		case 'P':
		case 'Q':
		case 'R':
		case 'S':
		case 'T':
		case 'U':
		case 'V':
		case 'W':
		case 'X':
		case 'Y':
		case 'Z':
		case 'a':
		case 'b':
		case 'c':
		case 'd':
		case 'e':
		case 'f':
		case 'g':
		case 'h':
		case 'i':
		case 'j':
		case 'k':
		case 'l':
		case 'm':
		case 'n':
		case 'o':
		case 'p':
		case 'q':
		case 'r':
		case 's':
		case 't':
		case 'u':
		case 'v':
		case 'w':
		case 'x':
		case 'y':
		case 'z':
		case '$':
		case '_': //java的标识符，只能以这些字符开头
			jt.scanIdentify()
			goto loop
		case '0': //0比较特殊，例如 0xF 0b10 等数字，需要单独处理
			//TODO husd
			break
		case '1':
		case '2':
		case '3':
		case '4':
		case '5':
		case '6':
		case '7':
		case '8':
		case '9':
			jt.scanNumber()
			goto loop
		case '.':
		case ',':
		case ';':
		case '(':
		case ')':
		case '{':
		case '}':
		case '[':
		case ']':
		case '/':
		case '\'':
		case '"':
			fmt.Println("遇到了特殊符号，单独处理")
		default:
			fmt.Println("默认逻辑----------------------")
		}
	}
	t := Token{}
	return &t
}

func (jt *JavaTokenizer) scanIdentify() {

	isIdentify := false
	reader := jt.reader
	start := reader.bp - 1
	reader.ReadRune() //读下一个字符
	for {
		switch reader.ch {
		case 'A':
		case 'B':
		case 'C':
		case 'D':
		case 'E':
		case 'F':
		case 'G':
		case 'H':
		case 'I':
		case 'J':
		case 'K':
		case 'L':
		case 'M':
		case 'N':
		case 'O':
		case 'P':
		case 'Q':
		case 'R':
		case 'S':
		case 'T':
		case 'U':
		case 'V':
		case 'W':
		case 'X':
		case 'Y':
		case 'Z':
		case 'a':
		case 'b':
		case 'c':
		case 'd':
		case 'e':
		case 'f':
		case 'g':
		case 'h':
		case 'i':
		case 'j':
		case 'k':
		case 'l':
		case 'm':
		case 'n':
		case 'o':
		case 'p':
		case 'q':
		case 'r':
		case 's':
		case 't':
		case 'u':
		case 'v':
		case 'w':
		case 'x':
		case 'y':
		case 'z':
		case '$':
		case '_':
		case '0':
		case '1':
		case '2':
		case '3':
		case '4':
		case '5':
		case '6':
		case '7':
		case '8':
		case '9':
			break
		case '\u0000':
		case '\u0001':
		case '\u0002':
		case '\u0003':
		case '\u0004':
		case '\u0005':
		case '\u0006':
		case '\u0007':
		case '\u0008':
		case '\u000E':
		case '\u000F':
		case '\u0010':
		case '\u0011':
		case '\u0012':
		case '\u0013':
		case '\u0014':
		case '\u0015':
		case '\u0016':
		case '\u0017':
		case '\u0018':
		case '\u0019':
		case '\u001B':
		case '\u007F':
			reader.ReadRune()
			continue
		case '\u001A': // EOI 也是一个有效的标识符 That's the Ctrl+Z control code.
			if reader.bp >= reader.size {
				var name string
				// 这个名字，可能需要缓存，因为关键字很多，没必要非得转字符串，把字节计算一下hash
				// 能提高下性能，这里要存一下符号表了，可以指定一下字节数组的位置和长度，就行了。
				name = string(reader.buf[start:reader.bp]) // 不仅仅有标识符的名字，还有其它属性 TODO husd
				//jt.tk = tokens.lookupKind(name)
				if compiler.DEBUG {
					fmt.Println("name is :", name)
				}
				return
			}
			reader.ReadRune()
			continue
		}
		// TODO husd 记录
		isIdentify = false
		fmt.Println("isIdentify :", isIdentify)
	}
}

func (jt *JavaTokenizer) scanNumber() {

	//TODO husd
}
