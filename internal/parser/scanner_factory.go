package parser

func GetScannerLexerFromFactory(path string) Scanner {

	return NewScannerLexer(path)
}
