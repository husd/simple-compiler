package util

/**
 * 定义了编译器的标准的name table，如果你要加新的name，也要在这里定义。
 * @author hushengdong
 */

type Names struct {

	// operators and punctuation
	asterisk    *Name // 星号
	comma       *Name // 逗号
	Empty       *Name // 空
	Hyphen      *Name // 连字符;连字号
	one         *Name // 1
	period      *Name // .
	semicolon   *Name // 分号
	slash       *Name // 斜线;斜杠;
	slashequals *Name // /=

	// keywords
	_class   *Name
	_default *Name
	_super   *Name
	_this    *Name

	// field and method *Names
	_name                  *Name
	addSuppressed          *Name
	any                    *Name
	append                 *Name
	clinit                 *Name
	clone                  *Name
	close                  *Name
	compareTo              *Name
	deserializeLambda      *Name
	desiredAssertionStatus *Name
	equals                 *Name
	Error                  *Name
	family                 *Name
	finalize               *Name
	forName                *Name
	getClass               *Name
	getClassLoader         *Name
	getComponentType       *Name
	getDeclaringClass      *Name
	getMessage             *Name
	hasNext                *Name
	hashCode               *Name
	init                   *Name
	initCause              *Name
	iterator               *Name
	length                 *Name
	next                   *Name
	ordinal                *Name
	serialVersionUID       *Name
	toString               *Name
	value                  *Name
	valueOf                *Name
	values                 *Name

	// class *Names
	java_io_Serializable          *Name
	java_lang_AutoCloseable       *Name
	java_lang_Class               *Name
	java_lang_Cloneable           *Name
	java_lang_Enum                *Name
	java_lang_Object              *Name
	java_lang_invoke_MethodHandle *Name

	// *Names of builtin classes
	Array  *Name
	Bound  *Name
	Method *Name

	// package *Names
	java_lang *Name

	// attribute *Names
	Annotation                           *Name
	AnnotationDefault                    *Name
	BootstrapMethods                     *Name
	Bridge                               *Name
	CharacterRangeTable                  *Name
	Code                                 *Name
	CompilationID                        *Name
	ConstantValue                        *Name
	Deprecated                           *Name
	EnclosingMethod                      *Name
	Enum                                 *Name
	Exceptions                           *Name
	InnerClasses                         *Name
	LineNumberTable                      *Name
	LocalVariableTable                   *Name
	LocalVariableTypeTable               *Name
	MethodParameters                     *Name
	RuntimeInvisibleAnnotations          *Name
	RuntimeInvisibleParameterAnnotations *Name
	RuntimeInvisibleTypeAnnotations      *Name
	RuntimeVisibleAnnotations            *Name
	RuntimeVisibleParameterAnnotations   *Name
	RuntimeVisibleTypeAnnotations        *Name
	Signature                            *Name
	SourceFile                           *Name
	SourceID                             *Name
	StackMap                             *Name
	StackMapTable                        *Name
	Synthetic                            *Name
	Value                                *Name
	Varargs                              *Name

	// members of java.lang.annotation.ElementType
	ANNOTATION_TYPE *Name
	CONSTRUCTOR     *Name
	FIELD           *Name
	LOCAL_VARIABLE  *Name
	METHOD          *Name
	PACKAGE         *Name
	PARAMETER       *Name
	TYPE            *Name
	TYPE_PARAMETER  *Name
	TYPE_USE        *Name

	// members of java.lang.annotation.RetentionPolicy
	CLASS   *Name
	RUNTIME *Name
	SOURCE  *Name

	// other identifiers
	T            *Name
	deprecated   *Name
	ex           *Name
	package_info *Name

	//lambda-related
	lambda         *Name
	metafactory    *Name
	altMetafactory *Name

	table *SharedNameTable
}

func (ns *Names) fromString(s string) *Name {

	return ns.table.fromString(s)
}

func InstanceNames(c *Context) *Names {

	ok, obj := c.Get(C_NAMES)
	if ok {
		return obj.(*Names)
	}
	return NewNames(c)
}

