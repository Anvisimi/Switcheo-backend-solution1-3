# Question 4

Provide 3 unique implementations of the following function in Golang or Rust.

<aside>
💡 Bonus points if you comment on the efficiency or complexity of each solution.

</aside>

**Input**: `n` - any integer

*Assuming this input will always produce a result lesser than `Number.MAX_SAFE_INTEGER`*.

**Output**: `return` - summation to `n`, i.e. `sum_to_n(5) === 1 + 2 + 3 + 4 + 5 === 15`.

```go
func sum_to_n_a(n int) int {
	// your code here
}

func sum_to_n_b(n int) int {
	// your code here
}

func sum_to_n_c(n int) int {
	// your code here
}
```

## Solution 
to run the code in sum_to_variations.go file navigate to the problem-4 directory and run the command `go run sum_to_variations.go`

## Efficiency Analysis
// sum_to_n_a: O(1) time, O(1) space - most efficient (uses mathematical formula)
// sum_to_n_b: O(n) time, O(1) space - linear time but constant space
// sum_to_n_c: O(n) time, O(n) space - linear time plus linear stack space due to recursion