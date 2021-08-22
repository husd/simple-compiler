package io

//CharSequence 只读的字符流 GO里没有字符
type CharSequence interface {
	/**
	 * length of the CharSequence
	 */
	Len() int
	/**
	 * charAt
	 */
	CharAt(index int) uint8
	/**
	 * 子序列
	 */
	SubCharSequence(start int, end int) string
}
