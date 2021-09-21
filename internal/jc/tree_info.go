package jc

import "husd.com/v0/ast_tree2"

/**
 *
 * @author hushengdong
 */

func treeTag2TreeKind(tag JCTreeTag) *ast_tree2.TreeType {

	switch tag {
	// Postfix expressions
	case TREE_TAG_POSTINC: // _ ++
		return ast_tree2.AST_TREE_KIND_POSTFIX_INCREMENT
	case TREE_TAG_POSTDEC: // _ --
		return ast_tree2.AST_TREE_KIND_POSTFIX_DECREMENT

	// Unary operators
	case TREE_TAG_PREINC: // ++ _
		return ast_tree2.AST_TREE_KIND_PREFIX_INCREMENT
	case TREE_TAG_PREDEC: // -- _
		return ast_tree2.AST_TREE_KIND_PREFIX_DECREMENT
	case TREE_TAG_POS: // +
		return ast_tree2.AST_TREE_KIND_UNARY_PLUS
	case TREE_TAG_NEG: // -
		return ast_tree2.AST_TREE_KIND_UNARY_MINUS
	case TREE_TAG_COMPL: // ~
		return ast_tree2.AST_TREE_KIND_BITWISE_COMPLEMENT
	case TREE_TAG_NOT: // !
		return ast_tree2.AST_TREE_KIND_LOGICAL_COMPLEMENT

	// Binary operators

	// Multiplicative operators
	case TREE_TAG_MUL: // *
		return ast_tree2.AST_TREE_KIND_MULTIPLY
	case TREE_TAG_DIV: // /
		return ast_tree2.AST_TREE_KIND_DIVIDE
	case TREE_TAG_MOD: // %
		return ast_tree2.AST_TREE_KIND_REMAINDER

	// Additive operators
	case TREE_TAG_PLUS: // +
		return ast_tree2.AST_TREE_KIND_PLUS
	case TREE_TAG_MINUS: // -
		return ast_tree2.AST_TREE_KIND_MINUS

	// Shift operators
	case TREE_TAG_SL: // <<
		return ast_tree2.AST_TREE_KIND_LEFT_SHIFT
	case TREE_TAG_SR: // >>
		return ast_tree2.AST_TREE_KIND_RIGHT_SHIFT
	case TREE_TAG_USR: // >>>
		return ast_tree2.AST_TREE_KIND_UNSIGNED_RIGHT_SHIFT

	// Relational operators
	case TREE_TAG_LT: // <
		return ast_tree2.AST_TREE_KIND_LESS_THAN
	case TREE_TAG_GT: // >
		return ast_tree2.AST_TREE_KIND_GREATER_THAN
	case TREE_TAG_LE: // <=
		return ast_tree2.AST_TREE_KIND_LESS_THAN_EQUAL
	case TREE_TAG_GE: // >=
		return ast_tree2.AST_TREE_KIND_GREATER_THAN_EQUAL

	// Equality operators
	case TREE_TAG_EQ: // ==
		return ast_tree2.AST_TREE_KIND_EQUAL_TO
	case TREE_TAG_NE: // !=
		return ast_tree2.AST_TREE_KIND_NOT_EQUAL_TO

	// Bitwise and logical operators
	case TREE_TAG_BITAND: // &
		return ast_tree2.AST_TREE_KIND_AND
	case TREE_TAG_BITXOR: // ^
		return ast_tree2.AST_TREE_KIND_XOR
	case TREE_TAG_BITOR: // |
		return ast_tree2.AST_TREE_KIND_OR

	// Conditional operators
	case TREE_TAG_AND: // &&
		return ast_tree2.AST_TREE_KIND_CONDITIONAL_AND
	case TREE_TAG_OR: // ||
		return ast_tree2.AST_TREE_KIND_CONDITIONAL_OR

	// Assignment operators
	case TREE_TAG_MUL_ASG: // *=
		return ast_tree2.AST_TREE_KIND_MULTIPLY_ASSIGNMENT
	case TREE_TAG_DIV_ASG: // /=
		return ast_tree2.AST_TREE_KIND_DIVIDE_ASSIGNMENT
	case TREE_TAG_MOD_ASG: // %=
		return ast_tree2.AST_TREE_KIND_REMAINDER_ASSIGNMENT
	case TREE_TAG_PLUS_ASG: // +=
		return ast_tree2.AST_TREE_KIND_PLUS_ASSIGNMENT
	case TREE_TAG_MINUS_ASG: // -=
		return ast_tree2.AST_TREE_KIND_MINUS_ASSIGNMENT
	case TREE_TAG_SL_ASG: // <<=
		return ast_tree2.AST_TREE_KIND_LEFT_SHIFT_ASSIGNMENT
	case TREE_TAG_SR_ASG: // >>=
		return ast_tree2.AST_TREE_KIND_RIGHT_SHIFT_ASSIGNMENT
	case TREE_TAG_USR_ASG: // >>>=
		return ast_tree2.AST_TREE_KIND_UNSIGNED_RIGHT_SHIFT_ASSIGNMENT
	case TREE_TAG_BITAND_ASG: // &=
		return ast_tree2.AST_TREE_KIND_AND_ASSIGNMENT
	case TREE_TAG_BITXOR_ASG: // ^=
		return ast_tree2.AST_TREE_KIND_XOR_ASSIGNMENT
	case TREE_TAG_BITOR_ASG: // |=
		return ast_tree2.AST_TREE_KIND_OR_ASSIGNMENT

	// Null check (implementation detail), for example, __.getClass()
	case TREE_TAG_NULLCHK:
		return ast_tree2.AST_TREE_KIND_OTHER

	case TREE_TAG_ANNOTATION:
		return ast_tree2.AST_TREE_KIND_ANNOTATION
	case TREE_TAG_TYPE_ANNOTATION:
		return ast_tree2.AST_TREE_KIND_TYPE_ANNOTATION

	default:
		return ast_tree2.AST_TREE_KIND_NIL
	}
}

