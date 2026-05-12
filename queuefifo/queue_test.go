package queuefifo

import (
	"testing"
)

var size int

var queues []IQueue

func createQueues(size int) {
	linkedListQueue := &LinkedListQueue{}

	arrayQueue := &ArrayQueue{}
	(*arrayQueue).Init(size)

	queues = []IQueue{arrayQueue, linkedListQueue}
}

func deleteQueues() {
	queues = nil
}

func setupTest() func() {
	size = 10
	createQueues(size)

	return func() {
		deleteQueues()
	}
}

func TestEnqueue(t *testing.T) {
	defer setupTest()()
	for _, queue := range queues {
		for i := 0; i < 2*size; i++ {
			queue.Enqueue(i)
			if queue.Size() != i+1 {
				t.Errorf("%T size = %d, expected %d", queue, queue.Size(), i+1)
			}
		}
	}
}

func TestDequeue(t *testing.T) {
	defer setupTest()()
	for _, queue := range queues {
		for i := 0; i < size; i++ {
			queue.Enqueue(i)
		}
		for i := 0; i < size; i++ {
			val, err := queue.Dequeue()

			if err != nil {
				t.Errorf("%T unexpected error: %v", queue, err)
			}

			if val != i {
				t.Errorf("%T dequeued %d, expected %d", queue, val, i)
			}

			if queue.Size() != size-i-1 {
				t.Errorf("%T size = %d, expected %d", queue, queue.Size(), size-i-1)
			}
		}
	}
}

func TestDequeueEmptyQueue(t *testing.T) {
	defer setupTest()()
	for _, queue := range queues {
		_, err := queue.Dequeue()
		if err == nil {
			t.Errorf("%T expected error on dequeue from empty queue", queue)
		}
	}
}

func TestCircularEnqueueAfterDequeue(t *testing.T) {
	defer setupTest()()

	for _, queue := range queues {

		// enche parcialmente
		for i := 0; i < size; i++ {
			queue.Enqueue(i)
		}

		// remove alguns
		for i := 0; i < size-2; i++ {
			queue.Dequeue()
		}

		// agora deve sobrar: [size-2, size-1]

		// adiciona mais elementos (deve usar espaço circular)
		for i := size; i < size+6; i++ {
			queue.Enqueue(i)
		}

		// agora sequência esperada:
		// size-2, size-1, size, size+1, ..., size+5

		for i := size - 2; i < size+6; i++ {
			val, err := queue.Front()

			if err != nil {
				t.Errorf("%T unexpected error: %v", queue, err)
			}

			if val != i {
				t.Errorf("%T got %d, expected %d", queue, val, i)
			}

			queue.Dequeue()
		}

		if queue.Size() != 0 {
			t.Errorf("%T expected empty queue, got size %d", queue, queue.Size())
		}
	}
}

func TestFront(t *testing.T) {
	defer setupTest()()
	for _, queue := range queues {
		for i := 0; i < size; i++ {
			queue.Enqueue(i)
			val, err := queue.Front()
			if err != nil {
				t.Errorf("%T unexpected error: %v", queue, err)
			}
			if val != 0 {
				t.Errorf("%T front = %d, expected 0", queue, val)
			}

		}
	}
}

func TestFrontEmptyQueue(t *testing.T) {
	defer setupTest()()
	for _, queue := range queues {
		_, err := queue.Front()
		if err == nil {
			t.Errorf("%T expected error on front from empty queue", queue)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	defer setupTest()()
	for _, queue := range queues {
		queue.Enqueue(0)
		empty := queue.IsEmpty()
		if empty {
			t.Errorf("%T should not be empty", queue)
		}
	}
}

func TestIsEmptyOnEmptyQueue(t *testing.T) {
	defer setupTest()()
	for _, queue := range queues {
		if !queue.IsEmpty() {
			t.Errorf("%T should be empty", queue)
		}
	}
}

func TestSize(t *testing.T) {
	defer setupTest()()
	for _, queue := range queues {
		queue.Enqueue(0)
		if queue.Size() != 1 {
			t.Errorf("%T size = %d, expected 1", queue, queue.Size())
		}
	}
}

func TestSizeEmptyQueue(t *testing.T) {
	defer setupTest()()
	for _, queue := range queues {
		if queue.Size() != 0 {
			t.Errorf("%T size = %d, expected 0", queue, queue.Size())
		}
	}
}
