package compiler

//编译阶段
type CompileState int

const (
	INIT     CompileState = 0
	PARSE    CompileState = 1
	ENTER    CompileState = 2
	PROC     CompileState = 3
	ATTR     CompileState = 4
	FLOW     CompileState = 5
	TRAN     CompileState = 6
	UNLAMBDA CompileState = 7
	LOWER    CompileState = 8
	GENERATE CompileState = 9
)

const DEBUG bool = true

const DEBUG_TOKEN bool = DEBUG && true

const DEBUG_SCAN_RUNE bool = DEBUG && true
