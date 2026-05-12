package linkedlist

import "errors"

// Node representa cada elemento (nó) da lista
type Node struct {
	value int
	next  *Node // Ponteiro para o próximo endereço de memória
}

// LinkedList representa a lista simplesmente encadeada
type LinkedList struct {
	head *Node // Ponteiro para o primeiro elemento
	size int   // Armazena o tamanho para evitar percorrer a lista sempre
}

// Init inicializa a lista encadeada
func (list *LinkedList) Init() {
	list.head = nil
	list.size = 0
}

// Size retorna a quantidade de elementos (útil para os arquivos de teste)
func (list *LinkedList) Size() int {
	return list.size
}

// funcao add:
func (list *LinkedList) Add(val int) {
	newNode := &Node{value: val, next: nil}

	// quando a lista está vazia:
	if list.head == nil {
		list.head = newNode
	} else {
		// quando a lista já tem gente (precisamos achar o último)
		current := list.head

		// enquanto o "próximo" não for nulo, eu continuo andando
		for current.next != nil {
			current = current.next
		}

		// achei o último! Agora o next dele aponta para o novo
		current.next = newNode
	}

	list.size++
}

// funcao add on index:
func (list *LinkedList) AddOnIndex(val int, index int) error {
	// validação de limites dos indices
	if index < 0 || index > list.size {
		return errors.New("índice fora dos limites")
	}

	// caso especial: inserção no início
	if index == 0 {
		newNode := &Node{value: val, next: list.head}
		list.head = newNode
		list.size++
		return nil
	}

	// caso geral: percorrer até o nó anterior (index - 1)
	newNode := &Node{value: val}
	current := list.head

	// o for estilo "while" caminha até o nó imediatamente antes do alvo
	for i := 0; i < index-1; i++ {
		current = current.next
	}

	// a troca de mãos (A ordem aqui é relevante)
	newNode.next = current.next // O novo aponta para o que vem depois
	current.next = newNode      // O anterior agora aponta para o novo

	list.size++
	return nil
}

// funcao get:
func (list *LinkedList) Get(index int) (int, error) {
	// validação de limites dos indices
	if index < 0 || index >= list.size {
		return 0, errors.New("índice fora dos limites")
	}

	// current começa na cabeça
	current := list.head

	// current percorre a lista até chegar no índice desejado
	for i := 0; i < index; i++ {
		current = current.next
	}

	// retorna o valor encontrado
	return current.value, nil
}

// funcao set:
func (list *LinkedList) Set(val int, index int) error {
	// validação de segurança
	if index < 0 || index >= list.size {
		return errors.New("índice fora dos limites")
	}

	// caminhamo até o nó alvo
	current := list.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	// atualiza o valor e retornam o sucesso
	current.value = val
	return nil
}

// funcao remove:
func (list *LinkedList) Remove(index int) error {
	// validação por segurança
	if index < 0 || index >= list.size {
		return errors.New("índice fora dos limites")
	}

	// caso especial: remover o primeiro elemento (index 0)
	if index == 0 {
		// o head simplesmente pula para o próximo nó
		// o nó antigo fica "órfão" e o Garbage Collector limpa sozinho, não precisa dar free() como em C/C++
		list.head = list.head.next
	} else {
		// caso geral: percorrer até o nó ANTERIOR ao que será removido
		current := list.head
		for i := 0; i < index-1; i++ {
			current = current.next
		}

		// a manobra: o anterior aponta para o "neto", pulando o alvo
		// current.next é o nó que queremos remover
		// current.next.next é o nó que vem depois dele
		current.next = current.next.next
	}

	// decrementa o tamanho
	list.size--
	return nil
}
