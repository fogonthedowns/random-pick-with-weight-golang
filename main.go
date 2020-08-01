package main

import "math/rand"

type Solution struct {
	prefixSums []int
	totalSum   int
}

// input [3, 14, 1, 7]
// sums [3, (3+14), (3+14+1), (3+14+1+7)]
// output [3, 17, 18, 25]
func Constructor(w []int) Solution {
	s := Solution{prefixSums: make([]int, len(w))}

	sum := 0
	for i := 0; i < len(w); i++ {
		sum += w[i]
		s.prefixSums[i] = sum
	}
	s.totalSum = sum
	return s
}

// the condition to apply binary search on a list is that the list should be sorted, either in ascending or descending order
// the input is sorted
func (s *Solution) PickIndex() int {
	// max of [3, 17, 18, 25]
	// 25
	max := s.prefixSums[len(s.prefixSums)-1]
	// rand  [0,25)
	target := rand.Intn(max)
	low := 0
	high := len(s.prefixSums) - 1

	// 0 < 3
	for low < high {
		var mid int
		mid = low + (high-low)/2
		if target >= s.prefixSums[mid] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
