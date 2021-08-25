package parser

import (
	"io/ioutil"
	"unicode/utf8"
)

// 实际为utf8解析 utf8 without bom
type UnicodeReader struct {
	buf  []byte //所有的数组
	size int    // 数组的大小
	pos  int    // 当前读到那个位置了
	ch   rune   // 当前的位置的rune
}

func NewUnicodeReader(buf []byte) *UnicodeReader {

	reader := UnicodeReader{}
	reader.pos = 0
	reader.size = len(buf)
	reader.buf = buf

	return &reader
}

func NewUnicodeReaderFromFile(path string) *UnicodeReader {

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		panic("读取文件错误：" + path)
	}
	reader := UnicodeReader{}
	reader.pos = 0
	reader.size = len(buf)
	reader.buf = buf
	return &reader
}

//调用这个方法之后，会移动指针到下一个位置
func (reader *UnicodeReader) ReadRune() (rune, int) {

	pos := reader.pos
	res, count := reader.CurrentRune()
	reader.pos = pos + count
	reader.ch = res
	return res, count
}

func (reader *UnicodeReader) CurrentRune() (rune, int) {

	currentByte := reader.CurrentByte()
	succ, count := utf8Start(currentByte)
	if !succ {
		panic("解析utf8编码失败")
	}
	pos := reader.pos
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
	}
	return res, count
}

func (reader *UnicodeReader) CurrentByte() uint8 {

	reader.checkPos(reader.pos)
	return reader.buf[reader.pos]
}

func (reader *UnicodeReader) CurrentPos() int {

	return reader.pos
}

func (reader *UnicodeReader) ByteAt(pos int) uint8 {

	reader.checkPos(pos)
	return reader.buf[pos]
}

func (reader *UnicodeReader) checkPos(pos int) {

	if pos >= reader.size {
		panic("out of index")
	}
}

func (reader *UnicodeReader) SubByteArray(start int, end int) []byte {

	return reader.buf[start:end]
}

//读取下一个字符
func (reader *UnicodeReader) ScanNextChar() {

	//javac里处理了原生的unicode \uFFFE 这里我们不处理这样的字符了
	reader.ReadRune()
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
