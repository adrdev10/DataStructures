package queue

//Queue represents a priorityQueue
type Queue []*Order

//Order is the type assigned to the priority queue
type Order struct {
	priority int
	quantity int
	product  string
}

//NewQueue creates a new queue object
func (order *Order) NewOrder(priority, quantity int, product string) *Order {
	return &Order{
		priority: priority,
		quantity: quantity,
		product:  product,
	}
}

//Add adds new element to the queue
func (queue *Queue) Add(order *Order) {
	if len(*queue) == 0 {
		*queue = append(*queue, order)
	}
	var appended = false
	for i, addedOrder := range *queue {
		if addedOrder.priority < order.priority {
			*queue = append((*queue)[:i], append(Queue{order}, (*queue)[i:]...)...)
			appended = true
		}
	}
	if !appended {
		*queue = append(*queue, order)
	}
}
