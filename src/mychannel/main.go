package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

var once sync.Once
var wg sync.WaitGroup

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
	ch1 := make(chan int64, 100)
	ch2 := make(chan int64, 100)
	//go push(ch1)
	//go sum(ch1, ch2)
	//for i := range ch2 {
	//	fmt.Println(i)
	//}
	//for i := 0; i < 3; i++ {
	//	// 开启三个goroutinue
	//	go worker(ch1, ch2, i)
	//}
	//// 添加五个任务到任务列表
	//for i := 0; i < 5; i++ {
	//	ch1 <- i
	//}
	//// 关闭信道
	//close(ch1)
	//for i := 0; i < 5; i++ {
	//	fmt.Println(<-ch2)
	//}
	go job(ch1)
	for i := 0; i < 24; i++ {
		wg.Add(1)
		go result(ch1, ch2)
	}
	wg.Wait()
	close(ch2)
	for i := range ch2 {
		fmt.Println(i)
	}
}

// 工作池
func worker(job <-chan int, work chan<- int, i int) {
	// 获取任务添加结果到结果通道
	for i := range job {
		work <- i * i
	}
}

//使用goroutine和channel实现一个计算int64随机数各位数和的程序。
//开启一个goroutine循环生成int64类型的随机数，发送到jobChan
//开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
//主goroutine从resultChan取出结果并打印到终端输出

func job(jobchan chan<- int64) {
	defer wg.Done()
	wg.Add(1)
	for i := 0; i < 24; i++ {
		num := rand.Int63()
		jobchan <- num
	}
	close(jobchan)
}

func result(jobchan <-chan int64, results chan<- int64) {
	defer wg.Done()
	for i := range jobchan {
		int64Num := i
		lenNum := len(strconv.FormatInt(int64Num, 10))
		var sum int64 = 0
		for i := 1; i <= lenNum; i++ {
			last := int64Num % 10
			sum = sum + last
			int64Num = int64Num / 10
		}
		results <- sum
	}
}

func findMinMoves(machines []int) int {
	leng := len(machines)
	all := 0
	for _, v := range machines {
		all += v
	}
	if all%leng != 0 {
		return -1
	}
	ints := make([]int, leng)
	div := all / leng
	for i, v := range machines {
		ints[i] = v - div
	}
	ans := ^int(^uint(0) >> 1)
	for i := 0; i < len(ints); i++ {
		num := ints[i]
		if num < 0 {
			num *= -1
		}
		if num > ans {
			ans = num
		}
		if ints[i+1] > ans {
			ans = ints[i+1]
		}
		ints[i+1] += ints[i]
	}
	if ans > ints[leng-1] {
		return ans
	} else {
		return ints[leng-1]
	}
}
