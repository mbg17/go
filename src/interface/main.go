package main

import "fmt"

// 定义一个接口
// 实现了以下方法都是该接口
type worker interface {
	working()
	drinking()
}

type shoppper interface {
	shopping()
}

type people interface {
	worker
	shoppper
}

func shop(shoppper shoppper) {
	shoppper.shopping()
}

func work(worker worker) {
	worker.working()
}

func drink(worker worker) {
	worker.drinking()
}
func p(people2 people) {
	fmt.Println("我是个人")
	people2.working()
	people2.shopping()
	people2.drinking()
}

type me struct {
	name string
	age  int
}

func (me me) working() {
	fmt.Printf("%s在工作\n", me.name)
}

func (me me) drinking() {
	fmt.Printf("%s在工作喝水\n", me.name)
}

func (me me) shopping() {
	fmt.Printf("%s在购物\n", me.name)
}

func main() {
	me := me{
		name: "luyuan",
		age:  18,
	}
	shop(me)
	work(me)
	drink(me)
	p(me)
}
