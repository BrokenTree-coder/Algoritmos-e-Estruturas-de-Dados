package selectionSort

type ISelection interface {
	SelectionSort_OutOfPlace(v []int) []int
	SelectionSort_InPlace(v []int) []int
}
