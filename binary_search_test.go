package binarysearch

import "testing"

func TestSortIntSlice(t *testing.T) {
	slice := []int{8, 4, 6, 2, 9, 1, 3, 5}
	SortIntSlice(slice)
	prevInt := -1
	for _, v := range slice {
		if v < prevInt {
			t.Error("Sort to increasing slice failed")
		}
		prevInt = v
	}
}

func TestBinarySearchFind(t *testing.T) {
	slice := []int{1, 2, 3, 4, 6, 7, 8, 9, 10}
	if found := BinarySearchFind(slice, 3); !found {
		t.Errorf("target %d not found in slice %v", 3, slice)
	}

	if found := BinarySearchFind(slice, 5); found {
		t.Errorf("target %d found in slice %v", 5, slice)
	}

	if found := BinarySearchFind(slice, 11); found {
		t.Errorf("target %d found in slice %v", 11, slice)
	}

	if found := BinarySearchFind(slice, 0); found {
		t.Errorf("target %d found in slice %v", 0, slice)
	}
}

func TestSliceInsert(t *testing.T) {
	slice := []int{1, 2, 3, 4, 6, 7, 8, 9, 10}
	resultSlice, _ := SliceInsert(slice, 4, 5)
	checkSlice(resultSlice, t)
	resultSlice, _ = SliceInsert(resultSlice, 10, 11)
	checkSlice(resultSlice, t)
}

func checkSlice(slice []int, t *testing.T) {
	for i := 1; i <= len(slice); i++ {
		if i != slice[i-1] {
			t.Error("insertion failes")
		}
	}
}

func TestBinarySearchInsert(t *testing.T) {
	slice := []int{1, 2, 3, 4, 6, 7, 8, 9, 10}
	resultSlice := BinarySearchInsert(slice, 5)
	checkSlice(resultSlice, t)

	resultSlice = BinarySearchInsert(resultSlice, 11)
	checkSlice(resultSlice, t)
}
