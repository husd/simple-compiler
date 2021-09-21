package util

import (
	"reflect"
	"testing"
)

/**
 * 这个是单元测试的辅助工具类
 * @author hushengdong
 */
func AssertTrue(t *testing.T, msg string, b bool) {

	if !b {
		t.Errorf("msg: %v excepted: true, got: false ", msg)
	}
}

func AssertFalse(t *testing.T, msg string, b bool) {

	if b {
		t.Errorf("msg: %v excepted: false, got: true ", msg)
	}
}

func AssertEquals(t *testing.T, msg string, want, got interface{}) {
	switch want.(type) {
	case int, int8, int16, int32, int64, float32, float64, string:
		if want != got {
			goto printError
		}
	}
	if !reflect.DeepEqual(want, got) {
		goto printError
	}
	return
printError:
	t.Errorf("msg: %v 期望值:    %v     实际值:    %v    ", msg, want, got)
}
