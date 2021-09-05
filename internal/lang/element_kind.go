package lang

/**
 *
 * @author hushengdong
 */
type ElementKind interface {
	IsClass() bool
	IsInterface() bool
	IsField() bool
}
