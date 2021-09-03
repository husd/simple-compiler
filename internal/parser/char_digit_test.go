package parser

import "testing"

/**
 *
 * @author hushengdong
 */
func Test_digitRuneRadix16(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{"测试16进制", args{'a'}, 10},
		{"测试16进制", args{'A'}, 10},
		{"测试16进制", args{'b'}, 11},
		{"测试16进制", args{'1'}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digitRuneRadix16(tt.args.r); got != tt.want {
				t.Errorf("digitRuneRadix16() arg:=%v = %v, want %v", tt.args, got, tt.want)
			} else {
				t.Logf("digitRuneRadix16() success arg:=%v ", tt.args)
			}
		})
	}
}
