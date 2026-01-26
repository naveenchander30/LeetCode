package LeetCode

func findMissingRepeatingNumbers(nums []int) []int {
	n := len(nums)
	sumN := n * (n + 1) / 2
	sumNSq := n * (n + 1) * (2*n + 1) / 6

	sumArr := 0
	sumArrSq := 0
	
	for _, num := range nums {
		sumArr += num
		sumArrSq += num * num
	}

	diff := sumN - sumArr                     
	diffSq := sumNSq - sumArrSq

	sumXY := diffSq / diff

	y := (diff + sumXY) / 2
	x := y - diff

	return []int{x, y}
}