package LeetCode

func sortColors(nums []int) {
	countZero := 0
	countOne := 0
	countTwo := 0

	for _, num := range nums {
		if num == 0 {
			countZero++
		}
		if num == 1 {
			countOne++
		}
		if num == 2 {
			countTwo++
		}
	}

	for i := 0; i < countZero; i++ {
		nums[i] = 0
	}
	for i := countZero; i < countZero+countOne; i++ {
		nums[i] = 1
	}
	for i := countZero + countOne; i < len(nums); i++ {
		nums[i] = 2
	}

}
