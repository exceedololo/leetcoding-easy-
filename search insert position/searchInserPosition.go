func searchInsert(nums []int, target int)int{
	for i := 0; i < len(nums) ; i++{
		if nums[len(nums)-1] < target{
			return len(nums)
		}else if nums[i] < target && nums[i+1] > target{
			return (i+1)
		}else if nums[i]==target{
			return i
		}
	}
	return 0
}
