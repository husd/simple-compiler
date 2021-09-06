package lang

/**
 * 作用域修饰符
 * See JLS sections 8.1.1, 8.3.1, 8.4.3, 8.8.3, and 9.1.1.
 * https://docs.oracle.com/javase/specs/jls/se8/html/jls-8.html#jls-8.3.1.1
 * @author hushengdong
 */
type Modifier int

const (
	PUBLIC int = iota
	PROTECTED
	PRIVATE
	ABSTRACT
	DEFAULT // since 1.8
	STATIC
	FINAL
	TRANSIENT
	VOLATILE
	SYNCHRONIZED
	NATIVE
	STRICTFP
)
