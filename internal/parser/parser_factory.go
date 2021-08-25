package parser

func GetParserFromFactory(path string) Parser {

	return NewJavacParser(path)
}
