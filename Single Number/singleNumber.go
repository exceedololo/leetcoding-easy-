package Single_Number

func singleNumber(nums []int) int {
	singleMap := make(map[int]int)
	for _, num := range nums {
		if _, ok := singleMap[num]; ok {
			singleMap[num]++
		} else {
			singleMap[num] = 1
		}
	}
	for key, value := range singleMap {
		if value == 1 {
			return key
		}
	}
	return 0
}
