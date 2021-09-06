package tree

/**
 *
 * @author hushengdong
 */
type LineMap struct {
	start *[]int // start position of each line
}

const firstLine = int(1)

func (l *LineMap) GetStartPosition(line int64) int {

	return (*l.start)[int(line)-firstLine]
}

func (l *LineMap) GetPosition(line int64, column int64) int64 {

	panic("implement me")
}

func (l *LineMap) GetLineNumber(pos int64) int64 {

	// 2分法 获取到第几行
	panic("implement me")
}

func (l *LineMap) GetColumnNumber(pos int64) int64 {

	panic("implement me")
}
