package Contains_Duplicates

func containsDuplicates(nums []int) bool {
	numsSlice := nums[:]
	freq := make(map[int]int)
	for _, i := range numsSlice {
		freq[i]++
	}
	for _, k := range freq {
		if k > 1 {
			return true
		}
	}
	return false
}
