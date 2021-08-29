package parser

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
		break
	case 8:
		res = digitRuneRadix8(r)
		break
	case 10:
		res = digitRuneRadix10(r)
		break
	case 16:
		res = digitRuneRadix16(r)
		break
	default:
		res = -1
	}
	return res
}

// 2进制只有 0-1 是合法的
func digitRuneRadix2(r rune) rune {

	if r == 0 || r == 1 {
		return r
	}
	return -1
}

// 8进制是 0-7 合法
func digitRuneRadix8(r rune) rune {

	if r >= 0 && r <= 7 {
		return r
	}
	return -1
}

// 10进制只有 0-9 是合法的
func digitRuneRadix10(r rune) rune {

	if r >= 0 && r <= 9 {
		return r
	}
	return -1
}

// 16进制只有 0-9 a-z A-Z 是合法的
// A-Z 的ASCII值对应于 65 -90
// a-z 的ASCII值对应于 97 - 122
func digitRuneRadix16(r rune) rune {

	if r >= 0 && r <= 9 {
		return r
	}
	if r >= 65 && r <= 90 {
		return r
	}
	if r >= 97 && r <= 122 {
		return r
	}
	return -1
}
