package parser

import "husd.com/v0/io"

/**
 * 这个类主要是解析字符的，unicode编码 这里本来要有一个各种编码的reader
 * 为了降低难度，先不写这个reader了，直接调用 CharSequence 来解析
 *
 * 可以把源代码里的 String name = "xiaoming 你好" 分解成:
 * String
 * name
 * =
 * "xiaoming 你好"
 *
 * 可以看出，并不是简单的按某个字符分割，还需要过滤出注释 包括单行注释和多行注释。
 *
 */

type JavaTokenizer struct {
	seq *io.CharSequence
}
