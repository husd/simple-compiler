package util

import (
	"strconv"
)

/**
 * 把字符串转换为相应进制的数字
 * @author hushengdong
 */

func String2int(str string, radix int) (n int64, err error) {

	return string2num(str, radix, 32)
}

func String2long(str string, radix int) (n int64, err error) {

	return string2num(str, radix, 64)
}

func string2num(str string, radix int, bitSize int) (n int64, err error) {

	if radix == 2 ||
		radix == 8 ||
		radix == 10 ||
		radix == 16 {
		n, err := strconv.ParseInt(str, radix, bitSize)
		return n, err
	}
	// 这样会返回一个错误
	return strconv.ParseInt(str, radix, -1)
}
