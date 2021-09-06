package parser

import (
	"fmt"
	"husd.com/v0/compiler"
	"husd.com/v0/util"
	"io/ioutil"
	"unicode/utf8"
)

/**
 * 实际为utf8解析 utf8 without bom 这里主要实现了从字节数组中解析utf8编码，然后
 * 提供了一个切片，缓存了当前识别到的字符，切片可以自动扩容。
 * @author hushengdong
 */
type UnicodeReader struct {
	buf              []byte //所有的数组
	size             int    // 数组的大小
	bp               int    // 当前读到那个位置了 byte position 这个位置还没被读取
	ch               rune   // 最后一次读到的 rune
	chLen            int    // 最后一次读到的 rune 占用的字节数
	lastConversionBp int    // 最后一次转换的unicode的位置

	sbuf    []byte // 切片，所有的已扫描到的数据 写个固定的数组
	spos    int    // 已扫描的数据的长度
	utf8Buf []byte //

	lineNum int // 多少行
	linePos int // 位置

	end bool //读到了末尾
}

func NewUnicodeReader(bufPoint *[]byte) *UnicodeReader {

	reader := UnicodeReader{}
	buf := (*bufPoint)
	reader.buf = buf
	reader.size = len(buf)
	reader.bp = 0
	reader.ch = rune(-1)
	reader.chLen = 0
	reader.lastConversionBp = -1
	reader.lineNum = 1
	reader.linePos = 0

	const SBUF_LEN = 1
	const SBUF_MAX = 1
	sbuf := make([]byte, SBUF_LEN, SBUF_MAX)
	reader.sbuf = sbuf
	reader.spos = 0
	reader.utf8Buf = make([]byte, 4, 4)

	reader.end = false

	reader.ScanRune()
	return &reader
}

func NewUnicodeReaderFromFile(path string) *UnicodeReader {

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		panic("读取文件错误：" + path)
	}
	return NewUnicodeReader(&buf)
}

//查看下一个rune，不移动指针
func (reader *UnicodeReader) seenNextRune() (rune, bool) {

	pos := reader.bp
	if pos >= reader.size {
		return -1, false
	}
	succ, res, _ := reader.runeAt(pos)
	return res, succ
}

//调用这个方法之后，会移动指针到下一个位置
func (reader *UnicodeReader) ScanRune() bool {

	pos := reader.bp
	if pos >= reader.size {
		reader.end = true
		reader.ch = -1 // TODO husd 考虑返回-1是否合适 目的是为了switch能走到default
		return false
	}
	succ, res, count := reader.runeAt(pos)
	if !succ {
		return false
	}
	reader.chLen = count
	reader.bp = pos + count
	reader.ch = res
	//这里表示读到了类似 \uFF41 这样的字符，就需要尝试看看是不是转换unicode
	//if res == '\\' {
	//	reader.convertUnicode()
	//}
	if compiler.DEBUG_SCAN_RUNE {
		fmt.Println("----debug----第[", reader.lineNum, "]行 pos :", reader.linePos, "  ScanRune ch : ", string(reader.ch))
	}
	if reader.ch == Layout_char_lf {
		reader.lineNum++
		reader.linePos = 0
	} else {
		reader.linePos++
	}
	return true
}

func (reader *UnicodeReader) peekChar() rune {

	_, res, _ := reader.runeAt(reader.bp)
	return res
}

func (reader *UnicodeReader) runeAt(pos int) (bool, rune, int) {

	if pos >= reader.size {
		return false, 0, 0
	}
	currentByte := reader.ByteAt(pos)
	succ, count := utf8Start(currentByte)
	if !succ {
		panic("解析utf8编码失败")
	}
	var res int32
	switch count {
	case 1:
		res, _ = utf8.DecodeRune(reader.buf[pos : pos+1])
	case 2:
		res, _ = utf8.DecodeRune(reader.buf[pos : pos+2])
	case 3:
		res, _ = utf8.DecodeRune(reader.buf[pos : pos+3])
	case 4:
		res, _ = utf8.DecodeRune(reader.buf[pos : pos+4])
	default:
		return false, 0, 0
	}
	return true, res, count
}

func (reader *UnicodeReader) currentByte() uint8 {

	return reader.ByteAt(reader.bp)
}

func (reader *UnicodeReader) CurrentPos() int {

	return reader.bp
}

