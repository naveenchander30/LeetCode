package LeetCode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	sortedArray:= make([]int, 0, totalLen)
	i, j, k := 0, 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			sortedArray = append(sortedArray, nums1[i])
			i++
		} else {
			sortedArray = append(sortedArray, nums2[j])
			j++
		}
		k++
	}

	for i < len(nums1) {
		sortedArray = append(sortedArray, nums1[i])
		i++
		k++
	}

	for j < len(nums2) {
		sortedArray = append(sortedArray, nums2[j])
		j++
		k++
	}

	if totalLen%2 == 0 {
		mid1 := totalLen / 2
		mid2 := mid1 - 1
		return float64(sortedArray[mid1]+sortedArray[mid2]) / 2.0
	} else {
		mid := totalLen / 2
		return float64(sortedArray[mid])
	}
}