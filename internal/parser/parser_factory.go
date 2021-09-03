package parser

import "husd.com/v0/util"

/**
 *
 * @author hushengdong
 */
func GetParserFromFactory(path string, context *util.Context) Parser {

	parser := NewJavacParser(path, context)
	return parser
}