func (reader *UnicodeReader) CurrentRune() rune {

	return reader.ch
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

/** Convert an ASCII digit from its base (8, 10, or 16)
 *  to its value.
 */
// 把当前的 ch ，转换为对应的值，base值的是进值 例如 8 10 16 例如 \uFF41 转换为：0x10
// bp只是记录错误日志使用 如果当前字符不是数字，那么就返回-1
// 词法分析器
const digitStr string = "0123456789abcdef"

func (reader *UnicodeReader) digit(bp int, base int) rune {

	ch := reader.ch
	res := digitRune(ch, base)
	if res >= 0 && ch > 0x7f {
		fmt.Println("读到了无效的ascii数字，请检查ch :", ch, " bp: ", bp+1, " base:", base)
		// TODO husd DEBUG
		reader.ch = rune(digitStr[res])
	}
	return res
}

//Append a character 记录以下读到的字符
func (reader *UnicodeReader) putRune(scan bool) {

	if reader.bp <= 0 {
		return
	}
	spos := reader.spos
	reader.ensureCapacity(spos, reader.chLen)
	start := reader.bp - reader.chLen
	switch reader.chLen {
	case 1:
		reader.sbuf[spos+0] = reader.buf[start]
	case 2:
		reader.sbuf[spos+0] = reader.buf[start]
		reader.sbuf[spos+1] = reader.buf[start+1]
	case 3:
		reader.sbuf[spos+0] = reader.buf[start]
		reader.sbuf[spos+1] = reader.buf[start+1]
		reader.sbuf[spos+2] = reader.buf[start+2]
	case 4:
		reader.sbuf[spos+0] = reader.buf[start]
		reader.sbuf[spos+1] = reader.buf[start+1]
		reader.sbuf[spos+2] = reader.buf[start+2]
		reader.sbuf[spos+3] = reader.buf[start+3]
	}
	reader.spos = spos + reader.chLen
	if scan {
		reader.ScanRune()
	}
}

func (reader *UnicodeReader) putChar(r rune) {

	reader.putRuneChar(r, false)
}

func (reader *UnicodeReader) putRuneChar(r rune, scan bool) {

	spos := reader.spos
	if r >= 0 && r <= 255 {
		reader.ensureCapacity(spos, 1)
		reader.sbuf[spos] = uint8(r)
		reader.spos = spos + 1
	} else {
		//这里要动态判断r有几个字节，所以再解析回来utf8编码
		//由于这里长度固定不超过4个，不会触发切片的扩容，所以可以直接把reader.utf8Buf传过去
		//不需要传递指针
		num := utf8.EncodeRune(reader.utf8Buf, r)
		reader.ensureCapacity(spos, num)
		switch num {
		case 1:
			reader.sbuf[spos+0] = reader.utf8Buf[0]
		case 2:
			reader.sbuf[spos+0] = reader.utf8Buf[0]
			reader.sbuf[spos+1] = reader.utf8Buf[1]
		case 3:
			reader.sbuf[spos+0] = reader.utf8Buf[0]
			reader.sbuf[spos+1] = reader.utf8Buf[1]
			reader.sbuf[spos+2] = reader.utf8Buf[2]
		case 4:
			reader.sbuf[spos+0] = reader.utf8Buf[0]
			reader.sbuf[spos+1] = reader.utf8Buf[1]
			reader.sbuf[spos+2] = reader.utf8Buf[2]
			reader.sbuf[spos+3] = reader.utf8Buf[3]
		}
		reader.spos = reader.spos + num
	}
	if scan {
		reader.ScanRune()
	}
}

func (reader *UnicodeReader) name() *util.Name {

	n := util.Name{}
	n.NameStr = string(reader.sbuf[0:reader.spos])

	return &n
}

func (reader *UnicodeReader) isUnicode() bool {

	return reader.lastConversionBp == reader.bp
}

func (reader *UnicodeReader) skipChar() {

	pos := reader.bp
	succ, _, count := reader.runeAt(reader.bp + reader.chLen)
	if succ {
		reader.bp = pos + count
		reader.chLen = count
		//reader.ch 这个不变，还是原来的值
	}
}

func (reader *UnicodeReader) hasNext() bool {

	return reader.bp < reader.size
}

/** Scan surrogate pairs.  If 'ch' is a high surrogate and
 *  the next character is a low surrogate, then put the low
 *  surrogate in 'ch', and return the high surrogate.
 *  otherwise, just return 0.
 */
func (reader *UnicodeReader) scanSurrogates() rune {

	const surrogatesSupported bool = false

	if surrogatesSupported {
		//先不管
	}
	return 0
}

// 是否是 0 110 1110 11110 这样的开头的格式 我们读取的是字节数组
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
	// 11110XXX
	if (b >> 3) == 30 {
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

/**
 * spos 当前写到切片的什么位置了
 * need 本次需要写入多个
 */
func (reader *UnicodeReader) ensureCapacity(spos int, need int) {

	currentCap := cap(reader.sbuf)
	if spos+need > currentCap {
		newCap := calcNewLength(currentCap, spos+need)
		newSbuf := make([]byte, newCap, newCap) //len设置为cap，这样才可以在任意位置写入
		copy(newSbuf, reader.sbuf)
		reader.sbuf = newSbuf
	}
}

func calcNewLength(len int, max int) int {

	for len < max+1 {
		len = len * 2
	}
	return len
}

func (reader *UnicodeReader) reachEnd() bool {

	return reader.end
}
