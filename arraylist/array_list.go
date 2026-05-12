package arraylist

import "errors"

type ArrayList struct {
	v    []int
	size int
}

func (list *ArrayList) Init(capacity int) {
	list.v = make([]int, capacity)
	list.size = 0
}

// função resize pra quando for necessário usar no add:
func (list *ArrayList) resize() {
	// definindo a nova capacidade (dobro da atual)
	newCapacity := len(list.v) * 2
	if newCapacity == 0 {
		newCapacity = 1
	}

	// alocando o novo array na memória
	newArray := make([]int, newCapacity)

	// copiando os elementos do array antigo para o novo
	// Aqui usamos um loop simples O(n)
	for i := 0; i < list.size; i++ {
		newArray[i] = list.v[i]
	}

	// substituíndo o array antigo pelo novo
	list.v = newArray
}

// função add:
func (list *ArrayList) Add(val int) {
	if list.size == len(list.v) {
		list.resize() // Aquela mesma lógica de dobrar o tamanho
	}
	list.v[list.size] = val
	list.size++
}

// função add on index:
func (list *ArrayList) AddOnIndex(val int, index int) error {
	if index < 0 || index > list.size {
		return errors.New("índice fora dos limites")
	}

	if list.size == len(list.v) {
		list.resize()
	}

	// A mágica (ou o pesadelo) acontece aqui:
	// Tem que empurrar todo mundo do 'index' para frente
	for i := list.size; i > index; i-- {
		list.v[i] = list.v[i-1]
	}

	list.v[index] = val
	list.size++
	return nil
}

// função get:
func (list *ArrayList) Get(index int) (int, error) {
	// Validação de segurança: o índice existe na parte "viva" da lista?
	if index < 0 || index >= list.size {
		return 0, errors.New("índice fora dos limites")
	}

	// Acesso direto
	return list.v[index], nil
}

// função set:
func (list *ArrayList) Set(val int, index int) error {
	if index < 0 || index >= list.size {
		return errors.New("índice fora dos limites")
	}

	list.v[index] = val
	return nil
}

// função remove:
func (list *ArrayList) Remove(index int) error {
	// Verificação de limites
	if index < 0 || index >= list.size {
		return errors.New("índice fora dos limites")
	}

	// Loop crescente para puxar os elementos para a esquerda
	// Começando no índice removido e indo até o penúltimo elemento
	for i := index; i < list.size-1; i++ {
		list.v[i] = list.v[i+1]
	}

	// zerando a última posição (boa prática)
	list.v[list.size-1] = 0

	// diminuindo o tamanho lógico
	list.size--

	return nil
}

func (list *ArrayList) Size() int {
	return list.size
}
