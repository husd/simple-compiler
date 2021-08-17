package lexical

import "testing"

func TestHandleError(t *testing.T) {
	type args struct {
		token *TokenTag
		msg   string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestCheckLexicalError(t *testing.T) {
	var tests = [...]struct {
		name    string
		args    string
		wantB   bool
		wantMsg string
	}{
		{"测试空", "", false, "无效的token，以数字开头:"},
		{"0abc数字开头", "0abc", false, "无效的token，以数字开头:0abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotB, gotMsg := CheckLexicalError(tt.args)
			if gotB != tt.wantB {
				t.Errorf("CheckLexicalError() gotB = %v, want %v", gotB, tt.wantB)
			}
			if gotMsg != tt.wantMsg {
				t.Errorf("CheckLexicalError() gotMsg = %v, want %v", gotMsg, tt.wantMsg)
			}
		})
	}
}
