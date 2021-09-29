package ast_tree2

/**
 *
 * @author hushengdong
 */
type CompilationUnitTreeV2 interface {
	GetTreeType() TreeType

	//-

	GetPackageName() ExpressionTreeV2
	//List<? extends ImportTree> getImports();
	GetTypeDecls() *[]TreeV2
	//JavaFileObject getSourceFile();

	/**
	 * Gets the line map for this compilation unit, if available.
	 * Returns null if the line map is not available.
	 * @return the line map for this compilation unit
	 */
	//LineMap getLineMap();
}
