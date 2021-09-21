package parser

import (
	"fmt"
	"testing"
)

/**
 *
 * @author hushengdong
 */

func Test_abstractPerson_eat(t *testing.T) {

	man := NewMan()
	man.eat()
	man.run()

	woman := NewWomen()
	woman.eat()
	woman.run()

	eatAndRun(man)
	eatAndRun(woman)
}

const (
	a, b = iota + 1, iota + 2
	c, d
	e, f
)

func Test_iota(t *testing.T) {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
