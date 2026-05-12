package stacklinkedlist

import "errors"

// Node representa um elemento na pilha
type Node struct {
	value int
	next  *Node
}

// Stack representa a estrutura da pilha
type Stack struct {
	top  *Node // O topo é por onde tudo entra e sai
	size int
}

// Init inicializa a pilha
func (s *Stack) Init() {
	s.top = nil
	s.size = 0
}

// Size retorna o tamanho atual
func (s *Stack) Size() int {
	return s.size
}

// funcao IsEmpty:
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

// funcao push:
func (s *Stack) Push(val int) {
	// o novo elemento é inicializado com o 'next' dele já apontando para quem era o topo anterior
	// já que ele sempre será introduzido no topo da pilha
	// se a pilha estiver vazia, o 's.top' é nil, então o next do novo elemento estará apontando pra nil, oque é de fato o correto
	// com essa implementação, os casos da pilha estar vazia ou preenchida são a mesma coisa
	newNode := &Node{value: val, next: s.top}

	// ponteiro do topo aponta pro novo elemento
	s.top = newNode

	// incrementação do tamanho
	s.size++

}

// funcao pop:
func (s *Stack) Pop() (int, error) {
	// se a pilha estiver vazia, não tem oque sair dela
	if s.IsEmpty() {
		return 0, errors.New("A pilha está vazia")
	}

	// caso a pilha nao esteja vazia
	saida := s.top.value
	s.top = s.top.next

	// decrementa o tamanho da pilha
	s.size--

	return saida, nil
}

// funcao peek:
func (s *Stack) Peek() (int, error) {
	// se a pilha estiver vazia, nao tem oque espiar
	if s.IsEmpty() {
		return 0, errors.New("A pilha está vazia")
	}

	// caso a pilha nao esteja vazia
	return s.top.value, nil
}
