package util

type SharedNameTable struct {
	HashMask  int     //The mask to be used for hashing
	NameArray []*Name //The hash table for names.
	Bytes     []*byte //The shared byte array holding all encountered names.

	nc int //The number of filled bytes in `names'.

}

func NewSharedNameTable(c *Context) *SharedNameTable {

	snt := SharedNameTable{}

	return &snt
}

/**
 * 这个是符号表的一部分内容
 * 这里有一个计算hash的方法 如果hash冲突了，就向后找，但是具体的性能怎么样，还需要再测试一下
 * TODO husd 测试性能怎么样
 * 这个方法可以看出来，设计的一个通用的准则，就是设计一个底层方法，拥有多个参数，低级方法，然后再封装高级方法来暴露，
 * 低级方法往往不会暴露，或者不推荐直接使用。
 */
func (snt *SharedNameTable) fromUtf8Shared(bytes *[]byte, start int, length int) *Name {

	h := hash(bytes, start, length) & snt.HashMask
	n := snt.NameArray[h]

	//TODO husd 后续再考虑使用数组来
	return n
}

//还有一个简单的策略，就是存储一个hashmap，直接来表示这个表
func (snt *SharedNameTable) fromString(s string) *Name {

	b := []byte(s)
	return snt.fromUtf8Shared(&b, 0, len(b))
}

func hash(byteArray *[]byte, start int, length int) int {

	h := 0
	off := start
	for i := 0; i < length; i++ {
		h = (h << 5) - h + int((*byteArray)[off])
		off++
	}
	return h
}
