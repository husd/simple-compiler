package lexical

import "fmt"

// 这里主要是进行错误处理的，发生了错误的话，都需要到这里进行处理。
// 只是简单按照 编译原理 书上说的模块划分来规划代码结构，暂时先这样，如果过程中有不太对的地方，就再及时调整。

func HandleError(token *TokenTag, msg string) {

	//这里处理错误的方式，就是先简单的打印一下错误信息
	fmt.Errorf("error happened line: %d token: %s , error msg: %s", token.LineNum, token.Token, msg)
}
