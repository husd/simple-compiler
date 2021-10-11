package parser

import (
	"fmt"
	"testing"
)

/**
 *
 * @author hushengdong
 */
type ATree interface {
	name() string
}

type ATree01 struct {
}

func NewATree01() ATree {

	return &ATree01{}
}

func NewATree02() ATree {

	return &ATree02{}
}

func (at *ATree01) name() string {
	return "ATree01"
}

func (at *ATree01) name01() string {
	return "ATree010101"
}

type ATree02 struct {
}

func (at *ATree02) name() string {
	return "ATree-02"
}

func (at *ATree02) name02() string {
	return "ATree020202"
}

func TestATree(t *testing.T) {

	fmt.Println("test.....")
	at01 := NewATree01()
	at02 := NewATree02()

	printName(at01)
	printName(at02)
}

func printName(at ATree) {

	fmt.Println("atree is:", at.name())
}
