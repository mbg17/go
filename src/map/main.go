package main

import (
	"fmt"
	"strings"
)

// 统计字符串出现的次数
func main() {
	split := strings.Split("how do you do", " ")
	countMap := map[string]int{}
	for _, s := range split {
		v, _ := countMap[s]
		countMap[s] = v + 1
	}
	for k, v := range countMap {
		fmt.Println(k, v)
	}
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s) //[1,2,3]
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)         // [1,3]
	fmt.Printf("%+v\n", m["q1mi"]) //[1,3,3]
}
