package stack

import (
	"strconv"
)

type Element struct {
	elementValue int
}

type Stack struct {
	elements     []*Element
	elementCount int
}

func (stack *Stack) NewStack() {
	stack.elements = make([]*Element, 0)
}

func (element *Element) String() string {
	return strconv.Itoa(element.elementValue)
}
