package jc

/**
 *
 * @author hushengdong
 */
type JCTreeTag int

const (
	NO_TAG   JCTreeTag = iota //For methods that return an invalid tag if a given condition is not met
	TOPLEVEL                  // Toplevel nodes of type TopLevel representing entire source files.
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
	ASSIGN           /** Assignment expressions of type Assign.*/
	TYPECAST         /** Type cast expressions of type TypeCast.*/
	TYPETEST         /** Type test expressions of type TypeTest.*/
	INDEXED          /** Indexed array expressions of type Indexed.*/
	SELECT           /** Selections of type Select.*/
	REFERENCE        /** Member references of type Reference.*/
	IDENT            /** Simple identifiers of type Ident.*/
	LITERAL          /** Literals of type Literal.*/
	TYPEIDENT        /** Basic type identifiers of type TypeIdent.*/
	TYPEARRAY        /** Array types of type TypeArray.*/
	TYPEAPPLY        /** Parameterized types of type TypeApply.*/
	TYPEUNION        /** Union types of type TypeUnion.*/
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
