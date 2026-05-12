package stackarray

type IStack_Array interface {
	Push(val int) error
	Pop() (int, error)
	Peek() (int, error)
	IsEmpty() bool
	Size() int
	IsFull() bool
}
