package queuefifo

import "errors"

type LinkedListQueue struct {
	front *Node
	rear  *Node
	size  int
}

type Node struct {
	val  int
	next *Node
}

// Enqueue: adiciona no final
func (queue *LinkedListQueue) Enqueue(val int) {
	newNode := &Node{val: val, next: nil}

	if queue.IsEmpty() {
		queue.front = newNode // Fila vazia: o novo nó é o primeiro
	} else {
		queue.rear.next = newNode // Fila com itens: o último atual aponta pro novo
	}

	queue.rear = newNode // O novo nó SEMPRE será o último
	queue.size++
}

// Dequeue: remove do início
func (queue *LinkedListQueue) Dequeue() (int, error) {
	if queue.IsEmpty() {
		return 0, errors.New("a fila está vazia")
	}

	val := queue.front.val
	queue.front = queue.front.next // O início passa a ser o próximo nó
	queue.size--

	// Se a remoção esvaziou a fila, o ponteiro rear precisa ser limpo
	// para não ficar apontando para um endereço de memória "fantasma"
	if queue.IsEmpty() {
		queue.rear = nil
	}

	return val, nil
}

// Front: Espia o primeiro elemento sem remover
func (queue *LinkedListQueue) Front() (int, error) {
	if queue.IsEmpty() {
		return 0, errors.New("a fila está vazia")
	}
	return queue.front.val, nil
}

func (queue *LinkedListQueue) IsEmpty() bool {
	return queue.size == 0
}

func (queue *LinkedListQueue) Size() int {
	return queue.size
}
