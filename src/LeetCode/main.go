package main

import (
	"fmt"
	"sort"
)

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
	a := 1
	b := 2
	a = a ^ b
	b = b ^ a
	a = a ^ b
	fmt.Println(1 << 2)
	fmt.Println(-1 << 2)
	fmt.Println(a, b)
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

func checkValidString(s string) bool {
	lefts := []int{}
	stars := []int{}
	for i, c := range s {
		if c == ' ' {
			continue
		}
		if c == '(' {
			lefts = append(lefts, i)
		} else if c == '*' {
			stars = append(stars, i)
		} else {
			if len(lefts) > 0 {
				lefts = lefts[:len(lefts)-1]
			} else if len(stars) > 0 {
				stars = stars[:len(stars)-1]
			} else {
				return false
			}
		}
	}
	for len(lefts) > 0 && len(stars) > 0 {
		l := lefts[0]
		s := stars[0]
		if l > s {
			return false
		}
		lefts = lefts[:len(lefts)-1]
		stars = stars[:len(stars)-1]
	}
	return len(lefts) == 0
}
func numberOfBoomerangs(points [][]int) int {
	ans := 0
	for i := 0; i < len(points); i++ {
		ansMap := make(map[int]int)
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}
			x := points[i][0] - points[j][0]
			y := points[i][1] - points[j][1]
			point := x*x + y*y
			ansMap[point] = ansMap[point] + 1
		}
		for _, v := range ansMap {
			ans += v * (v - 1)
		}
	}
	return ans
}

func findLongestWord(s string, dictionary []string) string {
	ans := ""
	max := -1
	sort.Strings(dictionary)
	for _, d := range dictionary {
		sI := 0
		dI := 0
		for sI < len(s) && dI < len(d) {
			if s[sI] == d[dI] {
				dI++
			}
			sI++
		}
		if dI == len(d) && len(d) > max {
			max = len(d)
			ans = d
		}
	}
	return ans
}

func isStraight(nums []int) bool {
	zero := 0
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			zero++
			continue
		}
		if nums[i+1]-nums[i] != 0 {
			if nums[i+1]-nums[i]-1 > zero || nums[i+1] == nums[i] {
				return false
			}
			zero -= nums[i+1] - nums[i] - 1
		}
	}
	return true
}

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	a := (ax2 - ax1) * (ay2 - ay1)
	b := (bx2 - bx1) * (by2 - by1)
	x1 := 0
	x2 := 0
	y1 := 0
	y2 := 0
	if (by2 < ay2 && by2 > ay1) || ay2 == by2 {
		y1 = by2
	}
	if (ay2 < by2 && ay2 > by1) || ay2 == by2 {
		y1 = ay2
	}
	if (by1 > ay1 && by1 < ay2) || ay1 == by1 {
		y2 = by1
	}
	if (ay1 > by1 && ay1 < by2) || ay1 == by1 {
		y2 = ay1
	}
	if (bx1 > ax1 && bx1 < ax2) || bx1 == ax1 {
		x1 = bx1
	}
	if (ax1 > bx1 && ax1 < bx2) || bx1 == ax1 {
		x1 = ax1
	}
	if (bx2 < ax2 && bx2 > ax1) || bx2 == ax2 {
		x2 = bx2
	}
	if (ax2 < bx2 && ax2 > bx1) || ax2 == bx2 {
		x2 = ax2
	}
	return a + b - (y1-y2)*(x2-x1)
}

func destCity(paths [][]string) string {
	pathMap := make(map[string]string)
	allPath := make(map[string]string)
	for _, v := range paths {
		pathMap[v[0]] = v[1]
		allPath[v[0]] = ""
		allPath[v[1]] = ""
	}
	ans := ""
	for k, _ := range allPath {
		city := pathMap[k]
		if city == "" {
			ans = k
			break
		}
	}
	return ans
}

func thirdMax(nums []int) int {
	ans := len(nums) - 1
	num := 1
	sort.IntSlice.Sort(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 || nums[i] == nums[i+1] {
			continue
		} else {
			ans = i
			num++
		}
		if num == 3 {
			return nums[ans]
		}
	}
	return nums[len(nums)-1]
}
