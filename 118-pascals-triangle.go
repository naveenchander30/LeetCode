package LeetCode

func generate(numRows int) [][]int {
	var res [][]int
	for i := 0; i < numRows; i++ {
		temp := make([]int, 0, i+1)
		temp = append(temp, 1)
		t1 := 0
		t2 := 1
		for len(temp) < cap(temp)-1 {
			temp = append(temp, res[i-1][t1]+res[i-1][t2])
			t1++
			t2++
		}
		if i != 0 {
			temp = append(temp, 1)
		}
		res = append(res, temp)

	}
	return res
}
