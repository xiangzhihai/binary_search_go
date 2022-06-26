package binarysearch

import (
	"errors"
	"sort"
)

func SortIntSlice(slice []int) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}

func BinarySearchFind(sortedSlice []int, target int) bool {
	// Exclude larger and smaller
	l := 0
	r := len(sortedSlice) - 1
	for l < r {
		m := (l + r) >> 1
		numToCompare := sortedSlice[m]
		if target == numToCompare {
			return true
		} else if target > numToCompare {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return false
}

func BinarySearchInsert(sortedSlice []int, target int) []int {
	// Find the index of the value that is just larger to
	// target

	if len(sortedSlice) == 0 {
		sortedSlice[0] = target
	} else {
		l := 0
		r := len(sortedSlice) - 1
		for l < r {
			m := (l + r) >> 1
			numToCompare := sortedSlice[m]
			if target < numToCompare {
				r = m
			} else {
				l = m + 1
			}
		}
		if sortedSlice[l] <= target {
			l++
		}

		sortedSlice, _ = SliceInsert(sortedSlice, l, target)
	}

	return sortedSlice
}

func SliceInsert(slice []int, index, value int) ([]int, error) {
	if index < 0 || index > len(slice) {
		return nil, errors.New("index cannot be less than 0")
	}

	if index == len(slice) {
		return append(slice, value), nil
	}

	slice = append(slice[:index+1], slice[index:]...)
	slice[index] = value

	return slice, nil
}
