package stacklinkedlist

type IStack_LinkedList interface {
	Push(val int)
	Pop() (int, error)
	Peek() (int, error)
	IsEmpty() bool
	Size() int
}
