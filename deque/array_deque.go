package deque

import "errors"

type ArrayDeque struct {
	v     []int
	front int // indice do primeiro elemento da array
	back  int // indice do ultimo elemento da array
	size  int
}

// capacity define a capacidade inicial do vetor
func (deque *ArrayDeque) Init(capacity int) {
	deque.front = 0
	deque.back = capacity - 1
	deque.size = 0
	deque.v = make([]int, capacity) // Aloca o slice com o tamanho certo
}

func (deque *ArrayDeque) reSize() {
	// dobrando a capacidade atual
	newCapacity := len(deque.v) * 2

	// se a capacidade inicial fosse 0, define um valor padrão (ex: 1) para evitar multiplicar por zero
	if newCapacity == 0 {
		newCapacity = 1
	}

	// criando o novo array maior
	newArray := make([]int, newCapacity)

	// "desenrolando" a fila antiga para dentro do novo array
	for i := 0; i < deque.size; i++ {
		// encontra o índice correto no array antigo respeitando a circularidade
		oldIndex := (deque.front + i) % len(deque.v)

		// coloca no novo array de forma alinhada, começando do 0
		newArray[i] = deque.v[oldIndex]
	}

	// atualizando os ponteiros da struct
	deque.v = newArray
	deque.front = 0             // O início volta a ser o índice 0
	deque.back = deque.size - 1 // O final é logo após o último elemento inserido
}

// quando o vetor estiver cheio, sua capacidade deve ser duplicada
// os elementos sao copiados respeitando a ordem lógica do deque e reajuste front e back
func (deque *ArrayDeque) PushBack(val int) {
	// se estiver cheio, dobra de tamanho antes de inserir
	if deque.size == len(deque.v) {
		deque.reSize()
	}

	// Matemática Circular: Avança o back
	// Se back era -1 (vazio), vira 0. Se era o último índice, volta para 0.
	deque.back = (deque.back + 1) % len(deque.v)

	// insere o valor e atualiza o tamanho
	deque.v[deque.back] = val
	deque.size++
}

// quando o vetor estiver cheio, sua capacidade deve ser duplicada
// os elementos sao copiados respeitando a ordem lógica do deque e reajuste front e back
func (deque *ArrayDeque) PushFront(val int) {
	// se estiver cheio, redimensiona (único caso O(n))
	if deque.size == len(deque.v) {
		deque.reSize()
	}

	// recua o front circularmente (sem loop/for)
	// somando len(deque.v) para garantir que o resultado antes do % nunca seja negativo
	deque.front = (deque.front - 1 + len(deque.v)) % len(deque.v)

	// insere na nova posição e atualiza o tamanho
	deque.v[deque.front] = val
	deque.size++
}

func (deque *ArrayDeque) PopFront() (int, error) {
	if deque.IsEmpty() {
		return 0, errors.New("o deque está vazio")
	}

	// salva o valor que está na frente
	saida := deque.v[deque.front]

	// limpeza opcional (boa prática para evitar memory leaks de objetos grandes)
	deque.v[deque.front] = 0

	// avança o front circularmente (+1) assim como o pushback
	deque.front = (deque.front + 1) % len(deque.v)

	// atualiza o tamanho
	deque.size--

	return saida, nil
}

func (deque *ArrayDeque) PopBack() (int, error) {
	if deque.IsEmpty() {
		return 0, errors.New("o deque está vazio")
	}

	// salva o valor que está na frente
	saida := deque.v[deque.back]

	// limpeza opcional (boa prática para evitar memory leaks de objetos grandes)
	deque.v[deque.back] = 0

	// recua o back circularmente (-1) assim como o pushfront
	deque.back = (deque.back - 1 + len(deque.v)) % len(deque.v)

	// atualiza o tamanho
	deque.size--

	return saida, nil
}

func (deque *ArrayDeque) Front() (int, error) {
	// verificação se o deque está vazio
	if deque.IsEmpty() {
		return 0, errors.New("O deque está vazio")
	}

	return deque.v[deque.front], nil
}

func (deque *ArrayDeque) Back() (int, error) {
	// verificação se o deque está vazio
	if deque.IsEmpty() {
		return 0, errors.New("O deque está vazio")
	}

	return deque.v[deque.back], nil
}

func (deque *ArrayDeque) IsEmpty() bool {
	return deque.size == 0
}

func (deque *ArrayDeque) Size() int {
	return deque.size
}
