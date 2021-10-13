package parser

/**
 * 长度计算器
 * @author hushengdong
 */

func calcNewLength(len int, max int) int {

	for len < max+1 {
		len = len * 2
	}
	return len
}
