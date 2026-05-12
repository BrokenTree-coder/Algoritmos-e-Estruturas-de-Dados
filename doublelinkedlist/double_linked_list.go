package doublelinkedlist

import "errors"

// Node representa cada elemento da double linked list
type Node struct {
	value int
	next  *Node // Ponteiro para o próximo
	prev  *Node // Ponteiro para o anterior
}

// DoublwLinkedList representa a lista com cabeça e cauda
type DoublyLinkedList struct {
	head *Node // Início da lista
	tail *Node // Fim da lista
	size int
}

// Init inicializa a lista vazia
func (list *DoublyLinkedList) Init() {
	list.head = nil
	list.tail = nil
	list.size = 0
}

// Size retorna a quantidade de elementos
func (list *DoublyLinkedList) Size() int {
	return list.size
}

// funcao add:
func (list *DoublyLinkedList) Add(val int) {
	// Cria o novo elemento (next e prev já nascem como nil por padrão no Go,
	// mas deixei explícito aqui para bater com a lógica)
	newNode := &Node{value: val, next: nil, prev: nil}

	if list.size == 0 {
		// Caso: Lista Vazia
		// Tanto o head quanto o tail devem apontar pro mesmo elemento
		list.head = newNode
		list.tail = newNode
	} else {
		// Caso Geral: Lista já tem elementos
		// O prev do elemento novo deve apontar pro tail anterior
		newNode.prev = list.tail

		// O next do tail anterior deve apontar pro elemento novo
		list.tail.next = newNode

		// O tail da lista deve apontar pro novo elemento
		list.tail = newNode
	}

	// Incrementação do size
	list.size++
}

// funcao add on index:
func (list *DoublyLinkedList) AddOnIndex(val int, index int) error {
	// Validação de segurança
	if index < 0 || index > list.size {
		return errors.New("índice fora dos limites")
	}

	// Caso Especial: Inserção no fim (Exatamente o Add)
	if index == list.size {
		list.Add(val)
		return nil
	}

	// Caso Especial: Inserção no início (index 0)
	if index == 0 {
		newNode := &Node{value: val, next: list.head, prev: nil}
		// O prev do antigo head agora aponta pro novo nó
		if list.head != nil {
			list.head.prev = newNode
		}
		// Atualiza o ponteiro principal da lista
		list.head = newNode
		list.size++
		return nil
	}

	// Caso Geral: Inserção no meio

	// Otimização: Decide de qual lado é mais rápido começar a buscar o índice
	var posterior *Node
	if index < list.size/2 {
		// Começa do início e vai para frente
		posterior = list.head
		for i := 0; i < index; i++ {
			posterior = posterior.next
		}
	} else {
		// Começa do fim e vai para trás
		posterior = list.tail
		for i := list.size - 1; i > index; i-- {
			posterior = posterior.prev
		}
	}

	// Como o 'posterior' está na posição onde queremos inserir,
	// o 'anterior' é simplesmente o nó que vem antes dele.
	anterior := posterior.prev

	// Cria o novo nó já com os engates de ida e volta
	newNode := &Node{value: val, prev: anterior, next: posterior}

	// Finaliza a troca das setas
	anterior.next = newNode
	posterior.prev = newNode

	list.size++
	return nil
}

// funcao get:
func (list *DoublyLinkedList) Get(index int) (int, error) {
	// Validação de segurança
	if index < 0 || index >= list.size {
		return 0, errors.New("índice fora dos limites")
	}

	// A otimização de busca
	var current *Node
	if index < list.size/2 {
		// Busca da esquerda para a direita
		current = list.head
		for i := 0; i < index; i++ {
			current = current.next
		}
	} else {
		// Busca da direita para a esquerda
		current = list.tail
		for i := list.size - 1; i > index; i-- {
			current = current.prev
		}
	}

	// Retorna o valor
	return current.value, nil
}

// funcao set:
func (list *DoublyLinkedList) Set(val int, index int) error {
	// Validação de segurança
	if index < 0 || index >= list.size {
		return errors.New("índice fora dos limites")
	}

	// A mesma otimização de busca
	var current *Node
	if index < list.size/2 {
		current = list.head
		for i := 0; i < index; i++ {
			current = current.next
		}
	} else {
		current = list.tail
		for i := list.size - 1; i > index; i-- {
			current = current.prev
		}
	}

	// Atualiza o valor e retorna sucesso
	current.value = val
	return nil
}

// funcao remove:
func (list *DoublyLinkedList) Remove(index int) error {
	// Validação de segurança
	if index < 0 || index >= list.size {
		return errors.New("índice fora dos limites")
	}

	// Caso Especial: Removendo o ÚNICO elemento da lista
	if list.size == 1 {
		list.head = nil
		list.tail = nil
		list.size--
		return nil
	}

	// Caso Especial: Removendo do início (index 0)
	if index == 0 {
		remover := list.head
		list.head = remover.next // Head pula pro segundo elemento
		list.head.prev = nil     // O novo primeiro não tem ninguém atrás dele

		remover.next = nil // Limpando o ponteiro (pra nao ter ninguem apontando pra ele e o Go limpar ele sozinho)
		list.size--
		return nil
	}

	// Caso Especial: Removendo do fim (index == size - 1)
	if index == list.size-1 {
		remover := list.tail
		list.tail = remover.prev // Tail recua pro penúltimo elemento
		list.tail.next = nil     // O novo último não aponta pra ninguém à frente

		remover.prev = nil // Limpando o ponteiro
		list.size--
		return nil
	}

	// Caso Geral: Removendo do meio da lista

	// Otimização de busca para achar o nó que será removido
	var current *Node
	if index < list.size/2 {
		current = list.head
		for i := 0; i < index; i++ {
			current = current.next
		}
	} else {
		current = list.tail
		for i := list.size - 1; i > index; i-- {
			current = current.prev
		}
	}

	// A troca de ponteiros (Isolando o current)
	anterior := current.prev
	posterior := current.next

	anterior.next = posterior
	posterior.prev = anterior

	// Limpando os ponteiros do elemento removido para o Garbage Collector
	current.prev = nil
	current.next = nil

	list.size--
	return nil
}
