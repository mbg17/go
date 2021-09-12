package main

import (
	"fmt"
	"sync"
)

// 用goroutine 实现并发编程
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go f1()
	go f2()
	fmt.Println("main")
	wg.Wait()
}

func f1() {
	fmt.Println("f1")
	wg.Done()
}

func f2() {
	fmt.Println("f2")
	wg.Done()
}
