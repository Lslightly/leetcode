package main

// TODO: TLE
func longestConsecutive(nums []int) int {
	input := make(map[int]bool)
	ks := make(map[int]bool)
	for _, num := range nums {
		input[num] = true
		if _, ok := input[num-1]; !ok { // only add num where num-1 does not exist in nums
			ks[num] = true
		}
	}
	maxPossibleLen := len(input)
	for l := 2; l <= maxPossibleLen; l++ {
		dels := make([]int, 0, len(input)/2)
		exist := false
		for k := range ks {
			if _, ok := input[k+l-1]; !ok {
				dels = append(dels, k)
			} else {
				exist = true
			}
		}
		if !exist { // all buckets[k+l-1] do not exist
			return l - 1
		}
		for _, k := range dels {
			delete(ks, k)
		}
	}
	return maxPossibleLen
}
