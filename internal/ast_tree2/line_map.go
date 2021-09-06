package ast_tree2

/**
 *
 * @author hushengdong
 */
type LineMap interface {

	/**
	 * Find the start position of a line.
	 */
	GetStartPosition(line int64) int
	/**
	 * Find the position corresponding to a (line,column).
	 *
	 * @param   line    line number (beginning at 1)
	 * @param   column  tab-expanded column number (beginning 1)
	 *
	 * @return  position of character
	 * @throws  IndexOutOfBoundsException
	 *           if {@code line < 1}
	 *           if {@code line > no. of lines}
	 */
	GetPosition(line int64, column int64) int64

	/**
	 * Find the line containing a position; a line termination
	 * character is on the line it terminates.
	 *
	 * @param   pos  character offset of the position
	 * @return the line number of pos (first line is 1)
	 */
	GetLineNumber(pos int64) int64

	/**
	 * Find the column for a character position.
	 * Tab characters preceding the position on the same line
	 * will be expanded when calculating the column number.
	 *
	 * @param  pos   character offset of the position
	 * @return       the tab-expanded column number of pos (first column is 1)
	 */
	GetColumnNumber(pos int64) int64
}
