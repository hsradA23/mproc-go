package main

import (
	"fmt"
	"mproc/message"
)

func sep_thread(inp chan interface{}) {
	stack := []int{}
	for rec := range inp {
		switch v := rec.(type) {
		case message.Pop:
			if len(stack) > 0 {
				v.ReturnTo <- stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			} else {
				panic("Stack underflow")
			}
		case message.Push:
			stack = append(stack, v.Val)
		case message.Shutdown:
			close(inp)
		default:
			panic("Unknown type recieved.")
		}
	}
}

func pop(inp chan interface{}) int {
	get_from := make(chan int)
	m := message.Pop{ReturnTo: get_from}
	inp <- m
	rec := <-get_from
	close(get_from)
	return rec
}

func push(inp chan interface{}, val int) {
	m := message.Push{Val: val}
	inp <- m
}

func main() {
	com := make(chan interface{})
	go sep_thread(com)

	push(com, 1)
	push(com, 2)
	push(com, 3)

	got := pop(com)
	fmt.Println("Got value: ", got)
	got = pop(com)
	fmt.Println("Got value: ", got)
	got = pop(com)
	fmt.Println("Got value: ", got)

	// got = pop(com) // Will panic
	// fmt.Println("Got value: ", got)

	m := message.Shutdown{}
	com <- m
}
