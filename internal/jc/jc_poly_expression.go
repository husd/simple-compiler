package jc

/**
 *
 * @author hushengdong
 */
type AbstractJCPolyExpression struct {
	*AbstractJCExpression
	PolyKind polyKind
}

type polyKind int

const (
	/** poly expression to be treated as a standalone expression */
	POLY_KIND_STANDALONE polyKind = 1
	/** true poly expression */
	POLY_KIND_POLY = 2
)

//这里并没有实现抽象类的方法，需要它的子类去实现
func NewJCPolyExpression() *AbstractJCPolyExpression {

	res := &AbstractJCPolyExpression{}
	res.AbstractJCExpression = NewJCExpression()
	return res
}
