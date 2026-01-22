package arrays

import (
	"testing"
)

func TestTwoSums(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		target   int
		expected []int
	}{
		{
			name:     "Test 1: Basic case",
			numbers:  []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1}, // 2 + 7 = 9
		},
		{
			name:     "Test 2: Non-adjacent elements",
			numbers:  []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2}, // 2 + 4 = 6
		},
		{
			name:     "Test 3: Duplicate values",
			numbers:  []int{3, 3},
			target:   6,
			expected: []int{0, 1}, // 3 + 3 = 6
		},
		{
			name:     "Test 4: Middle elements",
			numbers:  []int{1, 5, 3, 7, 2},
			target:   8,
			expected: []int{1, 2}, // 5 + 3 = 8
		},
		{
			name:     "Test 5: All negative numbers",
			numbers:  []int{-1, -2, -3, -4, -5},
			target:   -8,
			expected: []int{2, 4}, // -3 + (-5) = -8
		},
		{
			name:     "Test 6: Negative and positive",
			numbers:  []int{-3, 4, 3, 90},
			target:   0,
			expected: []int{0, 2}, // -3 + 3 = 0
		},
		{
			name:     "Test 7: Zeros in array",
			numbers:  []int{0, 4, 3, 0},
			target:   0,
			expected: []int{0, 3}, // 0 + 0 = 0
		},
		{
			name:     "Test 8: Larger array",
			numbers:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target:   19,
			expected: []int{8, 9}, // 9 + 10 = 19
		},
		{
			name:     "Test 9: Minimum size array",
			numbers:  []int{100, 200},
			target:   300,
			expected: []int{0, 1}, // 100 + 200 = 300
		},
		{
			name:     "Test 10: Non-adjacent in mixed array",
			numbers:  []int{5, 75, 25, 50, 10},
			target:   85,
			expected: []int{1, 4}, // 75 + 10 = 85
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TwoSums(tt.numbers, tt.target)

			// Check if result has 2 elements
			if len(result) != 2 {
				t.Errorf("Expected 2 indices, got %d", len(result))
				return
			}

			// Verify the sum equals target
			if tt.numbers[result[0]]+tt.numbers[result[1]] != tt.target {
				t.Errorf("Indices %v do not sum to target %d", result, tt.target)
				return
			}

			// Check if indices match expected (order matters based on implementation)
			if result[0] != tt.expected[0] || result[1] != tt.expected[1] {
				// Alternative: check if it's a valid solution even if order differs
				sum := tt.numbers[result[0]] + tt.numbers[result[1]]
				if sum == tt.target {
					t.Logf("Different valid solution found: got %v, expected %v", result, tt.expected)
				} else {
					t.Errorf("Expected %v, got %v", tt.expected, result)
				}
			}
		})
	}
}

func TestTwoSums_EdgeCases(t *testing.T) {
	t.Run("Empty array", func(t *testing.T) {
		result := TwoSums([]int{}, 5)
		if len(result) != 0 {
			t.Errorf("Expected empty slice for empty input, got %v", result)
		}
	})

	t.Run("Single element equals target", func(t *testing.T) {
		result := TwoSums([]int{5}, 5)
		expected := []int{0}
		if len(result) != 1 || result[0] != expected[0] {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Single element not equals target", func(t *testing.T) {
		result := TwoSums([]int{5}, 10)
		if len(result) != 0 {
			t.Errorf("Expected empty slice, got %v", result)
		}
	})

	t.Run("No valid pair exists", func(t *testing.T) {
		result := TwoSums([]int{1, 2, 3}, 100)
		if len(result) != 0 {
			t.Errorf("Expected empty slice when no solution, got %v", result)
		}
	})
}

// Benchmark test
func BenchmarkTwoSums(b *testing.B) {
	numbers := []int{2, 7, 11, 15, 20, 25, 30, 35, 40, 45}
	target := 65

	for i := 0; i < b.N; i++ {
		TwoSums(numbers, target)
	}
}
