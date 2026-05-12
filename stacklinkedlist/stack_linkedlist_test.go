package stacklinkedlist

import "testing"

func TestStack(t *testing.T) {
	s := &Stack{}
	s.Init()

	// Teste 1: Push e Size
	s.Push(10)
	s.Push(20)
	s.Push(30)
	if s.Size() != 3 {
		t.Errorf("Esperava tamanho 3, obteve %d", s.Size())
	}

	// Teste 2: Peek (Deve ver o 30 sem remover)
	val, _ := s.Peek()
	if val != 30 {
		t.Errorf("Peek incorreto: esperava 30, obteve %d", val)
	}
	if s.Size() != 3 {
		t.Errorf("Peek alterou o tamanho da pilha!")
	}

	// Teste 3: Pop (Ordem LIFO)
	val, _ = s.Pop() // Sai o 30
	if val != 30 {
		t.Errorf("Pop incorreto: esperava 30, obteve %d", val)
	}

	val, _ = s.Pop() // Sai o 20
	if val != 20 {
		t.Errorf("Pop incorreto: esperava 20, obteve %d", val)
	}

	// Teste 4: Pilha Vazia
	s.Pop() // Esvazia o 10
	_, err := s.Pop()
	if err == nil {
		t.Error("Deveria retornar erro ao dar Pop em pilha vazia")
	}
}
