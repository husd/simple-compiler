package ast_tree

/**
 *
 * @author hushengdong
 */

func treeTag2TreeKind(tag AstTreeTag) *AstTreeNodeKind {

	switch tag {
	// Postfix expressions
	case POSTINC: // _ ++
		return TREE_POSTFIX_INCREMENT
	case POSTDEC: // _ --
		return TREE_POSTFIX_DECREMENT

	// Unary operators
	case PREINC: // ++ _
		return TREE_PREFIX_INCREMENT
	case PREDEC: // -- _
		return TREE_PREFIX_DECREMENT
	case POS: // +
		return TREE_UNARY_PLUS
	case NEG: // -
		return TREE_UNARY_MINUS
	case COMPL: // ~
		return TREE_BITWISE_COMPLEMENT
	case NOT: // !
		return TREE_LOGICAL_COMPLEMENT

	// Binary operators

	// Multiplicative operators
	case MUL: // *
		return TREE_MULTIPLY
	case DIV: // /
		return TREE_DIVIDE
	case MOD: // %
		return TREE_REMAINDER

	// Additive operators
	case PLUS: // +
		return TREE_PLUS
	case MINUS: // -
		return TREE_MINUS

	// Shift operators
	case SL: // <<
		return TREE_LEFT_SHIFT
	case SR: // >>
		return TREE_RIGHT_SHIFT
	case USR: // >>>
		return TREE_UNSIGNED_RIGHT_SHIFT

	// Relational operators
	case LT: // <
		return TREE_LESS_THAN
	case GT: // >
		return TREE_GREATER_THAN
	case LE: // <=
		return TREE_LESS_THAN_EQUAL
	case GE: // >=
		return TREE_GREATER_THAN_EQUAL

	// Equality operators
	case EQ: // ==
		return TREE_EQUAL_TO
	case NE: // !=
		return TREE_NOT_EQUAL_TO

	// Bitwise and logical operators
	case BITAND: // &
		return TREE_AND
	case BITXOR: // ^
		return TREE_XOR
	case BITOR: // |
		return TREE_OR

	// Conditional operators
	case AND: // &&
		return TREE_CONDITIONAL_AND
	case OR: // ||
		return TREE_CONDITIONAL_OR

	// Assignment operators
	case MUL_ASG: // *=
		return TREE_MULTIPLY_ASSIGNMENT
	case DIV_ASG: // /=
		return TREE_DIVIDE_ASSIGNMENT
	case MOD_ASG: // %=
		return TREE_REMAINDER_ASSIGNMENT
	case PLUS_ASG: // +=
		return TREE_PLUS_ASSIGNMENT
	case MINUS_ASG: // -=
		return TREE_MINUS_ASSIGNMENT
	case SL_ASG: // <<=
		return TREE_LEFT_SHIFT_ASSIGNMENT
	case SR_ASG: // >>=
		return TREE_RIGHT_SHIFT_ASSIGNMENT
	case USR_ASG: // >>>=
		return TREE_UNSIGNED_RIGHT_SHIFT_ASSIGNMENT
	case BITAND_ASG: // &=
		return TREE_AND_ASSIGNMENT
	case BITXOR_ASG: // ^=
		return TREE_XOR_ASSIGNMENT
	case BITOR_ASG: // |=
		return TREE_OR_ASSIGNMENT

	// Null check (implementation detail), for example, __.getClass()
	case NULLCHK:
		return TREE_OTHER

	case ANNOTATION:
		return TREE_ANNOTATION
	case TYPE_ANNOTATION:
		return TREE_TYPE_ANNOTATION

	default:
		return TREE_NIL
	}
}
