package main

import (
	"fmt"
	"os"
)

/**
 * 这里写一个简单的入口，输入的内容是文件列表，只能是java的源文件
 * 这里先不加参数
 * @author hushengdong
 */
func main() {

	args := os.Args
	len := len(args)
	if len <= 0 {
		fmt.Println("请输入要编译的文件的路径")
		return
	}
	// compiler(args[1:])

	f := []string{"D:\\test.java"}
	//f := []string{"/tmp/test.java"}
	compiler(f)

	// v3 := lexical.WordLexicalV3{}
	// for _,path := range f {
	//	v3.LexicalAnalysis(path)
	// }
}
