package parser

type names struct {

	// keywords
	_class   name
	_default name
	_super   name
	_this    name
}

type name struct {
	Name  string
	Table map[string]string
}
