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

func (stack *Stack) Push(element *Element) {
	stack.elements = append(stack.elements[:stack.elementCount], element)
	stack.elementCount++
}

func (stack *Stack) Pop() *Element {
	if stack.elementCount == 0 {
		return nil
	}

	var length = len(stack.elements)
	var element = stack.elements[length-1]

	if length > 1 {
		stack.elements = stack.elements[:length-1]
	} else {
		stack.elements = stack.elements[0:]
	}
	stack.elementCount = len(stack.elements)
	return element

}
