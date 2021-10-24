package parser

/**
* 运算符优先级
* @author hushengdong
 */
type PrecOp int

const notExpression PrecOp = -1 // not an expression
const noPrec PrecOp = 0         // no enclosing expression
const assignPrec PrecOp = 1
const assignopPrec PrecOp = 2
const condPrec PrecOp = 3
const orPrec PrecOp = 4
const andPrec PrecOp = 5
const bitorPrec PrecOp = 6
const bitxorPrec PrecOp = 7
const bitandPrec PrecOp = 8
const eqPrec PrecOp = 9
const ordPrec PrecOp = 10
const shiftPrec PrecOp = 11
const addPrec PrecOp = 12
const mulPrec PrecOp = 13
const prefixPrec PrecOp = 14
const postfixPrec PrecOp = 15
const precCount PrecOp = 16

func opPrec(op TreeNodeTag) PrecOp {

	switch op {
	case pos,
		neg,
		not,
		compl,
		preinc,
		predec:
		return prefixPrec
	case postinc,
		postdec,
		nullchk:
		return postfixPrec
	case assign_:
		return assignPrec
	case bitor_asg,
		bitxor_asg,
		bitand_asg,
		sl_asg,
		sr_asg,
		usr_asg,
		plus_asg,
		minus_asg,
		mul_asg,
		div_asg,
		mod_asg:
		return assignopPrec
	case or:
		return orPrec
	case and:
		return andPrec
	case eq,
		ne:
		return eqPrec
	case lt,
		gt,
		le,
		ge:
		return ordPrec
	case bitor:
		return bitorPrec
	case bitxor:
		return bitxorPrec
	case bitand:
		return bitandPrec
	case sl,
		sr,
		usr:
		return shiftPrec
	case plus,
		minus:
		return addPrec
	case mul,
		div,
		mod:
		return mulPrec
	case typetest_:
		return ordPrec
	default:
		panic("无效的tree_node_tag :")
	}
}
