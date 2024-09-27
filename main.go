package main

import (
	"fmt"
	"mproc/message"
	"strconv"
	"time"
)

var count int

func sep_thread(inp chan interface{}) {
	stack := []int{}
	dec := make(map[int]string)

	for rec := range inp {
		count += 1
		switch v := rec.(type) {
		// stack messages
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
			fmt.Printf("Handled %d requests", count)
			close(inp)

		// Dict messages
		case message.Set:
			dec[v.Key] = v.Val
		case message.Get:
			v.ReturnTo <- dec[v.Key]

		default:
			panic("Unknown type recieved.")
		}
	}
}

func pop(inp chan interface{}) int {
	get_from := make(chan int)
	defer close(get_from)
	m := message.Pop{ReturnTo: get_from}
	inp <- m
	return <-get_from
}

func push(inp chan interface{}, val int) {
	m := message.Push{Val: val}
	inp <- m
}

func set(inp chan interface{}, key int, val string){
	m := message.Set{Key: key, Val: val}
	inp <- m
}

func get(inp chan interface{}, key int) string{

	get_from := make(chan string)
	defer close(get_from)
	m := message.Get{Key: key, ReturnTo: get_from}
	inp <- m
	return <-get_from
}


func update_dict(inp chan interface{}){
	for i := 0; i < 100_000_000_000; i++ {
		val, err := strconv.Atoi(get(inp, 1))
		if err != nil{
			panic("Cannot parse")
		}
		str := strconv.Itoa(val+1)
		set(inp, i, str)
	}
}

func main() {
	com := make(chan interface{})
	go sep_thread(com)
	count = 0

	// push(com, 1)
	// push(com, 2)
	// push(com, 3)

	// got := pop(com)
	// fmt.Println("Got value: ", got)
	// got = pop(com)
	// fmt.Println("Got value: ", got)
	// got = pop(com)
	// fmt.Println("Got value: ", got)

	// // got = pop(com) // Will panic
	// // fmt.Println("Got value: ", got)

	// m := message.Shutdown{}
	// com <- m

	set(com, 1, "1")
	go update_dict(com)
	go update_dict(com)
	go update_dict(com)
	go update_dict(com)
	go update_dict(com)
	go update_dict(com)
	go update_dict(com)
	go update_dict(com)
	go update_dict(com)
	go update_dict(com)
	go update_dict(com)
	go update_dict(com)
	go update_dict(com)
	go update_dict(com)
	go update_dict(com)
	go update_dict(com)
	go update_dict(com)
	go update_dict(com)
	go update_dict(com)
	go update_dict(com)
	go update_dict(com)
	go update_dict(com)
	go update_dict(com)
	go update_dict(com)
	go update_dict(com)
	go update_dict(com)
	go update_dict(com)
	go update_dict(com)
	time.Sleep(time.Second*1)
	// fmt.Println("Got value:", get(com, 2))

	fmt.Printf("Handled %d requests\n", count)
}
