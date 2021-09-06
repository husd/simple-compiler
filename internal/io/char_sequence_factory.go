package io

/**
 *
 * @author hushengdong
 */
// GetCharSequenceFromFactory 工厂类 有了这个工厂类，以后可以有多种策略，来获取这个CharSequence。
func GetCharSequenceFromFactory(path string) *CharSequence {

	//默认的文件类型的字符
	charSeq := NewStringCharSequence(path)
	return &charSeq
}
