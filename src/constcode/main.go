package main

//常量练习
func main() {
	// 默认赋值
	const name = "luyuan"
	const age = 18
	const isMan = true
	println(name, age, isMan)
	// 批量赋值，不能先声明后赋值
	const (
		name1  = "ceshi"
		age1   = 20
		isMan1 = false
	)
	println(name1, age1, isMan1)
	// 计数器 遇到const重置为0
	const (
		n1 = iota
		n2 // 不赋值默认和之前一样
		n3
		_ // 跳过
		n4
	)
	println(n1, n2, n3, n4)

}
