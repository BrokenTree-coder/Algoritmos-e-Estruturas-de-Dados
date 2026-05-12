package linkedlist

type ILinkedList interface {
	Add(val int)
	AddOnIndex(val int, index int) error
	Get(index int) (int, error)
	Set(val int, index int) error
	Remove(index int) error
	Size() int
}
