package parser

/**
 * 树的类型
 * @author hushengdong
 */
var array [92]string = [92]string{}

func init() {

	array[assign_] = "="
}

type TreeNodeTag int

const no_tag TreeNodeTag = 0            // for methods that return an invalid tag if a given condition is not met
const toplevel TreeNodeTag = 1          // toplevel nodes of type toplevel representing entire source files.
const import_ TreeNodeTag = 2           //
const classdef TreeNodeTag = 3          //
const methoddef TreeNodeTag = 4         //
const vardef TreeNodeTag = 5            // variable definitions of type vardef.
const skip TreeNodeTag = 6              // the no-op statement ";" of type skip
const block TreeNodeTag = 7             // blocks of type block.
const doloop TreeNodeTag = 8            //
const whileloop TreeNodeTag = 9         //
const forloop TreeNodeTag = 10          //
const foreachloop TreeNodeTag = 11      //
const labelled TreeNodeTag = 12         //
const switch_ TreeNodeTag = 13          //
const case_ TreeNodeTag = 14            //
const synchronized TreeNodeTag = 15     //
const try TreeNodeTag = 16              //
const catch TreeNodeTag = 17            //
const condexpr TreeNodeTag = 18         //
const if_ TreeNodeTag = 19              //
const exec TreeNodeTag = 20             //
const break_ TreeNodeTag = 21           //
const continue_ TreeNodeTag = 22        //
const return_ TreeNodeTag = 23          //
const throw TreeNodeTag = 24            //
const assert TreeNodeTag = 25           //
const apply TreeNodeTag = 26            //
const newclass TreeNodeTag = 27         //
const newarray TreeNodeTag = 28         //
const lambda TreeNodeTag = 29           //
const parens TreeNodeTag = 30           //
const assign_ TreeNodeTag = 31          // assignment expressions of type assign.*
const typecast TreeNodeTag = 32         // type cast expressions of type typecast.*
const typetest TreeNodeTag = 33         // type test expressions of type typetest.*
const indexed TreeNodeTag = 34          // indexed array expressions of type indexed.*
const select_ TreeNodeTag = 35          // selections of type select.*
const reference TreeNodeTag = 36        // member references of type reference.*
const ident TreeNodeTag = 37            // 标识符号 *
const literal TreeNodeTag = 38          // 字面量 *
const typeident TreeNodeTag = 39        // basic type identifiers of type typeident.*
const typearray TreeNodeTag = 40        // array types of type typearray.*
const typeapply TreeNodeTag = 41        // parameterized types of type typeapply.*
const typeunion TreeNodeTag = 42        // union types of type typeunion.*
const typeintersection TreeNodeTag = 43 // intersection types of type typeintersection.
const typeparameter TreeNodeTag = 44    // formal type parameters of type typeparameter.
const wildcard TreeNodeTag = 45         // type argument.
const typeboundkind TreeNodeTag = 46    //
const annotation TreeNodeTag = 47       //
const type_annotation TreeNodeTag = 48  //
const modifiers TreeNodeTag = 49        //
const annotated_type TreeNodeTag = 50   //
const erroneous TreeNodeTag = 51        // 错误
const pos TreeNodeTag = 52              // +
const neg TreeNodeTag = 53              // -
const not TreeNodeTag = 54              // !
const compl TreeNodeTag = 55            // ~
const preinc TreeNodeTag = 56           // ++ _
const predec TreeNodeTag = 57           // -- _
const postinc TreeNodeTag = 58          // _ ++
const postdec TreeNodeTag = 59          // _ --
const nullchk TreeNodeTag = 60          //
const or TreeNodeTag = 61               // ||
const and TreeNodeTag = 62              // &&
const bitor TreeNodeTag = 63            // |
const bitxor TreeNodeTag = 64           // ^
const bitand TreeNodeTag = 65           // &

// 66 - 71 是比较运算符号
const eq TreeNodeTag = 66 // ==
const ne TreeNodeTag = 67 // !=
const lt TreeNodeTag = 68 // <
const gt TreeNodeTag = 69 // >
const le TreeNodeTag = 70 // <=
const ge TreeNodeTag = 71 // >=

const sl TreeNodeTag = 72         // <<
const sr TreeNodeTag = 73         // >>
const usr TreeNodeTag = 74        // >>>
const plus TreeNodeTag = 75       // +
const minus TreeNodeTag = 76      // -
const mul TreeNodeTag = 77        // *
const div TreeNodeTag = 78        //
const mod TreeNodeTag = 79        // %
const bitor_asg TreeNodeTag = 80  // |=
const bitxor_asg TreeNodeTag = 81 // ^=
const bitand_asg TreeNodeTag = 82 // &=
const sl_asg TreeNodeTag = 83     // <<=
const sr_asg TreeNodeTag = 84     // >>=
const usr_asg TreeNodeTag = 85    // >>>=
const plus_asg TreeNodeTag = 86   // +=
const minus_asg TreeNodeTag = 87  // -=
const mul_asg TreeNodeTag = 88    // *=
const div_asg TreeNodeTag = 89    //
const mod_asg TreeNodeTag = 90    // %=
const letexpr TreeNodeTag = 91    // ala scheme a synthetic let expression of type letexpr.
