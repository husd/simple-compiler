package util

import "unsafe"

func IsLittleEndian() bool {
	var i int32 = 0x01020304
	u := unsafe.Pointer(&i)
	pb := (*byte)(u) //强制转byte
	b := *pb
	return b == 0x04
}
