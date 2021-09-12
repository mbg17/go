package main

import "fmt"

// 实现一个worker池
// 实现一个计算平方的channel
func push(ch1 chan<- int) {
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}
func sum(ch1 <-chan int, ch2 chan<- int) {
	for i := range ch1 {
		ch2 <- i * i
	}
	close(ch2)
}

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	//go push(ch1)
	//go sum(ch1, ch2)
	//for i := range ch2 {
	//	fmt.Println(i)
	//}
	for i := 0; i < 3; i++ {
		// 开启三个goroutinue
		go worker(ch1, ch2, i)
	}
	// 添加五个任务到任务列表
	for i := 0; i < 5; i++ {
		ch1 <- i
	}
	// 关闭信道
	close(ch1)
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch2)
	}
}

// 工作池
func worker(job <-chan int, work chan<- int, i int) {
	// 获取任务添加结果到结果通道
	for i := range job {
		work <- i * i
	}
}
