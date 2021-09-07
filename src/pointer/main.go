package main

import "fmt"

/*
 &a 获取指针
 *a 获取值
 *a = 100 通过指针修改值
*/
func main() {
	a := 10
	fmt.Println(a)
	changNum(a)
	fmt.Println(a)
	changePointer(&a)
	fmt.Println(a)
	// 初始化slice
	arrSlice := make([]int, 0)
	fmt.Println(arrSlice)
	// 切片追加元素
	arrSlice = append(arrSlice, 1, 2, 3, 4)
	fmt.Println(arrSlice)
	// 初始化map
	myMap := make(map[string][]int)
	fmt.Println(myMap)
	// map添加元素
	myMap["luyuan"] = []int{1, 2, 3, 4}
	fmt.Println(myMap)
	// 指针初始化
	num := new(int) // 返回的是指针
	*num = 100
	fmt.Println(num)
	fmt.Println(*num) // 获取指针的值
}

func changNum(num int) {
	num = 100
}

func changePointer(num *int) {
	*num = 200
}
