package code

/**
 *
 * @author hushengdong
 */
type TypeTag struct {
	numberClass int
	superClass  int
	primitive   bool
}

//TODO
var TYPE_TAG_INT = &TypeTag{1, 1, true}
var TYPE_TAG_BOOL = &TypeTag{0, 0, true}
var TYPE_TAG_CHAR = &TypeTag{3, 3, true}
