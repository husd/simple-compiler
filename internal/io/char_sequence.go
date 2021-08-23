package io

import "fmt"

//CharSequence 只读的字符流 GO里没有字符
type CharSequence interface {
	/**
	 * length of the CharSequence
	 */
	Len() int
	/**
	 * charAt 为了降低难度，这里默认只支持数字 字母 汉字 三种类型
	 * 默认是unicode编码
	 */
	CharAt(index int) rune
	/**
	 * 读取byte
	 */
	ByteAt(index int) uint8
	/**
	 * 子序列
	 */
	SubCharSequence(start int, end int) string
}

func checkScope(start int, end int, max int) {

	valid := start >= 0 && end >= start && max >= end
	if !valid {
		panic(fmt.Sprintf("out of index start:%d end:%d max:%d", start, end, max))
	}
}
