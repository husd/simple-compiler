package ast_tree

/**
 *
 * @author hushengdong
 */

/**
 * 修饰符 例如： private public protected static 等
 */
type JCModifiers struct {
	Flags       int64
	Annotations []JCAnnotation
}
