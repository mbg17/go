package main

import "fmt"

func main() {
	f := outer("luyuan")
	f()

	x, y := calc(100, 200)
	fmt.Println(x, y)

	f1, f2 := caleFunc(100)
	fmt.Println(f1(200), f2(300))
}

// 闭包函数
func outer(name string) func() {
	return func() {
		fmt.Println(name)
	}
}

// 多返回值函数
func calc(a, b int) (int, int) {
	add := a + b
	sub := a - b
	return add, sub
}

// 多函数返回
func caleFunc(num int) (func(int) int, func(int) int) {
	sub := func(x int) int {
		return num - x
	}
	add := func(y int) int {
		return num + y
	}
	return sub, add
}
