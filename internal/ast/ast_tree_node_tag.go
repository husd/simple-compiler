package ast

/**
 * 树的类型
 * @author hushengdong
 */
var tree_node_tag_array [92]string = [92]string{}

func init() {

	tree_node_tag_array[Tree_node_tag_assign] = "="
}

type TreeNodeTag int

const Tree_node_tag_no_tag TreeNodeTag = 0            // for methods that return an invalid tag if a given condition is not met
const Tree_node_tag_toplevel TreeNodeTag = 1          // toplevel nodes of type toplevel representing entire source files.
const Tree_node_tag_import TreeNodeTag = 2            //
const Tree_node_tag_classdef TreeNodeTag = 3          //
const Tree_node_tag_methoddef TreeNodeTag = 4         //
const Tree_node_tag_vardef TreeNodeTag = 5            // variable definitions of type vardef.
const Tree_node_tag_skip TreeNodeTag = 6              // the no-op statement ";" of type skip
const Tree_node_tag_block TreeNodeTag = 7             // blocks of type block.
const Tree_node_tag_doloop TreeNodeTag = 8            //
const Tree_node_tag_whileloop TreeNodeTag = 9         //
const Tree_node_tag_forloop TreeNodeTag = 10          //
const Tree_node_tag_foreachloop TreeNodeTag = 11      //
const Tree_node_tag_labelled TreeNodeTag = 12         //
const Tree_node_tag_switch TreeNodeTag = 13           //
const Tree_node_tag_case TreeNodeTag = 14             //
const Tree_node_tag_synchronized TreeNodeTag = 15     //
const Tree_node_tag_try TreeNodeTag = 16              //
const Tree_node_tag_catch TreeNodeTag = 17            //
const Tree_node_tag_condexpr TreeNodeTag = 18         //
const Tree_node_tag_if TreeNodeTag = 19               //
const Tree_node_tag_exec TreeNodeTag = 20             //
const Tree_node_tag_break TreeNodeTag = 21            //
const Tree_node_tag_continue TreeNodeTag = 22         //
const Tree_node_tag_return TreeNodeTag = 23           //
const Tree_node_tag_throw TreeNodeTag = 24            //
const Tree_node_tag_assert TreeNodeTag = 25           //
const Tree_node_tag_apply TreeNodeTag = 26            //
const Tree_node_tag_newclass TreeNodeTag = 27         //
const Tree_node_tag_newarray TreeNodeTag = 28         //
const Tree_node_tag_lambda TreeNodeTag = 29           //
const Tree_node_tag_parens TreeNodeTag = 30           //
const Tree_node_tag_assign TreeNodeTag = 31           // assignment expressions of type assign.*
const Tree_node_tag_typecast TreeNodeTag = 32         // type cast expressions of type typecast.*
const Tree_node_tag_typetest TreeNodeTag = 33         // type test expressions of type typetest.*
const Tree_node_tag_indexed TreeNodeTag = 34          // indexed array expressions of type indexed.*
const Tree_node_tag_select TreeNodeTag = 35           // selections of type select.*
const Tree_node_tag_reference TreeNodeTag = 36        // member references of type reference.*
const Tree_node_tag_ident TreeNodeTag = 37            // 标识符号 *
const Tree_node_tag_literal TreeNodeTag = 38          // 字面量 *
const Tree_node_tag_typeident TreeNodeTag = 39        // basic type identifiers of type typeident.*
const Tree_node_tag_typearray TreeNodeTag = 40        // array types of type typearray.*
const Tree_node_tag_typeapply TreeNodeTag = 41        // parameterized types of type typeapply.*
const Tree_node_tag_typeunion TreeNodeTag = 42        // union types of type typeunion.*
const Tree_node_tag_typeintersection TreeNodeTag = 43 // intersection types of type typeintersection.
const Tree_node_tag_typeparameter TreeNodeTag = 44    // formal type parameters of type typeparameter.
const Tree_node_tag_wildcard TreeNodeTag = 45         // type argument.
const Tree_node_tag_typeboundkind TreeNodeTag = 46    //
const Tree_node_tag_annotation TreeNodeTag = 47       //
const Tree_node_tag_type_annotation TreeNodeTag = 48  //
const Tree_node_tag_modifiers TreeNodeTag = 49        //
const Tree_node_tag_annotated_type TreeNodeTag = 50   //
const Tree_node_tag_erroneous TreeNodeTag = 51        // 错误
const Tree_node_tag_pos TreeNodeTag = 52              // +
const Tree_node_tag_neg TreeNodeTag = 53              // -
const Tree_node_tag_not TreeNodeTag = 54              // !
const Tree_node_tag_compl TreeNodeTag = 55            // ~
const Tree_node_tag_preinc TreeNodeTag = 56           // ++ _
const Tree_node_tag_predec TreeNodeTag = 57           // -- _
const Tree_node_tag_postinc TreeNodeTag = 58          // _ ++
const Tree_node_tag_postdec TreeNodeTag = 59          // _ --
const Tree_node_tag_nullchk TreeNodeTag = 60          //
const Tree_node_tag_or TreeNodeTag = 61               // ||
const Tree_node_tag_and TreeNodeTag = 62              // &&
const Tree_node_tag_bitor TreeNodeTag = 63            // |
const Tree_node_tag_bitxor TreeNodeTag = 64           // ^
const Tree_node_tag_bitand TreeNodeTag = 65           // &
const Tree_node_tag_eq TreeNodeTag = 66               // ==
const Tree_node_tag_ne TreeNodeTag = 67               // !=
const Tree_node_tag_lt TreeNodeTag = 68               // <
const Tree_node_tag_gt TreeNodeTag = 69               // >
const Tree_node_tag_le TreeNodeTag = 70               // <=
const Tree_node_tag_ge TreeNodeTag = 71               // >=
const Tree_node_tag_sl TreeNodeTag = 72               // <<
const Tree_node_tag_sr TreeNodeTag = 73               // >>
const Tree_node_tag_usr TreeNodeTag = 74              // >>>
const Tree_node_tag_plus TreeNodeTag = 75             // +
const Tree_node_tag_minus TreeNodeTag = 76            // -
const Tree_node_tag_mul TreeNodeTag = 77              // *
const Tree_node_tag_div TreeNodeTag = 78              //
const Tree_node_tag_mod TreeNodeTag = 79              // %
const Tree_node_tag_bitor_asg TreeNodeTag = 80        // |=
const Tree_node_tag_bitxor_asg TreeNodeTag = 81       // ^=
const Tree_node_tag_bitand_asg TreeNodeTag = 82       // &=
const Tree_node_tag_sl_asg TreeNodeTag = 83           // <<=
const Tree_node_tag_sr_asg TreeNodeTag = 84           // >>=
const Tree_node_tag_usr_asg TreeNodeTag = 85          // >>>=
const Tree_node_tag_plus_asg TreeNodeTag = 86         // +=
const Tree_node_tag_minus_asg TreeNodeTag = 87        // -=
const Tree_node_tag_mul_asg TreeNodeTag = 88          // *=
const Tree_node_tag_div_asg TreeNodeTag = 89          //
const Tree_node_tag_mod_asg TreeNodeTag = 90          // %=
const Tree_node_tag_letexpr TreeNodeTag = 91          // ala scheme a synthetic let expression of type letexpr.
