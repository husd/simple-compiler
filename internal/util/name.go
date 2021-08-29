package util

type Name struct {
	NameStr string
	index   int // 符号表的索引
}

func (name *Name) getIndex() int {

	return name.index
}
