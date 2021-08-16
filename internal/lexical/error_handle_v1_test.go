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
