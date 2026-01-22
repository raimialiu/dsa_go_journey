package arrays

import "slices"

/*
Given an array of integers nums and an integer target, return indixes of the two numbers such that they
add up to target. Each input has exactly one solution
*/
func TwoSums(numbers []int, target int) []int {
	answers := make([]int, 0)
	// given: { 1, 2, 3, 4, 5, 6 }, target = 9
	// 1 + 2, 1 + 3, 1 + 4, 1 + 5, 1+ 6
	// 2 + 1, 2 + 2, 2 + 3, 2 + 4, 2 + 5, 2 + 6
	// 3 + 1, 3 + 2
	if len(numbers) == 0 {
		return make([]int, 0)
	}

	if len(numbers) == 1 && numbers[0] == target {
		return []int{0}
	}

	i := 0
	poolbackup := numbers[:]
	for {
		if i == len(numbers) {
			return answers
		}
		front := numbers[i:]
		end := numbers[:i]
		pool := make([]int, 0)
		front = append(front, end...)
		pool = append(pool, front...)

		indexValue := numbers[i]
		for v := 0; v < len(pool); v++ {
			current := pool[v]
			sum := indexValue + current
			if sum == target {
				actualIndex := slices.Index(poolbackup, current)
				answers = append(answers, i, actualIndex)
				return answers
			}
		}

		i += 1
	}
}
