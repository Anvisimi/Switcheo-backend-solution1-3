package main

import (
	"fmt"
)

// Efficiency analysis:
// sum_to_n_a: O(1) time, O(1) space - most efficient (uses mathematical formula)
// sum_to_n_b: O(n) time, O(1) space - linear time but constant space
// sum_to_n_c: O(n) time, O(n) space - linear time plus linear stack space due to recursion

func sum_to_n_a(n int) int {

	// formula
	return n * (n + 1) / 2
}

func sum_to_n_b(n int) int {
	// iterative
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	return total
}

func sum_to_n_c(n int) int {
	// recursive
	if n == 1 {
		return 1
	}
	return n + sum_to_n_c(n-1)
}

func main() {
	fmt.Println(sum_to_n_a(150))
	fmt.Println(sum_to_n_b(160))
	fmt.Println(sum_to_n_c(170))
}
