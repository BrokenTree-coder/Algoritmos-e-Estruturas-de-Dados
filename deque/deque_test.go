package deque

import "testing"

var size int
var deques []IDeque

func createDeques(size int) {
	linked := &DoublyLinkedListDeque{}

	// Para ativar array também:
	array := &ArrayDeque{}
	array.Init(size)

	deques = []IDeque{array, linked}
}

func deleteDeques() {
	deques = nil
}

func setupTest() func() {
	size = 10
	createDeques(size)
	return func() {
		deleteDeques()
	}
}

func TestPushBack(t *testing.T) {
	defer setupTest()()

	for _, d := range deques {
		for i := 0; i < 2*size; i++ {
			d.PushBack(i)
			if d.Size() != i+1 {
				t.Errorf("%T size = %d, expected %d", d, d.Size(), i+1)
			}
		}
	}
}

func TestPushFront(t *testing.T) {
	defer setupTest()()

	for _, d := range deques {
		for i := 0; i < 2*size; i++ {
			d.PushFront(i)
			if d.Size() != i+1 {
				t.Errorf("%T size = %d, expected %d", d, d.Size(), i+1)
			}
		}
	}
}

func TestPopFront(t *testing.T) {
	defer setupTest()()

	for _, d := range deques {
		for i := 0; i < size; i++ {
			d.PushBack(i)
		}

		for i := 0; i < size; i++ {
			val, err := d.PopFront()
			if err != nil {
				t.Errorf("%T unexpected error: %v", d, err)
			}
			if val != i {
				t.Errorf("%T got %d, expected %d", d, val, i)
			}
			if d.Size() != size-i-1 {
				t.Errorf("%T size = %d, expected %d", d, d.Size(), size-i-1)
			}
		}
	}
}

func TestPopBack(t *testing.T) {
	defer setupTest()()

	for _, d := range deques {
		for i := 0; i < size; i++ {
			d.PushBack(i)
		}

		for i := 0; i < size; i++ {
			val, err := d.PopBack()
			if err != nil {
				t.Errorf("%T unexpected error: %v", d, err)
			}
			expected := size - i - 1
			if val != expected {
				t.Errorf("%T got %d, expected %d", d, val, expected)
			}
			if d.Size() != size-i-1 {
				t.Errorf("%T size = %d, expected %d", d, d.Size(), size-i-1)
			}
		}
	}
}

func TestPopEmpty(t *testing.T) {
	defer setupTest()()

	for _, d := range deques {
		if _, err := d.PopFront(); err == nil {
			t.Errorf("%T expected error on PopFront", d)
		}
		if _, err := d.PopBack(); err == nil {
			t.Errorf("%T expected error on PopBack", d)
		}
	}
}

func TestFront(t *testing.T) {
	defer setupTest()()

	for _, d := range deques {
		for i := 0; i < size; i++ {
			d.PushBack(i)
			val, err := d.Front()
			if err != nil {
				t.Errorf("%T unexpected error: %v", d, err)
			}
			if val != 0 {
				t.Errorf("%T front = %d, expected 0", d, val)
			}
		}
	}
}

func TestBack(t *testing.T) {
	defer setupTest()()

	for _, d := range deques {
		for i := 0; i < size; i++ {
			d.PushBack(i)
			val, err := d.Back()
			if err != nil {
				t.Errorf("%T unexpected error: %v", d, err)
			}
			if val != i {
				t.Errorf("%T back = %d, expected %d", d, val, i)
			}
		}
	}
}

func TestFrontBackEmpty(t *testing.T) {
	defer setupTest()()

	for _, d := range deques {
		if _, err := d.Front(); err == nil {
			t.Errorf("%T expected error on Front", d)
		}
		if _, err := d.Back(); err == nil {
			t.Errorf("%T expected error on Back", d)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	defer setupTest()()

	for _, d := range deques {
		if !d.IsEmpty() {
			t.Errorf("%T should be empty", d)
		}
		d.PushBack(1)
		if d.IsEmpty() {
			t.Errorf("%T should not be empty", d)
		}
	}
}

func TestSize(t *testing.T) {
	defer setupTest()()

	for _, d := range deques {
		if d.Size() != 0 {
			t.Errorf("%T size should be 0", d)
		}
		d.PushBack(1)
		if d.Size() != 1 {
			t.Errorf("%T size should be 1", d)
		}
	}
}

func TestSingleElement(t *testing.T) {
	defer setupTest()()

	for _, d := range deques {
		d.PushBack(42)

		val, err := d.PopFront()
		if err != nil {
			t.Errorf("%T unexpected error: %v", d, err)
		}
		if val != 42 {
			t.Errorf("expected 42, got %d", val)
		}

		if !d.IsEmpty() {
			t.Errorf("%T should be empty after removing single element", d)
		}
	}
}

func TestMixedOperations(t *testing.T) {
	defer setupTest()()

	for _, d := range deques {
		d.PushBack(1)
		d.PushFront(2)
		d.PushBack(3)

		val, err := d.PopFront() // 2
		if err != nil {
			t.Errorf("%T unexpected error: %v", d, err)
		}
		if val != 2 {
			t.Errorf("expected 2, got %d", val)
		}

		val, err = d.PopBack() // 3
		if err != nil {
			t.Errorf("%T unexpected error: %v", d, err)
		}
		if val != 3 {
			t.Errorf("expected 3, got %d", val)
		}

		val, err = d.PopFront() // 1
		if err != nil {
			t.Errorf("%T unexpected error: %v", d, err)
		}
		if val != 1 {
			t.Errorf("expected 1, got %d", val)
		}
	}
}

func TestPushCircularRight(t *testing.T) {
	defer setupTest()()

	for _, d := range deques {
		for i := 0; i < size; i++ {
			d.PushBack(i)
		}
		for i := 0; i < size-2; i++ {
			d.PopFront()
		}

		for i := 10; i < 16; i++ {
			d.PushBack(i)
		}

		for i := 8; i < 16; i++ {
			val, err := d.PopFront()
			if err != nil {
				t.Errorf("%T unexpected error: %v", d, err)
			}
			if val != i {
				t.Errorf("expected %d, got %d", i, val)
			}
		}

		if d.Size() != 0 {
			t.Errorf("%T should be empty", d)
		}
	}
}

func TestPushCircularLeft(t *testing.T) {
	defer setupTest()()

	for _, d := range deques {
		for i := 0; i < size; i++ {
			d.PushBack(i)
		}
		for i := 0; i < 2; i++ {
			d.PopBack()
		}

		for i := -1; i > -3; i-- {
			d.PushFront(i)
		}

		for i := -2; i < 8; i++ {
			val, err := d.PopFront()
			if err != nil {
				t.Errorf("%T unexpected error: %v", d, err)
			}
			if val != i {
				t.Errorf("expected %d, got %d", i, val)
			}
		}

		if d.Size() != 0 {
			t.Errorf("%T should be empty", d)
		}
	}
}
