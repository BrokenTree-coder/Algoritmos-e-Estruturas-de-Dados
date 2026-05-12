package deque

type IDeque interface {
	PushFront(value int)
	PushBack(value int)
	PopFront() (int, error)
	PopBack() (int, error)
	Front() (int, error)
	Back() (int, error)
	IsEmpty() bool
	Size() int
}
