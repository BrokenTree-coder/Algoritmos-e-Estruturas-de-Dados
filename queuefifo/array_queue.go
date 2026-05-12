package queuefifo

import "errors"

type ArrayQueue struct {
	v     []int // O array (slice) que vai guardar os dados
	front int   // O índice do primeiro elemento
	rear  int   // O índice onde o PRÓXIMO elemento será inserido
	size  int   // A quantidade atual de elementos na fila
}

// Init inicializa o array com a capacidade máxima fornecida
func (queue *ArrayQueue) Init(capacity int) {
	queue.v = make([]int, capacity) // Aloca o slice com o tamanho certo
	queue.front = 0
	queue.rear = 0
	queue.size = 0
}

// Enqueue adiciona um valor no final da fila circular

// função auxiliar privada para dobrar o tamanho da fila
func (queue *ArrayQueue) resize() {
	// dobra a capacidade atual
	newCapacity := len(queue.v) * 2

	// Se a capacidade inicial fosse 0, define um valor padrão (ex: 1) para evitar multiplicar por zero
	if newCapacity == 0 {
		newCapacity = 1
	}

	// cria o novo array maior
	newArray := make([]int, newCapacity)

	// "Desenrolamos" a fila antiga para dentro do novo array
	for i := 0; i < queue.size; i++ {
		// Encontra o índice correto no array antigo respeitando a circularidade
		oldIndex := (queue.front + i) % len(queue.v)

		// Coloca no novo array de forma alinhada, começando do 0
		newArray[i] = queue.v[oldIndex]
	}

	// atualiza os ponteiros da struct
	queue.v = newArray
	queue.front = 0         // O início volta a ser o índice 0
	queue.rear = queue.size // O final é logo após o último elemento inserido
}

// Enqueue em si:
func (queue *ArrayQueue) Enqueue(val int) {
	// Verificação de segurança: se a fila estiver cheia, não pode inserir
	if queue.size == len(queue.v) {
		queue.resize()
	}

	// Coloca o valor na posição atual do rear
	queue.v[queue.rear] = val

	// A mágica da fila circular:
	// Se o rear chegou no final do array, o módulo (%) faz ele voltar para o índice 0
	queue.rear = (queue.rear + 1) % len(queue.v)

	queue.size++
}

func (queue *ArrayQueue) Dequeue() (int, error) {
	// vrificação de segurança (Underflow)
	if queue.IsEmpty() {
		return 0, errors.New("a fila está vazia (underflow)")
	}

	// captura o valor que está na frente da fila
	val := queue.v[queue.front]

	// avança o ponteiro front de forma circular
	// Se front for 4 e o tamanho do array for 5, (4+1)%5 = 0
	queue.front = (queue.front + 1) % len(queue.v)

	// atualiza o contador de elementos
	queue.size--

	return val, nil
}

func (queue *ArrayQueue) Front() (int, error) {
	// sempre verificar se há algo para ler
	if queue.IsEmpty() {
		return 0, errors.New("a fila está vazia")
	}

	// apenas retornar o valor apontado pelo front
	// não altera o queue.front nem o queue.size
	return queue.v[queue.front], nil
}

// IsEmpty retorna true se a fila não tiver elementos
func (queue *ArrayQueue) IsEmpty() bool {
	return queue.size == 0
}

func (queue *ArrayQueue) Size() int {
	return queue.size
}
