package LeetCode

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	curSum := 0

	for _, num := range nums {
		curSum = max(curSum, 0)
		curSum += num
		maxSum = max(maxSum, curSum)
	}

	return maxSum
}
