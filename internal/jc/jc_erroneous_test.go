package jc

import (
	"fmt"
	"testing"
)

/**
 *
 * @author hushengdong
 */

func TestNewJCError(_ *testing.T) {

	t := NewJCError()
	fmt.Println(t.getTreeType())
	fmt.Println(t.getTag())
}
