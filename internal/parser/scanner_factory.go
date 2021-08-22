package parser

import "husd.com/v0/io"

func GetScannerLexerFromFactory(sequence *io.CharSequence) Scanner {

	return NewScannerLexer(sequence)
}
