package parser

import (
	"fmt"
	"husd.com/v0/util"
	"io/ioutil"
	"unicode/utf8"
)

// 实际为utf8解析 utf8 without bom
type UnicodeReader struct {
	buf              []byte //所有的数组
	size             int    // 数组的大小
	bp               int    // 当前读到那个位置了 byte position
	ch               rune   // 当前的位置的rune
	chLen            int    // 当前的位置的rune占用的字节数
	lastConversionBp int    // 最后一次转换的unicode的位置

	schStart int // 缓存的字符的开始
	schEnd   int // 结束
}

func NewUnicodeReader(buf []byte) *UnicodeReader {

	reader := UnicodeReader{}

	reader.buf = buf
	reader.size = len(buf)
	reader.bp = 0
	reader.ch = rune(-1)
	reader.chLen = 0
	reader.lastConversionBp = -1

	reader.schStart = 0
	reader.schEnd = 0

	return &reader
}

func NewUnicodeReaderFromFile(path string) *UnicodeReader {

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		panic("读取文件错误：" + path)
	}
	reader := UnicodeReader{}
	reader.bp = 0
	reader.size = len(buf)
	reader.buf = buf
	return &reader
}

//调用这个方法之后，会移动指针到下一个位置
func (reader *UnicodeReader) ReadRune() (bool, rune, int) {

	pos := reader.bp
	succ, res, count := reader.CurrentRune()
	if !succ {
		return succ, res, count
	}
	reader.bp = pos + count
	reader.ch = res
	//这里表示读到了类似 \uFF41 这样的字符，就需要尝试看看是不是转换unicode
	//if res == '\\' {
	//	reader.convertUnicodeText()
	//}
	return true, res, count
}

func (reader *UnicodeReader) CurrentRune() (bool, rune, int) {

	if reader.bp >= reader.size {
		return false, -1, -1
	}
	currentByte := reader.CurrentByte()
	succ, count := utf8Start(currentByte)
	if !succ {
		panic("解析utf8编码失败")
	}
	pos := reader.bp
	var res int32
	switch count {
	case 1:
		res, _ = utf8.DecodeRune(reader.buf[pos : pos+1])
		break
	case 2:
		res, _ = utf8.DecodeRune(reader.buf[pos : pos+2])
		break
	case 3:
		res, _ = utf8.DecodeRune(reader.buf[pos : pos+3])
		break
	case 4:
		res, _ = utf8.DecodeRune(reader.buf[pos : pos+4])
		break
	default:
		return false, -1, -1
	}
	return true, res, count
}

func (reader *UnicodeReader) CurrentByte() uint8 {

	reader.checkPos(reader.bp)
	return reader.buf[reader.bp]
}

func (reader *UnicodeReader) CurrentPos() int {

	return reader.bp
}

func (reader *UnicodeReader) ByteAt(pos int) uint8 {

	reader.checkPos(pos)
	return reader.buf[pos]
}

func (reader *UnicodeReader) checkPos(pos int) {

	if pos >= reader.size {
		panic(fmt.Sprintf("out of index pos is : %d", pos))
	}
}

func (reader *UnicodeReader) SubByteArray(start int, end int) []byte {

	return reader.buf[start:end]
}

//读取下一个字符
func (reader *UnicodeReader) ScanNextChar() {

	//javac里处理了原生的unicode \uFF41 这里我们不处理这样的字符了
	reader.ReadRune()
}

/** Convert an ASCII digit from its base (8, 10, or 16)
 *  to its value.
 */
// 把当前的 ch ，转换为对应的值，base值的是进值 例如 8 10 16 例如 \uFF41 转换为：0x10
func (reader *UnicodeReader) digit(bp int, base int) rune {

	//TODO husd
	panic("请实现这个方法 UnicodeReader.digit")
}

//Append a character 记录以下读到的字符
func (reader *UnicodeReader) putRune(scan bool) {

	//TODO husd 确定下schStart的值
	reader.schEnd = reader.schEnd + reader.chLen
	// 是否需要一个sbuf sp呢？
	if scan {
		reader.ReadRune()
	}
}

func (reader *UnicodeReader) name() *util.Name {

	n := util.Name{}
	n.NameStr = string(reader.buf[0:reader.schEnd])

	return &n
}

// 是否是 0 10 110 1110 这样的开头的格式 如果是
func utf8Start(b uint8) (bool, int) {

	// 0XXXXXXX
	if (b >> 7) == 0 {
		return true, 1
	}
	// 110XXXXX
	if (b >> 5) == 6 {
		return true, 2
	}
	// 1110XXXX
	if (b >> 4) == 14 {
		return true, 3
	}
	// 11110XXXX
	if (b >> 4) == 30 {
		return true, 4
	}
	return false, 0
}

// 10XXXXXX
func utf8Start10(b uint8) bool {

	return (b >> 6) == 2
}

func checkUtf8Start10(b uint8) {

	if !utf8Start10(b) {
		panic("无效的utf8字符")
	}
}
