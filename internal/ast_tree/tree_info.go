package ast_tree

import "husd.com/v0/jc"

/**
 *
 * @author hushengdong
 */

func treeTag2TreeKind(tag jc.AstTreeTag) *AstTreeNodeKind {

	switch tag {
	// Postfix expressions
	case jc.POSTINC: // _ ++
		return TREE_POSTFIX_INCREMENT
	case jc.POSTDEC: // _ --
		return TREE_POSTFIX_DECREMENT

	// Unary operators
	case jc.PREINC: // ++ _
		return TREE_PREFIX_INCREMENT
	case jc.PREDEC: // -- _
		return TREE_PREFIX_DECREMENT
	case jc.POS: // +
		return TREE_UNARY_PLUS
	case jc.NEG: // -
		return TREE_UNARY_MINUS
	case jc.COMPL: // ~
		return TREE_BITWISE_COMPLEMENT
	case jc.NOT: // !
		return TREE_LOGICAL_COMPLEMENT

	// Binary operators

	// Multiplicative operators
	case jc.MUL: // *
		return TREE_MULTIPLY
	case jc.DIV: // /
		return TREE_DIVIDE
	case jc.MOD: // %
		return TREE_REMAINDER

	// Additive operators
	case jc.PLUS: // +
		return TREE_PLUS
	case jc.MINUS: // -
		return TREE_MINUS

	// Shift operators
	case jc.SL: // <<
		return TREE_LEFT_SHIFT
	case jc.SR: // >>
		return TREE_RIGHT_SHIFT
	case jc.USR: // >>>
		return TREE_UNSIGNED_RIGHT_SHIFT

	// Relational operators
	case jc.LT: // <
		return TREE_LESS_THAN
	case jc.GT: // >
		return TREE_GREATER_THAN
	case jc.LE: // <=
		return TREE_LESS_THAN_EQUAL
	case jc.GE: // >=
		return TREE_GREATER_THAN_EQUAL

	// Equality operators
	case jc.EQ: // ==
		return TREE_EQUAL_TO
	case jc.NE: // !=
		return TREE_NOT_EQUAL_TO

	// Bitwise and logical operators
	case jc.BITAND: // &
		return TREE_AND
	case jc.BITXOR: // ^
		return TREE_XOR
	case jc.BITOR: // |
		return TREE_OR

	// Conditional operators
	case jc.AND: // &&
		return TREE_CONDITIONAL_AND
	case jc.OR: // ||
		return TREE_CONDITIONAL_OR

	// Assignment operators
	case jc.MUL_ASG: // *=
		return TREE_MULTIPLY_ASSIGNMENT
	case jc.DIV_ASG: // /=
		return TREE_DIVIDE_ASSIGNMENT
	case jc.MOD_ASG: // %=
		return TREE_REMAINDER_ASSIGNMENT
	case jc.PLUS_ASG: // +=
		return TREE_PLUS_ASSIGNMENT
	case jc.MINUS_ASG: // -=
		return TREE_MINUS_ASSIGNMENT
	case jc.SL_ASG: // <<=
		return TREE_LEFT_SHIFT_ASSIGNMENT
	case jc.SR_ASG: // >>=
		return TREE_RIGHT_SHIFT_ASSIGNMENT
	case jc.USR_ASG: // >>>=
		return TREE_UNSIGNED_RIGHT_SHIFT_ASSIGNMENT
	case jc.BITAND_ASG: // &=
		return TREE_AND_ASSIGNMENT
	case jc.BITXOR_ASG: // ^=
		return TREE_XOR_ASSIGNMENT
	case jc.BITOR_ASG: // |=
		return TREE_OR_ASSIGNMENT

	// Null check (implementation detail), for example, __.getClass()
	case jc.NULLCHK:
		return TREE_OTHER

	case jc.ANNOTATION:
		return TREE_ANNOTATION
	case jc.TYPE_ANNOTATION:
		return TREE_TYPE_ANNOTATION

	default:
		return TREE_NIL
	}
}
