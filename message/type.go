package message

type Push struct{
	Val int
}


type Pop struct{
	ReturnTo chan int
}

type Shutdown struct{}

// --------------------

type Get struct{
	Key int
	ReturnTo chan string
}
type Set struct{
	Key int
	Val string
}
