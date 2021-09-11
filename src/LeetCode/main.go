package main

import "fmt"

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	//arr := [5]int{1, 3, 5, 7, 8}
	//fmt.Println(sum(arr))
	//fmt.Println(arr)
	left := dispatchCoin()
	fmt.Println("剩下：", left)
	fmt.Println(distribution)
}

//a. 名字中每包含1个'e'或'E'分1枚金币
//b. 名字中每包含1个'i'或'I'分2枚金币
//c. 名字中每包含1个'o'或'O'分3枚金币
//d: 名字中每包含1个'u'或'U'分4枚金币
func dispatchCoin() int {
	for _, p := range users {
		ans := 0
		for _, c := range p {
			if coins > 0 {
				if c == 'e' || c == 'E' {
					ans += 1
					coins -= 1
				} else if c == 'i' || c == 'I' {
					ans += 2
					coins -= 2
				} else if c == 'o' || c == 'O' {
					ans += 3
					coins -= 3
				} else if c == 'u' || c == 'U' {
					ans += 4
					coins -= 4
				}
			}
		}
		distribution[p] = ans
	}
	return coins
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

func chalkReplacer(chalk []int, k int) int {
	sum := 0
	for _, v := range chalk {
		sum += v
	}
	k %= sum
	ans := 0
	for i, v := range chalk {
		k -= v
		if k < 0 {
			ans = i
			break
		}
	}
	return ans
}

//写了十几篇题解了，都没有超过10个赞的。今天有机会！
//遍历的代价太大了，而且反正要求的是二进制不连续的1；
//那就用dfs从1开始，按位增。如果当前结尾是1的话就补0；
//如果是0的话就补0或者1；
var ans int = 0
var g_n int

func findIntegers(n int) int {
	g_n = n
	ans = 1
	dfs(1)
	return ans
}

func dfs(n int) {
	if n > g_n {
		return
	}
	ans++
	if n&1 == 1 {
		dfs(n << 1)
	} else {
		dfs(n << 1)
		dfs((n << 1) + 1)
	}
}
