package tool

type JavaFileSrc int

const (
	java  JavaFileSrc = 1
	class JavaFileSrc = 2
	html  JavaFileSrc = 3
	other JavaFileSrc = 4
)

type JavaFileObj struct {
	src  JavaFileSrc // 0java 1class 2html 3 ""
	path string      //路径
}
