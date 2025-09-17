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

// Package main demonstrates the usage of sorting functions in the utils package.
package main

import (
	"fmt"
	"log"

	"github.com/likexian/whois/utils"
)

func main() {
	// Sample data to sort
	originalData := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Original data: %v\n\n", originalData)

	// Demonstrate BubbleSort
	fmt.Println("=== Bubble Sort Demo ===")
	bubbleData := make([]int, len(originalData))
	copy(bubbleData, originalData)

	fmt.Printf("Before BubbleSort (Ascending): %v\n", bubbleData)
	err := utils.BubbleSort(bubbleData, utils.Ascending)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("After BubbleSort (Ascending):  %v\n", bubbleData)
	fmt.Printf("Is sorted ascending: %v\n", utils.IsSorted(bubbleData, utils.Ascending))

	// Reset and sort descending
	copy(bubbleData, originalData)
	fmt.Printf("Before BubbleSort (Descending): %v\n", bubbleData)
	err = utils.BubbleSort(bubbleData, utils.Descending)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("After BubbleSort (Descending):  %v\n", bubbleData)
	fmt.Printf("Is sorted descending: %v\n\n", utils.IsSorted(bubbleData, utils.Descending))

	// Demonstrate QuickSort
	fmt.Println("=== Quick Sort Demo ===")
	quickData := make([]int, len(originalData))
	copy(quickData, originalData)

	fmt.Printf("Before QuickSort (Ascending): %v\n", quickData)
	err = utils.QuickSort(quickData, utils.Ascending)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("After QuickSort (Ascending):  %v\n", quickData)

	// Reset and sort descending
	copy(quickData, originalData)
	fmt.Printf("Before QuickSort (Descending): %v\n", quickData)
	err = utils.QuickSort(quickData, utils.Descending)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("After QuickSort (Descending):  %v\n\n", quickData)

	// Demonstrate BuiltinSort
	fmt.Println("=== Built-in Sort Demo ===")
	builtinData := make([]int, len(originalData))
	copy(builtinData, originalData)

	fmt.Printf("Before BuiltinSort (Ascending): %v\n", builtinData)
	err = utils.BuiltinSort(builtinData, utils.Ascending)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("After BuiltinSort (Ascending):  %v\n", builtinData)

	// Reset and sort descending
	copy(builtinData, originalData)
	fmt.Printf("Before BuiltinSort (Descending): %v\n", builtinData)
	err = utils.BuiltinSort(builtinData, utils.Descending)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("After BuiltinSort (Descending):  %v\n\n", builtinData)

	// Demonstrate BuiltinSortStrings
	fmt.Println("=== Built-in String Sort Demo ===")
	stringData := []string{"banana", "apple", "cherry", "date"}
	fmt.Printf("Original strings: %v\n", stringData)

	stringDataCopy := make([]string, len(stringData))
	copy(stringDataCopy, stringData)
	fmt.Printf("Before BuiltinSortStrings (Ascending): %v\n", stringDataCopy)
	err = utils.BuiltinSortStrings(stringDataCopy, utils.Ascending)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("After BuiltinSortStrings (Ascending):  %v\n", stringDataCopy)

	copy(stringDataCopy, stringData)
	fmt.Printf("Before BuiltinSortStrings (Descending): %v\n", stringDataCopy)
	err = utils.BuiltinSortStrings(stringDataCopy, utils.Descending)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("After BuiltinSortStrings (Descending):  %v\n\n", stringDataCopy)

	// Demonstrate error handling
	fmt.Println("=== Error Handling Demo ===")
	emptyData := []int{}
	fmt.Printf("Trying to sort empty slice: %v\n", emptyData)
	err = utils.BubbleSort(emptyData, utils.Ascending)
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}
}