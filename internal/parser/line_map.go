package parser

// 记录字符的绝对位置和行号的关系
type lineMap struct {
	lineMap map[int]int
}

func NewLineMap() *lineMap {

	m := lineMap{}
	m.lineMap = make(map[int]int)
	return &m
}

//GetLineNum 字符位置所在的行号
func (lm *lineMap) GetLineNum(num int) int {

	return lm.lineMap[num]
}
