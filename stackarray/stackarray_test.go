package stackarray

import (
	"testing"
)

func TestStackArray(t *testing.T) {
	capacity := 3
	s := &StackArray{}
	s.Init(capacity)

	if !s.IsEmpty() {
		t.Error("Pilha deveria estar vazia ao iniciar")
	}

	// 1. Enchendo a capacidade inicial
	s.Push(10)
	s.Push(20)
	s.Push(30)

	// 2. Testando o Crescimento Dinâmico (reSize)
	// Agora o Push(40) NÃO deve retornar erro e a capacidade deve dobrar
	err := s.Push(40)
	if err != nil {
		t.Errorf("Erro inesperado: a pilha deveria ter crescido, mas deu erro: %v", err)
	}

	if s.capacity != 6 {
		t.Errorf("Capacidade incorreta após resize: esperado 6, obteve %d", s.capacity)
	}

	// 3. Testando Peek (O topo agora é o 40!)
	val, err := s.Peek()
	if err != nil {
		t.Errorf("Erro inesperado no Peek: %v", err)
	}
	if val != 40 {
		t.Errorf("Peek incorreto: esperado 40, obteve %d", val)
	}

	// 4. Testando Pop (LIFO: sai o 40 primeiro)
	val, _ = s.Pop()
	if val != 40 {
		t.Errorf("Pop incorreto: esperado 40, obteve %d", val)
	}

	// 5. Esvaziando a pilha (Restam 30, 20, 10)
	s.Pop() // Tira o 30
	s.Pop() // Tira o 20
	s.Pop() // Tira o 10

	if !s.IsEmpty() {
		t.Error("Pilha deveria estar vazia após remover os 4 elementos")
	}

	// 6. Testando Underflow
	_, err = s.Pop()
	if err == nil {
		t.Error("Deveria retornar erro ao dar Pop em pilha vazia")
	}
}
