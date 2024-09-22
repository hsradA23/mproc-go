package message

type Push struct{
	Val int
}

type Pop struct{
	ReturnTo chan int
}

type Shutdown struct{}
