package main

import (
	"fmt"

	"github.com/xdragon1015/DataStructures/linear/queue"
)

func main() {
	var q queue.Queue
	q = make(queue.Queue, 0)
	or1 := &queue.Order{}

	quantity := 10
	priority := 20
	orderName := "Food"

	or1 = or1.NewOrder(priority, quantity, orderName)
	quantity2 := 100
	priority2 := 200
	orderName2 := "Beer"

	or2 := &queue.Order{}
	or2 = or2.NewOrder(priority2, quantity2, orderName2)

	q.Add(or1)
	q.Add(or2)

	for _, val := range q {
		fmt.Println(*val)
	}

}
