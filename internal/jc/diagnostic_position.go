package jc

/**
 * DiagnosticPosition提供有关文件中位置的信息
 * 这引起了一个诊断。它总是定义一个“首选”位置
 * 它最准确地定义了诊断的位置，也可能
 * 提供跨越该位置的相关树节点。
 */
type DiagnosticPosition interface {

	/** Gets the tree node, if any, to which the diagnostic applies. */
	getTree() *JCTree
	/** If there is a tree node, get the start position of the tree node.
	 *  Otherwise, just returns the same as getPreferredPosition(). */
	getStartPosition() int
	/** Get the position within the file that most accurately defines the
	 *  location for the diagnostic. */
	getPreferredPosition() int
	/** If there is a tree node, and if endPositions are available, get
	 *  the end position of the tree node. Otherwise, just returns the
	 *  same as getPreferredPosition(). */
	getEndPosition(endPosTable *EndPosTable) int
}
