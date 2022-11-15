package main

import "fmt"

// Struct for queue

type queue struct {

	values []int

//values slice inside the struct

}

func (q *queue) enqueue(i int) {
	
//enqueue is a method for append the values to slice.
	
q.values = append(q.values, i)

}

func main() {   //main function

	myqueue := queue{}
	myqueue.enqueue(1);myqueue.enqueue(2);myqueue.enqueue(3);
	myqueue.enqueue(4);myqueue.enqueue(5);fmt.Println(myqueue)
	myqueue.dequeue();fmt.Println(myqueue)
	myqueue.dequeue();fmt.Println(myqueue)
	myqueue.dequeue();fmt.Println(myqueue)

}

//method dequeue for remove elements from the slice 
 
func (q *queue) dequeue()int{
	
	toremove:=q.values[0]
	q.values=q.values[1:]
	return toremove

}