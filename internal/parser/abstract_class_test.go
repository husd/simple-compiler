package parser

import "testing"

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
}