func NewNames(c *Context) *Names {

	ns := &Names{}
	ns.table = NewSharedNameTable(c)

	// operators and punctuation
	ns.asterisk = ns.fromString("*")
	ns.comma = ns.fromString(",")
	ns.Empty = ns.fromString("")
	ns.Hyphen = ns.fromString("-")
	ns.one = ns.fromString("1")
	ns.period = ns.fromString(".")
	ns.semicolon = ns.fromString(";")
	ns.slash = ns.fromString("/")
	ns.slashequals = ns.fromString("/=")

	// keywords
	ns._class = ns.fromString("class")
	ns._default = ns.fromString("default")
	ns._super = ns.fromString("super")
	ns._this = ns.fromString("this")

	// field and method *Names
	ns._name = ns.fromString("name")
	ns.addSuppressed = ns.fromString("addSuppressed")
	ns.any = ns.fromString("<any>")
	ns.append = ns.fromString("append")
	ns.clinit = ns.fromString("<clinit>")
	ns.clone = ns.fromString("clone")
	ns.close = ns.fromString("close")
	ns.compareTo = ns.fromString("compareTo")
	ns.deserializeLambda = ns.fromString("$deserializeLambda$")
	ns.desiredAssertionStatus = ns.fromString("desiredAssertionStatus")
	ns.equals = ns.fromString("equals")
	ns.Error = ns.fromString("<error>")
	ns.family = ns.fromString("family")
	ns.finalize = ns.fromString("finalize")
	ns.forName = ns.fromString("forName")
	ns.getClass = ns.fromString("getClass")
	ns.getClassLoader = ns.fromString("getClassLoader")
	ns.getComponentType = ns.fromString("getComponentType")
	ns.getDeclaringClass = ns.fromString("getDeclaringClass")
	ns.getMessage = ns.fromString("getMessage")
	ns.hasNext = ns.fromString("hasNext")
	ns.hashCode = ns.fromString("hashCode")
	ns.init = ns.fromString("<init>")
	ns.initCause = ns.fromString("initCause")
	ns.iterator = ns.fromString("iterator")
	ns.length = ns.fromString("length")
	ns.next = ns.fromString("next")
	ns.ordinal = ns.fromString("ordinal")
	ns.serialVersionUID = ns.fromString("serialVersionUID")
	ns.toString = ns.fromString("toString")
	ns.value = ns.fromString("value")
	ns.valueOf = ns.fromString("valueOf")
	ns.values = ns.fromString("values")

	// class *Names
	ns.java_io_Serializable = ns.fromString("java.io.Serializable")
	ns.java_lang_AutoCloseable = ns.fromString("java.lang.AutoCloseable")
	ns.java_lang_Class = ns.fromString("java.lang.Class")
	ns.java_lang_Cloneable = ns.fromString("java.lang.Cloneable")
	ns.java_lang_Enum = ns.fromString("java.lang.Enum")
	ns.java_lang_Object = ns.fromString("java.lang.Object")
	ns.java_lang_invoke_MethodHandle = ns.fromString("java.lang.invoke.MethodHandle")

	// *Names of builtin classes
	ns.Array = ns.fromString("Array")
	ns.Bound = ns.fromString("Bound")
	ns.Method = ns.fromString("Method")

	// package *Names
	ns.java_lang = ns.fromString("java.lang")

	// attribute *Names
	ns.Annotation = ns.fromString("Annotation")
	ns.AnnotationDefault = ns.fromString("AnnotationDefault")
	ns.BootstrapMethods = ns.fromString("BootstrapMethods")
	ns.Bridge = ns.fromString("Bridge")
	ns.CharacterRangeTable = ns.fromString("CharacterRangeTable")
	ns.Code = ns.fromString("Code")
	ns.CompilationID = ns.fromString("CompilationID")
	ns.ConstantValue = ns.fromString("ConstantValue")
	ns.Deprecated = ns.fromString("Deprecated")
	ns.EnclosingMethod = ns.fromString("EnclosingMethod")
	ns.Enum = ns.fromString("Enum")
	ns.Exceptions = ns.fromString("Exceptions")
	ns.InnerClasses = ns.fromString("InnerClasses")
	ns.LineNumberTable = ns.fromString("LineNumberTable")
	ns.LocalVariableTable = ns.fromString("LocalVariableTable")
	ns.LocalVariableTypeTable = ns.fromString("LocalVariableTypeTable")
	ns.MethodParameters = ns.fromString("MethodParameters")
	ns.RuntimeInvisibleAnnotations = ns.fromString("RuntimeInvisibleAnnotations")
	ns.RuntimeInvisibleParameterAnnotations = ns.fromString("RuntimeInvisibleParameterAnnotations")
	ns.RuntimeInvisibleTypeAnnotations = ns.fromString("RuntimeInvisibleTypeAnnotations")
	ns.RuntimeVisibleAnnotations = ns.fromString("RuntimeVisibleAnnotations")
	ns.RuntimeVisibleParameterAnnotations = ns.fromString("RuntimeVisibleParameterAnnotations")
	ns.RuntimeVisibleTypeAnnotations = ns.fromString("RuntimeVisibleTypeAnnotations")
	ns.Signature = ns.fromString("Signature")
	ns.SourceFile = ns.fromString("SourceFile")
	ns.SourceID = ns.fromString("SourceID")
	ns.StackMap = ns.fromString("StackMap")
	ns.StackMapTable = ns.fromString("StackMapTable")
	ns.Synthetic = ns.fromString("Synthetic")
	ns.Value = ns.fromString("Value")
	ns.Varargs = ns.fromString("Varargs")

	// members of java.lang.annotation.ElementType
	ns.ANNOTATION_TYPE = ns.fromString("ANNOTATION_TYPE")
	ns.CONSTRUCTOR = ns.fromString("CONSTRUCTOR")
	ns.FIELD = ns.fromString("FIELD")
	ns.LOCAL_VARIABLE = ns.fromString("LOCAL_VARIABLE")
	ns.METHOD = ns.fromString("METHOD")
	ns.PACKAGE = ns.fromString("PACKAGE")
	ns.PARAMETER = ns.fromString("PARAMETER")
	ns.TYPE = ns.fromString("TYPE")
	ns.TYPE_PARAMETER = ns.fromString("TYPE_PARAMETER")
	ns.TYPE_USE = ns.fromString("TYPE_USE")

	// members of java.lang.annotation.RetentionPolicy
	ns.CLASS = ns.fromString("CLASS")
	ns.RUNTIME = ns.fromString("RUNTIME")
	ns.SOURCE = ns.fromString("SOURCE")

	// other identifiers
	ns.T = ns.fromString("T")
	ns.deprecated = ns.fromString("deprecated")
	ns.ex = ns.fromString("ex")
	ns.package_info = ns.fromString("package-info")

	//lambda-related
	ns.lambda = ns.fromString("lambda$")
	ns.metafactory = ns.fromString("metafactory")
	ns.altMetafactory = ns.fromString("altMetafactory")

	c.Put(C_NAMES, ns)
	return ns
}
