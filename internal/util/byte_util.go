package util

import "unsafe"

/**
 * 这个类暂时还没用上，后续可能会涉及到字节操作的一些顺序的问题
 * 通用类型的方法，先放着
 * @author hushengdong
 */
func IsLittleEndian() bool {
	var i int32 = 0x01020304
	u := unsafe.Pointer(&i)
	pb := (*byte)(u) //强制转byte
	b := *pb
	return b == 0x04
}
