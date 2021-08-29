package parser

func GetParserFromFactory(path string) Parser {

	parser := NewJavacParser(path)
	return parser
}
