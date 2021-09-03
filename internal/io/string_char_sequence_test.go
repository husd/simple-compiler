package io

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

/**
 *
 * @author hushengdong
 */
func TestStringCharSequence_CharAt(t *testing.T) {
	type fields struct {
		length  int
		content string
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   rune
	}{
		{"测试汉字是否正确", fields{4, "你好中国"}, args{0}, '你'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := StringCharSequence{
				length:  tt.fields.length,
				content: tt.fields.content,
			}
			if got := f.CharAt(tt.args.index); got != tt.want {
				t.Errorf("CharAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringAt(t *testing.T) {

	// 内置的len()函数，返回的字符串编码的长度，是byte数量
	//unicode unit 并不是byte，而是不确定的，utf8是一种变长的字符编码，所以
	//想要拿到实际的字面量的长度，需要调用另外的一个方法
	str := "你好世界123"
	fmt.Println("字节数量：", len(str))                              // 12 + 3 = 15 个 4个汉字占12个字节
	fmt.Println("unicode unit数量：", utf8.RuneCountInString(str)) // 7个长度
	utf8.DecodeRuneInString(str)
}
