package parser

import (
	"husd.com/v0/jc"
	"husd.com/v0/lang"
	"husd.com/v0/util"
)

/**
 * 符号表 存储了一些预先定义好的常量或者操作符号
 * @author hushengdong
 */
type SymbolTable struct {
	byteType    *lang.JCPrimitiveType
	charType    *lang.JCPrimitiveType
	shortType   *lang.JCPrimitiveType
	intType     *lang.JCPrimitiveType
	longType    *lang.JCPrimitiveType
	floatType   *lang.JCPrimitiveType
	doubleType  *lang.JCPrimitiveType
	booleanType *lang.JCPrimitiveType

	symbol jc.Symbol //符号

	tokenTable []Token // 符号表的数据
	tpos       int     // 符号表目前写到哪个位置了
}

func InstanceSymbolTable(c *util.Context) *SymbolTable {

	ok, obj := c.Get(util.C_SYMBOL_TABLE)
	if ok {
		return obj.(*SymbolTable)
	}
	return NewSymbolTable(c)
}

func NewSymbolTable(c *util.Context) *SymbolTable {

	res := &SymbolTable{}

	res.byteType = lang.NewJCPrimitiveType(lang.TYPE_KIND_BYTE)
	res.charType = lang.NewJCPrimitiveType(lang.TYPE_KIND_CHAR)
	res.shortType = lang.NewJCPrimitiveType(lang.TYPE_KIND_SHORT)
	res.intType = lang.NewJCPrimitiveType(lang.TYPE_KIND_INT)
	res.longType = lang.NewJCPrimitiveType(lang.TYPE_KIND_LONG)
	res.floatType = lang.NewJCPrimitiveType(lang.TYPE_KIND_FLOAT)
	res.doubleType = lang.NewJCPrimitiveType(lang.TYPE_KIND_DOUBLE)
	res.booleanType = lang.NewJCPrimitiveType(lang.TYPE_KIND_BOOLEAN)

	res.tokenTable = make([]Token, 512, 512)
	res.tpos = 0

	c.Put(util.C_SYMBOL_TABLE, res)
	return res
}

func (st *SymbolTable) PutToken(t Token) {

	st.ensureCapacity(st.tpos, 1)
	st.tokenTable[st.tpos] = t
	t.SetSymbolTableIndex(st.tpos)
	st.tpos++
}

func (st *SymbolTable) GetTokenByIndex(index int) Token {

	if index >= st.tpos {
		return nil
	}
	return st.tokenTable[index]
}

func (st *SymbolTable) ensureCapacity(spos int, need int) {

	currentCap := cap(st.tokenTable)
	if spos+need > currentCap {
		newCap := calcNewLength(currentCap, spos+need)
		newSbuf := make([]Token, newCap, newCap) //len设置为cap，这样才可以在任意位置写入
		copy(newSbuf, st.tokenTable)
		st.tokenTable = newSbuf
	}
}
