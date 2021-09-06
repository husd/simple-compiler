package parser

/**
 *
 * @author hushengdong
 */
// 字面量是否允许下划线 例如： int a = 124_23 jdk1.7之后可以这么做了
const allowUnderscoresInLiterals bool = true

//默认 十六进制浮点文本 是false 意思就是 0x1.0p1 是否是一个有效的16进制浮点数字
const hexFloatsWork bool = false
const allowHexFloats bool = false

const allowBinaryLiterals bool = true

const allowAssert = false
