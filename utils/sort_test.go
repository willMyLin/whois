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
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		direction SortDirection
		expected  []int
		wantErr   bool
		err       error
	}{
		{
			name:      "ascending order",
			input:     []int{64, 34, 25, 12, 22, 11, 90},
			direction: Ascending,
			expected:  []int{11, 12, 22, 25, 34, 64, 90},
			wantErr:   false,
		},
		{
			name:      "descending order",
			input:     []int{64, 34, 25, 12, 22, 11, 90},
			direction: Descending,
			expected:  []int{90, 64, 34, 25, 22, 12, 11},
			wantErr:   false,
		},
		{
			name:      "already sorted ascending",
			input:     []int{1, 2, 3, 4, 5},
			direction: Ascending,
			expected:  []int{1, 2, 3, 4, 5},
			wantErr:   false,
		},
		{
			name:      "already sorted descending",
			input:     []int{5, 4, 3, 2, 1},
			direction: Descending,
			expected:  []int{5, 4, 3, 2, 1},
			wantErr:   false,
		},
		{
			name:      "single element",
			input:     []int{42},
			direction: Ascending,
			expected:  []int{42},
			wantErr:   false,
		},
		{
			name:      "empty slice",
			input:     []int{},
			direction: Ascending,
			expected:  []int{},
			wantErr:   true,
			err:       ErrEmptySlice,
		},
		{
			name:      "duplicates ascending",
			input:     []int{3, 1, 4, 1, 5, 9, 2, 6, 5},
			direction: Ascending,
			expected:  []int{1, 1, 2, 3, 4, 5, 5, 6, 9},
			wantErr:   false,
		},
		{
			name:      "duplicates descending",
			input:     []int{3, 1, 4, 1, 5, 9, 2, 6, 5},
			direction: Descending,
			expected:  []int{9, 6, 5, 5, 4, 3, 2, 1, 1},
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of input to avoid modifying the test data
			data := make([]int, len(tt.input))
			copy(data, tt.input)

			err := BubbleSort(data, tt.direction)

			if tt.wantErr {
				if err == nil {
					t.Errorf("BubbleSort() expected error, got nil")
				}
				if tt.err != nil && !errors.Is(err, tt.err) {
					t.Errorf("BubbleSort() error = %v, want %v", err, tt.err)
				}
				return
			}

			if err != nil {
				t.Errorf("BubbleSort() unexpected error = %v", err)
				return
			}

			if !reflect.DeepEqual(data, tt.expected) {
				t.Errorf("BubbleSort() = %v, want %v", data, tt.expected)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		direction SortDirection
		expected  []int
		wantErr   bool
		err       error
	}{
		{
			name:      "ascending order",
			input:     []int{64, 34, 25, 12, 22, 11, 90},
			direction: Ascending,
			expected:  []int{11, 12, 22, 25, 34, 64, 90},
			wantErr:   false,
		},
		{
			name:      "descending order",
			input:     []int{64, 34, 25, 12, 22, 11, 90},
			direction: Descending,
			expected:  []int{90, 64, 34, 25, 22, 12, 11},
			wantErr:   false,
		},
		{
			name:      "already sorted ascending",
			input:     []int{1, 2, 3, 4, 5},
			direction: Ascending,
			expected:  []int{1, 2, 3, 4, 5},
			wantErr:   false,
		},
		{
			name:      "already sorted descending",
			input:     []int{5, 4, 3, 2, 1},
			direction: Descending,
			expected:  []int{5, 4, 3, 2, 1},
			wantErr:   false,
		},
		{
			name:      "single element",
			input:     []int{42},
			direction: Ascending,
			expected:  []int{42},
			wantErr:   false,
		},
		{
			name:      "empty slice",
			input:     []int{},
			direction: Ascending,
			expected:  []int{},
			wantErr:   true,
			err:       ErrEmptySlice,
		},
		{
			name:      "duplicates ascending",
			input:     []int{3, 1, 4, 1, 5, 9, 2, 6, 5},
			direction: Ascending,
			expected:  []int{1, 1, 2, 3, 4, 5, 5, 6, 9},
			wantErr:   false,
		},
		{
			name:      "duplicates descending",
			input:     []int{3, 1, 4, 1, 5, 9, 2, 6, 5},
			direction: Descending,
			expected:  []int{9, 6, 5, 5, 4, 3, 2, 1, 1},
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of input to avoid modifying the test data
			data := make([]int, len(tt.input))
			copy(data, tt.input)

			err := QuickSort(data, tt.direction)

			if tt.wantErr {
				if err == nil {
					t.Errorf("QuickSort() expected error, got nil")
				}
				if tt.err != nil && !errors.Is(err, tt.err) {
					t.Errorf("QuickSort() error = %v, want %v", err, tt.err)
				}
				return
			}

			if err != nil {
				t.Errorf("QuickSort() unexpected error = %v", err)
				return
			}

			if !reflect.DeepEqual(data, tt.expected) {
				t.Errorf("QuickSort() = %v, want %v", data, tt.expected)
			}
		})
	}
}

