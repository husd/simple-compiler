package util

import (
	"testing"
)

/**
 * 测试自定义上下文
 * @author hushengdong
 */

func TestContext_Get(t *testing.T) {

	c := NewContext()

	const name = "TestObj"
	var real = &TestObj{"xiaoming", 10}
	c.Put(name, real)
	var obj *TestObj
	ok, testObj := c.Get(name)
	if ok {
		obj = testObj.(*TestObj)
	} else {
		obj = real
	}
	AssertEquals(t, "测试context", "xiaoming", obj.Name)
	AssertEquals(t, "测试context", 10, obj.Age)
	AssertEquals(t, "测试context", real, obj)
}

type TestObj struct {
	Name string
	Age  int
}
