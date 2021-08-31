package parser

import "husd.com/v0/util"

func GetParserFromFactory(path string, context *util.Context) Parser {

	parser := NewJavacParser(path, context)
	return parser
}
