package LeetCode
import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	newIntervals := make([][]int, 0)
	currentInterval := intervals[0]
	for i := 0; i < len(intervals)-1; i++ {
		if currentInterval[1] < intervals[i+1][0] {
			newIntervals = append(newIntervals, currentInterval)
			currentInterval = intervals[i+1]
		} else {
			currentInterval[1] = max(currentInterval[1], intervals[i+1][1])
		}
	}
	newIntervals = append(newIntervals, currentInterval)
	return newIntervals
}
