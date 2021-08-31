package util

import (
	"reflect"
	"testing"
)

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

	if !reflect.DeepEqual(want, got) {
		t.Errorf("msg: %v 期望值:    %v     实际值:    %v    ", msg, want, got)
	}
}
