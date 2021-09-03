package parser

/**
 * 定义了一些 特殊字符的unicode码
 * @author hushengdong
 */

const Layout_char_tabInc = rune(8)
const Layout_char_diag = rune(4)
const Layout_char_details = rune(4)
const Layout_char_tab = rune(0x9)  //tab
const Layout_char_lf = rune(0xA)   //换行
const Layout_char_ff = rune(0xC)   //换页
const Layout_char_cr = rune(0xD)   //回车
const Layout_char_eoi = rune(0x1A) //双引号 用来识别字符串 "
