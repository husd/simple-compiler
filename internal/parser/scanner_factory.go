package parser

import "husd.com/v0/util"

func GetScannerLexerFromFactory(path string, context *util.Context) *Scanner {

	return NewScannerLexer(path, context)
}
