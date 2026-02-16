package LeetCode

func mergeSortB(nums []int, left, right int) int {
	count := 0
	if left < right {
		mid := left + (right-left)/2
		count += mergeSortB(nums, left, mid)
		count += mergeSortB(nums, mid+1, right)
		count += mergeAndCount(nums, left, mid, right)
	}
	return count
}

func mergeAndCount(nums []int, left, mid, right int) int {
	leftArrSize := mid - left + 1
	rightArrSize := right - mid

	leftArr := make([]int, leftArrSize)
	rightArr := make([]int, rightArrSize)

	for i := 0; i < leftArrSize; i++ {
		leftArr[i] = nums[left+i]
	}
	for j := 0; j < rightArrSize; j++ {
		rightArr[j] = nums[mid+1+j]
	}

	pairs := 0
	j := 0
	for i := 0; i < leftArrSize; i++ {
		for j < rightArrSize && leftArr[i] > 2*rightArr[j] {
			j++
		}
		pairs += j
	}


	i, j, k := 0, 0, left
	for i < leftArrSize && j < rightArrSize {
		if leftArr[i] <= rightArr[j] {
			nums[k] = leftArr[i]
			i++
		} else {
			nums[k] = rightArr[j]
			j++
		}
		k++
	}

	for i < leftArrSize {
		nums[k] = leftArr[i]
		i++
		k++
	}

	for j < rightArrSize {
		nums[k] = rightArr[j]
		j++
		k++
	}
	return pairs
}

func reversePairs(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return mergeSortB(nums, 0, len(nums)-1)
}