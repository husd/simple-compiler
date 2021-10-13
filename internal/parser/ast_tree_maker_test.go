package parser

import (
	"fmt"
	"testing"
)

/**
 *
 * @author hushengdong
 */

func TestDummyTreeNode(t *testing.T) {

	dummyNode := NewDummyTreeNode()
	dummyNode.Append(GetEmptyTreeNode())
	dummyNode.Append(GetEmptyTreeNode())
	// dummyNode.Append(GetEmptyTreeNode())

	fmt.Println(dummyNode)
	fmt.Println(dummyNode.GetFirstChildren())
	fmt.Println("------")
}
