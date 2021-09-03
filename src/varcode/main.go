package main

// 变量必须使用否则会报错
func main() {
	//标准赋值
	var name string = "陆远"
	var age int = 10
	var isMan bool = true
	println(name, age, isMan)
	// 批量赋值+变量推导
	var (
		name1  = "ceshi"
		age1   = 10
		isMan1 = false
	)
	println(name1, age1, isMan1)
	// 匿名变量
	name2 := "haha"
	age2 := 18
	isMan3 := true
	println(name2, age2, isMan3)
}
