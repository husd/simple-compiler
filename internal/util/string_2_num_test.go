package util

import "testing"

/**
 *
 * @author hushengdong
 */

func TestString2int(t *testing.T) {

	n, _ := String2int("123", 10, 32)
	AssertEquals(t, "测试10进制", 123, int(n))

	n, _ = String2int("11101010", 2, 32)
	AssertEquals(t, "测试2进制", 234, int(n))

	n, _ = String2int("34", 8, 32)
	AssertEquals(t, "测试8进制", 28, int(n))

	n, _ = String2int("43d234", 16, 32)
	AssertEquals(t, "测试16进制", 4444724, int(n))

}
