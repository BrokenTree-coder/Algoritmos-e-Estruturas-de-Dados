package deque

import "errors"

type DoublyLinkedListDeque struct {
	front *Node
	back  *Node
	size  int
}

type Node struct {
	prev *Node
	val  int
	next *Node
}

func (deque *DoublyLinkedListDeque) PushFront(val int) {
	newNode := &Node{val: val}

	// Deque Vazio (Cláusula de Guarda)
	if deque.IsEmpty() {
		deque.front = newNode
		deque.back = newNode
		deque.size++
		return
	}

	// Caso Geral: Pelo menos um elemento já existe
	newNode.next = deque.front
	deque.front.prev = newNode
	deque.front = newNode
	deque.size++
}

func (deque *DoublyLinkedListDeque) PushBack(val int) {
	newNode := &Node{val: val}

	if deque.IsEmpty() {
		deque.front = newNode
		deque.back = newNode
		deque.size++
		return
	}

	newNode.prev = deque.back
	deque.back.next = newNode
	deque.back = newNode
	deque.size++
}

func (deque *DoublyLinkedListDeque) PopFront() (int, error) {
	if deque.IsEmpty() {
		return 0, errors.New("o deque está vazio")
	}

	saida := deque.front.val

	// Caso Especial: Apenas UM elemento
	// Se front e back são o mesmo, a lista vai sumir/ficar vazia
	if deque.size == 1 {
		deque.front = nil
		deque.back = nil
		deque.size = 0
		return saida, nil
	}

	// Caso Geral: Mais de um elemento
	// agora sabendo que o next não é 'nil', podemos manipular esses ponteiros sem medo de dar algo como 'nil.prev = nil', um erro grave
	// que aconteceria se o tamanho do deque fosse 1
	deque.front.next.prev = nil
	deque.front = deque.front.next
	deque.size--

	return saida, nil
}

func (deque *DoublyLinkedListDeque) PopBack() (int, error) {
	if deque.IsEmpty() {
		return 0, errors.New("o deque está vazio")
	}

	saida := deque.back.val

	if deque.size == 1 {
		deque.front = nil
		deque.back = nil
		deque.size = 0
		return saida, nil
	}

	deque.back.prev.next = nil
	deque.back = deque.back.prev
	deque.size--

	return saida, nil
}

func (deque *DoublyLinkedListDeque) Front() (int, error) {
	if deque.IsEmpty() {
		return 0, errors.New("O deque está vazio")
	}

	return deque.front.val, nil
}

func (deque *DoublyLinkedListDeque) Back() (int, error) {
	if deque.IsEmpty() {
		return 0, errors.New("O deque está vazio")
	}

	return deque.back.val, nil
}

func (deque *DoublyLinkedListDeque) IsEmpty() bool {
	return deque.size == 0
}

func (deque *DoublyLinkedListDeque) Size() int {
	return deque.size
}
