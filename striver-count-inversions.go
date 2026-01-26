package LeetCode

func mergeSort(nums []int, left, right int) int {
	count := 0
	if left < right {
		mid := left + (right-left)/2
		count += mergeSort(nums, left, mid)
		count += mergeSort(nums, mid+1, right)
		count += mergeI(nums, left, mid, right)
	}
	return count
}

func mergeI(nums []int, left, mid, right int) int {
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
	
	i, j, k := 0, 0, left
	inversionCount := 0
	for i < leftArrSize && j < rightArrSize {
		if leftArr[i] <= rightArr[j] {
			nums[k] = leftArr[i]
			i++
		} else {
			nums[k] = rightArr[j]
			inversionCount += (leftArrSize - i)
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
	return inversionCount
}

func numberOfInversions(nums []int) int64 {
	if len(nums)<2{
		return 0
	}
	count := mergeSort(nums, 0, len(nums)-1)
	return int64(count)
}
