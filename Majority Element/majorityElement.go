package Majority_Element

func majorityElement(nums []int)int{
	majElem := len(nums) / 2
	majorityMap := make(map[int]int)
	for _, num := range nums{
		_, ok := majorityMap[num]
		if ok{
			majorityMap[num] += 1
		}else{
			majorityMap[num] = 1
		}
	}
	elements := []int{}
	for k,m := majorityMap{
		if m > majElem{
			elements = append(elements, k)
		}
	}
	return elements[0]
}
