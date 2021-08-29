package jc

import "container/list"

//JCCompilationUnit java源代码里的所有东西，都在这里有 这里先写一个空方法
type JCCompilationUnit struct {

	//包 JCAnnotation
	Annotations list.List
}
