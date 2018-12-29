package main

import "fmt"

/**
Given an array A of non-negative integers, half of the integers in A are odd, and half of the integers are even.

Sort the array so that whenever A[i] is odd, i is odd; and whenever A[i] is even, i is even.

You may return any answer array that satisfies this condition.



Example 1:

Input: [4,2,5,7]
Output: [4,5,2,7]
Explanation: [4,7,2,5], [2,5,4,7], [2,7,4,5] would also have been accepted.


Note:

2 <= A.length <= 20000
A.length % 2 == 0
0 <= A[i] <= 1000
 */

func main() {
	fmt.Println(sortArrayByParityII([]int{4,2,5,7}))
}

func sortArrayByParityII(A []int) []int {
	var (
		n    = len(A)
		i, j = 0, 1
	)
	for i<n && j<n {
		for i<n && A[i]&1 == i&1 {
			i+=2
		}
		for j<n && A[j]&1 == j&1 {
			j+=2
		}
		if j<n && i<n {
			A[i], A[j] = A[j], A[i]
			i+=2
			j+=2
		}
	}
	return A
}