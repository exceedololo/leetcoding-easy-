func twoSum(nums []int, target int) []int {
	mappa := make(map[int]int)
	for i, num := range nums{
		mappa[num] = i
	}
	for i, num := range nums{
		diff := target - num
		if j, ok := mappa[diff]; ok && j != i{
			return []int{j, i}
		}
	}
	return []int{}
}