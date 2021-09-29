package parser

import (
	"fmt"
	"husd.com/v0/code"
	"husd.com/v0/compiler"
	"husd.com/v0/util"
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
 * @author hushengdong
 */

type JavaTokenizer struct {
	context *util.Context
	reader  *UnicodeReader // reader
	source  code.JVersion  // jdk版本
	tk      tokenKind      // 当前token的类型
	name    *util.Name     // identify name

	tokenFactory *Tokens // token工厂类

	errPos int // 错误地址
	radix  int // 进制 set by nextToken().
}

func NewJavaTokenizer(path string, c *util.Context) *JavaTokenizer {

	v := code.JDK8

	javaTokenizer := JavaTokenizer{}
	javaTokenizer.reader = NewUnicodeReaderFromFile(path)
	javaTokenizer.source = v
	javaTokenizer.tk = TOKEN_KIND_DEF
	javaTokenizer.context = c
	javaTokenizer.tokenFactory = InstanceTokens(c)

	return &javaTokenizer
}

func NewJavaTokenizerWithString(str string, c *util.Context) *JavaTokenizer {

	v := code.JDK8

	javaTokenizer := JavaTokenizer{}
	b := []byte(str)
	javaTokenizer.reader = NewUnicodeReader(&b)
	javaTokenizer.source = v
	javaTokenizer.tk = TOKEN_KIND_ERROR
	javaTokenizer.context = c
	javaTokenizer.tokenFactory = InstanceTokens(c)

	return &javaTokenizer
}

func (jt *JavaTokenizer) readToken() Token {

	pos := 0
	reader := jt.reader
	reader.spos = 0 // reset spos
	endPos := 0
loop:
	for {
		pos = reader.bp
		switch reader.ch {
		case '\t', ' ', Layout_char_ff: // 空格(32) tab(9) 换页(12)等
			for {
				reader.ScanRune()
				if !(reader.ch == '\t' || reader.ch == ' ' || reader.ch == Layout_char_ff) {
					break
				}
			}
		case Layout_char_lf: // 换行
			reader.ScanRune()
		case Layout_char_cr: // 回车
			reader.ScanRune()
			if reader.ch == Layout_char_lf { //有的操作系统是： CRLF
				reader.ScanRune()
			}
		case
			'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
			'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
			'U', 'V', 'W', 'X', 'Y', 'Z',
			'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
			'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
			'u', 'v', 'w', 'x', 'y', 'z',
			'$', '_':
			//java的标识符，只能以这些字符开头
			jt.scanIdentify()
			break loop
		case '0': //0比较特殊，例如 0xF 0b10 等数字，需要单独处理
			reader.ScanRune()
			if reader.ch == 'X' || reader.ch == 'x' {
				//16进制
				reader.ScanRune()
				jt.skipUnderLine()
				if reader.ch == '.' {
					jt.scanHexFractionAndSuffix(pos, false)
				} else if reader.digit(pos, 16) < 0 {
					jt.lexError(pos, "无效的16进制数字")
				} else {
					jt.scanNumber(pos, 16)
				}
			} else if reader.ch == 'b' || reader.ch == 'B' {
				if !allowBinaryLiterals {
					jt.lexError(pos, "不允许使用二进制字面量", jt.source)
				}
				reader.ScanRune()
				jt.skipUnderLine()
				if reader.digit(pos, 2) < 0 {
					jt.lexError(pos, "无效的二进制数字")
				} else {
					jt.scanNumber(pos, 2)
				}
			} else {
				reader.putChar('0')
				if reader.ch == '_' {
					savedPos := reader.bp
					for {
						reader.ScanRune()
						if reader.ch != '_' {
							break
						}
					}
					if reader.digit(pos, 10) < 0 {
						jt.lexError(savedPos, "无效的下划线")
					}
				}
				jt.scanNumber(pos, 8)
			}
			break loop
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			jt.scanNumber(pos, 10)
			break loop
		case '.':
			//Java里对于点的这种开头的数字
			reader.ScanRune()
			if reader.ch >= '0' && reader.ch <= '9' {
				reader.putChar('.')
				jt.scanFractionAndSuffix(pos)
			} else if reader.ch == '.' {
				savedPos := reader.bp
				reader.putChar('.')
				reader.putRuneChar('.', true)
				if reader.ch == '.' {
					reader.ScanRune()
					reader.putChar('.')
					jt.tk = TOKEN_KIND_ELLIPSIS
				} else {
					jt.lexError(savedPos, "无效的点 . ")
				}
			} else {
				jt.tk = TOKEN_KIND_DOT
			}
			break loop
		case ',':
			reader.ScanRune()
			jt.tk = TOKEN_KIND_COMMA
			break loop
		case ';':
			reader.ScanRune()
			jt.tk = TOKEN_KIND_SEMI
			break loop
		case '(':
			reader.ScanRune()
			jt.tk = TOKEN_KIND_LPAREN
			break loop
		case ')':
			reader.ScanRune()
			jt.tk = TOKEN_KIND_RPAREN
			break loop
		case '{':
			reader.ScanRune()
			jt.tk = TOKEN_KIND_LBRACE
			break loop
		case '}':
			reader.ScanRune()
			jt.tk = TOKEN_KIND_RBRACE
			break loop
		case '[':
			reader.ScanRune()
			jt.tk = TOKEN_KIND_LBRACKET
			break loop
		case ']':
			reader.ScanRune()
			jt.tk = TOKEN_KIND_RBRACKET
			break loop
		case '/': // 这里涉及到了注释的解析
			reader.ScanRune()
			if reader.ch == '/' { //读到了单行注释
				for {
					reader.ScanRune()
					if reader.ch == Layout_char_cr ||
						reader.ch == Layout_char_lf ||
						reader.reachEnd() { // 一直读到换行或者结束
						reader.ScanRune() // 读下个字符
						break
					}
				}
				/**    这里要注意，读完了注释内容，这个时候要继续去读下一个内容，如下这种情况：
				 *      int a // 1234
				 *       = 4;
				 */
				goto loop
			} else if reader.ch == '*' { //读到了多行注释
				for {
					reader.ScanRune()
					if reader.ch == '*' {
						// 看下一个是不是 /
						nextRune, succ := reader.seenNextRune()
						if succ && nextRune == '/' { // 找到了结束符号
							reader.ScanRune() // 跳过 结束符号
							reader.ScanRune() // 查看下一个
							/**
							 * 这里同单行注释是一样的，找到了结束符号，就需要继续循环
							 */
							goto loop
						}
					} else {
						//不是 * 就继续找
					}
					if reader.reachEnd() {
						goto loop
					}
				}
				if reader.reachEnd() {
					//找到结束，都没有找到多行注释的结束符号，直接报错
					jt.lexError(pos, "多行注释没有找到结束符号")
					break loop
				}
			} else if reader.ch == '=' {
				// 对于Java语法来说，就是遇到了 /= 这种符号
				jt.tk = TOKEN_KIND_SLASHEQ
				reader.ScanRune()
			} else {
				jt.tk = TOKEN_KIND_SLASH //单纯就是 /
			}
			break loop
		case '\'': //遇到了单引号
			reader.ScanRune()
			if reader.ch == '\'' { // ''
				jt.lexError(pos, "空 '' ")
			} else {
				if reader.ch == Layout_char_cr ||
					reader.ch == Layout_char_lf {
					jt.lexError(pos, "无效的换行")
				}
				jt.scanLitChar(pos)
				if reader.ch == '\'' { // 找到了 ''
					reader.ScanRune()
					jt.tk = TOKEN_KIND_CHAR_LITERAL
				} else {
					jt.lexError(pos, "单引号没有关闭")
				}
			}
			break loop
		case '"':
			reader.ScanRune()
			for reader.ch != '"' &&
				reader.ch != Layout_char_cr &&
				reader.ch != Layout_char_lf &&
				!reader.reachEnd() {
				jt.scanLitChar(pos)
			}
			if reader.ch == '"' {
				jt.tk = TOKEN_KIND_STRING_LITERAL
				reader.ScanRune()
			} else {
				jt.lexError(pos, "字符串没有关闭 只有一个双引号")
			}
			break loop
		default:
			if isSpecialOperator(reader.ch) {
				jt.scanOperator()
			} else {
				//走到这里，默认就假设是标识符了 ，按标识符解析
				isJavaIdentifyPart := false
				if reader.ch < '\u0080' { // 小于 \u0080 的都是ascii码，已经处理过了，所以肯定不是标识符
					isJavaIdentifyPart = false
				} else {
					//TODO husd 判断是不是标识符
					isJavaIdentifyPart = true
				}
				if reader.reachEnd() {
					jt.tk = TOKEN_KIND_EOF
					pos = reader.size
					if isJavaIdentifyPart {
						jt.scanIdentify()
					}
				} else if isJavaIdentifyPart {
					jt.scanIdentify()
				} else if reader.ch == Layout_char_eoi ||
					reader.reachEnd() { //JTS 3.5 主要说的EOI的问题
					jt.tk = TOKEN_KIND_EOF
					pos = reader.size
				} else {
					jt.lexError(pos, "无效的字符:", reader.ch)
					reader.ScanRune()
				}
			}
			break loop
		}
	}
	endPos = reader.bp
	tag := GetTokenKindTag(jt.tk)
	switch tag {
	case TOKEN_TAG_DEFAULT:
		return newDefaultToken(jt.tk, reader.lineNum, reader.linePos, pos, endPos)
	case TOKEN_TAG_STRING:
		return newStringToken(jt.tk, reader.lineNum, reader.linePos, reader.name().NameStr, pos, endPos)
	case TOKEN_TAG_NUMERIC:
		return newNumericToken(jt.tk, reader.lineNum, reader.linePos, reader.name().NameStr, jt.radix, pos, endPos)
	case TOKEN_TAG_NAMED:
		return newNamedToken(jt.tk, reader.lineNum, reader.linePos, jt.name, pos, endPos)
	default:
		panic(fmt.Sprintf("错误的tokenKind pos:%v endPos:%v", pos, endPos))
	}
}

func isSpecialOperator(ch rune) bool {

	switch ch {
	case
		'!', '%', '&', '*', '?',
		'+', '-', ':', '<', '=',
		'>', '^', '|', '~', '@':
		return true
	default:
		return false
	}
	return false
}

/**
 * 扫描标识符 Java里的变量等 关键字也属于这里识别的
 */
func (jt *JavaTokenizer) scanIdentify() {

	reader := jt.reader
	isIdentify := false
	var high rune
	reader.putRune(true)
loop:
	for {
		switch reader.ch {
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
			'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
			'U', 'V', 'W', 'X', 'Y', 'Z',
			'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
			'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
			'u', 'v', 'w', 'x', 'y', 'z',
			'$', '_',
			'0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			goto next
		case '\u0000', '\u0001', '\u0002', '\u0003',
			'\u0004', '\u0005', '\u0006', '\u0007',
			'\u0008', '\u000E', '\u000F', '\u0010',
			'\u0011', '\u0012', '\u0013', '\u0014',
			'\u0015', '\u0016', '\u0017',
			'\u0018', '\u0019', '\u001B':
			reader.ScanRune()
			goto loop
		case '\u001A': // EOI 也是一个有效的标识符 That's the Ctrl+Z control code.
			if reader.bp >= reader.size {
				// 这个名字，可能需要缓存，因为关键字很多，没必要非得转字符串，把字节计算一下hash
				// 能提高下性能，这里要存一下符号表了，可以指定一下字节数组的位置和长度，就行了。
				jt.name = reader.name() // 不仅仅有标识符的名字，还有其它属性
				jt.tk = jt.tokenFactory.lookupTokenKind(jt.name)
				if compiler.DEBUG {
					fmt.Println("EOI name is :", jt.name)
				}
				return
			}
			reader.ScanRune()
			goto loop
		default:
			if reader.ch < '\u0080' {
				isIdentify = false //ascii已经考虑过了，直接结束
			} else {
				if isIdentifierIgnorable(reader.ch) {
					reader.ScanRune()
					goto loop
				} else {
					high = reader.scanSurrogates()
					if high != 0 {
						//reader.putChar(high)
						//temp := toCodePoint(high,reader.ch)
						//isIdentify  = isJavaIdentifierPart(temp)
					} else {
						isIdentify = isJavaIdentifierPart(reader.ch)
					}
				}
				isIdentify = false
			}
			if !isIdentify { //直接读取结束了。
				jt.name = reader.name()
				jt.tk = jt.tokenFactory.lookupTokenKind(jt.name)
				return //end
			}
		}
		// break 之后，执行到这里了，扫描下一个字符
	next:
		reader.putRune(true)
	}
}

/**
 * 词法分析要分析出来源代码里的 数字，目前支持 2进制 8进制 10进制 16进制的数字
 * 词法分析阶段，没有把数字解析为对应的10进制数字，而是保留了字面量和进制
 * 这也符合词法分析阶段的功能，只做解析，而不处理数据。
 */
func (jt *JavaTokenizer) scanNumber(pos int, radix int) {

	jt.radix = radix
	//对于8进制，按10进制规则来，因为有可能是一个单精度数字的字面量 0x1.0p1
	digitRadix := radix
	if radix == 8 {
		digitRadix = 10
	}
	seenDigit := false
	reader := jt.reader
	if reader.digit(pos, digitRadix) >= 0 {
		seenDigit = true
		jt.scanDigits(pos, digitRadix)
	}
	if radix == 16 && reader.ch == '.' {
		jt.scanHexFractionAndSuffix(pos, seenDigit)
	} else if seenDigit && radix == 16 && (reader.ch == 'p' || reader.ch == 'P') {
		jt.scanHexExponentAndSuffix(pos)
	} else if digitRadix == 10 && reader.ch == '.' {
		reader.putRune(true)
		jt.scanFractionAndSuffix(pos)
	} else if digitRadix == 10 &&
		(reader.ch == 'e' || reader.ch == 'E' ||
			reader.ch == 'f' || reader.ch == 'F' ||
			reader.ch == 'd' || reader.ch == 'D') {
		jt.scanFractionAndSuffix(pos)
	} else {
		if reader.ch == 'l' || reader.ch == 'L' {
			reader.ScanRune()
			jt.tk = TOKEN_KIND_LONG_LITERAL
		} else {
			jt.tk = TOKEN_KIND_INT_LITERAL
		}
	}
}

func (jt *JavaTokenizer) skipUnderLine() {

	if jt.reader.ch == '_' {
		jt.lexError(jt.reader.bp, "warn:数字里有无效的下划线")
		for jt.reader.ch == '_' {
			jt.reader.ScanRune()
		}
	}
}

func (jt *JavaTokenizer) lexError(bp int, msg ...interface{}) {

	start := bp - 20
	end := bp + 20

	if start < 0 {
		start = 0
	}
	if end > jt.reader.size {
		end = jt.reader.size
	}

	fmt.Println("---------------- error/warn -------------词法解析错误，位置：", bp, " msg:", msg, " 前后词语:", string(jt.reader.buf[start:end]))
	jt.tk = TOKEN_KIND_ERROR
	jt.errPos = bp
}

// 16进制小数部分和后缀 d或者f    float a = 10.3f  double b = 102.3d
func (jt *JavaTokenizer) scanHexFractionAndSuffix(pos int, seenDigit bool) {

	jt.radix = 16
	reader := jt.reader
	reader.putRune(true)
	jt.skipUnderLine()
	if reader.digit(pos, 16) >= 0 {
		seenDigit = true
		jt.scanDigits(pos, 16)
	}
	if !seenDigit {
		jt.lexError(pos, "无效的16进制数字")
	} else {
		jt.scanHexExponentAndSuffix(pos)
	}
}

// 16进制小数部分
func (jt *JavaTokenizer) scanHexExponentAndSuffix(pos int) {

	reader := jt.reader
	if reader.ch == 'p' || reader.ch == 'P' {
		reader.putRune(true)
		jt.skipUnderLine()
		if reader.ch == '+' || reader.ch == '-' {
			reader.putRune(true)
		}
		jt.skipUnderLine()
		if reader.ch >= 0 && reader.ch <= 9 {
			jt.scanDigits(pos, 10)
			if !allowHexFloats {
				jt.lexError(pos, "不允许十六进制浮点文本")
			} else if !hexFloatsWork {
				jt.lexError(pos, "不允许十六进制浮点文本")
			}
		} else {
			jt.lexError(pos, "fp.lt 格式不正确")
		}
	} else {
		jt.lexError(pos, "fp.lt 格式不正确")
	}
	if reader.ch == 'f' || reader.ch == 'F' {
		reader.putRune(true)
		jt.tk = TOKEN_KIND_FLOAT_LITERAL
		jt.radix = 16
	} else {
		if reader.ch == 'd' || reader.ch == 'D' {
			reader.putRune(true)
			jt.tk = TOKEN_KIND_DOUBLE_LITERAL
			jt.radix = 16
		}
	}
}

// 扫描数字 注意区分 scanNumber number的含义更丰富
func (jt *JavaTokenizer) scanDigits(pos int, radix int) {

	reader := jt.reader
	var res rune
	var savedPos int
	var haxNext bool = true
	for {
		if reader.ch != '_' {
			reader.putRune(false)
		} else {
			if !allowUnderscoresInLiterals {
				jt.lexError(pos, "error:数字不支持下划线")
			}
			res = reader.ch
			savedPos = reader.bp
		}
		haxNext = reader.ScanRune()
		if !((reader.digit(pos, radix) >= 0 || reader.ch == '_') && haxNext) {
			break
		}
	}
	if res == '_' {
		jt.lexError(savedPos, "warn:无效的下划线")
	}
}

// read next char or string literal and copy into sbuf
func (jt *JavaTokenizer) scanLitChar(pos int) {

	reader := jt.reader
	if reader.ch == '\\' {
		if reader.peekChar() == '\\' && !reader.isUnicode() {
			reader.skipChar()
			reader.putRuneChar('\\', true)
		} else {
			reader.ScanRune()
			switch reader.ch {
			case '0', '1', '2', '3', '4', '5', '6', '7':
				leadCh := reader.ch
				oct := reader.digit(pos, 8)
				reader.ScanRune()
				if reader.ch >= '0' && reader.ch <= '7' { //解析八进制的数字
					oct = oct*8 + reader.digit(pos, 8)
					reader.ScanRune()
					if leadCh <= '3' && reader.ch >= '0' && reader.ch <= '7' {
						oct = oct*8 + reader.digit(pos, 8)
						reader.ScanRune()
					}
				}
				reader.putChar(oct)
			case 'b':
				reader.putRuneChar('\b', true)
			case 't':
				reader.putRuneChar('\t', true)
			case 'n':
				reader.putRuneChar('\n', true)
			case 'f':
				reader.putRuneChar('\f', true)
			case 'r':
				reader.putRuneChar('\r', true)
			case '\'':
				reader.putRuneChar('\'', true)
			case '"':
				reader.putRuneChar('"', true)
			case '\\':
				reader.putRuneChar('\\', true)
			default:
				jt.lexError(reader.bp, "illegal.esc.char")
			}
		}
	} else if reader.bp != reader.size {
		reader.putRune(true)
	}
}

// 扫描操作符号 需要尽可能的长
func (jt *JavaTokenizer) scanOperator() {

	reader := jt.reader
	for {
		reader.putRune(false)
		//看看操作符号前面是什么
		newName := reader.name()
		tempTk := jt.tokenFactory.lookupTokenKind(newName)
		if tempTk == TOKEN_KIND_IDENTIFIER { //?
			reader.spos--
			break
		}
		jt.tk = tempTk
		reader.ScanRune()
		if !isSpecialOperator(reader.ch) {
			break
		}
	}
}

// 读取小数部分和单精度小数的d 或者 f 的后缀
// float a = 10.4f double b = 10.00d
func (jt *JavaTokenizer) scanFractionAndSuffix(pos int) {

	jt.radix = 10
	jt.scanFraction(pos)
	reader := jt.reader
	if reader.ch == 'f' || reader.ch == 'F' {
		reader.putRune(true)
		jt.tk = TOKEN_KIND_FLOAT_LITERAL
	} else {
		if reader.ch == 'd' || reader.ch == 'D' {
			reader.putRune(true)
		}
		jt.tk = TOKEN_KIND_DOUBLE_LITERAL
	}
}

//查找小数部分
func (jt *JavaTokenizer) scanFraction(pos int) {

	jt.skipUnderLine()
	reader := jt.reader
	if '0' <= reader.ch && reader.ch <= '9' {
		jt.scanDigits(pos, 10)
	}
	spos := reader.spos
	if reader.ch == 'e' || reader.ch == 'E' { // double c1 = 120.0e32; 这种
		reader.putRune(true)
		jt.skipUnderLine()
		if reader.ch == '+' || reader.ch == '-' {
			reader.putRune(true)
		}
		jt.skipUnderLine()
		if '0' <= reader.ch && reader.ch <= '9' {
			jt.scanDigits(pos, 10)
			return
		}
		jt.lexError(pos, "无效的.fp.lit pos:", pos)
		reader.spos = spos
	}
}
