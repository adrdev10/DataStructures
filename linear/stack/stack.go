package stack

import (
	"strconv"
)

type Element struct {
	ElementValue int
}

type Stack struct {
	Elements     []*Element
	elementCount int
}

func (stack *Stack) NewStack() {
	stack.Elements = make([]*Element, 0)
}

func (element *Element) String() string {
	return strconv.Itoa(element.ElementValue)
}

func (stack *Stack) Push(element *Element) {
	stack.Elements = append(stack.Elements[:stack.elementCount], element)
	stack.elementCount++
}

func (stack *Stack) Pop() *Element {
	if stack.elementCount == 0 {
		return nil
	}

	var length = len(stack.Elements)
	var element = stack.Elements[length-1]

	if length > 1 {
		stack.Elements = stack.Elements[:length-1]
	} else {
		stack.Elements = stack.Elements[:0]
	}
	stack.elementCount = len(stack.Elements)
	return element

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
