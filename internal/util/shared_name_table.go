package util

/**
 *
 * @author hushengdong
 */
type SharedNameTable struct {
	HashMask  int     //The mask to be used for hashing
	NameArray []*Name //The hash table for names.
	Bytes     []*byte //The shared byte array holding all encountered names.

	HashMap map[string]*Name // 符号表
	nc      int              //The number of filled bytes in `names'.
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

/**
 * 目前先用一个hashmap来存。
 * 后续再优化 把key优化为字节数组，计算该字节数组的hash值，这样就省去了把
 * 字节数组转换为string的这一个步骤。TODO
 */
func (snt *SharedNameTable) fromString(s string) *Name {

	if n, ok := snt.HashMap[s]; ok {
		return n
	}
	n := &Name{}

	n.NameStr = s
	n.Index = 0 //TODO husd

	//put it to map
	snt.HashMap[s] = n
	return n
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
