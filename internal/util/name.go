package util

/**
 *
 * @author hushengdong
 */
type Name struct {
	NameStr string
	Index   int // 符号表的索引

}

func (name *Name) getIndex() int {

	return name.Index
}
