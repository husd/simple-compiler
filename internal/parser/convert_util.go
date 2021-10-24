package parser

/**
 * token kind 转换为 tree_tag
 * @auth or hushengdong
 */

func unOpTag(tk TokenKind) TreeNodeTag {

	switch tk {
	case PLUS:
		return pos
	case SUB:
		return neg
	case BANG:
		return not
	case TILDE:
		return compl
	case PLUSPLUS:
		return preinc
	case SUBSUB:
		return predec
	default:
		return no_tag
	}
}

// tk 转换为 TreeNodeTag 运算符号相关的 tk
func toOpTag(tk TokenKind) TreeNodeTag {

	switch tk {
	case BARBAR:
		return or
	case AMPAMP:
		return and
	case BAR:
		return bitor
	case BAREQ:
		return bitor_asg
	case CARET:
		return bitxor
	case CARETEQ:
		return bitxor_asg
	case AMP:
		return bitand
	case AMPEQ:
		return bitand_asg
	case EQEQ:
		return eq
	case BANGEQ:
		return ne
	case LT:
		return lt
	case GT:
		return gt
	case LTEQ:
		return le
	case GTEQ:
		return ge
	case LTLT:
		return sl
	case LTLTEQ:
		return sl_asg
	case GTGT:
		return sr
	case GTGTEQ:
		return sr_asg
	case GTGTGT:
		return usr
	case GTGTGTEQ:
		return usr_asg
	case PLUS:
		return plus
	case PLUSEQ:
		return plus_asg
	case SUB:
		return minus
	case SUBEQ:
		return minus_asg
	case STAR:
		return mul
	case STAREQ:
		return mul_asg
	case SLASH:
		return div
	case SLASHEQ:
		return div_asg
	case PERCENT:
		return mod
	case PERCENTEQ:
		return mod_asg
	case INSTANCEOF:
		return typetest
	default:
		return no_tag
	}
}

func getTreeTypeByTreeNodeTag(tag TreeNodeTag) TreeType {

	switch tag {
	// postfix expressions
	case postinc: // _ ++
		return tt_postfix_increment
	case postdec: // _ --
		return tt_postfix_decrement

	// unary operators
	case preinc: // ++ _
		return tt_prefix_increment
	case predec: // -- _
		return tt_prefix_decrement
	case pos: // +
		return tt_unary_plus
	case neg: // -
		return tt_unary_minus
	case compl: // ~
		return tt_bitwise_complement
	case not: // !
		return tt_logical_complement

	// binary operators

	// multiplicative operators
	case mul: // *
		return tt_multiply
	case div: // /
		return tt_divide
	case mod: // %
		return tt_remainder

	// additive operators
	case plus: // +
		return tt_plus
	case minus: // -
		return tt_minus

	// shift operators
	case sl: // <<
		return tt_left_shift
	case sr: // >>
		return tt_right_shift
	case usr: // >>>
		return tt_unsigned_right_shift

	// relational operators
	case lt: // <
		return tt_less_than
	case gt: // >
		return tt_greater_than
	case le: // <=
		return tt_less_than_equal
	case ge: // >=
		return tt_greater_than_equal

	// equality operators
	case eq: // ==
		return tt_equal_to
	case ne: // !=
		return tt_not_equal_to

	// bitwise and logical operators
	case bitand: // &
		return tt_and
	case bitxor: // ^
		return tt_xor
	case bitor: // |
		return tt_or

	// conditional operators
	case and: // &&
		return tt_conditional_and
	case or: // ||
		return tt_conditional_or

	// assignment operators
	case mul_asg: // *=
		return tt_multiply_assignment
	case div_asg: // /=
		return tt_divide_assignment
	case mod_asg: // %=
		return tt_remainder_assignment
	case plus_asg: // +=
		return tt_plus_assignment
	case minus_asg: // -=
		return tt_minus_assignment
	case sl_asg: // <<=
		return tt_left_shift_assignment
	case sr_asg: // >>=
		return tt_right_shift_assignment
	case usr_asg: // >>>=
		return tt_unsigned_right_shift_assignment
	case bitand_asg: // &=
		return tt_and_assignment
	case bitxor_asg: // ^=
		return tt_xor_assignment
	case bitor_asg: // |=
		return tt_or_assignment

	// null check (implementation detail), for example, __.getclass()
	case nullchk:
		return tt_other

	case annotation:
		return tt_annotation
	case type_annotation:
		return tt_type_annotation

	default:
		return tt_nil
	}
}
