package jc

/**
 *
 * @author hushengdong
 */
type JCTreeTag int

const (
	TREE_TAG_NO_TAG   JCTreeTag = iota //For methods that return an invalid tag if a given condition is not met
	TREE_TAG_TOPLEVEL                  // Toplevel nodes of type TopLevel representing entire source files.
	TREE_TAG_IMPORT
	TREE_TAG_CLASSDEF
	TREE_TAG_METHODDEF
	TREE_TAG_VARDEF // Variable definitions of type VarDef.
	TREE_TAG_SKIP   //The no-op statement ";" of type Skip
	TREE_TAG_BLOCK  //Blocks of type Block.
	TREE_TAG_DOLOOP
	TREE_TAG_WHILELOOP
	TREE_TAG_FORLOOP
	TREE_TAG_FOREACHLOOP
	TREE_TAG_LABELLED
	TREE_TAG_SWITCH
	TREE_TAG_CASE
	TREE_TAG_SYNCHRONIZED
	TREE_TAG_TRY
	TREE_TAG_CATCH
	TREE_TAG_CONDEXPR
	TREE_TAG_IF
	TREE_TAG_EXEC
	TREE_TAG_BREAK
	TREE_TAG_CONTINUE
	TREE_TAG_RETURN
	TREE_TAG_THROW
	TREE_TAG_ASSERT
	TREE_TAG_APPLY
	TREE_TAG_NEWCLASS
	TREE_TAG_NEWARRAY
	TREE_TAG_LAMBDA
	TREE_TAG_PARENS
	TREE_TAG_ASSIGN           /** Assignment expressions of type Assign.*/
	TREE_TAG_TYPECAST         /** Type cast expressions of type TypeCast.*/
	TREE_TAG_TYPETEST         /** Type test expressions of type TypeTest.*/
	TREE_TAG_INDEXED          /** Indexed array expressions of type Indexed.*/
	TREE_TAG_SELECT           /** Selections of type Select.*/
	TREE_TAG_REFERENCE        /** Member references of type Reference.*/
	TREE_TAG_IDENT            /** 标识符号 */
	TREE_TAG_LITERAL          /** 字面量 */
	TREE_TAG_TYPEIDENT        /** Basic type identifiers of type TypeIdent.*/
	TREE_TAG_TYPEARRAY        /** Array types of type TypeArray.*/
	TREE_TAG_TYPEAPPLY        /** Parameterized types of type TypeApply.*/
	TREE_TAG_TYPEUNION        /** Union types of type TypeUnion.*/
	TREE_TAG_TYPEINTERSECTION //Intersection types of type TypeIntersection.
	TREE_TAG_TYPEPARAMETER    //Formal type parameters of type TypeParameter.
	TREE_TAG_WILDCARD         //Type argument.
	TREE_TAG_TYPEBOUNDKIND
	TREE_TAG_ANNOTATION
	TREE_TAG_TYPE_ANNOTATION
	TREE_TAG_MODIFIERS
	TREE_TAG_ANNOTATED_TYPE
	TREE_TAG_ERRONEOUS //错误
	TREE_TAG_POS       // +
	TREE_TAG_NEG       // -
	TREE_TAG_NOT       // !
	TREE_TAG_COMPL     // ~
	TREE_TAG_PREINC    // ++ _
	TREE_TAG_PREDEC    // -- _
	TREE_TAG_POSTINC   // _ ++
	TREE_TAG_POSTDEC   // _ --
	TREE_TAG_NULLCHK
	TREE_TAG_OR         // ||
	TREE_TAG_AND        // &&
	TREE_TAG_BITOR      // |
	TREE_TAG_BITXOR     // ^
	TREE_TAG_BITAND     // &
	TREE_TAG_EQ         // ==
	TREE_TAG_NE         // !=
	TREE_TAG_LT         // <
	TREE_TAG_GT         // >
	TREE_TAG_LE         // <=
	TREE_TAG_GE         // >=
	TREE_TAG_SL         // <<
	TREE_TAG_SR         // >>
	TREE_TAG_USR        // >>>
	TREE_TAG_PLUS       // +
	TREE_TAG_MINUS      // -
	TREE_TAG_MUL        // *
	TREE_TAG_DIV        // /
	TREE_TAG_MOD        // %
	TREE_TAG_BITOR_ASG  // |=
	TREE_TAG_BITXOR_ASG // ^=
	TREE_TAG_BITAND_ASG // &=
	TREE_TAG_SL_ASG     // <<=
	TREE_TAG_SR_ASG     // >>=
	TREE_TAG_USR_ASG    // >>>=
	TREE_TAG_PLUS_ASG   // +=
	TREE_TAG_MINUS_ASG  // -=
	TREE_TAG_MUL_ASG    // *=
	TREE_TAG_DIV_ASG    // /=
	TREE_TAG_MOD_ASG    // %=
	TREE_TAG_LETEXPR    // ala scheme A synthetic let expression of type LetExpr.
)

func GetNumberOfOperator() int {

	return int(TREE_TAG_MOD-TREE_TAG_POS) + 1
}
