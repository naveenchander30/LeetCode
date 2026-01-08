package LeetCode

func nextPermutation(nums []int) {
	cutIndex := -1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			cutIndex = i
			break
		}
	}

	if cutIndex != -1 {
		for i := len(nums) - 1; i >= 0; i-- {
			if nums[i] > nums[cutIndex] {
				temp := nums[i]
				nums[i] = nums[cutIndex]
				nums[cutIndex] = temp
				break
			}
		}
	}

	start := cutIndex + 1
	end := len(nums) - 1

	for start < end {
		temp := nums[start]
		nums[start] = nums[end]
		nums[end] = temp
		start++
		end--
	}
}
