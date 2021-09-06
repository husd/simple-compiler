package lang

/**
 * Element 是抽象意义上的源代码里的元素，例如包、类、接口、方法等等，可能是静态或者非静态
 * 需要注意使用equals方法比较
 * @author hushengdong
 */
type Element interface {
	/**
	 * 元素的类型
	 */
	GetTypeMirror() TypeMirror
	/**
	 * 是否相等，就看名字是不是一样就可以了
	 */
	GetName() string
	/**
	 * 元素的类型
	 */
	GetElementKind() *ElementKindEnum

	Element_()
}
