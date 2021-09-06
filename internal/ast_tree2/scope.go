package ast_tree2

/**
 * Interface for determining locally available program elements, such as
 * local variables and imports.
 * Upon creation, a Scope is associated with a given program position;
 * for example, a {@linkplain Tree tree node}. This position may be used to
 * infer an enclosing method and/or class.
 *
 * <p>A Scope does not itself contain the details of the elements corresponding
 * to the parameters, methods and fields of the methods and classes containing
 * its position. However, these elements can be determined from the enclosing
 * elements.
 *
 * <p>Scopes may be contained in an enclosing scope. The outermost scope contains
 * those elements available via "star import" declarations; the scope within that
 * contains the top level elements of the compilation unit, including any named
 * imports.
 *
 * @author hushengdong
 */
type Scope interface {

	/**
	 * Returns the enclosing scope.
	 */
	getEnclosingScope() Scope

	/**
	 * Returns the innermost type element containing the position of this scope
	 */
	// getEnclosingClass() TypeElement

	/**
	 * Returns the innermost executable element containing the position of this scope.
	 */
	// getEnclosingMethod() ExecutableElement

	/**
	 * Returns the elements directly contained in this scope.
	 */
	// getLocalElements() *[]Element
}
