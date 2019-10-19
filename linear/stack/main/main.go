package main

import (
	"fmt"

	"github.com/xdragon1015/DataStructures/linear/stack"
)

func main() {
	var stk = &stack.Stack{}
	stk.NewStack()
	var element1 = &stack.Element{
		ElementValue: 10,
	}
	// var element2 = &stack.Element{
	// 	ElementValue: 10,
	// }
	// var element3 = &stack.Element{
	// 	ElementValue: 10,
	// }
	// var element4 = &stack.Element{
	// 	ElementValue: 10,
	// }

	stk.Push(element1)

	fmt.Println("Stack caontains:", stk.Elements)

	stk.Pop()
	fmt.Println("Stack caontains:", stk.Elements)
}
