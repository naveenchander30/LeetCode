package LeetCode

func majorityElement(nums []int) []int {
    cnt1,cnt2:=0,0
	var element1,element2 int

	for _,num:=range nums{
		if cnt1==0 && num!=element2{
			element1=num
			cnt1++
		}else if cnt2==0 && num!=element1{
			element2=num
			cnt2++
		}else if num==element1{
			cnt1++
		}else if num==element2{
			cnt2++
		}else{
			cnt1--
			cnt2--
		}
	}
	cnt1,cnt2=0,0
	for _,num:=range nums{
		switch num {
		case element1:
			cnt1++
		case element2:
			cnt2++
		}
	}
	result:=make([]int,0)
	n:=len(nums)
	if cnt1>n/3{
		result=append(result,element1)
	}
	if cnt2>n/3{
		result=append(result,element2)
	}
	return result
}