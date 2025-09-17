/*
 * Copyright 2014-2024 Li Kexian
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Go module for domain and ip whois information query
 * https://www.likexian.com/
 */

package utils

import (
	"errors"
	"sort"
)

var (
	// ErrEmptySlice is returned when an empty slice is passed to sorting functions
	ErrEmptySlice = errors.New("utils: slice is empty")
)

// SortDirection represents the direction of sorting
type SortDirection int

const (
	// Ascending sorts in ascending order
	Ascending SortDirection = iota
	// Descending sorts in descending order
	Descending
)

// BubbleSort implements bubble sort algorithm for integer slices.
// It sorts the slice in-place and returns an error if the slice is empty.
//
// Time complexity: O(n²)
// Space complexity: O(1)
//
// Example usage:
//
//	data := []int{64, 34, 25, 12, 22, 11, 90}
//	err := BubbleSort(data, Ascending)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(data) // Output: [11 12 22 25 34 64 90]
func BubbleSort(data []int, direction SortDirection) error {
	if len(data) == 0 {
		return ErrEmptySlice
	}

	n := len(data)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			shouldSwap := false
			if direction == Ascending {
				shouldSwap = data[j] > data[j+1]
			} else {
				shouldSwap = data[j] < data[j+1]
			}

			if shouldSwap {
				data[j], data[j+1] = data[j+1], data[j]
				swapped = true
			}
		}
		// If no swapping occurred, the array is already sorted
		if !swapped {
			break
		}
	}

	return nil
}

// QuickSort implements quick sort algorithm for integer slices.
// It sorts the slice in-place and returns an error if the slice is empty.
//
// Time complexity: O(n log n) average, O(n²) worst case
// Space complexity: O(log n) average due to recursion
//
// Example usage:
//
//	data := []int{64, 34, 25, 12, 22, 11, 90}
//	err := QuickSort(data, Ascending)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(data) // Output: [11 12 22 25 34 64 90]
func QuickSort(data []int, direction SortDirection) error {
	if len(data) == 0 {
		return ErrEmptySlice
	}

	quickSortRecursive(data, 0, len(data)-1, direction)
	return nil
}

// quickSortRecursive is the recursive helper function for QuickSort
func quickSortRecursive(data []int, low, high int, direction SortDirection) {
	if low < high {
		// Partition the array and get the pivot index
		pivotIndex := partition(data, low, high, direction)

		// Recursively sort elements before and after partition
		quickSortRecursive(data, low, pivotIndex-1, direction)
		quickSortRecursive(data, pivotIndex+1, high, direction)
	}
}

// partition rearranges the array so that elements smaller than the pivot
// are on the left and elements greater than the pivot are on the right
func partition(data []int, low, high int, direction SortDirection) int {
	// Choose the rightmost element as pivot
	pivot := data[high]
	i := low - 1 // Index of smaller element

	for j := low; j < high; j++ {
		shouldSwap := false
		if direction == Ascending {
			shouldSwap = data[j] < pivot
		} else {
			shouldSwap = data[j] > pivot
		}

		if shouldSwap {
			i++
			data[i], data[j] = data[j], data[i]
		}
	}

	// Place pivot in the correct position
	data[i+1], data[high] = data[high], data[i+1]
	return i + 1
}

// BuiltinSort demonstrates the usage of Go's built-in sort package.
// It provides examples for sorting integer slices in both ascending and descending order.
//
// Example usage:
//
//	data := []int{64, 34, 25, 12, 22, 11, 90}
//	err := BuiltinSort(data, Ascending)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(data) // Output: [11 12 22 25 34 64 90]
func BuiltinSort(data []int, direction SortDirection) error {
	if len(data) == 0 {
		return ErrEmptySlice
	}

	if direction == Ascending {
		// Sort in ascending order using sort.Ints
		sort.Ints(data)
	} else {
		// Sort in descending order using sort.Slice
		sort.Slice(data, func(i, j int) bool {
			return data[i] > data[j]
		})
	}

	return nil
}

// BuiltinSortStrings demonstrates sorting string slices using Go's built-in sort package.
// It returns an error if the slice is empty.
//
// Example usage:
//
//	data := []string{"banana", "apple", "cherry", "date"}
//	err := BuiltinSortStrings(data, Ascending)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(data) // Output: [apple banana cherry date]
func BuiltinSortStrings(data []string, direction SortDirection) error {
	if len(data) == 0 {
		return ErrEmptySlice
	}

	if direction == Ascending {
		sort.Strings(data)
	} else {
		sort.Slice(data, func(i, j int) bool {
			return data[i] > data[j]
		})
	}

	return nil
}

// IsSorted checks if an integer slice is sorted in the specified direction.
// It returns true if the slice is sorted, false otherwise.
// An empty slice is considered sorted.
//
// Example usage:
//
//	data := []int{1, 2, 3, 4, 5}
//	fmt.Println(IsSorted(data, Ascending))  // Output: true
//	fmt.Println(IsSorted(data, Descending)) // Output: false
func IsSorted(data []int, direction SortDirection) bool {
	if len(data) <= 1 {
		return true
	}

	for i := 1; i < len(data); i++ {
		if direction == Ascending {
			if data[i-1] > data[i] {
				return false
			}
		} else {
			if data[i-1] < data[i] {
				return false
			}
		}
	}

	return true
}
