package dynamic

func ArrayStepper(nums []int, i int, memo map[int]int) bool {
	if i >= len(nums) {
		return true
	}

	if _, ok := memo[i]; ok {
		return false
	}

	for j := 1; j <= nums[i]; j++ {
		if ArrayStepper(nums, i + j, memo) {
			return true
		}
	}

	return false
} 
