package stackarray

import "errors"

// StackArray representa a pilha implementada com um array (slice)
type StackArray struct {
	elements []int
	capacity int
	top      int // Índice do elemento que está no topo (-1 se estiver vazia)
}

// Init inicializa a pilha com uma capacidade fixa
func (s *StackArray) Init(cpcty int) {
	s.elements = make([]int, cpcty)
	s.capacity = cpcty
	s.top = -1 // Começa em -1 para indicar que não há ninguém
}

// Size retorna a quantidade de elementos atuais
func (s *StackArray) Size() int {
	// se o top está no índice 0, tem 1 elemento.
	return s.top + 1
}

// funcao IsFull:
func (s *StackArray) IsFull() bool {
	// se o top está no indice igual a capacidade - 1, significa que a array está cheia/completamente preenchida
	return s.top == s.capacity-1
}

// funcao IsEmpty:
func (s *StackArray) IsEmpty() bool {
	// se o top esta no indice -1, que é como ele é inicializado vazio, significa que a matriz está vazia
	return s.top == -1
}

func (s *StackArray) reSize() {
	// define a nova capacidade como o dobro da atual
	newCapacity := len(s.elements) * 2

	// caso a pilha tenha começado com capacidade 0, inicializa com 1
	if newCapacity == 0 {
		newCapacity = 1
	}

	// aloca um novo slice com a nova capacidade
	newElements := make([]int, newCapacity)

	// copia os elementos da pilha antiga para a nova
	// como a pilha cresce linearmente a partir do índice 0, o copy resolve perfeitamente (diferente do deque, por array, que é circular)
	copy(newElements, s.elements)

	// substitui o slice antigo pelo novo redimensionado
	s.elements = newElements
	s.capacity = newCapacity
}

// funcao push:
func (s *StackArray) Push(val int) error {
	// se a pilha/array estiver cheia
	if s.IsFull() {
		s.reSize()
	}

	// caso nao esteja cheia

	// incrementando o ponteiro pro elemento no "topo", pra fazer o push diretamente no index
	s.top++

	// inserindo diretamente no index
	s.elements[s.top] = val

	return nil
}

// funcao pop:
func (s *StackArray) Pop() (int, error) {
	// caso a pilha esteja vazia, não tem oque tirar
	if s.IsEmpty() {
		return 0, errors.New("A pilha está vazia")
	}

	// caso nao esteja vazia

	// armazenando oque eu quero tirar
	saida := s.elements[s.top]

	// zerando aquele indice da matriz pra "esvaziar" a pilha
	s.elements[s.top] = 0

	// movendo o indice que indica o topo pro elemento anterior
	s.top = s.top - 1

	return saida, nil
}

// funcao peek:
func (s *StackArray) Peek() (int, error) {
	// se a pilha estiver vazia, nao tem oque espiar
	if s.IsEmpty() {
		return 0, errors.New("A pilha está vazia")
	}

	// se nao esta vazia
	return s.elements[s.top], nil
}
