package main

import "fmt"

func main() {
	arr := [5]int{1, 3, 5, 7, 8}
	fmt.Println(sum(arr))
	fmt.Println(arr)
}

func twoSum(nums []int, target int) []int {
	ans := []int{0, 0}
	flag := false
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				ans[0] = i
				ans[1] = j
				flag = !flag
				break
			}
		}
		if flag {
			break
		}
	}
	return ans
}

func sum(nums [5]int) int {
	for index := range nums {
		if index == 0 {
			continue
		}
		nums[index] = nums[index] + nums[index-1]
	}
	return nums[len(nums)-1]
}
func rand10() int {
	ans := (rand7()-1)*7 + rand7()
	for ans > 40 {
		ans = (rand7()-1)*7 + rand7()
	}
	return ans%10 + 1
}

func rand7() int {
	return 1
}
