package parser

/**
* 运算符优先级
* @author hushengdong
 */

const notExpression int = -1 // not an expression
const noPrec int = 0         // no enclosing expression
const assignPrec int = 1
const assignopPrec int = 2
const condPrec int = 3
const orPrec int = 4
const andPrec int = 5
const bitorPrec int = 6
const bitxorPrec int = 7
const bitandPrec int = 8
const eqPrec int = 9
const ordPrec int = 10
const shiftPrec int = 11
const addPrec int = 12
const mulPrec int = 13
const prefixPrec int = 14
const postfixPrec int = 15
const precCount int = 16

func opPrec(op TreeNodeTag) int {

	switch op {
	case Tree_node_tag_pos,
		Tree_node_tag_neg,
		Tree_node_tag_not,
		Tree_node_tag_compl,
		Tree_node_tag_preinc,
		Tree_node_tag_predec:
		return prefixPrec
	case Tree_node_tag_postinc,
		Tree_node_tag_postdec,
		Tree_node_tag_nullchk:
		return postfixPrec
	case Tree_node_tag_assign:
		return assignPrec
	case Tree_node_tag_bitor_asg,
		Tree_node_tag_bitxor_asg,
		Tree_node_tag_bitand_asg,
		Tree_node_tag_sl_asg,
		Tree_node_tag_sr_asg,
		Tree_node_tag_usr_asg,
		Tree_node_tag_plus_asg,
		Tree_node_tag_minus_asg,
		Tree_node_tag_mul_asg,
		Tree_node_tag_div_asg,
		Tree_node_tag_mod_asg:
		return assignopPrec
	case Tree_node_tag_or:
		return orPrec
	case Tree_node_tag_and:
		return andPrec
	case Tree_node_tag_eq,
		Tree_node_tag_ne:
		return eqPrec
	case Tree_node_tag_lt,
		Tree_node_tag_gt,
		Tree_node_tag_le,
		Tree_node_tag_ge:
		return ordPrec
	case Tree_node_tag_bitor:
		return bitorPrec
	case Tree_node_tag_bitxor:
		return bitxorPrec
	case Tree_node_tag_bitand:
		return bitandPrec
	case Tree_node_tag_sl,
		Tree_node_tag_sr,
		Tree_node_tag_usr:
		return shiftPrec
	case Tree_node_tag_plus,
		Tree_node_tag_minus:
		return addPrec
	case Tree_node_tag_mul,
		Tree_node_tag_div,
		Tree_node_tag_mod:
		return mulPrec
	case Tree_node_tag_typetest:
		return ordPrec
	default:
		panic("无效的tree_node_tag :")
	}
}
