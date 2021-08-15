package lexical

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWordLexical_LexicalAnalysis(t *testing.T) {

	wordLexer := WordLexical{}

	singleComment := "  abc e;// ..... 单行注释    "
	res := wordLexer.LexicalAnalysis(singleComment)
	got := [2]string{}
	for i, j := res.Front(), 0; i != nil; i, j = i.Next(), j+1 {
		got[j] = fmt.Sprintf("%v", i.Value)
	}
	want := [...]string{"abc", "e;"}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("excepted:%v, got:%v", want, got)
	}
}
