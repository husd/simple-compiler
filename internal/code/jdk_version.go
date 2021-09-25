package code

/**
 *
 * @author hushengdong
 */

// 定义了jdk版本 能做的一些事情 这里仅仅做个示意 不做那么多复杂内容
type JVersion int

const (
	JDK1 JVersion = 1
	JDK2 JVersion = 2
	JDK3 JVersion = 3
	JDK4 JVersion = 4
	JDK5 JVersion = 5
	JDK6 JVersion = 6
	JDK7 JVersion = 7
	JDK8 JVersion = 8
)

/**
 * 允许foreach的条件 JDK5 之后的版本
 */
func AllowForeach(version JVersion) bool {

	return version >= JDK5
}

//十六进制浮点文本
func AllowHexFloats(v JVersion) bool {
	return false
}

//先不支持泛型
func AllowGenerics(v JVersion) bool {

	return false
}