/** Operator precedences values.
 */
const notExpression int = -1 // not an expression
const noPrec = 0             // no enclosing expression
const assignPrec = 1
const assignopPrec = 2
const condPrec = 3
const orPrec = 4
const andPrec = 5
const bitorPrec = 6
const bitxorPrec = 7
const bitandPrec = 8
const eqPrec = 9
const ordPrec = 10
const shiftPrec = 11
const addPrec = 12
const mulPrec = 13
const prefixPrec = 14
const postfixPrec = 15
const precCount = 16

/** Map operators to their precedence levels.
 */
func OpPrec(tt JCTreeTag) int {

	switch tt {
	case TREE_TAG_POS, TREE_TAG_NEG, TREE_TAG_NOT, TREE_TAG_COMPL,
		TREE_TAG_PREINC,
		TREE_TAG_PREDEC:
		return prefixPrec
	case TREE_TAG_POSTINC,
		TREE_TAG_POSTDEC,
		TREE_TAG_NULLCHK:
		return postfixPrec
	case TREE_TAG_ASSIGN:
		return assignPrec
	case TREE_TAG_BITOR_ASG,
		TREE_TAG_BITXOR_ASG,
		TREE_TAG_BITAND_ASG,
		TREE_TAG_SL_ASG,
		TREE_TAG_SR_ASG,
		TREE_TAG_USR_ASG,
		TREE_TAG_PLUS_ASG,
		TREE_TAG_MINUS_ASG,
		TREE_TAG_MUL_ASG,
		TREE_TAG_DIV_ASG,
		TREE_TAG_MOD_ASG:
		return assignopPrec
	case TREE_TAG_OR:
		return orPrec
	case TREE_TAG_AND:
		return andPrec
	case TREE_TAG_EQ,
		TREE_TAG_NE:
		return eqPrec
	case TREE_TAG_LT,
		TREE_TAG_GT,
		TREE_TAG_LE,
		TREE_TAG_GE:
		return ordPrec
	case TREE_TAG_BITOR:
		return bitorPrec
	case TREE_TAG_BITXOR:
		return bitxorPrec
	case TREE_TAG_BITAND:
		return bitandPrec
	case TREE_TAG_SL,
		TREE_TAG_SR,
		TREE_TAG_USR:
		return shiftPrec
	case TREE_TAG_PLUS,
		TREE_TAG_MINUS:
		return addPrec
	case TREE_TAG_MUL,
		TREE_TAG_DIV,
		TREE_TAG_MOD:
		return mulPrec
	case TREE_TAG_TYPETEST:
		return ordPrec
	default:
		panic("invalid tree_tag")
	}
}
