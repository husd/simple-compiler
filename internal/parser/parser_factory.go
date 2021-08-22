package parser

import "husd.com/v0/io"

func GetParserFromFactory(sequence *io.CharSequence) Parser {

	return NewJavacParser(sequence)
}
