package tree

/**
 *
 * @author hushengdong
 */
type LineMap struct {
	start *[]int // start position of each line start[3]= 100 表示 第3行的开头是 pos = 100
	max   int    // 切片的长度 一共有多少行
	/**
	 * 因为编译器会频繁的访问这个MAP，而且每次pos都是递增的，所以可以记录一下最后一次的访问结果
	 * 这样可以加快2分查找的速度
	 */
	lastPos  int
	lastLine int
}

const firstLine = int(1)
const firstColumn = 1

func NewLineMap(charArr *[]rune, max int) *LineMap {

	l := &LineMap{}

	c := 0 // 数组的下标，多少行
	i := 0 // pos 默认是0 表示从0开始
	lineBuf := make([]int, max, max)
	for i < max {
		lineBuf[c] = i
		c++
		for { // 尝试读取下一个换行
			ch := (*charArr)[i]
			if ch == '\r' || ch == '\n' { // 读到了换行，跳出循环
				if ch == '\r' && (i+1) < max && (*charArr)[(i+1)] == '\n' {
					i = i + 2 // 如果读到了 \r\n ，就表示要直接读2个字符，所以i = i+2
				} else {
					i++
				}
			} else if ch == '\t' { // 记录制表符
				// setTabPosition(i)
			}
			i++
			if i >= max {
				break
			}
		}
	}

	l.start = &lineBuf
	l.max = max
	return l
}

func (l *LineMap) GetStartPosition(line int) int {

	return (*l.start)[line-firstLine]
}

func (l *LineMap) GetPosition(line int, column int) int {

	startPos := (*l.start)[line-firstLine]
	return startPos + column - firstColumn
}

func (l *LineMap) GetLineNumber(pos int) int {

	lastPos := l.lastPos
	lastLine := l.lastLine

	if pos == lastPos {
		return lastLine
	}
	l.lastPos = pos
	// 返回结果在 low ~ high 标准的2分查找
	low := 0
	high := l.max - 1
	for low <= high {
		mid := low + (high-low)>>1
		midVal := (*l.start)[mid]
		if pos > midVal {
			low = midVal + 1
		} else if pos < midVal {
			high = midVal + 1
		} else {
			// find it
			l.lastLine = mid + 1
			return lastLine
		}
	}
	l.lastLine = low
	return lastLine // pos is on this line
}

func (l *LineMap) GetColumnNumber(pos int) int {

	return pos - (*l.start)[l.GetLineNumber(pos)-firstLine] + firstColumn
}
