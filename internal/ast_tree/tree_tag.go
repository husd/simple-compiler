package ast_tree

/**
 *
 * @author hushengdong
 */
type AstTreeTag int

const (
	NO_TAG   AstTreeTag = iota //For methods that return an invalid tag if a given condition is not met
	TOPLEVEL                   // Toplevel nodes of type TopLevel representing entire source files.
	IMPORT
	CLASSDEF
	METHODDEF
	VARDEF // Variable definitions of type VarDef.
	SKIP   //The no-op statement ";" of type Skip
	BLOCK  //Blocks of type Block.
	DOLOOP
	WHILELOOP
	FORLOOP
	FOREACHLOOP
	LABELLED
	SWITCH
	CASE
	SYNCHRONIZED
	TRY
	CATCH
	CONDEXPR
	IF
	EXEC
	BREAK
	CONTINUE
	RETURN
	THROW
	ASSERT
	APPLY
	NEWCLASS
	NEWARRAY
	LAMBDA
	PARENS

	/** Assignment expressions of type Assign.
	 */
	ASSIGN

	/** Type cast expressions of type TypeCast.
	 */
	TYPECAST

	/** Type test expressions of type TypeTest.
	 */
	TYPETEST

	/** Indexed array expressions of type Indexed.
	 */
	INDEXED

	/** Selections of type Select.
	 */
	SELECT

	/** Member references of type Reference.
	 */
	REFERENCE

	/** Simple identifiers of type Ident.
	 */
	IDENT

	/** Literals of type Literal.
	 */
	LITERAL

	/** Basic type identifiers of type TypeIdent.
	 */
	TYPEIDENT

	/** Array types of type TypeArray.
	 */
	TYPEARRAY

	/** Parameterized types of type TypeApply.
	 */
	TYPEAPPLY

	/** Union types of type TypeUnion.
	 */
	TYPEUNION
	TYPEINTERSECTION //Intersection types of type TypeIntersection.
	TYPEPARAMETER    //Formal type parameters of type TypeParameter.
	WILDCARD         //Type argument.
	TYPEBOUNDKIND
	ANNOTATION
	TYPE_ANNOTATION
	MODIFIERS
	ANNOTATED_TYPE
	ERRONEOUS
	POS     // +
	NEG     // -
	NOT     // !
	COMPL   // ~
	PREINC  // ++ _
	PREDEC  // -- _
	POSTINC // _ ++
	POSTDEC // _ --
	NULLCHK
	OR         // ||
	AND        // &&
	BITOR      // |
	BITXOR     // ^
	BITAND     // &
	EQ         // ==
	NE         // !=
	LT         // <
	GT         // >
	LE         // <=
	GE         // >=
	SL         // <<
	SR         // >>
	USR        // >>>
	PLUS       // +
	MINUS      // -
	MUL        // *
	DIV        // /
	MOD        // %
	BITOR_ASG  // |=
	BITXOR_ASG // ^=
	BITAND_ASG // &=
	SL_ASG     // <<=
	SR_ASG     // >>=
	USR_ASG    // >>>=
	PLUS_ASG   // +=
	MINUS_ASG  // -=
	MUL_ASG    // *=
	DIV_ASG    // /=
	MOD_ASG    // %=
	LETEXPR    // ala scheme A synthetic let expression of type LetExpr.
)

func GetNumberOfOperator() int {

	return int(MOD-POS) + 1
}
