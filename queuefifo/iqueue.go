package queuefifo

type IQueue interface {
	Enqueue(value int)
	Dequeue() (int, error)
	Front() (int, error)
	IsEmpty() bool
	Size() int
}
