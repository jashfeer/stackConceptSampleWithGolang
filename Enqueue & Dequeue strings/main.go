package main

import "fmt"

type Queue []string

func (q *Queue) isEmptye() bool {
	return len(*q) == 0
}

func (q *Queue) Enqueue(new string) {
	*q = append(*q, new)
}
func (q *Queue) Dequeue() (string, bool) {
	if q.isEmptye() {
		return "", false
	} else {
		element := (*q)[0]
		*q = (*q)[1:]
		return element, true
	}
}

func main() {
	var myQueue Queue
	name,ok:=myQueue.Dequeue()
	fmt.Println(name,ok)
	fmt.Println(myQueue)
	myQueue.Enqueue("Rahul")
	myQueue.Enqueue("Raju")
	myQueue.Enqueue("Ramu")
	myQueue.Enqueue("Rajesh")
	myQueue.Enqueue("Rashid")
	fmt.Println(myQueue)
	name,ok=myQueue.Dequeue()
	fmt.Println(name,ok)
	fmt.Println(myQueue)

}
