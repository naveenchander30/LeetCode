package LeetCode

func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix) - 1
	for i := 0; i < len(matrix)-1; i++ {
		if target >= matrix[i][0] && target < matrix[i+1][0] {
			row = i
			break
		}
	}
	low:=0
	high:=len(matrix[0])-1
	for low<=high{
		mid:=low+(high-low)/2
		if matrix[row][mid]==target{
			return true
		}else if matrix[row][mid]<target{
			low=mid+1
		}else{
			high=mid-1
		}
	}
	return false
}