func TestBuiltinSort(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		direction SortDirection
		expected  []int
		wantErr   bool
		err       error
	}{
		{
			name:      "ascending order",
			input:     []int{64, 34, 25, 12, 22, 11, 90},
			direction: Ascending,
			expected:  []int{11, 12, 22, 25, 34, 64, 90},
			wantErr:   false,
		},
		{
			name:      "descending order",
			input:     []int{64, 34, 25, 12, 22, 11, 90},
			direction: Descending,
			expected:  []int{90, 64, 34, 25, 22, 12, 11},
			wantErr:   false,
		},
		{
			name:      "empty slice",
			input:     []int{},
			direction: Ascending,
			expected:  []int{},
			wantErr:   true,
			err:       ErrEmptySlice,
		},
		{
			name:      "single element",
			input:     []int{42},
			direction: Ascending,
			expected:  []int{42},
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of input to avoid modifying the test data
			data := make([]int, len(tt.input))
			copy(data, tt.input)

			err := BuiltinSort(data, tt.direction)

			if tt.wantErr {
				if err == nil {
					t.Errorf("BuiltinSort() expected error, got nil")
				}
				if tt.err != nil && !errors.Is(err, tt.err) {
					t.Errorf("BuiltinSort() error = %v, want %v", err, tt.err)
				}
				return
			}

			if err != nil {
				t.Errorf("BuiltinSort() unexpected error = %v", err)
				return
			}

			if !reflect.DeepEqual(data, tt.expected) {
				t.Errorf("BuiltinSort() = %v, want %v", data, tt.expected)
			}
		})
	}
}

func TestBuiltinSortStrings(t *testing.T) {
	tests := []struct {
		name      string
		input     []string
		direction SortDirection
		expected  []string
		wantErr   bool
		err       error
	}{
		{
			name:      "ascending order",
			input:     []string{"banana", "apple", "cherry", "date"},
			direction: Ascending,
			expected:  []string{"apple", "banana", "cherry", "date"},
			wantErr:   false,
		},
		{
			name:      "descending order",
			input:     []string{"banana", "apple", "cherry", "date"},
			direction: Descending,
			expected:  []string{"date", "cherry", "banana", "apple"},
			wantErr:   false,
		},
		{
			name:      "empty slice",
			input:     []string{},
			direction: Ascending,
			expected:  []string{},
			wantErr:   true,
			err:       ErrEmptySlice,
		},
		{
			name:      "single element",
			input:     []string{"hello"},
			direction: Ascending,
			expected:  []string{"hello"},
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of input to avoid modifying the test data
			data := make([]string, len(tt.input))
			copy(data, tt.input)

			err := BuiltinSortStrings(data, tt.direction)

			if tt.wantErr {
				if err == nil {
					t.Errorf("BuiltinSortStrings() expected error, got nil")
				}
				if tt.err != nil && !errors.Is(err, tt.err) {
					t.Errorf("BuiltinSortStrings() error = %v, want %v", err, tt.err)
				}
				return
			}

			if err != nil {
				t.Errorf("BuiltinSortStrings() unexpected error = %v", err)
				return
			}

			if !reflect.DeepEqual(data, tt.expected) {
				t.Errorf("BuiltinSortStrings() = %v, want %v", data, tt.expected)
			}
		})
	}
}

func TestIsSorted(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		direction SortDirection
		expected  bool
	}{
		{
			name:      "sorted ascending",
			input:     []int{1, 2, 3, 4, 5},
			direction: Ascending,
			expected:  true,
		},
		{
			name:      "sorted descending",
			input:     []int{5, 4, 3, 2, 1},
			direction: Descending,
			expected:  true,
		},
		{
			name:      "not sorted ascending",
			input:     []int{1, 3, 2, 4, 5},
			direction: Ascending,
			expected:  false,
		},
		{
			name:      "not sorted descending",
			input:     []int{5, 3, 4, 2, 1},
			direction: Descending,
			expected:  false,
		},
		{
			name:      "empty slice",
			input:     []int{},
			direction: Ascending,
			expected:  true,
		},
		{
			name:      "single element",
			input:     []int{42},
			direction: Ascending,
			expected:  true,
		},
		{
			name:      "duplicates sorted ascending",
			input:     []int{1, 1, 2, 3, 3, 4},
			direction: Ascending,
			expected:  true,
		},
		{
			name:      "duplicates sorted descending",
			input:     []int{4, 3, 3, 2, 1, 1},
			direction: Descending,
			expected:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsSorted(tt.input, tt.direction)
			if result != tt.expected {
				t.Errorf("IsSorted() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// Benchmark tests
func BenchmarkBubbleSort(b *testing.B) {
	data := []int{64, 34, 25, 12, 22, 11, 90, 88, 76, 50, 43}
	for i := 0; i < b.N; i++ {
		testData := make([]int, len(data))
		copy(testData, data)
		BubbleSort(testData, Ascending)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	data := []int{64, 34, 25, 12, 22, 11, 90, 88, 76, 50, 43}
	for i := 0; i < b.N; i++ {
		testData := make([]int, len(data))
		copy(testData, data)
		QuickSort(testData, Ascending)
	}
}

func BenchmarkBuiltinSort(b *testing.B) {
	data := []int{64, 34, 25, 12, 22, 11, 90, 88, 76, 50, 43}
	for i := 0; i < b.N; i++ {
		testData := make([]int, len(data))
		copy(testData, data)
		BuiltinSort(testData, Ascending)
	}
}
