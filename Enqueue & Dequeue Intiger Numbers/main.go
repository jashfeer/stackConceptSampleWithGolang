package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}
func (q *Queue) Dequeue() int {
	removeValue := q.items[0]
	q.items = q.items[1:]
	return removeValue
}

func main() {
	myQueue := Queue{}
	fmt.Println(myQueue)
	myQueue.Enqueue(10)
	myQueue.Enqueue(20)
	myQueue.Enqueue(30)
	myQueue.Enqueue(40)
	fmt.Println(myQueue)
	value:=myQueue.Dequeue()
	fmt.Println(myQueue)
	fmt.Println(value)
	myQueue.Enqueue(50)
	fmt.Println(myQueue)
}
