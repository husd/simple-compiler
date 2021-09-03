package parser

/**
 *
 * @author hushengdong
 */

// 如果 ch = 0x10 那么 会挨个解析这个数字 看其是不是符合对应的进制要求
// 例如 digitRune(3,2) = -1 因为2进制只有0和1
//  digitRune(0,2) = 0
// digitRune(1,2) = 1
// digitRune(a,16) = 9
// 这里我们简单处理
func digitRune(r rune, base int) rune {

	var res rune
	switch base {
	case 2:
		res = digitRuneRadix2(r)
	case 8:
		res = digitRuneRadix8(r)
	case 10:
		res = digitRuneRadix10(r)
	case 16:
		res = digitRuneRadix16(r)
	default:
		res = -1
	}
	return res
}

/**   2进制    ascii编码  10进制数
    00110000	48			0
	00110001	49			1
	00110010	50			2
	00110011	51			3
	00110100	52			4
	00110101	53			5
	00110110	54			6
	00110111	55			7
	00111000	56			8
	00111001	57			9
*/

// 注意： r 是unicode编码 ascii编码的 0-9 分别对应于 48 - 57
// 2进制只有 0-1 是合法的
func digitRuneRadix2(r rune) rune {

	if r == 48 || r == 49 {
		return r - 48
	}
	return -1
}

// 8进制是 0-7 合法 0的ascii编码是48 7的ascii编码是55
func digitRuneRadix8(r rune) rune {

	if r >= 48 && r <= 55 {
		return r - 48
	}
	return -1
}

// 10进制只有 0-9 是合法的
func digitRuneRadix10(r rune) rune {

	if r >= 48 && r <= 57 {
		return r - 48
	}
	return -1
}

// 16进制只有 0-9 a-z A-Z 是合法的
// A-Z 的ASCII值对应于 65 -90
// a-z 的ASCII值对应于 97 - 122
func digitRuneRadix16(r rune) rune {

	if r >= 48 && r <= 57 {
		return r - 48
	}
	if r >= 65 && r <= 90 {
		return r - 65 + 10
	}
	if r >= 97 && r <= 122 {
		return r - 97 + 10
	}
	return -1
}
